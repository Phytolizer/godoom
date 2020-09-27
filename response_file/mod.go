package response_file

import (
	"godoom/argv"
	error2 "godoom/errors"
	"godoom/misc"
	"os"
	"unicode"
)

func FindResponseFile() {
	var i int
	for i1, arg := range argv.MyArgs {
		i = i1
		if arg[0] == '@' {
			loadResponseFile(i, argv.MyArgs[i][1:])
		}
	}

	for {
		i = argv.CheckParmWithArgs("-response", 1)
		if i <= 0 {
			break
		}

		argv.MyArgs[i] = "-_"
		loadResponseFile(i+1, argv.MyArgs[i+1])
	}
}

func loadResponseFile(argvIndex int, filename string) {
	handle, err := os.Open(filename)
	if err != nil {
		println("\nNo such response file!")
		os.Exit(1)
	}

	println("Found response file ", filename, "!")

	size, err := misc.FileLength(handle)
	if err != nil {
		println("Error getting response file length: ", err)
		os.Exit(1)
	}

	file := make([]byte, size)

	for i := int64(0); i < size; {
		k, err := handle.Read(file)
		if err != nil {
			println("Error reading response file: ", err)
			os.Exit(1)
		}
		i += int64(k)
	}

	err = handle.Close()
	if err != nil {
		println("Error closing response file: ", err)
		os.Exit(1)
	}

	newargs := make([]string, argv.MaxArgvs)
	newargc := 0

	for i := 0; i < argvIndex; i += 1 {
		newargs[i] = argv.MyArgs[i]
		newargc += 1
	}

	for k := int64(0); k < size; {
		// skip space
		for ; k < size && unicode.IsSpace(rune(file[k])); k += 1 {
		}

		if k >= size {
			break
		}

		if file[k] == '"' {
			// skip beginning quote
			k += 1
			start := k

			// yoink the string
			for ; k < size && file[k] != '"' && file[k] != '\n'; k += 1 {
			}

			// if we hit \n or EOF, string was unterminated
			if k >= size || file[k] == '\n' {
				error2.Error("Quotes unclosed in response file '", filename, "'")
			}

			end := k

			newargs[newargc] = string(file[start:end])
			newargc += 1
		} else {
			start := k
			// yoink a token
			for ; k < size && !unicode.IsSpace(rune(file[k])); k += 1 {
			}
			end := k
			newargs[newargc] = string(file[start:end])
			newargc += 1
		}
	}

	for _, arg := range argv.MyArgs[argvIndex+1:] {
		newargs[newargc] = arg
		newargc += 1
	}

	argv.MyArgs = newargs
}
