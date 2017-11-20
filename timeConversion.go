package main 

import "fmt"
import "strconv"

var(
	jam string
	
)
func main() {
	fmt.Scan(&jam)
	fmt.Println(jam)
	var hasil string
	if jam[8] == 'P' || jam[8] == 'p'{
		a,_ := strconv.ParseInt(jam[0:2],10,64)
		if a==12{
			a=0
		}else{
			a = a+12
		}
		b := strconv.FormatInt(a,10)
		hasil = b+jam[2:8]
	}else{
		hasil=jam
	}
	
	

	fmt.Println(hasil[0:8])
}