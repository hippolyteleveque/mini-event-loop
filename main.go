package main

import "fmt"
import "time"

func main() {

	var callStack []string
	var callbackStack []string

	// main stack loop
	go func() {
		for {
			if len(callStack) > 0 {
				called := callStack[0]
				callStack = callStack[1:]
				fmt.Printf("Called: [%s]\n", called)
			}
		}
	}()

	// event loop
	go func() {
		for {
			if len(callStack) == 0 && len(callbackStack) > 0 {
				toPush := callbackStack[0]
				callbackStack = callbackStack[1:]
				callStack = append(callStack, toPush)
			}	
		}
	}()

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <- ticker.C:
				callbackStack = append(callbackStack, "Bonjour")
			}
		}
	}()

	for i := 0; ; i++ {
		callStack = append(callStack, "Hello")
		time.Sleep(3 * time.Second)
	}
}