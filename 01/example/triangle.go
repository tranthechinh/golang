package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Nhập 3 số tương ứng; a, b, c, kiểm tra xem 3 số vừa nhập có lập thành một tam giác hay không")

	reader := bufio.NewReader(os.Stdin)
	//Nhập a #0
	fmt.Print("Nhập a = ")
	str1, err := reader.ReadString('\n')
	str1 = strings.TrimSuffix(str1, "\n")
	a, err := strconv.ParseFloat(str1, 64)
	if a == 0 {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Nhập b
	fmt.Print("Nhập b = ")
	str2, err := reader.ReadString('\n')
	str2 = strings.TrimSuffix(str2, "\n")
	b, err := strconv.ParseFloat(str2, 64)
	if b == 0 {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Nhập c
	fmt.Print("Nhập c = ")
	str3, err := reader.ReadString('\n')
	str3 = strings.TrimSuffix(str3, "\n")
	c, err := strconv.ParseFloat(str3, 64)
	if c == 0 {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if a+b > c && a+c > b && b+c > a {
		if a*a+b*b == c*c || b*b+c*c == a*a || c*c+a*a == b*b {
			fmt.Println("Ba số mà bạn vừa nhập là độ dài 3 cạnh của một tam giác vuông")
		} else {
			if a == b && b == c {
				fmt.Println("Ba số mà bạn vừa nhập là độ dài 3 cạnh của một tam giác đều")
			} else {
				if a*a > b*b+c*c || b*b > a*a+c*c || c*c > a*a+b*b {
					fmt.Println("Ba số mà bạn vừa nhập là độ dài 3 cạnh của một tam giác tù")
				} else {
					fmt.Println("Ba số mà bạn vừa nhập là độ dài 3 cạnh của một tam giác nhọn")
				}
			}

		}

	} else {
		println("a = " + str1 + ", b = " + str2 + " và c = " + str3 + " mà bạn vừa nhập không phải là độ dài 3 cạnh của một tam giác")
	}

	//Có thể dùng switch case thay cho if else...

}
