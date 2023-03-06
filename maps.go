package main
import "fmt"
func main(){
	m1:=make(map[string]int)
	m1["k1"]=1
	m1["k2"]=2
	m1["k3"]=3
	fmt.Println(m1)
	v1:=m1["k2"]
	fmt.Println("getting value directly",v1)
	fmt.Println("len=",len(m1))
	delete(m1,"k2")
	fmt.Println(m1)
	_,prs:=m1["k2"]
	fmt.Println("geting invalid key",prs)
	_,ans:=m1["k1"]
	fmt.Println("valid",ans)
	n:=map[string]int{"k1":1,"k2":20}
	fmt.Println(n)
}