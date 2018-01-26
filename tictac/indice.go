package tictac

var indice = 0

func GetIndice() int {
	val := make(chan int)
	go inc(val)
	return <- val
}

func inc(val chan int) {
	indice ++
	val <- indice
}