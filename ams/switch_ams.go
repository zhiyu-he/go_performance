package main

/*
	go tool compile -S switch_ams.go
	only for-each
*/
func IsSomething(num int) bool {
	switch num {
	case 1, 3, 5, 7, 9:
		return true
	default:
		return false
	}
}

func IsSomething2(num int) bool {
	switch num {
	case 1, 10, 100, 1000, 10000, 100000:
		return true
	default:
		return false
	}
}

func main() {
	IsSomething(10)
}
