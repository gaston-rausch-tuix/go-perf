package goperf

import (
	"runtime"
	"time"
)

type GoPerf struct {
	StartTime      time.Time
	EndTime        time.Time
	MaxMemoryUsage uint64
}

func (g *GoPerf) Start() {
	g.StartTime = time.Now()
}

func (g *GoPerf) End() {
	g.EndTime = time.Now()
}

func (g *GoPerf) Duration() time.Duration {
	return g.EndTime.Sub(g.StartTime)
}

func (g *GoPerf) GetMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	if m.Alloc > g.MaxMemoryUsage {
		g.MaxMemoryUsage = m.Alloc
	}
	return m.Alloc
}

func (g *GoPerf) GetMaxMemoryUsage() uint64 {
	return g.MaxMemoryUsage
}
