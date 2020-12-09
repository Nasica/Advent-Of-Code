package main

import (
    "bufio"
    "os"
    "log"
    "strconv"
    "fmt"
    "sort"
)

func main() {
    input := readFile()
    result := searchForTotal(675280050, input)
    sort.Ints(result)
    finalAnswer := result[0] + result[len(result)-1]
    fmt.Printf("Array is - %d\n With the key being - %d\n", result, finalAnswer)
}

func searchForTotal(target int, pool []int) []int {
    var index int = 1
    var innerIndex int = 1
    var total int = pool[0]
    var retVal []int
    for index < len(pool) {
        for total < target {
            retVal = append(retVal, pool[innerIndex])
            total += pool[innerIndex]
            if total == target {
                return retVal
            }
            innerIndex++
        }
        retVal = retVal[:0]
        total = 0
        index++
        innerIndex = index
    }
    return retVal 
}

/*
 * Reads a file containing numbers on each line, passes those numbers
 * into an array/slice and returns the array in descending order
 */
func readFile() []int {
    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var list []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        temp, _ := strconv.Atoi(scanner.Text())
        list = append(list, temp)
    }
    return list
}
