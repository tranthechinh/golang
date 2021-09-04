package main

import (
	"bufio"
	"fmt"
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

}
