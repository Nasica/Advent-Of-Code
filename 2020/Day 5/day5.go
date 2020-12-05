package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "sort"
)

func main() {

    seats := readFile()
    var row int = 0
    var col int = 0
    var allSeats []int
    var highest int = 0
    for _, v := range seats {
        rowMax := 127
        rowMin := 0
        colMax := 7
        colMin := 0
        rowCounter := 64
        colCounter := 4
        rowID := v[:7]
        colID := v[7:]
        // Calculate the row
        for i := 0; i < len(rowID) - 1; i++ {
            if (string(rowID[i:i+1]) == "F") {
                rowMax = rowMax - rowCounter
            } else if (string(rowID[i:i+1]) == "B") {
                rowMin = rowMin + rowCounter
            }
            rowCounter = rowCounter / 2
        }
        if (string(rowID[len(rowID) - 1]) == "F") {
            row = rowMin * 8
        } else if (string(rowID[len(rowID) - 1]) == "B") {
            row = rowMax * 8
        }

        // Calculate the column
        for j := 0; j < len(colID) - 1; j++ {
            if (string(colID[j:j+1]) == "L") {
                colMax = colMax - colCounter
            } else if (string(colID[j:j+1]) == "R") {
                colMin = colMin + colCounter
            }
            colCounter = colCounter / 2
        }
        if (string(colID[len(colID) - 1]) == "L") {
            col = colMin 
        } else if (string(colID[len(colID) - 1]) == "R") {
            col = colMax
        }
        
        allSeats = append(allSeats, row + col)
        if (row + col > highest) {
            highest = row + col
        }
    }
    fmt.Printf("Highest seat number is - %d\n", highest)
    fmt.Printf("Which is Row:%d and Col:%d\n", highest / 8, highest % 8 )
    sort.Ints(allSeats)
    for i := 0; i < len(allSeats) - 1; i++ {
        if (allSeats[i + 1] - allSeats[i] > 1){
            fmt.Printf("Your seat ID is - %d\n", allSeats[i] + 1)
            fmt.Printf("Which is Row:%d and Col:%d\n", (allSeats[i] + 1) / 8, ((allSeats[i] + 1) % 8 ))
            break
        }
    }
} 


/*
 * Reads a file containing numbers on each line, passes those numbers
 * into an array/slice and returns the array in descending order
 */
func readFile() []string {
    file, err := os.Open("./seats.txt")
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
