package models

import (
	"time"
)

// Category 商品種類
type Category struct {
	Name string
}

// SubCate 品牌
type SubCate struct {
	Name string
}

// Item 商品
type Item struct {
	Category
	SubCate
	Name         string
	OriPrice     int
	SpecialPrice int
	CoolMoney    int
	Tags         []string
	Group        string
}

// DailyData 每日資料
type DailyData struct {
	TimeStamp time.Time
	Date      string
	Items     []Item
}
