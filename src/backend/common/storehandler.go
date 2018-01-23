package common

import (
	"github.com/coreos/etcd/client"
	"github.com/golang/glog"
)

type StoreHandler struct {
	Client client.Client
}

func CreateStoreHandler(endpoint string) (*StoreHandler,error) {

	cfg := client.Config{Endpoints: []string{endpoint}}
	c, err := client.New(cfg)
	if err != nil {
		glog.Error("Failed to connect to etcd backend")
	}

	storeHandler := StoreHandler{Client: c}

	return &storeHandler,err
}

func (storeHandler *StoreHandler) CreateDirectory(name string) error {

	//c := storeHandler.Client
	//kapi := client.NewKeysAPI(c)
	//resp, err := kapi.CreateInOrder(context.Background(), "/endpoint", endpoint.Name, &client.SetOptions{Dir: true})

	return nil
}
