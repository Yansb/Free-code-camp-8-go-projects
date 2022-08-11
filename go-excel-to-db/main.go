package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/yansb/go-excel-to-db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fileName := "./Excel-Data.xlsx"
	excelChan := make(chan []*xlsx.Cell)
	doneChan := make(chan bool)

	db := connectToDb()

	go readExcel(fileName, excelChan)
	go saveInDB(db, excelChan, doneChan)
	<-doneChan
}

func readExcel(fileName string, channel chan []*xlsx.Cell) {
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		panic(err)
	}

	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}
			channel <- row.Cells
		}
		close(channel)
	}

}

func saveInDB(db *gorm.DB, rows chan []*xlsx.Cell, done chan bool) {
	for i := range rows {
		db.Create(&models.Data{
			WO:       i[0].Value,
			District: i[1].Value,
			LeadTech: i[2].Value,
			Service:  i[3].Value,
			Techs:    i[4].Value,
			LbrHrs:   i[5].Value,
			PartsCst: i[6].Value,
			Payment:  i[7].Value,
		})
		fmt.Printf("saved %s in db: \n", i)
	}
	done <- true
}

func connectToDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=go port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Data{})
	if err != nil {
		panic(err)
	}

	return db
}
