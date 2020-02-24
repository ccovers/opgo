package ha

import (
	"lib/common/clog"
	"lib/common/cutil"
	"lib/context"
	"lib/etcd-client"
)

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type CfgParse func(cfg string) (interface{}, error)

type ModeMaster struct {
	ParseFunc   CfgParse
	Dir         string
	members     map[string]interface{}
	KeysAPI     client.KeysAPI
	member_lock *sync.RWMutex
}

func (m *ModeMaster) GetConfigString(dir string) (interface{}, bool) {
	var environment = os.Getenv("GOENV")
	if environment == "" {
		environment = "online"
	} else {
		environment = strings.ToLower(environment)
	}

	m.member_lock.RLock()
	defer m.member_lock.RUnlock()

	cfg, ok := m.members[dir+"/"+environment]
	return cfg, ok
}

func NewModeMaster(endpoints []string, dir string, parse_func CfgParse) (*ModeMaster, error) {
	var environment = os.Getenv("GOENV")
	if environment == "" {
		environment = "online"
	} else {
		environment = strings.ToLower(environment)
	}

	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		clog.Logger.Error("new etcd error: %s", err.Error())
		return nil, err
	}

	master := &ModeMaster{
		ParseFunc:   parse_func,
		Dir:         dir,
		members:     make(map[string]interface{}),
		KeysAPI:     client.NewKeysAPI(etcdClient),
		member_lock: new(sync.RWMutex),
	}

	rsp, err := master.KeysAPI.Get(context.Background(), dir+"/"+environment, nil)
	if err != nil {
		clog.Logger.Error("read config [%s] from etcd error: %v", dir+"/"+environment, err)
		return nil, err
	}
	if rsp.Node == nil {
		err = fmt.Errorf("初始化节点信息为空: %s", dir+"/"+environment)
		return nil, err
	}
	master.AddModeMaster(dir+"/"+environment, rsp.Node.Value)

	go master.WatchModeWorkers()
	return master, nil
}

func (m *ModeMaster) WatchModeWorkers() {
	watcher := m.KeysAPI.Watcher(m.Dir, &client.WatcherOptions{
		Recursive: true,
	})

	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			clog.Logger.Error("Error watch workers: %s", err.Error())
			break
		}

		if res.Action == "set" || res.Action == "update" {
			err := m.AddModeMaster(res.Node.Key, res.Node.Value)
			if err != nil {
				clog.Logger.Error("更新节点信息失败[%s]: %s", res.Node.Key, err.Error())
			} else {
				clog.Logger.Info("更新节点信息成功[%s]", res.Node.Key)
			}
		} else if res.Action == "delete" {
			m.DeleteModeMaster(res.Node.Key)
			clog.Logger.Error("删除节点: %s", res.Node.Key)
		}
	}
}

func (m *ModeMaster) AddModeMaster(nodeKey string, nodeValue string) error {
	cfg, err := m.ParseFunc(nodeValue)
	if err != nil {
		cutil.BusinessAlarm(fmt.Sprintf("更新节点信息失败[%s]: %s", nodeKey, err.Error()))
		return err
	}

	m.member_lock.Lock()
	defer m.member_lock.Unlock()
	m.members[nodeKey] = cfg
	return nil
}

func (m *ModeMaster) DeleteModeMaster(nodeKey string) {
	m.member_lock.Lock()
	defer m.member_lock.Unlock()
	delete(m.members, nodeKey)
}

func WriteToEtcd(endpoints []string, dir string, value string) error {
	var environment = os.Getenv("GOENV")
	if environment == "" {
		environment = "online"
	} else {
		environment = strings.ToLower(environment)
	}

	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		fmt.Printf("new etcd error: %s\n", err.Error())
		return err
	}

	keys_api := client.NewKeysAPI(etcdClient)

	_, err = keys_api.Set(context.Background(), dir+"/"+environment, value, &client.SetOptions{
		TTL: -1,
	})
	if err != nil {
		fmt.Printf("Error update workerInfo: %s\n", err.Error())
		return err
	}

	return nil
}
