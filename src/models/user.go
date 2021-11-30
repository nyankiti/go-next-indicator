package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// gorm:"unique" とすることでStructの定義の段階で値のバリデーションができる
	Email string `json:"email" gorm:"unique"`
	// Password と IsAmbassadorはjsonとして返す際に含めたくないので - としておく
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"-"`
	// omitemptyとすることで、フィールドの値が空の場合はjsonへとエンコーディングされる時に省略される
	// pointerを格納しているのは、値でomitemptyを使うと、値が0だった場合も省略されてしまうから。
	// 0も立派な値としてjsonで返却したい場合はpointerを使うと良い
	Revenue *float64 `json:"revenue,omitempty" gorm:"-"`
}

func (user *User) SetPassword(password string) {
	// 第二引数ではhashアルゴリズムを繰り返す回数を指定する
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

// userがadminかambassadorかによって同じ名前で違う結果を返すメソッドを実装したい
//（ポリモーフィズム）ので、以下のようにAliasesを作る

type Admin User

func (admin *Admin) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   admin.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AdminRevenue
		}
	}

	admin.Revenue = &revenue
}

type Ambassador User

func (ambassador *Ambassador) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   ambassador.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AmbassadorRevenue
		}
	}

	ambassador.Revenue = &revenue
}
