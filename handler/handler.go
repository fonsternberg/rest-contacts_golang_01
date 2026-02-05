package handler

import (
	"fmt"
	"log"

)

func Handler() {
	var method, id, name, phone string

	n, err := fmt.Scan(method, id, name, phone)
	if n > 4 || err != nil {
		log.Fatalln("Wrong format of response")
	}
	switch method {
	case "GET":
		runGet()
	case "PUT":
		runPut()
	case "DELETE":
		runDelete()
	case "POST":
		runPost()
	default:
		log.Fatalln("Undefined method")
	}
}