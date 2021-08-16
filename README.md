# korean-romanizer-go
korean-romanizer-go is a golang module that romanizes Korean text in Hangul into its alphabet equivalent.

It currently follows the [Revised Romanization of Korean](https://www.korean.go.kr/front_eng/roman/roman_01.do) rule developed by the National Institute of Korean Language, the official romanization system being used in the Republic of Korea.

This repository is a direct port of the original Python transliteration implementation
by the core contributor [Ilkyu Ju](https://github.com/osori), 
in the [osori/korean-romanizer](https://github.com/osori/korean-romanizer) GitHub repository.

## Usage

### Installation
#### Linux
```bash
wget https://github.com/srevinsaju/korean-romanizer-go/releases/latest/korean_romanizer_linux
```

### Basic Usage
```go
package main

import (
	k "github.com/srevinsaju/korean-romanizer-go"
	
)

func main() {
	r := k.NewRomanizer("안녕하세요")
	r.Romanize()
	# returns 'annyeonghaseyo'
}
```

### License
This is a GoLang port of the original python module. So respect its license.
See [LICENSE](./LICENSE) for more information.