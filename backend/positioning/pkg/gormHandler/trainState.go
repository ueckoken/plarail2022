package gormHandler

import (
	"gorm.io/gorm"
	"ueckoken/plarail2022-positioning/pkg/trainState"
)

type SQLHandler struct {
	Db *gorm.DB
}

func (s SQLHandler) Store(state trainState.State) {
	s.Db.Create(state)
}

func (s SQLHandler) FetchFromTrainID(trainID int) (state trainState.States) {
	s.Db.Where("train_id = ?", trainID).Order("fetched_time_stump asc").Find(&state.States)
	return state
}

func (s SQLHandler) FetchFromHallSensorName(hallID string) (state trainState.States) {
	s.Db.Where("hall_sensor_name = ?", hallID).Order("fetched_time_stump asc").Find(&state.States)
	return state
}

func (s SQLHandler) Fetch(trainID int, hallID string) (state trainState.States) {
	s.Db.Where("hall_sensor_name = ?", hallID).Where("train_id = ?", trainID).Order("fetched_time_stump asc").Find(&state.States)
	return state
}
