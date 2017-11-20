package main 

import "fmt"

func main(){
	var v byte
	var d[17] byte
	fmt.Scanf("%c", &v)
	dl:=0
	for v != '\r' && v != '\n'{
		d[dl]=v
		dl=dl+1
		fmt.Scanf("%c", &v)
	}
	i:=dl-1
	for i>=0 {
		fmt.Printf("%c",d[i])
		i=i-1
	}
}