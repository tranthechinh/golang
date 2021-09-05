package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i < 5; i++ {

		n := rand.Intn(100)
		reader := bufio.NewReader(os.Stdin)
		//Nhập số đoán
		fmt.Println("Lần đoán số: ", i)
		fmt.Print("Bạn đoán xem máy sẽ tự sinh ra số nào? Nhập số ở đây: ")
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
		fmt.Println("Số máy tính tự sinh là: ", n)
		fmt.Println("Số bạn đoán là: ", a)
		switch {

		case a < n:
			fmt.Println("Số bạn đoán nhỏ hơn số máy tự sinh")
			fmt.Println("-------------------------------------")
		case a > n:
			fmt.Println("Số bạn đoán lớn hơn số máy tự sinh")
			fmt.Println("-------------------------------------")
		default:
			fmt.Println("Chúc mừng bạn, bạn đã đoán đúng số máy sinh ra!")
			fmt.Println("-------------------------------------")
		}
	}
}
