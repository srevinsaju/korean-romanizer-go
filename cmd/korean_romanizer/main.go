package main

import (
	"bufio"
	"fmt"
	k "github.com/srevinsaju/korean-romanizer-go"
	"log"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		r := k.NewRomanizer(input.Text())
		romanized, err := r.Romanize()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(romanized)
	}
}
