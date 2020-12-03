package main
import (
  "fmt"
  "bufio"
  "os"
  "log"
  "strconv"
  "strings"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  i := 1
  total := 0

  if err != nil{
    fmt.Println("File reading error", err)
    return
  }
  defer input.Close()
  for scanner.Scan(){ //read each line
    init := (scanner.Text())
    stringSplits1 := strings.Split(init, " ")
    stringSplits2 := strings.Split(stringSplits1[0], "-")
    stringSplits3 := strings.Split(stringSplits1[1], ":")
    lower, err := strconv.Atoi(stringSplits2[0])
    upper, err := strconv.Atoi(stringSplits2[1])
    charToCheck := stringSplits3[0]
    valueString := stringSplits1[2]
    occurences := 0
    for i = 0; i < len(valueString); i++ {
      currentChar := string(valueString[i])
      if currentChar == charToCheck{
        occurences += 1
      }
    }
    if (occurences >= lower) && (occurences <= upper){
      total += 1
    }
    if err != nil{
      return
    }
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  fmt.Println(total)
}
