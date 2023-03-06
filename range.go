package main
import "fmt"
func main(){
	nums:=[]int{1,2,3}
	sum:=0
	for _,num :=range nums{
		sum+=num
	}
	fmt.Println("sum=",sum)
	for i,num:= range nums{
		if num==3{
			fmt.Println("index for 3 is",i)//we used _ above because we dont need index in go we can iterate wuth both value and index
		}
	}
	kvs:=map[string]string{"a":"apple","b":"ball"}
	for k,v:= range kvs{
		fmt.Printf("%s-> %s\n",k,v)
	}
	for k:=range kvs{
		fmt.Println("Key",k)
	}
	for i,c:=range "go"{
		fmt.Println(i,c)
	}
}