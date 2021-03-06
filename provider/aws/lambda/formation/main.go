package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lob/rack/provider/aws/lambda/formation/handler"
)

func die(err error) {
	fmt.Printf("die error: %s\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die(fmt.Errorf("must specify event as argument"))
	}

	data := []byte(os.Args[1])

	var req handler.Request

	err := json.Unmarshal(data, &req)

	if err != nil {
		die(err)
	}

	fmt.Printf("main req = %+v\n", req)

	err = handler.HandleRequest(req)

	if err != nil {
		fmt.Printf("main error: %s\n", err)
		return
	}
}
