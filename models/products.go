package models

type Products struct {
	Id    int     `json:"id"`
	Name  string  `json:"name" gorm:"column:name;not null"`
	Price float64 `json:"price" gorm:"column:price;not null`
}
