package main

import (
  "fmt"
  "os"
  "strconv"
)

// TODO: Add support for negative numbers
// binary remainder for division [10101011100011 / 1100101011 = 1101 R: 110110100]
func add(a *string, b *string, base int) *string {
  baseConvert := map[int]func(int64, int) string{16: strconv.FormatInt, 8: strconv.FormatInt, 2: strconv.FormatInt}
  if conv, ok := baseConvert[base]; ok {
    ia, _ := strconv.ParseInt(*a, base, 0)
    ib, _ := strconv.ParseInt(*b, base, 0)
    result := ia + ib
    str := conv(result, base) //str[2:]
    return &str
  } else {
    str := "Invalid base"
    return &str
  }
}

func subtract(a *string, b *string, base int) *string {
  baseConvert := map[int]func(int64, int) string{16: strconv.FormatInt, 8: strconv.FormatInt, 2: strconv.FormatInt}
  if conv, ok := baseConvert[base]; ok {
    ia, _ := strconv.ParseInt(*a, base, 64)
    ib, _ := strconv.ParseInt(*b, base, 64)
    result := ia - ib
    str := conv(result, base)
    return &str
  } else {
    str := "Invalid base"
    return &str
  }
}

func multiply(a *string, b *string, base int) *string {
  baseConvert := map[int]func(int64, int) string{16: strconv.FormatInt, 8: strconv.FormatInt, 2: strconv.FormatInt}
  if conv, ok := baseConvert[base]; ok {
    ia, _ := strconv.ParseInt(*a, base, 64)
    ib, _ := strconv.ParseInt(*b, base, 64)
    result := ia * ib
    str := conv(result, base)
    return &str
  } else {
    str := "Invalid base"
    return &str
  }
}

func divide(a *string, b *string, base int) *string {
  baseConvert := map[int]func(int64, int) string{16: strconv.FormatInt, 8: strconv.FormatInt, 2: strconv.FormatInt}
  if conv, ok := baseConvert[base]; ok {
    ia, _ := strconv.ParseInt(*a, base, 64)
    ib, _ := strconv.ParseInt(*b, base, 64)
    result := ia / ib
    str := conv(result, base)
    return &str
  } else {
    str := "Invalid base"
    return &str
  }
}

func main() {
  if len(os.Args) < 5 {
    fmt.Println("Usage: calc [OPERATOR] [BASE] [DATA1] [DATA2]")
    fmt.Println("e.g. calc add 2 100 110011: add 2 binary numbers 100 and 110011")
    os.Exit(0)
  }

  arg1 := os.Args[1]
  arg2, err := strconv.Atoi(os.Args[2])
  if err != nil {
    fmt.Println("heyy ur almost there! please use an integer(2 for binary, 8 for octal, 16 for hex)")
    os.Exit(0)
  }
  arg3 := os.Args[3]
  arg4 := os.Args[4]

  var result *string

  switch arg1 {
  case "add":
    result = add(&arg3, &arg4, arg2)
  case "subtract":
    result = subtract(&arg3, &arg4, arg2)
  case "multiply":
    result = multiply(&arg3, &arg4, arg2)
  case "divide":
    result = divide(&arg3, &arg4, arg2)
  default:
    fmt.Println("Invalid operator")
    os.Exit(0)
  }
  fmt.Println(*result)
}
