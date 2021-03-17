package handler

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func Pfile() {
	file, err := os.Open("./tmp/users.csv")
	if err != nil {
		log.Fatal(err)
	}

	records := make(chan []string)
	go func() {
		parser := csv.NewReader(file)

		defer close(records)
		for {
			record, err := parser.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			records <- record
		}
	}()
	err = dbInstance.CopyTable()
	if err != nil {
		log.Fatal(err)
	}
}
func Tfile() chan []string {
	file, err := os.Open("./tmp/lista_relevancia_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	file2, err := os.Open("./tmp/lista_relevancia_2.txt")
	if err != nil {
		log.Fatal(err)
	}

	files := []io.Reader{file, file2}
	records := make(chan []string)
	go func() {
		defer close(records)
		for _, file := range files {
			parser := csv.NewReader(file)

			for {
				record, err := parser.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				records <- record
			}
		}

	}()

	return records
}
