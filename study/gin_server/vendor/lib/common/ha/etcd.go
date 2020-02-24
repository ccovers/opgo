package ha

import (
	"context"
	"encoding/json"
	"fmt"
	"lib/etcd-client"
	"os"
	"strings"
)

var EtcdClient *KeysAPI
var defaultEtcdAddr []string = []string{"http://etcd.in.netwa.cn:2379"}

func init() {
	EtcdClient, _ = newDefaultClient(defaultEtcdAddr...)
}

type KeysAPI struct {
	addr []string
	api  client.KeysAPI
}

func NewEtcdClient(addrlist ...string) error {
	var err error
	EtcdClient, err = newDefaultClient(addrlist...)
	return err
}

func newDefaultClient(addrlist ...string) (*KeysAPI, error) {
	if len(addrlist) > 0 {
		defaultEtcdAddr = addrlist
	}
	cfg := client.Config{
		Endpoints:               defaultEtcdAddr,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: TTL,
	}
	c, err := client.New(cfg)
	if err != nil {
		fmt.Printf("new etcd api where addr[%+v] err:%+v", defaultEtcdAddr, err)
		return nil, err
	}

	ec := &KeysAPI{
		addr: addrlist,
		api:  client.NewKeysAPI(c),
	}
	return ec, nil
}
func (c *KeysAPI) Api() client.KeysAPI {
	return c.api
}

func (c *KeysAPI) GetServerConfig(servername string, cfg interface{}) error {
	var environment = os.Getenv("GOENV")
	if environment == "" {
		environment = "online"
	} else {
		environment = strings.ToLower(environment)
	}
	fmt.Println("environment", environment)

	rsp, err := c.api.Get(context.Background(), c.ConfigKey(servername, environment), nil)
	if err != nil {
		fmt.Printf("read config [%s:%s] from etcd error:%v", servername, environment, err)
		return err
	}

	if rsp.Node == nil {
		fmt.Printf("empty etcd node")
		return fmt.Errorf("empty etcd node")
	}
	return json.Unmarshal([]byte(rsp.Node.Value), cfg)

}

func (c *KeysAPI) GetValueStr(servername string) (string, error) {
	var environment = os.Getenv("GOENV")
	if environment == "" {
		environment = "online"
	} else {
		environment = strings.ToLower(environment)
	}
	fmt.Println("environment", environment)

	rsp, err := c.api.Get(context.Background(), c.ConfigKey(servername, environment), nil)
	if err != nil {
		fmt.Printf("read config [%s:%s] from etcd error:%v", servername, environment, err)
		return "", err
	}

	if rsp.Node == nil {
		fmt.Printf("empty etcd node")
		return "", fmt.Errorf("empty etcd node")
	}

	return rsp.Node.Value, nil
}

func (c *KeysAPI) ConfigKey(service, env string) string {
	return fmt.Sprintf("/config/%s/%s", service, env)
}

func (c *KeysAPI) PositionKey(service, group string) string {
	return fmt.Sprintf("/position/%s/%s", service, group)

}

func (c *KeysAPI) GetConfig(path string, cfg interface{}) error {
	rsp, err := c.api.Get(context.Background(), path, nil)
	if err != nil {
		fmt.Printf("read config [%s:%s] from etcd error:%v", path, environment, err)
		return err
	}

	if rsp.Node == nil {
		fmt.Printf("empty etcd node")
		return fmt.Errorf("empty etcd node")
	}
	return json.Unmarshal([]byte(rsp.Node.Value), cfg)
}
