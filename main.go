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
	var logFile string = "./source/LTE KPI Backup EAS 20220617_20220619.csv" // change target file here
	file, err := os.Open(logFile)
	defer file.Close()
	checkError(err)

	var listFileName = map[string]string{}
	var pathForWrite = map[string]string{}
	var headerColumn string
	// var folderProvince = map[string]string{}

	scanner := bufio.NewScanner(file)
	var counter int = 0
	for scanner.Scan() {

		// continue to next line (pass the first line)
		if counter == 0 {
			headerColumn = scanner.Text()
			counter += 1
			continue
		}


		provinceCode := strings.Split(scanner.Text(), ",")[2][:3]
		date := strings.Split(scanner.Text(), ",")[1][:8]
		listFileName[provinceCode] = getFileName(provinceCode, date)
		isCreate(provinceCode, date, pathForWrite, headerColumn)
		
		writeLine(pathForWrite[provinceCode], scanner.Text())
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

func isCreate(newProvince string, date string, pathForWrite map[string]string, headerColumn string) {

	var fileName string = getFileName(newProvince, date)
	var provinceFolder string = "./result/" + newProvince +"/"
	var pathWriterFile string = provinceFolder + fileName

	if _, err := os.Stat(provinceFolder); err != nil {
		os.Mkdir(provinceFolder, os.ModePerm)
	}

	if _, err := os.Stat(pathWriterFile); err != nil {
		os.Create(pathWriterFile)
		writeLine(pathWriterFile, headerColumn)
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
	return "LTE KPI Backup EAS " + provine + "&" + date + ".csv"
}
