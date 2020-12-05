package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
    "regexp"
)

func main (){
    passports := readFile()
    names := [7]string {
        "byr",
        "iyr",
        "eyr",
        "hgt",
        "hcl",
        "ecl",
        "pid",
    }
    var isValid bool = false
    var count int = 0
    var validCount int
    // Regex for Hair colour
    re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
    // Regex for Passport ID
    re_pid := regexp.MustCompile(`^[0-9]{9}$`)
    for _, v := range passports {
        
        //Check if the string contains all the correct fields
        validCount = 0
        for _, w := range names {
            if (strings.Contains(v, w)) {
                validCount++
            }
        }
        if (validCount != 7) {
            continue
        }

        //Check if the string is valid
        values := strings.Split(v, " ")
    Loop:
        for _, w := range values {
            switch w[:4] {
            case "byr:":
                birth, _ := strconv.Atoi(w[4:])
                if ((birth >= 1920) && (birth <= 2002)) {
                    isValid = true
                } else {
                    isValid = false
                }
            case "iyr:":
                issue, _ := strconv.Atoi(w[4:])
                if ((issue >= 2010) && (issue <= 2020)) {
                    isValid = true
                } else {
                    isValid = false
                }
            case "eyr:":
                expiry, _ := strconv.Atoi(w[4:])
                if ((expiry >= 2020) && (expiry <= 2030)) {
                    isValid = true
                } else {
                    isValid = false
                }
            case "hgt:":
                height, _ := strconv.Atoi(w[4:len(w) -2])
                if (w[len(w) - 2:] == "in") {
                    if ((height >= 59) && (height <= 76)) {
                        isValid = true
                    } else {
                        isValid = false
                    }
                } else if (w[len(w) - 2:] == "cm") {
                    if ((height >= 150) && (height <= 193)) {
                       isValid = true 
                    } else {
                       isValid = false
                    }
                } else {
                    isValid = false
                }
            case "hcl:":
                if (re.MatchString(w[4:])) {
                    isValid = true
                } else {
                    isValid = false
                }
            case "ecl:":
                if (w[4:] == "amb" || w[4:] == "blu" || w[4:] == "brn" || w[4:] == "gry" || w[4:] == "grn" || w[4:] == "hzl" || w[4:] == "oth"){
                    isValid = true
                } else {
                    isValid = false
                }
            case "pid:":
                if (re_pid.MatchString(w[4:])) {
                    isValid = true
                } else {
                    isValid = false
                }
            case "cid:":
                isValid = true
            }
            if !(isValid) {
                break Loop
            } else {
            }
        }
        if (isValid) {
            count++
        }
        isValid = false
    }
    fmt.Printf("Total number of valid passwords is - %d\n", count)
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
            list = append(list, line[:len(line) - 1])
            line = ""
            continue
        }
        line += temp + " "
    }
    return list
}
