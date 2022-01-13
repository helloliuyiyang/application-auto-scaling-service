package cache

import (
	"sync"

	"nanto.io/application-auto-scaling-service/pkg/search/types"
)

var (
	cache *FleetsCache
	once  sync.Once
)

type FleetsCache struct {
	fleetsMap map[string]*types.Fleet
	//mu          sync.Mutex
}

func GetFleetsCache() *FleetsCache {
	once.Do(func() {
		cache = &FleetsCache{
			fleetsMap: make(map[string]*types.Fleet),
		}
	})
	return cache
}

func (c *FleetsCache) AddFleet(fleet *types.Fleet) {
	c.fleetsMap[fleet.Id] = fleet
}

func (c *FleetsCache) Snapshot() []types.Fleet {
	fleets := []types.Fleet{}
	for _, f := range c.fleetsMap {
		fleets = append(fleets, *f)
	}
	return fleets
}
