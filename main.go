package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var logFile string = " " // change target file here
	file, err := os.Open(logFile)
	defer file.Close()
	checkError(err)

	var listFileName = map[string]string{}
	var pathForWrite = map[string]string{}

	scanner := bufio.NewScanner(file)
	var counter int = 0
	for scanner.Scan() {

		// continue to next line (pass the first line)
		if counter == 0 {
			counter += 1
			continue
		}

		

		provinceCode := strings.Split(scanner.Text(), ",")[2][:3]
		date := strings.Split(scanner.Text(), ",")[1][:8]
		listFileName[provinceCode] = getFileName(provinceCode, date)
		checkIsCreateAndOpenProvinceCSV(provinceCode, date, pathForWrite)
		writeLine(pathForWrite[provinceCode], scanner.Text())
		counter += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	fmt.Printf("main, execution time %s\n", time.Since(start))
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println("cannot read the file", err)
		panic(err)
	}
	return
}

func checkIsCreateAndOpenProvinceCSV(newProvince string, date string, pathForWrite map[string]string) {
	var path string = "./result/"
	var fileName string = getFileName(newProvince, date)
	var pathWriterFile string = path + fileName

	
	if _, err := os.Stat(pathWriterFile); err != nil {
		os.Create(pathWriterFile)
		pathForWrite[newProvince] = pathWriterFile
	}

	return
}

func writeLine(writeTo, line string) {
	w, err := os.OpenFile(writeTo, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	checkError(err)
	w.WriteString(line)
	w.WriteString("\n")
	return
}

func getFileName(provine string, date string) string {
	return "LTE KPI Backup EAS " + provine + "&" + date +".csv"
}