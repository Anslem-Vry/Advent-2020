package main
import (
  "fmt"
  "bufio"
  "os"
  "log"
  "strconv"
)

func main() {
  input, err := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  i, k, l, m := 0, 0, 0, 0
  j := 1
  var inp = new([200]int)

  if err != nil{
    fmt.Println("File reading error", err)
    return
  }
  defer input.Close()
  
  for scanner.Scan(){
    inp[i], err = strconv.Atoi(scanner.Text())
    i++
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  for i = 0; i < 200; i++{
    for j = 1; j < 200; j++ {
      if (inp[i] + inp[j]) < 2020 {
        m = 2020-(inp[i] + inp[j])
        fmt.Println(m)
          for l = 0; l < 200; l++{
            if ((l == i) || l == j){
              l++
            } else if (inp[l] == m){
              k = (inp[i] * inp[j] * inp[l])
            }
          }
      }
    }
  }
  fmt.Println("Contents of file:\n", k)
}