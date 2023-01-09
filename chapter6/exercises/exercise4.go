package main

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = 1
		i += 2
		return
	}
}

func main() {
	makeOddGenerator()
}
