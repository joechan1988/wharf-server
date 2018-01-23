package res

import (
	"wharf/wharf-server/src/backend/common"
	"github.com/coreos/etcd/client"
	"context"
	"github.com/golang/glog"
)


type Token struct {
	TokenString string
}

func GetToken(storeHandler *common.StoreHandler,token string) (*Token, error) {

	var tokenResp Token
	c := storeHandler.Client
	kapi := client.NewKeysAPI(c)

	resp,err:=kapi.Get(context.Background(),"/tokens/"+token,nil)

	if err!=nil{
		common.HandleError(err)
		tokenResp.TokenString = ""
	}else{
		glog.Info(resp.Node.Key)
		glog.Infof(resp.Node.Value)
		glog.Infof("%b",resp.Node.Dir)
		tokenResp.TokenString = resp.Node.Key
		//tokenResp = Token{
		//	TokenString: resp.Node.Key,
		//}
	}
	return &tokenResp, nil

}

func StoreToken() (error) {
	return nil
}

func GetTokenFromStore() (error) {
	return nil
}
