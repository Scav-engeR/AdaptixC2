package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

const (
	TYPE_LISTENER = "listener"

	TYPE_INTERNAL = "internal"
	TYPE_EXTERNAL = "external"
)

type Teamserver interface{}

type ModuleExtender struct {
	ts Teamserver
}

type ModuleInfo struct {
	ModuleName string
	ModuleType string
}

type ListenerInfo struct {
	ListenerType     string
	ListenerProtocol string
	ListenerName     string
	ListenerUI       string
}

type ListenerData struct {
	Name      string `json:"l_name"`
	Type      string `json:"l_type"`
	BindHost  string `json:"l_bind_host"`
	BindPort  string `json:"l_bind_port"`
	AgentHost string `json:"l_agent_host"`
	AgentPort string `json:"l_agent_port"`
	Status    string `json:"l_status"`
}

var ModuleObject ModuleExtender

////////////////////////////

const (
	SetType     = TYPE_EXTERNAL
	SetProtocol = "http"
	SetName     = "BeaconHTTP"
	SetUiPath   = "ui_listener.json"
)

////////////////////////////

func (m *ModuleExtender) InitPlugin(ts any) ([]byte, error) {
	var (
		buffer bytes.Buffer
		err    error
	)

	ModuleObject.ts = ts.(Teamserver)

	info := ModuleInfo{
		ModuleName: "BeaconHTTP",
		ModuleType: TYPE_LISTENER,
	}

	err = json.NewEncoder(&buffer).Encode(info)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (m *ModuleExtender) ListenerInit() ([]byte, error) {
	var (
		buffer bytes.Buffer
		err    error
	)

	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	uiPath := filepath.Join(dir, SetUiPath)
	listenerUI, err := os.ReadFile(uiPath)
	if err != nil {
		return nil, err
	}

	info := ListenerInfo{
		ListenerType:     SetType,
		ListenerProtocol: SetProtocol,
		ListenerName:     SetName,
		ListenerUI:       string(listenerUI),
	}

	err = json.NewEncoder(&buffer).Encode(info)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

////////////////////////////

func (m *ModuleExtender) ListenerValid(data string) error {
	var (
		err  error
		conf HTTPConfig
	)

	err = json.Unmarshal([]byte(data), &conf)
	if err != nil {
		return err
	}

	if conf.Host == "" {
		return errors.New("host is required")
	}

	portInt, err := strconv.Atoi(conf.Port)
	if err != nil {
		return errors.New("port must be an integer")
	}

	if portInt < 1 || portInt > 65535 {
		return errors.New("port must be in the range 1-65535")
	}

	matched, err := regexp.MatchString(`^/[a-zA-Z0-9]+(/[a-zA-Z0-9]+)*$`, conf.Uri)
	if err != nil || !matched {
		return errors.New("uri invalid")
	}

	return nil
}

func (m *ModuleExtender) ListenerStart(data string) ([]byte, error) {
	var (
		err          error
		conf         HTTPConfig
		listener     *HTTP
		listenerData ListenerData
		buffer       bytes.Buffer
	)

	err = json.Unmarshal([]byte(data), &conf)
	if err != nil {
		return nil, err
	}

	listener = NewConfigHttp()
	listener.Config = conf

	err = listener.Start()
	if err != nil {
		return nil, err
	}

	listenerData = ListenerData{
		BindHost:  listener.Config.Host,
		BindPort:  listener.Config.Port,
		AgentHost: listener.Config.Host,
		AgentPort: listener.Config.Port,
		Status:    "Listen",
	}

	if !listener.Active {
		listenerData.Status = "Closed"
	}

	err = json.NewEncoder(&buffer).Encode(listenerData)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
