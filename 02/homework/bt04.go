package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Bài 4 Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp")
	fmt.Println("Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:")
	fmt.Println("- Sắp xếp tên nhân viên tăng dần theo bảng chữ cái")
	fmt.Println("- Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)")
	fmt.Println("- Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	sortSliceWithFunc()
}

func sortSliceWithFunc() {
	people := []struct {
		Name      string
		Ratio     float64
		Allowance float64
	}{
		{"Gopher", 2.5, 100000},
		{"Alice", 2.6, 100000},
		{"Vera", 2.6, 110000},
		{"Bob", 2.7, 110000},
	}
	fmt.Println("Danh sách nhân viên ban đầu: ", people)
	fmt.Println()
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("Danh sách nhân viên sắp theo tên :", people)
	fmt.Println()

	sort.Slice(people, func(i, j int) bool { return people[i].Ratio*people[i].Allowance > people[j].Ratio*people[j].Allowance })
	fmt.Println("Danh sách nhân viên theo mức lương giảm dần:", people)
	fmt.Println()

	fmt.Println("Danh sách nhân viên có mức lương lớn thứ 2:", people[1])
	fmt.Println()
}
