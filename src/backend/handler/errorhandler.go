package handler

import "log"

func HandleFatalError(err error) {
	log.Fatalf("Error occurs: %s\n", err)

}

func HandlePanic() {
	if err := recover(); err != nil {
		log.Fatalf("Exception caught: %s\n",err)
	}
}
