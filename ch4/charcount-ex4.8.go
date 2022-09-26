package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	typeOfRune := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int // Количество длин кодировок UTF-8
	invalid := 0                    // Количество некорректных символов UTF-8
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Возврат руны,  байтов,  ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		case {
			unicode.isControl(r):
				typeOfRune["Control"]++
			unicode.isDigit(r):
				typeOfRune["Digit"]++
			unicode.isGraphic(r):
				typeOfRune["Graphic"]++
			unicode.isLetter(r):
				typeOfRune["Letter"]++
			unicode.isLower(r):
				typtypeOfRune["Lower"]
			unicode.isMark(r):
				typeOfRune["Mark"]
			unicode.IsNumber(r)
				typeOfRune["Number"]++
			unicode.IsNumber(r)
				typeOfRune{t}
		}

	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("M%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
}
