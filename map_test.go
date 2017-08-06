package main

import (
	"image/color"
	"sync"
	"testing"

	"github.com/200sc/go-dist/colorrange"
	"github.com/200sc/go-dist/intrange"

	"golang.org/x/sync/syncmap"
)

var (
	keys   = intrange.NewLinear(0, 100000)
	values = colorrange.NewLinear(color.RGBA{0, 0, 0, 0}, color.RGBA{254, 254, 254, 254})
	writes = 10000
	reads  = 10000
	v, ok  interface{}
)

func BenchmarkSyncMap(b *testing.B) {
	m := syncmap.Map{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < writes; j++ {
			m.Store(keys.Poll(), values.Poll())
		}
		for j := 0; j < reads; j++ {
			v, ok = m.Load(keys.Poll())
		}
	}
}

func BenchmarkRWLock(b *testing.B) {
	m := make(map[int]color.Color)
	l := sync.RWMutex{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < writes; j++ {
			l.Lock()
			m[keys.Poll()] = values.Poll()
			l.Unlock()
		}
		for j := 0; j < reads; j++ {
			l.RLock()
			v, ok = m[keys.Poll()]
			l.RUnlock()
		}
	}
}
