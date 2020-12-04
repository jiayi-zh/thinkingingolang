package mysql

import (
	"testing"
)

func Test(t *testing.T) {
	//// connect mysql
	//db, err := BuildMysqlConnect("192.168.9.102", 31004, "golang", "gPMg#W9fdBA%tsd9", "project_data_server")
	//if err != nil {
	//	return
	//}
	//tx := db.Begin()
	//tx1 := tx.Table("t_device_info")
	//tx2 := tx.Table("t_channel_info")
	//tx1.Select("id").Create(&Id{Id: 1})
	//tx2.Select("id").Create(&Id{Id: 2})
	//tx.Rollback()
	//fmt.Printf("tx: %p, tx1: %p, tx2: %p \n", tx, tx1, tx2)
}

type Id struct {
	Id int64 `json:"id" gorm:"id"`
}
