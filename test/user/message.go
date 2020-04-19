package user

type CreateArgs struct {
	UserName string
	Phone    string
	Password string
	Token    string
	Device   uint32
	IP       string
}

type CreateBack struct {
	User *User
}
