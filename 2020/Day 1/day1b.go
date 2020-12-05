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

    fmt.Printf("Searching for three numbers that add to %d...\n", target)

    for i, v := range nums {
        //As they are ordered, if the first number plus last two numbers is
        //greater than the target the first number cannot be in the list
        if (v + nums[len(nums) - iter] + nums[len(nums) - iter2] > target) {
            continue
        }
        for i < len(nums) - iter {
            for i < len(nums) - iter2 {
                if (v + nums[len(nums) - iter] + nums[len(nums) - iter2] == target) {
                    fmt.Printf("Success! Found three numbers that add to %d\n", target)
                    fmt.Printf("%d, %d, and %d\n", v, nums[len(nums) - iter], nums[len(nums) - iter2])
                    fmt.Printf("Puzzle solutions is these numbers multiplied\n")
                    fmt.Printf("Solution - %d\n", v * nums[len(nums) - iter] * nums[len(nums) - iter2])
                }
                iter2++
            }
            iter++
            iter2 = iter + 1
        }
        iter = 1
        iter2 = 2

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
