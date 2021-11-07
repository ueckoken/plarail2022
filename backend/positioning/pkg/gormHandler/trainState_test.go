package gormHandler

import (
	"fmt"
	"os"
	"testing"
	"time"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUpdate(t *testing.T) {
	fmt.Println("DB is: ", os.Getenv("DB"))
	d, err := gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		t.Errorf("%s", err)
	}
	db := SQLHandler{DB: d}
	db.AutoMigrate(&trainState.State{})
	ts := trainState.State{
		TrainId:          123,
		HallSensorName:   "sdkaa",
		FetchedTimeStump: time.Now(),
	}
	db.Store(ts)
	tsO := db.FetchFromTrainId(123)
	fmt.Println(len(tsO.States))
	if !(ts.TrainId==tsO.States[0].TrainId) && !(ts.HallSensorName==tsO.States[0].HallSensorName){
		t.Errorf("error")
	}
}
