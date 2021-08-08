package main

type laptopSize string

func (ls laptopSize) getSizeOfLaptop() {
	return ls
}

func main() {
	var te laptopSize = "hahah"
	te.getSizeOfLaptop()
}
