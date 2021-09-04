package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Bài 1: Giải phương trình bậc 2")
	fmt.Println("Nhập vào ba số a, b, c kiểu int hãy giải phương trình bậc 2. a*x*x+b*x+c=0")

	reader := bufio.NewReader(os.Stdin)
	//Nhập a #0
	fmt.Print("Nhập a = ")
	str1, err := reader.ReadString('\n')
	str1 = strings.TrimSuffix(str1, "\n")
	a, err := strconv.Atoi(str1)
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
	b, err := strconv.Atoi(str2)

	//Nhập c
	fmt.Print("Nhập c = ")
	str3, err := reader.ReadString('\n')
	str3 = strings.TrimSuffix(str3, "\n")
	c, err := strconv.Atoi(str3)

	fmt.Print("Phương trình bậc 2 là: " + str1 + "x^2 + " + str2 + "x + " + str3 + " = 0 \n")
	var delta int
	delta = (b * b) - 4*a*c

	switch {
	case delta == 0:
		{
			fmt.Println("Phương trình có nghiệm kép x=", -b/(2*a))
		}
	case delta < 0:
		{
			fmt.Println("Phương trình đã cho vô nghiệm")
		}
	case delta > 0:
		{
			fmt.Print("Phương trình tồn tại 2 nghiệm:") //, -b+math.Sqrt(delta)/(2*a), -b-math.Sqrt(delta)/(2*a))
			//fmt.Println("x1 = ", %(-b + math.Sqrt(delta)/2*a))
			//fmt.Println("x2 = ", %(-b - math.Sqrt(delta)/2*a))
		}
	}

}
