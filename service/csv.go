package service

import (
	"context"
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"storage/config"
	"storage/models"
)

func readCsvFile() []*models.Product {
	csvFile := config.GetCsvFile()
	file, error := os.Open(csvFile)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Successfully opened the CSV file")
	defer file.Close()

	var products []*models.Product

	if err := gocsv.UnmarshalFile(file, &products); err != nil {
		panic(err)
	}
	// read CSV file
	//fileReader := csv.NewReader(file)
	//records, error := fileReader.ReadAll()
	//if error != nil {
	//	fmt.Println(error)
	//}
	//fmt.Println(records)
	return products
}

func (s *ProductService) ProccessCsvData() {
	records := readCsvFile()
	err := s.SaveData(records)
	if err != nil {
		log.Println("save csv data in data base got error :", err)
	}

}

func (s *ProductService) SaveData(products []*models.Product) error {
	err := s.repo.Create(context.Background(), products)
	if err != nil {
		return err
	}

	return nil
}
