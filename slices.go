package main
import "fmt"
func main(){
	s:=make([]string,3)
	fmt.Println("empty",s)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set",s)
	fmt.Println(("get 2nd"),s[1])
	fmt.Println("length",len(s))
	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("copy",c)
	s=append(s,"aa")
	s=append(s, "sd")
	fmt.Println("appended",s)
	l:=s[2:4]
	fmt.Println("splicing splice",l)
	l=s[:4]
	fmt.Println("from first",l)
	l=s[2:]
	fmt.Println(("till last"),l)
	twoD:=make([][]int,3)
	for i:=0;i<3;i++{
		innerlen:=i+1
		twoD[i]=make([]int,innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d",twoD)
}