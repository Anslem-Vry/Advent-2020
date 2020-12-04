package main
import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  a := slopes(1, 1)
  a = a * slopes(3, 1)
  a = a * slopes(5, 1)
  a = a * slopes(7, 1)
  a = a * slopes(1, 2)
  fmt.Println(a)
}

func slopes(xScale int, yScale int) int {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  xPos := 0
  yPos := 0
  i := 0
  var inp [323] string
  trees := 0

  if err != nil{
    fmt.Println("File reading error", err)
    return 0
  }
  for scanner.Scan() {
    inp[i] = scanner.Text()
    i++
  }
  xReset := (len(inp[0])- 1) //reset bounds for "pattern repeat"
  for yPos = yScale; yPos < len(inp); yPos = (yPos + yScale) {
    xPos += xScale
    if (xPos > xReset) {
      xPos = xPos - (xReset + 1)
    }
    current := string(inp[yPos][xPos])
    if (current == "#") {
      trees +=1
    }
  }
  return trees
}