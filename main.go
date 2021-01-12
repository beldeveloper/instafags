package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	sdk "github.com/ahmdrz/goinsta/v2"
	"github.com/beldeveloper/instafags/controller"
	"github.com/beldeveloper/instafags/repository/goinsta"
)

func main() {
	client := sdk.New(os.Getenv("IF_USERNAME"), os.Getenv("IF_PASSWORD"))
	err := client.Login()
	if err != nil {
		log.Printf("error while login: %v\n", err)
		return
	}
	e := strings.Split(os.Getenv("IF_EXCEPTIONS"), ",")
	c := controller.NewController(goinsta.Instagram{Client: client}, e)
	fags, err := c.ListFags(os.Getenv("IF_ACCOUNT"))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Fags:")
	for _, u := range fags {
		fmt.Println(u.Username)
	}
}
