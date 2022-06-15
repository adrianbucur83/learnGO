package main

import "fmt"

const myConst int = 42
const (
	_ = iota //throw away this value, use built in incrementor to assign numeric values (iota)
	mere
	pere
	nuci
)

func main() {
	const myConst = "SADAS 20"
	println(myConst)
	fmt.Println(mere, pere, nuci)
}
