package main
import (
	"fmt"
	"math"
)

func main(){
	var sNum,lNum,res float64
	fmt.Println("Enter a Small Number")
	fmt.Scan(&sNum)
	fmt.Println("Enter a Large Number")
	fmt.Scan(&lNum)
	res= math.Mod(lNum,sNum)
	fmt.Println("The Reamainder is" + res)


}

