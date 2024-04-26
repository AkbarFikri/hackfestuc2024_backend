package entity

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"not null;unique"`
	Password string
	FullName string
}
