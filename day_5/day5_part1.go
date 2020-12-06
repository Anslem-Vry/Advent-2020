package main
import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  inp := []string{}
  i, j, bitVal, finalRow, finalColumn, seatId, highestId := 0, 0, 0, 0, 0, 0, 0
  rowBits := [7]int{}
  columnBits := [3]int{}
  divider := 1

  if err != nil{
    fmt.Println("File reading error", err)
    return
  }
  for scanner.Scan() {
    inp = append(inp, scanner.Text())
  }

  for i = 0; i < len(inp); i++ {
    for j = 0; j < len(inp[i]); j++ {
      if ((inp[i][j] == 'F') || (inp[i][j] == 'L')) {
        bitVal = 0
      } else if ((inp[i][j] == 'B') || (inp[i][j] == 'R')) {
        bitVal = 1
      } else {
        fmt.Println("Check why you're here, you shouldn't be")
      }
      if (j > 6) {
        columnBits[j-7] = bitVal
      } else {
        rowBits[j] = bitVal
      }
    }

    finalRow = 0
    finalColumn = 0
    for j = 0; j < len(rowBits); j++ {
      if (rowBits[j] == 1){
        finalRow = finalRow + (64/divider)
      }
      divider = divider * 2
    }
    divider = 1
    for j = 0; j < len(columnBits); j++ {
      if (columnBits[j] == 1){
        finalColumn = finalColumn + (4/divider)
      }
      divider = divider * 2
    }
    divider = 1
    seatId = ((finalRow * 8) + finalColumn)
    if (seatId > highestId){
      highestId = seatId
    }
    fmt.Println("sequence, row:", rowBits, "=", finalRow, "column:", columnBits, "=", finalColumn, "seat ID:", seatId, "Highest ID:", highestId)
  }
}