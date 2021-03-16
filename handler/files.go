package handler

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/PicPay/software-engineer-challenge/models"
)

func Pfile() [][]string {
	// file, _ := ioutil.ReadDir("./tmp/new-f.csv")
	file, err := os.Open("./tmp/new-f.csv")
	if err != nil {
		log.Fatal(err)
	}

	parser := csv.NewReader(file)
	record := make([][]string, 0)
	// defer parsing(records)
	for {
		records, err := parser.Read()
		// if records, err := parser.Read(); err != nil {
		// }
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		record = append(record, records)
		// fmt.Println(record)
	}
	return record
	// Parsing(record)
}

func Print_records(records [][]string) {
	log.Println(records)
	for record, a := range records {
		log.Println(record, a)
	}
}

//TODO: refactor it to use chan and concurrency 
func Parsing(records [][]string) {
	log.Println("entrou aqui")
	for _, record := range records {
		User := &models.User{ID: record[0], Name: record[1], Username: record[2]}
		log.Println(User)
		if err := dbInstance.AddUser(User); err != nil {
			log.Fatal(err)
		}
		// log.Println("going to the next")
		// fmt.Println(record[i])
	}
}
