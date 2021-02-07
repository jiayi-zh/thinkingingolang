package mysql

import (
	"testing"
)

func Test(t *testing.T) {
	// connect mysql
	db, err := BuildMysqlConnect("192.168.9.102", 12004, "golang", "gPMg#W9fdBA%tsd9", "project_data_server")
	if err != nil {
		return
	}

	tx := db.Begin()

	tx.Commit()

}

type Temp struct {
	Id    int64 `json:"id" gorm:"id"`
	Depth int   `json:"depth"`
}
