package anomalies

import (
	"github.com/graiendor/anomaly_detector/internal"
	"math"
)

func CalculateMean(d []internal.Report) float64 {
	length := float64(len(d))
	mean := 0.0
	for _, entry := range d {
		mean += entry.Frequency / length
	}
	return mean
}

func CalculateStandardDeviation(d []internal.Report, mean float64) float64 {
	devSquareSum := 0.0
	for _, entry := range d {
		devSquareSum += math.Pow(entry.Frequency-mean, 2.0)
	}
	return math.Sqrt(devSquareSum / float64(len(d)))
}
