package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/yansb/go-study-project/go-file-to-db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := connectToDb()

	start := time.Now()
	fileName := "./csv-data.csv"
	excelChan := make(chan []string)
	doneChan := make(chan bool)

	go readExcel(fileName, excelChan)
	go saveInDB(db, excelChan, doneChan)
	<-doneChan

	fmt.Println("time elapsed: ", time.Since(start))
}

func readExcel(fileName string, channel chan []string) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		panic("error opening file")
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic("error reading csv")
	}

	for idx, lines := range csvLines {
		fmt.Print(idx)
		channel <- lines
	}
	fmt.Printf("fechei o channel")
	close(channel)

}

func saveInDB(db *gorm.DB, rows chan []string, done chan bool) {
	data := make([]*models.Data, 0)

	for i := range rows {
		data = append(data, &models.Data{
			Idade:       i[0],
			Trabalho:    i[1],
			Numero:      i[2],
			Graduacao:   i[3],
			OutroNumero: i[4],
			EstadoCivil: i[5],
		})
		fmt.Println(i[0])
	}

	db.Create(&data)
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
