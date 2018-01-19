package backend

import (
	"github.com/spf13/pflag"
	"log"
	"os"
	"flag"
	"net/http"
	"wharf/wharf-server/src/backend/handler"
	"fmt"
	"net"
)

var (
	argPort = pflag.Int("port", 8080, "The port to listen to for incoming HTTP requests")
	argAddr = pflag.IP("advertise-address",net.IPv4(0,0,0,0),"Advertised address")
)

func main() {

	defer handler.HandlePanic()

	log.SetOutput(os.Stdout)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	apiHander, err := handler.CreateHTTPAPIHandler()

	if err != nil {
		log.Fatalf("Hander Error")
		panic("Failed")
	}

	http.Handle("/", apiHander)

	listenAddr := fmt.Sprintf("%s:%d", *argAddr, *argPort)

	log.Printf("Serving on HTTP address %s:%d",*argAddr,*argPort)

	go func() {
		handler.HandleFatalError(http.ListenAndServe(listenAddr, nil))
		}()

	select {}
}
