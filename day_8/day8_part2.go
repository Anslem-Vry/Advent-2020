package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  inp := []string{}

  if err != nil{
    fmt.Println("File reading error", err)
    return
  }
  for scanner.Scan() {
    inp = append(inp, scanner.Text())
  }

  commandRun := [637]int{}
  accu := 0
  jmpVal := 0
  var notInfiniteLoop bool

  for k := 0; k < len(inp); k++ {
    notInfiniteLoop = true
    command := strings.Split(inp[jmpVal], " ")
    process := command[0]
    
    for (notInfiniteLoop) {
      if (jmpVal == k) {
        switch process {
        case "jmp":
          process = "nop"
          break
        case "nop":
          process = "jmp"
          break
        }
      }

      switch process {
      case "acc":
        if (commandRun[jmpVal] == 0) {
          operate, _ := strconv.Atoi(command[1])
          commandRun[jmpVal] = 1
          fmt.Println("increasing accu by:", operate)
          accu += operate
          jmpVal += 1
        } else {
          break
        }
      case "jmp":
        if (commandRun[jmpVal] == 0) {
          operate, _ := strconv.Atoi(command[1])
          commandRun[jmpVal] = 1
          fmt.Println("increasing jmpVal by:", operate)
          jmpVal += operate
        } else {
          break
        }
      case "nop":
        if (commandRun[jmpVal] == 0) {
          commandRun[jmpVal] = 1
          jmpVal += 1
        } else {
          break
        }
      }
      if (jmpVal >= len(inp)) {
        fmt.Println("flip needed at:", jmpVal)
      } else if commandRun[jmpVal] == 1 {
        notInfiniteLoop = false
      }
    }
  }
  fmt.Println(accu)
}