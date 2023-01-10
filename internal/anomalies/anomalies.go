package anomalies

import (
	"fmt"
	"github.com/graiendor/anomaly_detector/cmd/postgresql"
	"github.com/graiendor/anomaly_detector/internal"
	"strconv"
)

func LogPredictions(data []internal.Report) {
	mean := CalculateMean(data)
	sd := CalculateStandardDeviation(data, mean)
	fmt.Println("Processed: " + strconv.Itoa(len(data)) +
		"\tMean: " + strconv.FormatFloat(mean, 'f', -1, 64) +
		"\tSTD: " + strconv.FormatFloat(sd, 'f', -1, 64))
}

func LogAnomalies(coefficient float64, data []internal.Report) {
	mean := CalculateMean(data)
	sd := CalculateStandardDeviation(data, mean)
	fmt.Println("ANOMALY ENTRIES:")
	reportServer := postgresql.ServerInit()
	for _, entry := range data {
		if entry.DetectAnomaly(mean, sd, coefficient) {
			fmt.Println(entry.SessionID + " " +
				strconv.FormatFloat(entry.Frequency, 'f', -1, 64) +
				" " + entry.Timestamp)
			postgresql.InsertEntry(reportServer, entry)
		}
	}
}
