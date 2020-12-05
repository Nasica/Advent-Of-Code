package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
    "strings"
)

func main() {
    contents := readFile()
    var validCount int = 0
    for _, s := range contents {
        password, ruleChar, first, last := splitString(s)
        if ((len(password) > first) && (len(password) > last)){
            if ((string(password[first]) == ruleChar || string(password[last]) == ruleChar) && (string(password[first]) != string(password[last]))){
                validCount++
            }
        }
    }
    fmt.Printf("Total valid passwords - %d\n", validCount)
}

func splitString(s string) (string, string, int, int) {
    s1 := strings.Split(s, ":")
    s1b := strings.Split(s1[0], " ")
    var password string = s1[1]
    var ruleChar string = s1b[1]
    s1c := strings.Split(s1b[0], "-")
    min, err := strconv.Atoi(s1c[0])
    max, err := strconv.Atoi(s1c[1])
    if (err != nil){
    }
    return password, ruleChar, min, max
}

/*
 * Reads a file containing numbers on each line, passes those numbers
 * into an array/slice and returns the array in descending order
 */
func readFile() []string {
    file, err := os.Open("./passwords.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var list []string
    var temp string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        temp = scanner.Text()
        list = append(list, temp)
    }
    return list
}

