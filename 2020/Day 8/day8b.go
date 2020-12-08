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
    beenTested bool
}


func main() {
    input := readFile()
    var instrArray []Instruction
    var stackPointer int = 0
    var total int = 0
    var inTesting = false
    for _, v := range input {
        temp := strings.Split(v, " ")
        numTemp, _ := strconv.Atoi(temp[1])
        tempStruct := Instruction{instr:temp[0], addr:numTemp, beenVisited:false, beenTested:false}
        instrArray = append(instrArray, tempStruct)
    }

    for stackPointer < len(instrArray) {
        // Reset the state, last test didn't work 
        if instrArray[stackPointer].beenVisited == true {
            stackPointer = 0
            total = 0
            inTesting = false
            for i := 0; i < len(instrArray); i++ {
                instrArray[i].beenVisited = false
            }
            continue
        }

        // This instruction has now been visitied
        instrArray[stackPointer].beenVisited = true
        // Check which instruction is to be executed
        switch instrArray[stackPointer].instr {
        case "nop":
            /*
             * If the system isn't conducting a test and this instruction hasn't been
             * tested yet, then treat this instruction as a JMP
             */
            if inTesting == false && instrArray[stackPointer].beenTested == false {
                inTesting = true
                instrArray[stackPointer].beenTested = true
                stackPointer += instrArray[stackPointer].addr
            } else {
                // Treat as a NOP, Do nothing
            }
        case "jmp":
            /*
             * Ifd the system isn't conducting a test and this instruction hasn't been
             * tested yet, then treat this instruction as a NOP
             */
            if inTesting == false && instrArray[stackPointer].beenTested == false {
                inTesting = true
                instrArray[stackPointer].beenTested = true
            } else {
                // Normal operation, treat as a JMP
                stackPointer += instrArray[stackPointer].addr
                continue
            }
        case "acc":
            total += instrArray[stackPointer].addr
        }
        stackPointer++
    }
    fmt.Printf("Accumumlator total is - %d\n", total)
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
