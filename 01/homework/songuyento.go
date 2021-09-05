package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Bài 3: Lập dãy số nguyên tố")
	fmt.Println("Nhập vào số nguyên dương N < 100,000 hãy trả về mảng các số nguyên tố <= N.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nhập giá trị N: ")
	str, err := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	n, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if n < 2 {
		fmt.Println("Số " + str + " không phải là số nguyên tố!")
		return
	}
	fmt.Printf("Các số nguyên tố: ")
	printPrimeNumbers(2, n)
}

func printPrimeNumbers(num1, num2 int) {
	if num1 < 2 || num2 < 2 {
		fmt.Println("Không phải là số nguyên tố!.")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
		}
		num1++
	}
	fmt.Println()
}
