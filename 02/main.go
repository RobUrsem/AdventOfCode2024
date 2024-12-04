package main

import (
	"02/reports"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	reportsData, err := reports.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	numSafeReports, err := reports.FindSafeReports(reportsData)
	fmt.Printf("Number of safe reports  : %v\n", numSafeReports)
}
