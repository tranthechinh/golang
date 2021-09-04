package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	if city != "\n" {
		fmt.Print("You live in " + city) //Nếu bạn nhập dữ liệu vào thì in dòng này
	} else {
		fmt.Println("Bye bye!!! :D") //Nếu bạn không nhập dữ liệu thì in dòng này
	}
}
