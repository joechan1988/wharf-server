package res

import (
	"testing"
	"wharf/wharf-server/src/backend/common"
)

func TestCreateEndpoint(t *testing.T) {

	storeHandler, err := common.CreateStoreHandler("http://127.0.0.1:2379")

	if err != nil {
		t.Error(err)
	}

	endpoint := Endpoint{
		Name: "test",
		Address: "http://127.0.0.1:8888",
		Type: "new",
	}

	err = CreateEndpoint(*storeHandler, endpoint)
	if err!=nil{
		t.Error(err)
	}

}
