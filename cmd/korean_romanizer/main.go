package main

import (
	"bufio"
	"fmt"
	k "github.com/srevinsaju/korean-romanizer-go"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		r := k.NewRomanizer(input.Text())
		romanized := r.Romanize()
		fmt.Println(romanized)
	}
}
