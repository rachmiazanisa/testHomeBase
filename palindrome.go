package main 
import "fmt"

var(
	kata byte
	p[17] byte
	sum int
	valid bool
	
)

func main() {
	fmt.Print("input kata :")
	fmt.Scanf("%c",&kata)
	i:=1
	for kata !='\n'{
		p[i] = kata
		i=i+1
		fmt.Scanf("%c",&kata)
	}
	j:=i-1
	k:=0

	for j>=i/2 {
		if(p[k]==p[j]){
			sum=sum+1
		}
		j=j-1
		k=k+1
	}

	if(sum == i/2){
		valid=true
	}
	fmt.Print(valid)
	
}