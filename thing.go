package thing

type Thinger interface {
	Thing(val string) string
}

func FuncToTest(thinger Thinger, val string) string {
	return thinger.Thing(val)
}
