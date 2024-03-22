package entity

type UserInput struct {
	Username string
	Name     string
	Lastname string
	Age      int
	Password string
}

type EncryptRequest struct {
	Text string
}
