package main

import (
	"strings"
	"fmt"
)

func main(){
	text:="/data/resourcecenter/rec64736251test222.png"
	comma := strings.Index(text, "/")
	pos := strings.Index(text[comma:], "rec")
	fmt.Println(text[comma+pos:])

}
