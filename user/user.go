package user

type User struct {
	ID int
	Name string
}

func Greet(u User) string{
	str := "Hi! I'm " + u.Name
	return str
}