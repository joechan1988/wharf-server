package res

// Load Balancer endpoints

import (
	"github.com/coreos/etcd/client"
	"context"
	"github.com/golang/glog"
	"wharf/wharf-server/src/backend/common"
	"reflect"
)

type Endpoint struct {
	Name       string
	Address    string
	Type       string //apiserver,etcd-clusterdata...
	Port       string
	SecurePort string
}

type EndpointList struct {
	List []Endpoint
}

func GetEndpointList() (*EndpointList, error) {
	return nil, nil
}

func CreateEndpoint(storeHandler common.StoreHandler, endpoint Endpoint) error {

	c := storeHandler.Client
	kapi := client.NewKeysAPI(c)

	t := reflect.TypeOf(endpoint)
	v := reflect.ValueOf(endpoint)

	for k := 0; k < t.NumField(); k++ {
		resp, err := kapi.Set(context.Background(),
			"/endpoint/"+endpoint.Name+"/"+t.Field(k).Name,
			v.Field(k).String(),
			nil)

		if err != nil {
			glog.V(1).Infof("Failed to set key: %s", endpoint.Name)
			return err
		} else {
			glog.Infof("Response: %q\n", resp)
		}
	}

	return nil

}
