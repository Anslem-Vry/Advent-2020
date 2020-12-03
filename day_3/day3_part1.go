package main
import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  xPos := 0
  yPos := 0
  i := 0
  var inp [323] string
  trees := 0

  if err != nil{
    fmt.Println("File reading error", err)
    return
  }
  for scanner.Scan() {
    inp[i] = scanner.Text()
    i++
  }
  xReset := (len(inp[0])- 1) //reset boundry for "pattern repeat"
  for yPos = 1; yPos < len(inp); yPos ++{
    xPos += 3
    if (xPos > xReset){
      xPos = xPos - (xReset + 1)
    }
    current := string(inp[yPos][xPos])
    if (current == "#"){
      trees +=1
    }
  }
  fmt.Println(trees)
}