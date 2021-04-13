package main

import (
	"fmt"
	"github.com/maskimko/ServiceStatus/pkg/service"
	"log"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS != "linux" {
		log.Fatal("Only GNU/Linux OS is supported")
	}
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <service unit name> <start|stop|status>", os.Args[0])
	}
	if len(os.Args[1]) == 0 {
		log.Fatalf("You must provide a service unit name\nUsage: %s <service unit name> <start|stop|status>", os.Args[0])
	}

	srv := service.NewDefaultService(os.Args[1])
	switch os.Args[2] {
	case "start":
		err := srv.Start()
		if err != nil {
			log.Fatalf("Failed to start service %s Error: %s", srv.Name, err.Error())
		}
	case "stop":
		err := srv.Stop()
		if err != nil {
			log.Fatalf("Failed to stop service %s Error: %s", srv.Name, err.Error())
		}
	case "status":
		status, err := srv.Status()
		if err != nil {
			log.Fatalf("Failed to check status. Error: %s", err.Error())
		}
		fmt.Printf("Service unit %s status:\n%s", srv.Name, status.String())
	default:
		log.Fatalf("Unknown command %s\nUsage: %s <service unit name> <start|stop|status>", os.Args[2], os.Args[0])

	}

}
