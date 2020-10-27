package main

// use this struct to check whether
// a User is premium or not
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
