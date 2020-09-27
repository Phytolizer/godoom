package deh

import fmt2 "fmt"

type formatArg int

const (
	FormatArgInvalid formatArg = iota
	FormatArgInt               = iota
	FormatArgFloat             = iota
	FormatArgChar              = iota
	FormatArgString            = iota
	FormatArgPtr               = iota
	FormatArgSavePos           = iota
)

func formatArgumentType(c uint8) formatArg {
	switch c {
	case 'd', 'i', 'o', 'u', 'x', 'X':
		return FormatArgInt
	case 'e', 'E', 'f', 'F', 'g', 'G', 'a', 'A':
		return FormatArgFloat
	case 'c', 'C':
		return FormatArgChar
	case 's', 'S':
		return FormatArgString
	case 'p':
		return FormatArgPtr
	case 'n':
		return FormatArgSavePos
	default:
		return FormatArgInvalid
	}
}

func nextFormatArgument(fmt string, rover *int) formatArg {
	for ; *rover < len(fmt); *rover += 1 {
		if fmt[*rover] == '%' {
			*rover += 1
			if fmt[*rover] != '%' {
				break
			}
		}
	}

	for ; *rover < len(fmt); *rover += 1 {
		argtype := formatArgumentType(fmt[*rover])
		if argtype != FormatArgInvalid {
			*rover += 1
			return argtype
		}
	}
	return FormatArgInvalid
}

func validArgumentReplacement(original formatArg, replacement formatArg) bool {
	// Chars are ints
	if original == FormatArgChar && replacement == FormatArgInt {
		return true
	}
	// Strings are pointers
	if original == FormatArgString && replacement == FormatArgPtr {
		return true
	}
	return original == replacement
}

func validFormatReplacement(fmt string, repl string) bool {
	rover1 := 0
	rover2 := 0

	for {
		argtype1 := nextFormatArgument(fmt, &rover1)
		argtype2 := nextFormatArgument(repl, &rover2)
		if argtype2 == FormatArgInvalid {
			// No more args to read from repl string
			break
		} else if argtype1 == FormatArgInvalid {
			// No more args in original string
			return false
		} else if !validArgumentReplacement(argtype1, argtype2) {
			return false
		}
	}
	return true
}

func formatStringReplacement(fmt string) string {
	repl := String(fmt)
	if !validFormatReplacement(fmt, repl) {
		println("WARNING: Unsafe dehacked replacement provided for printf format string: ", fmt)
		return fmt
	}
	return repl
}

func Print(fmt string, args ...interface{}) {
	repl := formatStringReplacement(fmt)

	fmt2.Print(repl)
	fmt2.Print(args...)
}
