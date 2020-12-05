package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
    "sort"
)

func main() {
    var target   int = 2020
    var nums []  int = readFile()
    var iter     int = 1
    var iter2    int = 2
    var success bool = false

    fmt.Printf("Searching for two numbers that add to %d...\n", target)

    for i, v := range nums {
        for i < len(nums) - iter {
            if (v + nums[len(nums) - iter] == target){
                success = true
                fmt.Printf("Numbers found! \n%d - %d\n", v, nums[len(nums) - iter])
                fmt.Printf("Puzzle solution is these two numbers multiplied\n")
                fmt.Printf("Answer - %d\n1", v * nums[len(nums) - iter])
                break
            } else if (v + nums[len(nums) - iter] < target) {
                iter++
            } else {
                break
            }
        }
        if success {
            break
        }
        iter = 1
    }
}

/*
 * Reads a file containing numbers on each line, passes those numbers
 * into an array/slice and returns the array in descending order
 */
func readFile() []int {
    file, err := os.Open("./numbers.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var numList []int
    var temp int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        temp, err = strconv.Atoi(scanner.Text())
        if (temp != 0) {
            numList = append(numList, temp)
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(numList)))
    return numList
}
