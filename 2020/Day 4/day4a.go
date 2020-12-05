package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
)

func main (){
    names :=[7]string{
        "byr",
        "iyr",
        "eyr",
        "hgt",
        "hcl",
        "ecl",
        "pid",
    }
    passports := readFile()
    var count int = 0
    var valid int = 0
    for _, v := range passports {
        for _, w := range names {
            if (strings.Contains(v, w)) {
                count++
            }
        }
        if (count == 7) {
            valid++
        }
        count = 0
    }
    fmt.Printf("%d\n", valid)
}

/*
 * Reads a file containing numbers on each line, passes those numbers
 * into an array/slice and returns the array in descending order
 */
func readFile() []string {
    file, err := os.Open("./passports.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var list []string
    var line string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        temp := scanner.Text()
        if (len(temp) == 0){
            line += "\n"
            list = append(list, line)
            line = ""
            continue
        }
        line += temp
    }
    return list
}
