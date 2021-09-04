package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Tính giai thừa của n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nhập giá trị n: ")
	str, err := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	n, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Giai thừa của "+str+" là: ", factorial(n))
}

func factorial(i int) int {

	if i == 1 {
		return 1
	}
	return i * factorial(i-1)
}
