package model

type User struct {
	UserName string
	Password string
	Email string
	Phone string
}

func (user User) GetUserName() string {
	return user.UserName;
}

func (user *User) SetUserName(username string) {
	user.UserName = username
}

func (user User) GetPassword() string {
	return user.Password
}

func (user *User) SetPassword(password string) {
	user.Password = password
}

func (user User) GetEmail() string {
	return user.Email
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user User) GetPhone() string {
	return user.Phone
}

func (user *User) SetPhone(phone string) {
	user.Phone = phone
}