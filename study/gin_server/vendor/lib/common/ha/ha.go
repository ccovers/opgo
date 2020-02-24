package ha

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"lib/context"
	"lib/etcd-client"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

var environment string

func init() {
	e := os.Getenv("GOENV")
	if e == "" {
		environment = "online"
	} else {
		environment = strings.ToLower(e)
	}
}

var slaveCheckReg *regexp.Regexp = regexp.MustCompile("Key already exists")
var TTL time.Duration = 10 * time.Second

type Node struct {
	ServerName string //服务名
	GroupId    string //组别,同一组里面选出一个master
}

//如下接口都不不能阻塞运行
type Server interface {
	OnMaster(msg string) error //拨测协程唤醒主服务
	OnSlave(msg string) error  //拨测协程停止主服务
	ServerName() string
	GroupId() string
	IsRunaway() (bool, string) //服务拨测协程，服务自己是否已经退出
}

type EtcdMutex struct {
	Key  string
	Val  string
	Ttl  time.Duration
	KApi client.KeysAPI
}

func (this *EtcdMutex) Lock() error {
	opt := &client.SetOptions{PrevExist: client.PrevExist, PrevValue: this.Val, TTL: this.Ttl}
	_, err := this.KApi.Set(context.TODO(), this.Key, this.Val, opt)
	return err
}

func (this *EtcdMutex) TryLock() error {
	opt := &client.SetOptions{PrevExist: client.PrevNoExist, TTL: this.Ttl}
	_, err := this.KApi.Set(context.TODO(), this.Key, this.Val, opt)
	return err
}

func (this *EtcdMutex) UnLock() error {
	opt := &client.SetOptions{PrevExist: client.PrevExist, PrevValue: this.Val, TTL: time.Nanosecond}
	_, err := this.KApi.Set(context.TODO(), this.Key, this.Val, opt)
	return err
}

type NodeStatus string

const (
	UnKnown NodeStatus = "unKnown"
	Master  NodeStatus = "master"
	Slave   NodeStatus = "slave"
)

type Stat struct {
	AsMasterCheckFailed int64
	AsSlaveCheckFaied   int64
	SwitchToMasterCnt   int64
	SwitchToSlaveCnt    int64
}

type HAServer struct {
	node      Node
	server    Server
	ServerId  string
	CurStatus NodeStatus
	stat      Stat
}

func (p HAServer) String() string {
	return fmt.Sprintf("[%v.%v] as %+v run with server_id [%v]", p.node.ServerName, p.node.GroupId, p.CurStatus, p.ServerId)
}
func serverId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func RunAsHAServer(server Server, etcd ...string) error {
	var err error
	node := Node{ServerName: server.ServerName(), GroupId: server.GroupId()}
	for loop := true; loop; loop = false {
		if len(node.ServerName) == 0 || len(node.GroupId) == 0 {
			err = fmt.Errorf("node info is err %+v", node)
			break
		}
		if server == nil {
			err = errors.New("server is nil")
			break
		}
	}
	if err != nil {
		log.Printf("RunAsHAServer Failed:%+v", err)
		return err
	}

	ha := &HAServer{node: node, server: server, ServerId: serverId(), CurStatus: Slave}
	go ha.run()
	return nil
}

func (this *HAServer) getKey() string {
	return fmt.Sprintf("/lock/%s/%s/%s", environment, this.node.ServerName, this.node.GroupId)
}

func (this *HAServer) switchToMaster(msg string) {
	log.Printf("%+v, switch to master,msg:%v", *this, msg)
	if err := this.server.OnMaster(msg); err != nil {
		log.Printf("%+v, switch to master, err:%+v", *this, err)
		return
	}
	this.stat.SwitchToMasterCnt++
	this.CurStatus = Master
}

func (this *HAServer) switchToSlave(msg string) {
	log.Printf("%+v, switch to slave, msg:%v", *this, msg)
	if err := this.server.OnSlave(msg); err != nil {
		log.Printf("%+v, switch to slave, err:%+v", *this, err)
	}
	this.stat.SwitchToSlaveCnt++
	this.CurStatus = Slave
	return
}

func (this *HAServer) runaway() (bool, string) {
	return this.server.IsRunaway()
}

const (
	MsgNoMaster       = "no master server "
	MasterCheckingMsg = "master checking err"
)

func (this *HAServer) run() {
	log.Printf("start %+v ", *this)
	interval := TTL / 3
	if interval < time.Second {
		interval = 3 * time.Second
	}
	tick := time.NewTicker(interval)
	defer tick.Stop()
	mutex := EtcdMutex{Key: this.getKey(), Val: this.ServerId, Ttl: TTL, KApi: EtcdClient.Api()}
	var master_lock_failed int = 0
	var err error
	sig := this.listenSignal()
	for {
		select {
		case <-tick.C:
			if this.CurStatus == Master {
				if err = mutex.Lock(); err != nil {
					master_lock_failed++
					log.Printf(" server checking failed, etcd_addr:%+v, err:%+v, failed_cnt:%+v, %+v", defaultEtcdAddr, err, master_lock_failed, *this)
					if master_lock_failed == 3 {
						mutex.UnLock()
						this.switchToSlave(fmt.Errorf("%v:%v", MasterCheckingMsg, err).Error())
						master_lock_failed = 0
					}
					this.stat.AsMasterCheckFailed++
				} else {
					if is_run, msg := this.runaway(); is_run {
						log.Printf(" ---runaway--- :%+v, msg:%v", *this, msg)
						this.switchToSlave(msg)
						mutex.UnLock()
						master_lock_failed = 0
						time.Sleep(10 * time.Second)
					}
					this.stat.AsMasterCheckFailed = 0
					log.Printf(" server checking success, %+v", *this)
				}
			} else {
				if err = mutex.TryLock(); err == nil {
					this.switchToMaster(MsgNoMaster)
				} else {
					if !slaveCheckReg.MatchString(err.Error()) {
						this.stat.AsSlaveCheckFaied++
						log.Printf(" server checking failed, etcd_addr:%+v, err:%+v, %+v", defaultEtcdAddr, err, *this)
					} else {
						this.stat.AsSlaveCheckFaied = 0
						log.Printf(" server checking success, %+v", *this)
					}
				}
			}
		case s := <-sig:
			msg := fmt.Sprintf("server get signal:%v, exit", s.String())
			log.Printf(msg)
			this.switchToSlave(msg)
			os.Exit(0)
		}
	}
	log.Printf("stop %+v ", *this)
}

func (this *HAServer) listenSignal(signals ...os.Signal) <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	if len(signals) == 0 {
		signals = append(signals, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR2)
	}
	signal.Notify(sig, signals...)
	return sig
}
