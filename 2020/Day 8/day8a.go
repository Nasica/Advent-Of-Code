package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"
)


type Instruction struct {
    instr string
    addr int
    beenVisited bool
}


func main() {
    input := readFile()

    var instrArray []Instruction
    var stackPointer int = 0
    var total int = 0
    
    // Setup array of Instruction struts
    for _, v := range input {
        temp := strings.Split(v, " ")
        numTemp, _ := strconv.Atoi(temp[1])
        tempStruct := Instruction{instr:temp[0],addr:numTemp,beenVisited:false}
        instrArray = append(instrArray, tempStruct)
    }

    // While the particular instruction hasn't been visited
    for instrArray[stackPointer].beenVisited == false {
        // This instruction has now been visited
        instrArray[stackPointer].beenVisited = true
        // Check which instruction is to be executed and execute it
        switch instrArray[stackPointer].instr {
        case "nop":
            // Don't do anything
        case "jmp":
            // Jump x instructions
            stackPointer += instrArray[stackPointer].addr
            continue
        case "acc":
            // Add to the accumulator
            total += instrArray[stackPointer].addr
        }
        // Go to next instruction
        stackPointer++
    }
    fmt.Printf("Accumulator total is - %d\n", total)
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
