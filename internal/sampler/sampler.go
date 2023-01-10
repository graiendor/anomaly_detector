package sampler

import (
	"math/rand"
)

const (
	minMean = -10
	maxMean = 10
	minSD   = 0.3
	maxSD   = 1.2
)

type Sampler interface {
	GetSample() float64
	GetStats() (float64, float64)
}

type Stats struct {
	Mean float64
	SD   float64
}

func (s *Stats) GetStats() (float64, float64) {
	return s.Mean, s.SD
}

func (s *Stats) GetSample() float64 {
	return rand.NormFloat64()*s.SD + s.Mean
}

func GenerateStats() (sampler Sampler) {
	return &Stats{
		Mean: minMean + rand.Float64()*(maxMean-minMean),
		SD:   minSD + rand.Float64()*(maxSD-minSD),
	}
}
