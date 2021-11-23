package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           uint `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	// gorm:"unique" とすることでStructの定義の段階で値のバリデーションができる
	Email        string `json:"email" gorm:"unique"`
	// Password と IsAmbassadorはjsonとして返す際に含めたくないので - としておく
	Password     []byte `json:"-"`
	IsAmbassador bool `json:"-"`
}

func (user *User) SetPassword(password string) {
	// 第二引数ではhashアルゴリズムを繰り返す回数を指定する
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
