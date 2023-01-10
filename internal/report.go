package internal

import "math"

type Report struct {
	ID        int64   `pg:"id,pk"`
	SessionID string  `pg:"session_id"`
	Frequency float64 `pg:"frequency"`
	Timestamp string  `pg:"timestamp"`
}

type Detector interface {
	DetectAnomaly(mean float64, sd float64, k float64) bool
}

func (data *Report) DetectAnomaly(mean float64, sd float64, k float64) bool {
	return math.Abs(data.Frequency-mean) > sd*k
}
