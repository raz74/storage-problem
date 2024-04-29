package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"storage/config"
	"storage/models"
	"strconv"
	"time"
)

const layout = "2006-01-02 03:04:05 -0700 MST"

func readCsvFile() [][]string {
	csvFile := config.GetCsvFile()
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("the csv file not found !", err)
	}
	fmt.Println("Successfully opened the CSV file")
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	//err = os.Remove(csvFile)
	//if err != nil {
	//	log.Println("remove csv file got err:", err)
	//}

	return records
}
func mapToProduct(records [][]string) []*models.Product {
	var products []*models.Product

	for _, record := range records {
		price, _ := strconv.ParseFloat(record[1], 64)
		exp, _ := time.Parse(layout, record[2])
		data := models.Product{
			Id:             record[0],
			Price:          price,
			ExpirationData: exp,
		}
		products = append(products, &data)
	}
	return products
}

func (s *ProductService) ProcessCsvData() {
	records := readCsvFile()
	products := mapToProduct(records)
	err := s.SaveData(products)
	if err != nil {
		log.Println("save csv data in data base got error :", err)
	}
	time.Sleep(30 * time.Minute)
}

func (s *ProductService) SaveData(products []*models.Product) error {
	err := s.repo.Create(context.Background(), products)
	if err != nil {
		return err
	}

	return nil
}
