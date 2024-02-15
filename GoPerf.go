package goperf

import (
	"time"
)

type GoPerf struct {
	StartTime time.Time
	EndTime   time.Time
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
