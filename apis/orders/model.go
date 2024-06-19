package orders

import "time"

type Order struct {
	Id        int64     `json:"id" gorm:"id"`
	StartTime time.Time `json:"startTime" gorm:"startTime"`
	EndTime   time.Time `json:"endTime" gorm:"endTime"`
	Price     float64   `json:"price" gorm:"price"`
	Location  string    `json:"location" gorm:"location"`
	Latitude  float64   `json:"latitude" gorm:"latitude"`
	Longitude float64   `json:"longitude" gorm:"longitude"`
	Remark    string    `json:"remark" gorm:"remark"`
	Amount    int64     `json:"amount" gorm:"amount"`
	DishName  string    `json:"dishName" gorm:"dishName"`
	Status    int64     `json:"status" gorm:"status"`
	IsSetMeal int       `json:"isSetMeal" gorm:"isSetMeal"`
}
