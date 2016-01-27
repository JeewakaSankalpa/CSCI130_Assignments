package main
import (
	"log"
	"os"
	"text/template"
	"fmt"
)
type loginInfo struct {
	userName string
	password string
	loginSuccess bool

}
func main(){
	logindata := loginInfo{userName:"Admin",password : "password" , loginSuccess: false}

	var user,passwrod string
	fmt.Println("Enter user Name : ")
	fmt.Scan(&user)
	fmt.Println("Enter Password : ")
	fmt.Scan(&passwrod)
	if user==logindata.userName && passwrod == logindata.password {
		logindata.loginSuccess = true
	}
	welcome, err := template.ParseFiles("welcome.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = welcome.Execute(os.Stdout, logindata)
	if err != nil {
		log.Fatalln(err)
	}

}
