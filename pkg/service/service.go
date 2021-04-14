package service

import "fmt"

//Service is a representation of a Linux systemd unit of type service
type Service interface {
	//Start starts the service (launches process)
	Start() error
	//Stop stops the service (terminates process)
	Stop() error
	//Status method return a Status object, and error if i is impossible to build up the status object
	Status() (*Status, error)
}

//Status type holds the most commonly used data about systemd service unit status
type Status struct {
	//IsActive is true if the service is active
	IsActive bool
	//IsRunning is true if the service is running
	IsRunning bool
	//IsLoaded is true if the service is loaded
	IsLoaded bool
	//IsEnabled is true if the service is enabled
	IsEnabled bool
	//MainPID is an object which holds a pid of main service process
	//and some information about executable name and its children
	MainPID PID
	//text field holds original systemctl status answer
	text string
}

//String method return a string representation of the PID object
func (s *Status) String() string {
	return fmt.Sprintf("Active: %t\n"+
		"Running: %t\n"+
		"Loaded: %t\n"+
		"Enabled: %t\n"+
		"Main PID:\n%s\n"+
		"Status text:\n%s", s.IsActive, s.IsRunning, s.IsLoaded, s.IsEnabled, s.MainPID.String(), s.text)
}
