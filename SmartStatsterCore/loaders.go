// loaders
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

func get(uri string) ([]byte, error) {
	statusCode, body, err := fasthttp.Get(nil, uri)
	if err != nil {
		return nil, err
	}
	fmt.Println(statusCode)
	return body, nil
}

func parseCsvParams() Params {
	var params Params = Params{"", ""}
	reader, file := loadAndReadCsv("params.csv", 2)
	defer file.Close()

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		switch {
		case record[0] == "baseurl":
			params.BaseUrl = record[1]
			break
		case record[0] == "token":
			params.Token = record[1]
		}
	}
	return params
}

func parseCsvPatients() [][]string {
	var patients [][]string
	reader, file := loadAndReadCsv("patients.csv", 5)
	defer file.Close()

	for {
		patient, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		patients = append(patients, patient)
	}
	return patients
}

func loadAndReadCsv(path string, fieldsNum int) (*csv.Reader, *os.File) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = fieldsNum
	reader.Comment = '#'
	return reader, file
}

func convint(strtoint string) int {
	if strtoint == "" {
		return 0
	}
	i, err := strconv.Atoi(strtoint)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
