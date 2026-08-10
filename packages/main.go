package main

import (
	"fmt"

	"github.com/Gagan2004bansal/LetsLearnGo/auth"
	"github.com/Gagan2004bansal/LetsLearnGo/user"
	"github.com/fatih/color"
)

func main() {

	auth.LoginWithCredientials("gagan", "password@12")

	session := auth.GetSession()
	fmt.Println("Session : ", session)

	user := user.User{
		Email: "bansalgagan2004@gmail.com",
		Name:  "Gagan Bansal",
	}

	color.Red(user.Email)
	color.Green(user.Name)
}
