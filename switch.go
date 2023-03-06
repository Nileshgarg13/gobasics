package main
import (
	"fmt"
	"time"
)
func main(){
	i:=2
	fmt.Println("Write",i,"as")
	switch i{
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	}
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("WEEKEND")
	default:
		fmt.Println("WEEKDAY")
	}
	t:=time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("BEFORE NOON")
	default:
		fmt.Println("AFTERNOON")
	}
	whatAmI := func(i interface{}){
		switch t :=i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a integer")
		default:
            fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}