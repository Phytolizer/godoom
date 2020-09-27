package system

func PrintBanner(msg string) {
	spaces := 35 - (len(msg) / 2)

	for i := 0; i < spaces; i += 1 {
		print(" ")
	}
	println(msg)
}
