package common

import "github.com/golang/glog"

func HandleFatalError(err error) {
	glog.Fatalf("Error occurs: %s\n", err)
}

func HandleError(err error){
	glog.Errorf("[ERR] %s\n",err)
}

func HandlePanic() {
	if err := recover(); err != nil {
		glog.Fatalf("Exception caught: %s\n",err)
	}
}
