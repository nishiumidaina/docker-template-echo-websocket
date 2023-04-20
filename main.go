package main

import (
	"docker-echo-template/controllers"
)

func main() {
	router := controller.Router()
    router.Start(":8000")
}