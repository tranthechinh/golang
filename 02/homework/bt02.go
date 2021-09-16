package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Bài 2 Cho 1 mảng các chuỗi. Viết function lọc ra các phần tử có độ dài lớn nhất.")
	fmt.Println("Ví dụ: findMaxLengthElement[aba, aa, ad, c, vcd] => [aba, vcd]")
	fmt.Println("-------------------------------------------------")
	inputarray(5)
}

func inputarray(n int) {
	input := []string{}
	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Nhập array[%d] = ", i)
		str1, err := reader.ReadString('\n')
		str1 = strings.TrimSuffix(str1, "\n")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		input = append(input, str1)
	}
	fmt.Println("Mảng bạn vừa nhập gồm các phần tử sau: ", input)
	findMaxLengthElement(input)
}

func findMaxLengthElement(arr []string) {
	maxnumber := len(arr[0])
	var newarr []string
	for i := 0; i < len(arr); i++ {
		if maxnumber <= len(arr[i]) {
			maxnumber = len(arr[i])
		}
	}
	for i := 0; i < len(arr); i++ {
		if maxnumber <= len(arr[i]) {
			newarr = append(newarr, arr[i])
		}
	}
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Các phần tử trong mảng có độ dài lớn nhất là %d gồm: %v", maxnumber, newarr)
	fmt.Println()
}
