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
		listFileName[provinceCode] = provinceCode
		checkIsCreateAndOpenProvinceCSV(provinceCode, pathForWrite)
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

func checkIsCreateAndOpenProvinceCSV(newProvince string, pathForWrite map[string]string) {
	var pathWriterFile string = "./result/" + newProvince + ".csv"
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
