package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func findLineEndingWith(lines []string, suffix string) string {
	for _, line := range lines {
		if strings.HasSuffix(line, suffix) {
			return line
		}
	}
	return "" // Return an empty string if no matching line is found.
}

func main() {
	cmd := exec.Command("sh", "get_active_threads.sh")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	fmt.Println(string(output))

	filePath := "input.txt" // Replace with your file path.
	lines, err := readFileLines(filePath)

	if os.IsNotExist(err) {
		log.Fatalf("Error reading file: %v", filePath)
	}

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// fmt.Println("File contents:")
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
	lookFor := "-intros"
	amatch := findLineEndingWith(lines, lookFor)
	if amatch != "" {
		fmt.Println(amatch)
	} else {
		fmt.Printf("Not found: %v\n", lookFor)
	}
}
