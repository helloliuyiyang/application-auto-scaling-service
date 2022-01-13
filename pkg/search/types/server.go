package types

import (
	"net"
)

const (
	defaultMaxProcess         = 20
	defaultProcessConcurrency = 10
)

type Server struct {
	Id string
	IP net.IP
	//NowConcurrency int

	MaxProcessNum int
	ProcessConfig *ProcessConfig
	//Processes     []*Process
}

type ProcessConfig struct {
	MaxConcurrency int64
}

func NewServer(id string, ip net.IP) *Server {
	return &Server{
		Id:            id,
		IP:            ip,
		MaxProcessNum: defaultMaxProcess,
		ProcessConfig: &ProcessConfig{MaxConcurrency: defaultProcessConcurrency},
		//Processes:     []*Process{},
	}
}
