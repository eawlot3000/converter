package main
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ==== FROM BINARY ====
func binary_decimal(b string) int {
	reverse := ""
	output := 0
	// only integer part
	for i := range b {
		reverse += string(b[len(b)-1-i])
		if b[len(b)-1-i] == '1' {
			output += 1 << uint(i)
		}
	}
	return output
}

func binary_octal(b string) string {
	dic := map[string]int{"001": 1, "010": 2, "011": 3, "100": 4, "101": 5, "110": 6, "111": 7}
  var out string = ""
	parts := strings.Split(strings.TrimRightFunc(b, func(r rune) bool { return r == '0' }), "000")
	if len(parts[0]) < 3 {
		parts[0] = strings.Repeat("0", 3-len(parts[0])) + parts[0]
	}
	for _, p := range parts {
		out += strconv.Itoa(dic[p])
	}
	return out
}

func binary_hexadecimal(b string) string {
	dic := map[string]string{"0000": "0", "0001": "1", "0010": "2", "0011": "3", "0100": "4", "0101": "5", "0110": "6", "0111": "7", "1000": "8", "1001": "9", "1010": "A", "1011": "B", "1100": "C", "1101": "D", "1110": "E", "1111": "F"}
	out := ""
	parts := strings.Split(strings.TrimRightFunc(b, func(r rune) bool { return r == '0' }), "0000")
	if len(parts[0]) < 4 {
		parts[0] = strings.Repeat("0", 4-len(parts[0])) + parts[0]
	}
	for _, p := range parts {
		out += dic[p]
	}
	return out
}

// ==== FROM DECIMAL ====
func decimal_binary(d int) string {
	out := ""
	for d > 0 {
		out = strconv.Itoa(d%2) + out
		d /= 2
	}
	return out
}

func decimal_octal(d int) string {
	out := ""
	for d > 0 {
		out = strconv.Itoa(d%8) + out
		d /= 8
	}
	return out
}

func decimal_hexadecimal(d int) string {
	dic := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F"}
	out := ""
	for d > 0 {
		out = dic[d%16] + out
		d /= 16
	}
	return out
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: convert [FROM] [TO] [DATA]")
		fmt.Println("e.g. convert b d 100110011")
		os.Exit(0)
	}
	arg1, arg2, arg3 := os.Args[1], os.Args[2], os.Args[3]
  if arg1 == "b" {
    if arg2 == "d" {
    		fmt.Println(binary_decimal(arg3))
    	} else if arg2 == "o" {
    		fmt.Println(binary_octal(arg3))
    	} else if arg2 == "h" {
    		fmt.Println(binary_hexadecimal(arg3))
    	}
    } else if arg1 == "d" {
    	decimalData, _ := strconv.Atoi(arg3)
    	if arg2 == "b" {
    		fmt.Println(decimal_binary(decimalData))
    	} else if arg2 == "o" {
    		fmt.Println(decimal_octal(decimalData))
    	} else if arg2 == "h" {
    		fmt.Println(decimal_hexadecimal(decimalData))
    	}
    } else {
    	fmt.Println("Invalid input")
    }
}
