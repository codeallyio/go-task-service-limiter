package main

// User struct might be needed to check
// if the user is premium or not ;)
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64
}

// HandleRequest runs the processes requested by users.
// Returns false if process has to be killed.
func HandleRequest(process func(), u *User) bool {
	process()
	return true
}

func main() {
	RunMockServer()
}
