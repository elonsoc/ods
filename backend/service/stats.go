package service

import (
	"github.com/sirupsen/logrus"
	statsd "github.com/smira/go-statsd"
)

func initStatsD(statsdURL string, log *logrus.Logger) *Stat {
	s := &Stat{}
	s.stat = statsd.NewClient(statsdURL, statsd.MetricPrefix("backend."))
	return s
}

type StatIFace interface {
	Increment(string, ...statsd.Tag)
	TimeElapsed(string, int64)
}

type Stat struct {
	stat *statsd.Client
}

func (s *Stat) Increment(stat string, tags ...statsd.Tag) {
	s.stat.Incr(stat, 1, tags...)
}

func (s *Stat) TimeElapsed(stat string, time int64) {
	s.stat.Timing(stat, time)
}
