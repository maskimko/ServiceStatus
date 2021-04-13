package ServiceStatus

import (
	"./pkg/service"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <service unit name> <start|stop|status>", os.Args[0])
	}
	if len(os.Args[1]) == 0 {
		log.Fatalf("You must provide a service unit name\nUsage: %s <service unit name> <start|stop|status>", os.Args[0])
	}
	srv := service.NewDefaultService(os.Args[1])
	switch os.Args[2] {
	case "start":
	case "stop":
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
