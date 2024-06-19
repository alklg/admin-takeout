package merchants

import (
	"time"
)

type Merchants struct {
	Id              int64     `json:"id"`
	Host            string    `json:"host"` // owner
	Name            string    `json:"name"` // name of store
	OrderCnt        int64     `json:"orderCnt"`
	Score           int64     `json:"score"`
	ScoreReal       float64   // ScoreReal = Score / 100
	Status          int       `json:"status"`
	StartTime       string    `json:"startTime"`
	EndTime         string    `json:"endTime"`
	Address         string    `json:"address"`
	PhotoUrl        string    `json:"photoUrl"`
	DayTotalAmount  float64   `json:"dayTotalAmount"`
	WeekTotalAmount float64   `json:"weekTotalAmount"`
	TotalAmount     float64   `json:"totalAmount"`
	CreateTime      time.Time `json:"createTime"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	Info            string    `json:"info"`
	Phone           string    `json:"phone"`
}

type OrderMerchant struct {
	Id   int       `json:"id"`
	Oid  int       `json:"oid"`
	Mid  int       `json:"mid"`
	Time time.Time `joso:"time"`
}
