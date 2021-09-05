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
	fmt.Println("Bài 1: Giải phương trình bậc 2")
	fmt.Println("Nhập vào ba số a, b, c kiểu int hãy giải phương trình bậc 2. a*x*x+b*x+c=0")

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

	//Nhập c
	fmt.Print("Nhập c = ")
	str3, err := reader.ReadString('\n')
	str3 = strings.TrimSuffix(str3, "\n")
	c, err := strconv.ParseFloat(str3, 64)

	fmt.Print("Phương trình bậc 2 là: " + str1 + "x^2 + " + str2 + "x + " + str3 + " = 0 \n")
	var delta float64
	delta = (b * b) - 4*a*c

	switch {
	case delta == 0:
		{
			fmt.Println("Phương trình có nghiệm kép x1=x2=", -b/(2*a))
		}
	case delta < 0:
		{
			fmt.Println("Phương trình đã cho vô nghiệm")
		}
	case delta > 0:
		{
			x1 := -b + math.Sqrt((b*b)-4*a*c)/(2*a)
			x2 := -b - math.Sqrt((b*b)-4*a*c)/(2*a)
			fmt.Println("Phương trình tồn tại 2 nghiệm:")
			fmt.Println("x1 = ", x1)
			fmt.Println("x2 = ", x2)
		}
	}

}
