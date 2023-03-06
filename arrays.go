package main
import "fmt"
func main(){
	var a [5]int
	fmt.Println("empty",a)
	a[4]=100
	fmt.Println("assigned ny index",a,"length=",len(a))
	b:=[5]int {1,2,3,4,5}
	fmt.Println("using shorthand",b)
	var twoD [5][5]int 
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			twoD[i][j]=j
		}
	}
	fmt.Println("2D:",twoD)
}