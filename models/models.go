package models

import "sync"

type DiscoveredBulb struct {
	IP  string
	Mac string
}

func NewBulb(ip string, mac string) DiscoveredBulb {
	return DiscoveredBulb{IP: ip, Mac: mac}
}

type BulbRegistry struct {
	BulbsByMac map[string]DiscoveredBulb
	*sync.RWMutex
}

func NewRegistry() BulbRegistry {
	return BulbRegistry{BulbsByMac: make(map[string]DiscoveredBulb), RWMutex: &sync.RWMutex{}}
}

func (r *BulbRegistry) Register(bulb DiscoveredBulb) {
	r.Lock()
	defer r.Unlock()
	r.BulbsByMac[bulb.Mac] = bulb
}

func (r *BulbRegistry) Bulbs() []*DiscoveredBulb {
	r.RLock()
	defer r.RUnlock()
	res := make([]*DiscoveredBulb, len(r.BulbsByMac))
	cnt := 0
	for _, v := range r.BulbsByMac {
		if cnt > (len(r.BulbsByMac) - 1) {
			break
		}
		res[cnt] = &v
		cnt++
	}
	return res
}
