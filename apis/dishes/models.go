package dishes

import "time"

type Dish struct {
	Id             int64     `gorm:"id" json:"id"`
	DishName       string    `gorm:"dishName" json:"dishName"`
	SaleCnt        int64     `gorm:"saleCnt" json:"saleCnt"`
	PhotoUrl       string    `gorm:"photoUrl" json:"photoUrl"`
	Price          float64   `gorm:"price" json:"price"`
	Discount       float64   `gorm:"discount" json:"discount"`
	Material       string    `gorm:"material" json:"material"`
	Taste          string    `gorm:"taste" json:"taste"`
	Morv           string    `gorm:"morv" json:"morv"`
	Corh           string    `gorm:"corh" json:"corh"`
	Method         string    `gorm:"method" json:"method"`
	SugarContained string    `gorm:"sugarContained" json:"sugarContained"`
	CreatedTime    time.Time `gorm:"createdTime" json:"createdTime"`
}

type DishMerchant struct {
	Id  int64 `gorm:"id" json:"id"`
	Did int64 `gorm:"did" json:"did"`
	Mid int64 `gorm:"mid" json:"mid"`
}

type Merchants struct {
	Id          int64     `gorm:"id" json:"id"`
	Host        string    `gorm:"host" json:"host"`
	Name        string    `gorm:"name" json:"name"`
	OrderCnt    int64     `gorm:"orderCnt" json:"orderCnt"`
	Score       int32     `gorm:"score" json:"score"`
	Status      int32     `gorm:"status" json:"status"`
	StartTime   float32   `gorm:"startTime" json:"startTime"`
	EndTime     float32   `gorm:"endTime" json:"endTime"`
	Address     string    `gorm:"address" json:"address"`
	PhotoUrl    string    `gorm:"photoUrl" json:"PhotoUrl"`
	TotalAmount string    `gorm:"totalAmount" json:"totalAmount"`
	CreateTime  time.Time `gorm:"create_time" json:"create_time"`
	Latitude    float64   `gorm:"latitude" json:"latitude"`
	Longitude   float64   `gorm:"longitude" json:"longitude"`
	Info        string    `gorm:"info" json:"info"`
	Phone       string    `gorm:"phone" json:"phone"`
}
