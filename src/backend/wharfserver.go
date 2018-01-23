package main

import (
	flag "github.com/spf13/pflag"
	"log"
	"os"
	goflag "flag"
	"net/http"
	"wharf/wharf-server/src/backend/handler"
	"fmt"
	"net"
	"wharf/wharf-server/src/backend/common"
	//"wharf/wharf-server/src/backend/res"
	"github.com/golang/glog"
)

var (
	argPort         = flag.Int("port", 8888, "The port to listen to for incoming HTTP requests")
	argAddr         = flag.IP("advertise-address", net.IPv4(0, 0, 0, 0), "Advertised address")
	argEtcdEndpoint = flag.String("etcd-endpoint", "http://127.0.0.1:2379", "The endpoint address of etcd store backend")
)

func main() {

	defer func(){
		common.HandlePanic()
		glog.Flush()
	}()

	log.SetOutput(os.Stdout)
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	goflag.CommandLine.Parse([]string{})  //override glog enforcing to parse goflag
	flag.Parse()
	flag.Set("logtostderr","true")

	storeHandler, err := common.CreateStoreHandler(*argEtcdEndpoint)
	if err != nil {
		//log.Fatalf("Failed to Create API Handler...")
		panic("Failed to Create Store Handler...")
	}

	apiHander, err := handler.CreateHTTPAPIHandler(storeHandler)
	if err != nil {
		//log.Fatalf("Failed to Create API Handler...")
		panic("Failed to Create API Handler...")
	}
	http.Handle("/", apiHander)


	//res.CreateEndpoint(*storeHandler, res.Endpoint{Name: "etcd-metadata", Address: "127.0.0.1", Type: "etcd-metadata"})

	listenAddr := fmt.Sprintf("%s:%d", *argAddr, *argPort)
	log.Printf("Serving on HTTP address %s:%d", *argAddr, *argPort)
	go func() {
		common.HandleFatalError(http.ListenAndServe(listenAddr, nil))
	}()

	select {}
}
