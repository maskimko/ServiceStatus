package examples

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
func GetSpawnedPIDs(serviceName string) ([]int, error) {
	srv := service.NewDefaultService(serviceName)
	status, err := srv.Status()
	if err != nil {
		return nil, err
	}
	var spawnedPIDs []int
	for _, p := range status.MainPID.Children {
		spawnedPIDs = append(spawnedPIDs, p.ID)
	}
	return spawnedPIDs, nil
}
