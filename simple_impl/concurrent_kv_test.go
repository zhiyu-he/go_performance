package simple_impl

import (
	"testing"
	"sync"
	"math/rand"
)


type CacheMap struct {
	cache map[int64]interface{}
	lock *sync.RWMutex
}

type Q struct {
	id int64
}

func (p *Q) GetKey() int64 {
	return p.id
}

func (p *Q) GetValue() interface{} {
	return p
}

var (
	kv *KV
	cacheMap *CacheMap
	dat []*Q
	dat2 []IOP
	objIdList []int64
)



func SetMultiValsToCacheMap(vals []*Q) {
	cacheMap.lock.Lock()
	defer cacheMap.lock.Unlock()
	for  _, v := range vals {
		cacheMap.cache[v.id] = v
	}
}

func SetMultiValsToConcurrentKV(vals []IOP) {
	req := &Req{
		op: EVENT_SET,
		values: vals,
	}
	kv.Exec(req)
}


func init() {
	kv = NewKV(100, 2)
	cacheMap = &CacheMap{
		cache: make(map[int64]interface{}, 100),
		lock: new(sync.RWMutex),
	}
	dat = make([]*Q, 0, 100)
	dat2 = make([]IOP, 0, 100)
	objIdList = make([]int64, 0, 100)
	for i := 0; i < 100; i++ {
		id := rand.Int63()
		objIdList = append(objIdList, id)
		q := &Q{
			id: id,
		}
		dat = append(dat, q)
		dat2 = append(dat2, q)

	}
	SetMultiValsToCacheMap(dat)
	SetMultiValsToConcurrentKV(dat2)
}

func QueryMultiKeys(ids []int64) {
	cacheMap.lock.RLock()
	defer cacheMap.lock.RUnlock()
	missIds := make([]int64, 0, 16)
	items := make([]*Q, 0, 16)
	for _, id := range ids {
		if v, ok := cacheMap.cache[id]; ok {
			items = append(items, v.(*Q))
		} else {
			missIds = append(missIds, id)
		}
	}
}

func QueryMultiKeysWithKV(ids []int64) {
	missIds := make([]int64, 0, 16)
	items := make([]*Q, 0, 16)
	query := func(cache map[int64]interface{}) {
		for _, id := range ids {
			if v, ok := cache[id]; ok {
				items = append(items, v.(*Q))
			} else {
				missIds = append(missIds, id)
			}
		}
	}
	req := &Req{
		op: EVENT_GET,
		query: query,
		signal: make(chan struct{}),
	}
	kv.Exec(req)
	<-req.signal
}

func TestKV(t *testing.T) {
	ids := objIdList[:16]
	QueryMultiKeys(ids)
	QueryMultiKeysWithKV(ids)
}


func BenchmarkQueryMultiKeys(b *testing.B) {
	ids := objIdList[:16]
	for i := 0; i < b.N; i++ {
		QueryMultiKeys(ids)
	}
}

func BenchmarkQueryMultiKeysWithKV(b *testing.B) {
	ids := objIdList[:16]
	for i := 0; i < b.N; i++ {
		QueryMultiKeysWithKV(ids)
	}
}


func BenchmarkQueryMultiKeysParallel(b *testing.B) {
	ids := objIdList[:16]
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			QueryMultiKeys(ids)
		}
	})
}

func BenchmarkQueryMultiKeysWithKVParallel(b *testing.B) {
	ids := objIdList[:16]
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			QueryMultiKeysWithKV(ids)
		}
	})
}
