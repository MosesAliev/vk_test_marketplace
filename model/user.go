package model

type User struct {
	Login    string `gorm:"primaryKey"`
	Password string
	Ads      []Ad `gorm:"foreignKey:UserLogin"`
}
