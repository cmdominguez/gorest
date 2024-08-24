package utils

func LogAlgo(msg string) {
	println(msg)
	logPrivado(msg)
}

// Privado

func logPrivado(msg string) {
	println("Privado", msg)
}
