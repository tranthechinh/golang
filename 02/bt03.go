package main

import (
	"fmt"
	"os"
)

func main() {
	inputarray(7)
}

func inputarray(n int) {
	fmt.Println("Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng")
	fmt.Println("Ví dụ: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]")
	fmt.Println("---------------------------------------------------------------------")
	input := []int{}
	for i := 0; i < n; i++ {
		fmt.Printf("Nhập array[%d] = ", i)
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Println("Bạn nhập dữ liệu không đúng!")
			return
		}
		input = append(input, number)
	}
	fmt.Println("Mảng bạn vừa nhập gồm các phần tử sau: ", input)
	removeDuplicate(input)
}

func removeDuplicate(a []int) (result []int) {
	keys := map[int]bool{} //hoặc có thể viết keys := make(map[int]bool) nhưng dài dòng
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	fmt.Printf("Các phần tử trong mảng sau khi remove những phần tử bị trùng lặp là: %d", result)
	fmt.Println()
	return
}
