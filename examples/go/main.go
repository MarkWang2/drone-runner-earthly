package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(HelloWorld())
}

// HelloWorld is a function that returns a string containing "hello world".
func HelloWorld() string {
	logrus.Info("hello world")
	return "hello world"
}
