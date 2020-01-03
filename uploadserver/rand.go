package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateCaptcha() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}

func main(){
	fmt.Print(CreateCaptcha())
}