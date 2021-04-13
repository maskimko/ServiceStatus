package service

import "fmt"

type Service interface {
	Start() error
	Stop() error
	Status() (*Status, error)
}

type Status struct {
	IsActive  bool
	IsRunning bool
	IsLoaded  bool
	IsEnabled bool
	MainPID   PID
	text      string
}

func (s *Status) String() string {
	return fmt.Sprintf("Active: %t\n"+
		"Running: %t\n"+
		"Loaded: %t\n"+
		"Enabled: %t\n"+
		"Main PID: %d\n"+
		"Status text:\n%s", s.IsActive, s.IsRunning, s.IsLoaded, s.IsEnabled, s.MainPID, s.text)
}

type PID struct {
	Id       int
	Children []PID
}
