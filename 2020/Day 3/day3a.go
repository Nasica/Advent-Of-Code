package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

func main() {
    const TREE string = "#"
    // The moves to take in the format [horizontal][vertical]
    var moves = [2][5]int {{1,3,5,7,1}, {1,1,1,1,2}}
    // Tracks horizontal position
    var xPos int = 0
    var treeCount int = 0
    var total int = 1
    // Get the mountain data
    var mountain []string = readFile()

    // For each entry in the 2d array moves
    for j := 0; j < 5; j++ {
        // For each line in the mountain definition
        for i, v := range mountain{
            // Skip over every moves[1][j]'th line
            if (i % moves[1][j] > 0){
                continue
            }
            if (string(v[xPos]) == TREE) {
                treeCount++
            }
            // Move horizontally
            xPos = xPos + moves[0][j]
            // Check for wrap around
            if (xPos >= len(v)){
                xPos = xPos - len(v)
            }
        }
        // Output running total and multiply the grand total
        fmt.Printf("Total Trees hit - %d\n", treeCount)
        total = total * treeCount
        // Reset for next loop iteration
        treeCount = 0
        xPos = 0
    }
    // Answer for part B
    fmt.Printf("Answer is tree counts multiplied - %d\n", total)
}


/*
 * Reads a file containing numbers on each line, passes those numbers
 * into an array/slice and returns the array in descending order
 */
func readFile() []string {
    file, err := os.Open("./mountain.txt")
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
