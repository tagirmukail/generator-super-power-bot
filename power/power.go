package power

import (
	"encoding/json"
	"generator-super-power-bot/models"
	"io/ioutil"
	"sync"
)

type PowersCache struct {
	sync.RWMutex
	Cache map[string]string // represent to map[power name]description
}

func NewPowersCache(path string) (*PowersCache, error) {
	var powers []models.Power

	var powCache = &PowersCache{
		Cache:make(map[string]string),
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, powers)
	if err != nil {
		return nil, err
	}

	for _, power := range powers {
		powCache.Cache[power.PowerName] = power.Description
	}
}
