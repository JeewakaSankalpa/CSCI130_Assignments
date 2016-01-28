package main
import "fmt"

func max(num ...int) int  {
	max :=0
	for _,temp:=range num  {
		if max<temp{
			max = temp
		}
	}
	return max
}
func main(){

	list:=[]int{50,40,35,25,26,34,89,85,45}

	fmt.Println(max(list...))
}
