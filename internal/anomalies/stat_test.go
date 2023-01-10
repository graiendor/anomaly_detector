package anomalies

import (
	r "github.com/graiendor/anomaly_detector/internal"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MyMockedObject struct {
	report []r.Report
	mock.Mock
}

func (m *MyMockedObject) CalculateMean() (bool, error) {
	args := m.Called(CalculateMean(m.report))
	return args.Bool(0), args.Error(1)
}

func (m *MyMockedObject) CalculateSTD() (bool, error) {
	args := m.Called(CalculateStandardDeviation(m.report, CalculateMean(m.report)))
	return args.Bool(0), args.Error(1)
}

func TestCalculateMean(t *testing.T) {
	m := new(MyMockedObject)
	m.report = []r.Report{
		{
			Frequency: 1.0,
		},
		{
			Frequency: 2.0,
		},
		{
			Frequency: 3.0,
		},
		{
			Frequency: 4.0,
		},
		{
			Frequency: 5.0,
		},
	}
	m.On("CalculateMean", 3.0).Return(true, nil)
	m.On("CalculateSTD", 1.4142135623730951).Return(true, nil)

	_, err := m.CalculateMean()
	if err != nil {
		return
	}
	_, err = m.CalculateSTD()
	if err != nil {
		return
	}

	m.AssertExpectations(t)
}
