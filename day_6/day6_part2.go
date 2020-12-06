package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  inp := []string{}
  group := ""
  allGroup := []string{}
  groupCharCheck := [26]int{}
  i, j, k, groupHigh, groupCount, totalCount := 0, 0, 0, 0, 0, 0

  if err != nil{
    fmt.Println("File reading error", err)
    return
  }
  for scanner.Scan() {
    inp = append(inp, scanner.Text())
  }
  for i = 0; i < len(inp); i++ {
    group = group + inp[i] + " "
    if (inp[i] =="") { //reached the end of a group
      k++
      allGroup = append(allGroup, group)
      group = ""
    }
  }
  for i = 0; i < k; i++ {
    groupLen := len(strings.Split(allGroup[i], " "))-2
    for j = 0; j < len(allGroup[i]); j++ {
      char := (allGroup[i][j] - 97)
      if !(allGroup[i][j] == ' ') {
        groupCharCheck[char] = groupCharCheck[char] + 1
      }
    }
    for j = 0; j < len(groupCharCheck); j++ {
      if groupCharCheck[j] > groupHigh {
        groupHigh = groupCharCheck[j]
      }
    }
    for j = 0; j < len(groupCharCheck); j++ {
      if groupCharCheck[j] == groupLen {
        groupCount += 1
      }
    }
    totalCount =  totalCount + groupCount
    groupCount = 0
    groupHigh = 0
    for j = 0; j < len(groupCharCheck); j++ { // Reset the group
      groupCharCheck[j] = 0
    }
  }
  fmt.Println(totalCount)
}