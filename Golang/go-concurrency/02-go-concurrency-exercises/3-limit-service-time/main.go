package main

type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64
}

func HandleRequest(process func(), u *User) bool {
	process()
	return true
}

func main() {
	RunMockServer()
}
