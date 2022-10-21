package main

import (
	"log"
	"os"
	"time"
	"ueckoken/plarail2022-positioning/internal"
	"ueckoken/plarail2022-positioning/pkg/gormHandler"
	"ueckoken/plarail2022-positioning/pkg/trainState"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbenv := os.Getenv("DB")
	log.Println("DB:", dbenv)
	var d *gorm.DB
	var err error
	for {
		d, err = gorm.Open(postgres.Open(dbenv), &gorm.Config{})
		if err != nil {
			log.Printf("error when connecting to sql: %s", err)
			time.Sleep(1 * time.Second)
			log.Panicf("re-connecting to sql")
			continue
		}
		break
	}
	db := gormHandler.SQLHandler{Db: d}
	if err := db.Db.AutoMigrate(&trainState.State{}); err != nil {
		log.Panicf("DB auto migrate failed, err=%s\n", err)
	}
	as, err := internal.NewApplicationStatus()
	if err != nil {
		log.Panicf("err=%s\n", err)
	}
	r := internal.NewPositionReceiver(db, as)
	r.StartPositionReceiver()
}
