package postgresql

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	r "github.com/graiendor/anomaly_detector/internal"
	"log"
)

func ServerInit() *pg.DB {
	reportServer := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "postgres",
	})
	err := createSchema(reportServer)
	if err != nil {
	}
	return reportServer
}

func InsertEntry(reportServer *pg.DB, report r.Report) {
	_, err := reportServer.Model(&report).Insert()
	if err != nil {
		log.Fatalf("Error inserting report: %v", err)
	}
}

func createSchema(reportServer *pg.DB) error {
	models := []interface{}{
		(*r.Report)(nil),
	}
	for _, model := range models {
		err := reportServer.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func FindEntry(reportServer *pg.DB, userToFind *r.Report) {
	err := reportServer.Model(userToFind).WherePK().Select()
	if err != nil {
		log.Fatalf("Error with insert")
	}
}
