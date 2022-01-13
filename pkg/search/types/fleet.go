package types

import (
	"github.com/pkg/errors"
)

type Fleet struct {
	Id   string
	Name string

	MaxConcurrency int
	NowConcurrency int

	MaxServerNum int

	// 非并发安全，目前只由一个goroutine处理
	ServersMap map[string]*Server
	//ServersMap sync.Map
	//mu         sync.Mutex
}

func NewFleet(id string, name string, maxServerNum int) *Fleet {
	return &Fleet{
		Id:             id,
		Name:           name,
		MaxConcurrency: 0,
		NowConcurrency: 0,
		MaxServerNum:   maxServerNum,
		ServersMap:     make(map[string]*Server),
	}
}

// 如果不使用锁，该方法非必要
func (f *Fleet) isServerExist(serverId string) bool {
	//f.mu.Lock()
	//defer f.mu.Unlock()

	_, ok := f.ServersMap[serverId]
	return ok
}

func (f *Fleet) AddServer(server *Server) error {
	if f.isServerExist(server.Id) {
		return errors.Errorf("server queue already has a server with id[%s]", server.Id)
	}
	if len(f.ServersMap) >= f.MaxServerNum {
		return errors.Errorf("server queue meets it max servers limit[%d]", f.MaxServerNum)
	}

	f.ServersMap[server.Id] = server
	return nil
}

func (f *Fleet) DeleteServer(serverId string) {
	delete(f.ServersMap, serverId)
}
