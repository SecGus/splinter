package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"flag"
)

func executeCommand(command string, input string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute command: %s\n", err)
	}
	// Execute command without adding an extra newline here, to manage it more precisely later
}

func printSection(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var printedLines int
	for scanner.Scan() {
		line := scanner.Text()
		// Check if the line is not just whitespace
		if strings.TrimSpace(line) != "" {
			fmt.Println(line)
			printedLines++
		}
	}
	// Add a newline only if any lines were printed to separate sections clearly without extra empty lines
	if printedLines > 0 {
		fmt.Println()
	}
}

func main() {
	var n int
	var command string
	var section int
	flag.IntVar(&n, "n", 2, "Number of sections to split the input into")
	flag.StringVar(&command, "c", "", "Command to execute on each section (optional)")
	flag.IntVar(&section, "s", -1, "Section to execute the command on (optional)")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" { // Skip entirely whitespace or empty lines when reading
			lines = append(lines, line)
		}
	}

	chunkSize := (len(lines) + n - 1) / n // Calculate size of each chunk
	for i := 0; i < n; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(lines) {
			end = len(lines)
		}
		chunk := lines[start:end]
		chunkText := strings.Join(chunk, "\n")
		
		// Check if a command is specified. If so, replace {{number}} and execute.
		if command != "" {
			modifiedCommand := strings.Replace(command, "{{number}}", fmt.Sprintf("%d", i+1), -1)
			if section == -1 || section == i+1 {
				executeCommand(modifiedCommand, chunkText)
			}
		} else {
			// If no command is specified, simply print the sections without empty lines.
			if section == -1 || section == i+1 {
				printSection(chunkText)
			}
		}
	}
}
