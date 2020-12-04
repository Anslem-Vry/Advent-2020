package main
import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
  "regexp"
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

        byr = validBirth(stringSplits1[l])

      } else if (strings.Count(stringSplits1[l], "iyr")) == 1 {

        iyr = validIssue(stringSplits1[l])

      } else if (strings.Count(stringSplits1[l], "eyr")) == 1 {

        eyr = validExp(stringSplits1[l])

      } else if (strings.Count(stringSplits1[l], "hgt")) == 1 {
        format := validHei(stringSplits1[l])
        if (format) {
          hgt = 1
        }
      } else if (strings.Count(stringSplits1[l], "hcl")) == 1 {
        format := validHcl(stringSplits1[l])
        if (format) {
          hcl = 1
        }
      } else if (strings.Count(stringSplits1[l], "ecl")) == 1 {
        ecl = validEye(stringSplits1[l])

      } else if (strings.Count(stringSplits1[l], "pid")) == 1 {
        format := validPid(stringSplits1[l])
        if (format) {
          pid = 1
        }
      }
    }
    if (byr + iyr + eyr + hgt + hcl + ecl + pid) == 7{
      total += 1
    }
  }
  fmt.Println(total)
}

func validBirth(input string) int {
  stringSplits2 := strings.Split(input, ":")
  check, _ := strconv.Atoi(stringSplits2[1])
  if ((check >= 1920) && (check <= 2002)){
    return 1
  }
  return 0
}

func validIssue(input string) int {
  stringSplits2 := strings.Split(input, ":")
  check, _ := strconv.Atoi(stringSplits2[1])
  if ((check >= 2010) && (check <= 2020)){
    return 1
  }
  return 0
}

func validExp(input string) int {
  stringSplits2 := strings.Split(input, ":")
  check, _ := strconv.Atoi(stringSplits2[1])
  if ((check >= 2020) && (check <= 2030)) {
    return 1
  } 
  return 0
}

func validHei(input string) bool {
  stringSplits2 := strings.Split(input, ":")
  regex := regexp.MustCompile("^([0-9]+)(cm|in)$")
  match := regex.FindStringSubmatch(stringSplits2[1])
  if (match != nil) {
    height, _ := strconv.Atoi(match[1])
    if match[2] == "cm" {
      if (height >= 150 && height <= 193) {
        return true
      }
    } else {
      if (height >= 59 && height <= 76) {
        return true
      }
    }
  }
  return false
}

func validHcl(input string) bool {
  stringSplits2 := strings.Split(input, ":")
  regex := regexp.MustCompile("^#[a-f|0-9]{6}$")
  match := regex.MatchString(stringSplits2[1])
  return match
}

func validEye(input string) int {
  options := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
  stringSplits2 := strings.Split(input, ":")
  for j := 0; j < len(options); j++ {
    if (stringSplits2[1] == options[j]) {
      return 1
    }
  }
  return 0
}

func validPid(input string) bool {
  stringSplits2 := strings.Split(input, ":")
  regex := regexp.MustCompile("^[0-9]{9}$")
  match := regex.MatchString(stringSplits2[1])
  return match
}