package internal

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MyMockedObject struct {
	mean float64
	sd   float64
	k    float64
	mock.Mock
}

func (m *MyMockedObject) DetectAnomaly() (bool, error) {
	rep := Report{Frequency: 1.238}
	args := m.Called(rep.DetectAnomaly(m.mean, m.sd, m.k))
	return args.Bool(0), args.Error(1)
}

func TestReport_DetectAnomaly(t *testing.T) {
	m := new(MyMockedObject)
	m.mean = 3
	m.sd = 0.9
	m.k = 3.0
	m.On("DetectAnomaly", false).Return(true, nil)

	_, err := m.DetectAnomaly()
	if err != nil {
		return
	}

	m.mean = 3
	m.sd = 0.3
	m.k = 3.0
	m.On("DetectAnomaly", true).Return(true, nil)
	_, err = m.DetectAnomaly()
	if err != nil {
		return
	}

	m.AssertExpectations(t)
}
