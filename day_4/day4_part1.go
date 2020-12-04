package main
import (
  "fmt"
  "bufio"
  "os"
  //"strconv"
  "strings"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)

  i := 0
  l := 0
  j := 0 // length of input
  k := 0 // total passports
  inp := []string{}
  process := ""
  total := 0
  row2 := []string{}
  if err != nil{
    fmt.Println("File reading error", err)
    return
  }
  for scanner.Scan() {
    inp = append(inp, scanner.Text())
  }
  j = len(inp)
  for i = 0; i < j; i++ {
    process = process + inp[i] + " "
    
    if (inp[i] =="") { //reached the end of a passport
      k++
      row2 = append(row2, process)
      process = ""
    }
  }
  for i = 0; i < k ; i++ {
    byr, iyr, eyr, hgt, hcl, ecl, pid := 0, 0, 0, 0, 0, 0, 0
    stringSplits1 := strings.Split(row2[i], " ")
    for l = 0; l < len(stringSplits1); l++ {
      if (strings.Count(stringSplits1[l], "byr")) == 1 {
        byr = 1
      } else if (strings.Count(stringSplits1[l], "iyr")) == 1 {
        iyr = 1
      } else if (strings.Count(stringSplits1[l], "eyr")) == 1 {
        eyr = 1
      } else if (strings.Count(stringSplits1[l], "hgt")) == 1 {
        hgt = 1
      } else if (strings.Count(stringSplits1[l], "hcl")) == 1 {
        hcl = 1
      } else if (strings.Count(stringSplits1[l], "ecl")) == 1 {
        ecl = 1
      } else if (strings.Count(stringSplits1[l], "pid")) == 1 {
        pid = 1
      }
    }
    if (byr + iyr + eyr + hgt + hcl + ecl + pid) == 7{
      total += 1
    }
  }
  fmt.Println(total)
}