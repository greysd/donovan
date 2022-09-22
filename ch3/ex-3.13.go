package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	fmt.Printf("Kilo : %T\n Mega: %T\n Giga: %T\n Tera: %T\n Peta: %T\n Exa: %T\n Zeta: %T\n Yota: %T\n",
		KiB, MiB, GiB, TiB, PiB, EiB, ZiB, YiB)
}
