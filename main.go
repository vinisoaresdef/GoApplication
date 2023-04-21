package main

import (
	"firstproject/API/config"
	"firstproject/API/router"
	"fmt"
)

func main() {

	//initialize config
	err := config.Init()
	if err != nil {
		fmt.Println("Error initializing config: ", err)
		return // exit the program
	}

	//initialize router
	router.Initialize()
}
