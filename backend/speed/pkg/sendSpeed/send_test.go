package sendSpeed

import (
	"net/http"
	"testing"
	"ueckoken/plarail2022-speed/pkg/storeSpeed"
	"ueckoken/plarail2022-speed/spec"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TrainMock struct {
	mock.Mock
}

func (t *TrainMock) GetTrain() storeSpeed.Train {
	return storeSpeed.Train{Name: spec.SendSpeed_TAKAO, Addr: "localhost:8080"}
}
func (t *TrainMock) SetSpeed(speed int32) error {
	return nil
}
func (t *TrainMock) GetSpeed() int32 {
	arg := t.Called()
	return int32(arg.Int(0))
}

func TestGetJson(t *testing.T) {
	s := NewSendSpeed(&http.Client{})
	tm := new(TrainMock)
	tm.On("GetSpeed").Return(100)
	s.Train = tm
	b, err := s.getJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `{"speed":100}`, string(b))
}
