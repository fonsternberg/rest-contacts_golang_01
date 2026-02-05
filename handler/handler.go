package handler

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Handler() {
	var method, id, name, phone string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())

	switch len(parts) {
	case 1:
		method = parts[0]
	case 2:
		method = parts[0]
		id = parts[1]
	case 4:
		method = parts[0]
		id = parts[1]
		name = parts[2]
		phone = parts[3]
	default:
		log.Fatalln("Wrong format of response")	
	}

	switch method {
	case "GET":
		if len(parts) == 2 {
			runGetOne()
		} else if len(parts) == 1 {
			runGetAll()
		}
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