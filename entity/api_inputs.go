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

type LoginRequest struct {
	Username string
	Password string
}
