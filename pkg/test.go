package pkg

import (
	"fmt"
	"time"
)


func Test(){
	for {
		fmt.Println("Hello World")
		time.Sleep(5 * time.Second)
	}
}