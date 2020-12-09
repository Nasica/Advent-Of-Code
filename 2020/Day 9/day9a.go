package main

import (
    "bufio"
    "os"
    "log"
    "strconv"
    "fmt"
)

func main() {
    input := readFile()
    var preamble int = 25
    var start int = 0

    for preamble < len(input) {
        if !(searchForAddition(input[preamble:preamble+1], input[start:preamble])) {
            break
        }
        start++
        preamble++
    }
    fmt.Printf("Answer is - %d\n", input[preamble])
}

func searchForAddition(target []int, pool []int) bool {
    var back int = 0
    var front int = 0
    for front < len(pool) {
        front = back + 1
        for front < len(pool) {
            if pool[back] + pool[front] == target[0] {
                return true
            }
            front++
        }
        back++
        front = back + 1
    }
    return false
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
