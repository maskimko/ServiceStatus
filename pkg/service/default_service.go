package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

const systemctl = "systemctl"

func NewDefaultService(name string) *DefaultService {
	return &DefaultService{Name: name}
}

type DefaultService struct {
	Name string
}

func (ds *DefaultService) Start() error {
	cmd := exec.Command(systemctl, "start", ds.Name)
	err := cmd.Run()
	return err
}

func (ds *DefaultService) Stop() error {
	cmd := exec.Command(systemctl, "stop", ds.getServiceUnitName())
	err := cmd.Run()
	return err
}

func (ds *DefaultService) Status() (*Status, error) {
	active, err := ds.getIsActive()
	if err != nil {
		log.Print("Failed to check active state")
		return nil, err
	}
	running, err := ds.getIsRunning()
	if err != nil {
		log.Print("Failed to check running state")
		return nil, err
	}
	enabled, err := ds.getIsEnabled()
	if err != nil {
		log.Print("Failed to check enabled state")
		return nil, err
	}
	loaded, err := ds.getIsLoaded()
	if err != nil {
		log.Print("Failed to check loaded state")
		return nil, err
	}
	pid, err := ds.getMainPid()
	if err != nil {
		log.Print("Failed to check main pid")
		return nil, err
	}
	text, err := ds.getStatusText()
	if err != nil {
		log.Print("Failed to read status info")
		return nil, err
	}
	status := Status{
		IsActive:  active,
		IsRunning: running,
		IsLoaded:  loaded,
		IsEnabled: enabled,
		MainPID:   *pid,
		text:      text,
	}
	return &status, nil
}

func (ds *DefaultService) getServiceUnitName() string {
	return fmt.Sprintf("%s.service", ds.Name)
}

func (ds *DefaultService) getIsRunning() (bool, error) {
	value, err := ds.showParam("SubState")
	if err != nil {
		return false, err
	}
	if value == "running" {
		return true, nil
	}
	return false, nil
}

func (ds *DefaultService) getIsActive() (bool, error) {
	value, err := ds.showParam("ActiveState")
	if err != nil {
		return false, err
	}
	if value == "active" {
		return true, nil
	}
	return false, nil
}

func (ds *DefaultService) getIsLoaded() (bool, error) {
	value, err := ds.showParam("LoadState")
	if err != nil {
		return false, err
	}
	if value == "loaded" {
		return true, nil
	}
	return false, nil
}

func (ds *DefaultService) getIsEnabled() (bool, error) {
	value, err := ds.showParam("UnitFileState")
	if err != nil {
		return false, err
	}
	if value == "enabled" {
		return true, nil
	}
	return false, nil
}

func (ds *DefaultService) getMainPid() (*PID, error) {
	value, err := ds.showParam("MainPID")
	if err != nil {
		return nil, err
	}
	pid, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &PID{Id: pid}, nil
}

func (ds *DefaultService) showParam(param string) (string, error) {
	var buf bytes.Buffer
	var errBuf bytes.Buffer
	cmd := exec.Command(systemctl, "show", ds.Name, "-p", param, "--value")
	cmd.Stdout = &buf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	state, err := ioutil.ReadAll(&buf)
	if err != nil {
		sErrOut, e := ioutil.ReadAll(&errBuf)
		if e != nil {
			return "", err
		}
		return "", fmt.Errorf("%w stderr: %s", string(sErrOut))
	}
	return strings.ToLower(strings.TrimSpace(string(state))), nil
}

func (ds *DefaultService) getStatusText() (string, error) {
	var buf bytes.Buffer
	cmd := exec.Command(systemctl, "status", ds.Name)
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		var canProceed = false
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() == 3 {
					canProceed = true
				}
			}
		}
		if !canProceed {
			return "", err
		}
	}
	status, err := ioutil.ReadAll(&buf)
	if err != nil {
		return "", err
	}
	return string(status), nil
}
