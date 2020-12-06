package main
import (
  "fmt"
  "bufio"
  "os"
  //"regexp"
  //"strconv"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  inp := []string{}
  group := ""
  allGroup := []string{}
  groupCharCheck := [26]int{}
  i, j, k, n, groupCount, totalCount := 0, 0, 0, 0, 0, 0 //k is total number of groups

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
    for j = 0; j < len(allGroup[i]); j++ {
      char := (allGroup[i][j] - 97)
      if !(allGroup[i][j] == ' ') {
        if groupCharCheck[char] == 0 {
          groupCharCheck[char] = 1
        }
      }
    }
    for n = 0; n < len(groupCharCheck); n++ {
      if groupCharCheck[n] == 1 {
        groupCount += 1
      }
    }
    totalCount += groupCount
    groupCount = 0
    for j = 0; j < len(groupCharCheck); j++ {
      groupCharCheck[j] = 0
    }

  }
  fmt.Println(totalCount)
}