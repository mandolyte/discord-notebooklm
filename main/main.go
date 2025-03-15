package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
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

func findLineEndingWith(lines []string, suffix string) []string {
	var asection []string
	for i := 0; i < len(lines); i++ {
		// fmt.Println(lines[i])
		if strings.HasSuffix(lines[i], suffix) {
			for j := i + 1; j < len(lines); j++ {
				if strings.HasPrefix(lines[j], " *") {
					asection = append(asection, lines[j])
				} else {
					break
				}
			}
		}
	}

	return asection
}

func writeLinesToFile(lines []string, filePath string) error {
	file, err := os.Create(filePath) // Create or overwrite the file
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n") // Append linefeed and write
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	start := time.Now() // Record the start time
	fmt.Println("Get active threads on Discord server")
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
	lookFor := "-bugs"
	fmt.Println("Look for bugs section")
	amatch := findLineEndingWith(lines, lookFor)
	if len(amatch) != 0 {
		writeLinesToFile(amatch, "bugs/input.txt")
	} else {
		fmt.Printf("Not found: %v\n", lookFor)
	}

	lookFor = "-feature-requests"
	fmt.Println("Look for feature-requests section")
	amatch = findLineEndingWith(lines, lookFor)
	if len(amatch) != 0 {
		writeLinesToFile(amatch, "feature-requests/input.txt")
	} else {
		fmt.Printf("Not found: %v\n", lookFor)
	}

	//
	// With the data inputs now ready we can run the extract scripts
	//
	fmt.Println("Start Discord server channel extracts")
	var wg sync.WaitGroup // Create a WaitGroup

	wg.Add(5) // Add 5 workers to the WaitGroup

	go worker(1, &wg, "runBugs.sh")            // Start worker 1 in a goroutine
	go worker(2, &wg, "runFeatureRequests.sh") // Start worker 2 in a goroutine
	go worker(3, &wg, "runGeneral.sh")         // Start worker 3 in a goroutine
	go worker(4, &wg, "runIntros.sh")          // Start worker 4 in a goroutine
	go worker(5, &wg, "runUseCases.sh")        // Start worker 5 in a goroutine

	wg.Wait() // Wait for all workers to finish

	fmt.Println("All workers completed.")

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Extraction took %s to run\n", elapsed)

}

func worker(id int, wg *sync.WaitGroup, script string) {
	defer wg.Done() // Signal completion when the worker finishes
	fmt.Printf("Worker %d started: %v\n", id, script)

	cmdBugs := exec.Command("sh", script)
	output, err := cmdBugs.Output()
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	fmt.Println(string(output))
	fmt.Printf("Worker %d finished: %v\n", id, script)

}
