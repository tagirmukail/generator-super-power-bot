package power

import (
	"encoding/json"
	"generator-super-power-bot/consts"
	"generator-super-power-bot/models"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"
)

type PowersCache struct {
	sync.RWMutex
	Powers []*models.Power // represent to map[power name]description
}

func NewPowersCache() (*PowersCache, error) {
	var powCache = &PowersCache{}

	data, err := ioutil.ReadFile(consts.POWERS_PATH)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &powCache.Powers)
	if err != nil {
		return nil, err
	}

	return powCache, nil
}

func (p *PowersCache) GetRandomPower() *models.Power {
	p.RLock()
	var min, max = 0, len(p.Powers)
	var randomIndex = rand.Intn(max-min) + min
	var randomPower = p.Powers[randomIndex]
	p.RUnlock()

	return randomPower
}

func (p *PowersCache) Update() {
	for {
		time.Sleep(10 * time.Minute)
		data, err := ioutil.ReadFile(consts.POWERS_PATH)
		if err != nil {
			log.Fatal(err)
		}

		p.Lock()
		defer p.Unlock()
		err = json.Unmarshal(data, &p.Powers)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Powers cache updated")
	}
}
