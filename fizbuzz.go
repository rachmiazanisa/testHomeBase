package main 

import "fmt"
var(
	angka,hasil2,hasil1 int
)
func main() {
	fmt.Print("Input Angka : ")
	fmt.Scanln(&angka)
	hasil1=angka%3
	hasil2=angka%5
	if(hasil1 == 0){
		fmt.Print("FIZZ ")
	}
	if(hasil2 == 0){
		fmt.Println("BUZZ")
	}
	if(hasil1 !=0) && (hasil2 !=0){
		fmt.Println(hasil1)
		fmt.Println(hasil2)
	}

}