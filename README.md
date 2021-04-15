# ServiceStatus

[![Go Report Card](https://goreportcard.com/badge/github.com/maskimko/ServiceStatus)](https://goreportcard.com/report/github.com/maskimko/ServiceStatus)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/maskimko/ServiceStatus)
[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## About


*ServiceStatus* is a simple and small Go library 
which wraps some basic control function of linux systemd units of type service

It allows _starting_, _stopping_ and _querying service status_

_main.go_ file introduces a CLI interface to this wrapper

>Usage: ServiceStatus <service_unit_name> <start|stop|status>


## Examples:


### CLI:
```$ ServiceStatus NetworkManager status
Service unit NetworkManager status:
Active: true
Running: true
Loaded: true
Enabled: true
Main PID:
412 NetworkManager
1499 dhclient

Status text:
● NetworkManager.service - Network Manager
Loaded: loaded (/lib/systemd/system/NetworkManager.service; enabled; vendor preset: enabled)
Active: active (running) since Wed 2021-04-14 11:24:35 EEST; 3h 45min ago
Docs: man:NetworkManager(8)
Main PID: 412 (NetworkManager)
Tasks: 4 (limit: 1135)
Memory: 7.6M
CGroup: /system.slice/NetworkManager.service
├─ 412 /usr/sbin/NetworkManager --no-daemon
└─1499 /sbin/dhclient -d -q -sf /usr/lib/NetworkManager/nm-dhcp-helper -pf /run/dhclient-enp0s3.pid -lf /var/lib/NetworkManager/dhclient-7473cee9-ad3c-476f-ab76-62148a8bec54-enp0s3.lease -cf /var/lib/NetworkManager/dhclient-enp0s3.conf enp0s3
```

### Golang
```go

import "github.com/maskimko/ServiceStatus/pkg/service"

// GetStatusOfSomeService returns a status of some systemd service unit
func GetStatusOfSomeService(serviceName string) *service.Status {
srv := service.NewDefaultService(serviceName)
status, _ := srv.Status()
return status
}

// StartSomeService starts some systemd service unit
func StartSomeService(serviceName string) {
srv := service.NewDefaultService(serviceName)
_ = srv.Start()
}

// GetSpawnedPIDs gets a list of spawned by service sub-process process ids
func GetSpawnedPIDs(serviceName string) ([]int,error) {
srv := service.NewDefaultService(serviceName)
status, err := srv.Status()
if err != nil {
return nil, err
}
var spawnedPIDs []int
for _, p := range status.MainPID.Children {
spawnedPIDs = append(spawnedPIDs, p.Id)
}
return spawnedPIDs, nil
}
```


