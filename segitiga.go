package main 

import "fmt"

var n int
func main() {
	fmt.Scanln(&n)
	var tab[10][10] int
	i:=0
	for i<n {
		j:=0
		for j<=i{
			if(j==0) || (j==i){
				tab[i][j] = 1
			}else{
				tab[i][j]=tab[i-1][j-1] + tab[i-1][j]
			}
			j=j+1
		}
		i=i+1
	}
	i=0
	for i<n {
		fmt.Print("\n")
		j:=0
		for j<=i {
			fmt.Print(tab[i][j])
			j=j+1
		}
		i=i+1
	}
}