package ha

import (
	"lib/common/clog"
	"lib/common/errcode"
	"lib/context"
	"lib/etcd-client"
)

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"sync"
	"time"
)

type Smaster struct {
	Dir         string
	members     map[string]*Member
	KeysAPI     client.KeysAPI
	member_lock *sync.RWMutex
}

type WorkerInfo struct {
	Name      string
	Addr      string
	InnerAddr string
	ConnCount int
}

type Member struct {
	InGroup bool
	WorkerInfo
}

type Worker struct {
	WorkFunc GenWorkInfo
	KeysAPI  client.KeysAPI
	Dir      string
}

func (m *Smaster) GetAvailableAddr() (string, string, bool) {
	var node *Member = nil

	m.member_lock.Lock()
	defer m.member_lock.Unlock()

	for _, v := range m.members {
		if !v.InGroup {
			continue
		}
		if nil == node {
			node = v
		} else if v.ConnCount < node.ConnCount {
			node = v
		}
	}
	//clog.Logger.Debug("members: [%v]", m.members)

	if node != nil {
		return node.Addr, node.InnerAddr, true
	} else {
		clog.Logger.Debug("node nim: dir=%s, members=[%v]", m.Dir, m.members)
		return "", "", false
	}
}

func (m *Smaster) ValidServerAddr(serveraddr string) bool {
	var valid bool = false

	m.member_lock.Lock()
	defer m.member_lock.Unlock()

	for _, v := range m.members {
		if !v.InGroup {
			continue
		}
		if strings.Compare(v.Addr, serveraddr) == 0 {
			valid = true
			break
		}
	}
	//clog.Logger.Debug("members: [%v]", m.members)
	return valid
}

func NewSmaster(endpoints []string, dir string) (*Smaster, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	master := &Smaster{
		members:     make(map[string]*Member),
		KeysAPI:     client.NewKeysAPI(etcdClient),
		Dir:         dir,
		member_lock: new(sync.RWMutex),
	}
	go master.WatchWorkers()

	return master, nil
}

func (m *Smaster) UpdateWorker(info *WorkerInfo, nodeKey string) {
	m.member_lock.Lock()
	defer m.member_lock.Unlock()

	node, exists := m.members[nodeKey]
	if exists {
		if node.InGroup != true ||
			info.Name != node.Name ||
			info.Addr != node.Addr ||
			info.InnerAddr != node.InnerAddr ||
			info.ConnCount != node.ConnCount {
			m.members[nodeKey].InGroup = true
			m.members[nodeKey].Name = info.Name
			m.members[nodeKey].Addr = info.Addr
			m.members[nodeKey].InnerAddr = info.InnerAddr
			m.members[nodeKey].ConnCount = info.ConnCount
			clog.Logger.Debug("更新节点: %s [%v]", nodeKey, m.members)
		}
	} else {
		member := Member{
			InGroup: true,
		}
		member.Name = info.Name
		member.Addr = info.Addr
		member.InnerAddr = info.InnerAddr
		member.ConnCount = info.ConnCount
		m.members[nodeKey] = &member
		clog.Logger.Debug("创建节点: %s [%v]", nodeKey, m.members)
	}
}

func (m *Smaster) WatchWorkers() {
	api := m.KeysAPI
	watcher := api.Watcher(m.Dir+"/", &client.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			clog.Logger.Error("Error watch workers: %s", err.Error())
			break
		}

		if res.Action == "expire" {
			func() {
				m.member_lock.Lock()
				defer m.member_lock.Unlock()
				_, ok := m.members[res.Node.Key]
				if ok {
					m.members[res.Node.Key].InGroup = false
					clog.Logger.Debug("%s过期", res.Node.Key)
				}
			}()
		} else if res.Action == "set" || res.Action == "update" {
			info := WorkerInfo{}
			err := json.Unmarshal([]byte(res.Node.Value), &info)
			if err != nil {
				clog.Logger.Error("更新的节点信息格式错误: %s %s", res.Node.Value, err.Error())
			} else {
				m.UpdateWorker(&info, res.Node.Key)
			}
		} else if res.Action == "delete" {
			func() {
				m.member_lock.Lock()
				defer m.member_lock.Unlock()
				delete(m.members, res.Node.Key)
				clog.Logger.Debug("删除节点: %s", res.Node.Key)
			}()
		}
	}
}

type GenWorkInfo func() WorkerInfo

func NewWorker(dir string, endpoints []string, info_func GenWorkInfo) (*Worker, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	if len(endpoints) == 0 || info_func == nil {
		return nil, errcode.InvalidParameterError
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	w := &Worker{
		Dir:      dir,
		WorkFunc: info_func,
		KeysAPI:  client.NewKeysAPI(etcdClient),
	}
	go w.HeartBeat()
	return w, nil
}

func (w *Worker) HeartBeat() {
	api := w.KeysAPI

	for {
		info := w.WorkFunc()

		key := w.Dir + "/" + base64.StdEncoding.EncodeToString([]byte(info.Name))
		value, _ := json.Marshal(info)

		_, err := api.Set(context.Background(), key, string(value), &client.SetOptions{
			TTL: time.Second * 10,
		})
		if err != nil {
			clog.Logger.Error("Error update workerInfo: ", err.Error())
		}
		time.Sleep(time.Second * 3)
	}
}

//获取可用的节点列表
func GetAvailableNodeMap(rootdir string, program string, endpoints []string) (map[string]int, error) {
	keymap := make(map[string]int)

	if len(rootdir) == 0 ||
		len(program) == 0 ||
		len(endpoints) == 0 {
		clog.Logger.Info("etcd取值参数错误")
		return keymap, errcode.InvalidParameterError
	}

	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		clog.Logger.Error("连接etcd失败: %s", err.Error())
		return keymap, err
	}

	api := client.NewKeysAPI(etcdClient)

	nodedir := rootdir + "/" + program
	res, err := api.Get(context.Background(), nodedir, &client.GetOptions{
		Recursive: true,
		Sort:      true,
	})
	if err != nil {
		clog.Logger.Error("etcd取数据失败: %s", err.Error())
		return keymap, err
	}

	if res.Node.Nodes != nil {
		for i := 0; i < len(res.Node.Nodes); i++ {
			keymap[res.Node.Nodes[i].Key] = i
		}
	}
	return keymap, nil
}

func GetServerDir(sname string) string {
	return "server_mgr/" + sname
}
