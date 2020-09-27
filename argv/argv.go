package argv

var MyArgs []string

const MaxArgvs = 100

func ParmExists(parm string) bool {
	for _, arg := range MyArgs {
		if arg == parm {
			return true
		}
	}
	return false
}

func CheckParm(parm string) (index int) {
	return CheckParmWithArgs(parm, 0)
}

func CheckParmWithArgs(parm string, args int) (index int) {
	if len(MyArgs) > 1 {
		for i, arg := range MyArgs[1 : len(MyArgs)-args+1] {
			if parm == arg {
				return i
			}
		}
	}
	return 0
}
