package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var rows []string
	for fileScanner.Scan() {
		rows = append(rows, fileScanner.Text())
	}

	longestQueue := 0
	queue := 0
	for i, _ := range rows {
		if rows[i] == "1" {
			queue++
		}
		if i+1 == len(rows) || rows[i] == "0" {
			if longestQueue < queue {
				longestQueue = queue
			}
			queue = 0
		}
	}

	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if _, err := writer.Write([]byte(strconv.Itoa(longestQueue))); err != nil {
		log.Fatal(err)
	}
	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}
}
