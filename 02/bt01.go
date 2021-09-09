package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Bài 1 Viết function tìm ra số lớn thứ nhì trong mảng các số.")
	fmt.Println("Ví dụ: max2Numbers([2, 1, 3, 4]) => 3")
	fmt.Println("-------------------------------------------------------------")
	inputarray(4)
}

func inputarray(n int) {
	array := []int{}
	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Nhập array[%d] = ", i)
		str1, err := reader.ReadString('\n')
		str1 = strings.TrimSuffix(str1, "\n")
		a, err := strconv.Atoi(str1)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		array = append(array, a)
	}
	fmt.Println("Mảng bạn vừa nhập gồm các phần tử sau:")
	fmt.Println(array)
	findmaxnumber(array)
	fmt.Println("-------------------------------------------------")
	findmax2number(array)
}

func findmaxnumber(mang1 []int) {
	maxnumber := mang1[0]
	for i := 0; i < len(mang1); i++ {
		if maxnumber < mang1[i] {
			maxnumber = mang1[i]
		}
	}
	fmt.Println("Phần tử lớn nhất trong mảng trên là: ", maxnumber)
}

func findmax2number(mang2 []int) {
	sort.Ints(mang2)
	fmt.Println("Phần tử trong mảng được sắp xếp theo thứ tự lớn dần như sau:", mang2)
	fmt.Println("==> SỐ LỚN THỨ 2 TRONG MẢNG CÁC PHẦN TỬ TRÊN LÀ: ", mang2[len(mang2)-2])
}
