package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

const LETTER_IN_ALPHA = 26

func main() {
    // ASCII value of 'a' to put character in range of 0-25
    const ASCII_CONST int = 97
    // Array to store counts of each occurance of a letter
    var checkSheet [LETTER_IN_ALPHA]int
    // The final result
    var total int = 0
    // Get input
    answers := readFile()
    for i, v := range answers {
        if v != "" {
            // For each letter increment its relevant spot in an array
            // of size 26
            for x := 0; x < len(v); x++ {
                checkSheet[int(v[x]) - ASCII_CONST]++
            }
        }
        // Check if end of a group of answers or last group in answer set
        if (v == "") || (i == len(answers) - 1) { 
            count(&checkSheet, &total)
        }
    }
    fmt.Printf("%d\n", total)
}

// Counts number of unique letters used 
func count(checkSheet *[LETTER_IN_ALPHA]int, total *int) {
    for i, w := range checkSheet {
        if w > 0 {
           *total++
        }
        checkSheet[i] = 0
    }
}

/*
 * Reads a file containing numbers on each line, passes those numbers
 * into an array/slice and returns the array in descending order
 */
func readFile() []string {
    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var list []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        list = append(list, scanner.Text())
    }
    return list
}
