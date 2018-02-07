package models

import (
	"time"
)

// Category 商品種類
type Category struct {
	Name string `json:"name" bson:"name"`
}

// SubCate 品牌
type SubCate struct {
	Name string `json:"name" bson:"name"`
}

// Item 商品
type Item struct {
	Category     `bson:"category"`
	SubCate      `bson:"sub_cate"`
	Name         string   `json:"name" bson:"name"`
	OriPrice     int      `json:"oriPrice" bson:"ori_price"`
	SpecialPrice int      `json:"specialPrice" bson:"special_price"`
	CoolMoney    int      `json:"CoolMoney" bson:"cool_money"`
	Tags         []string `json:"Tags" bson:"tags"`
	Group        string   `json:"Group" bson:"group"`
	Date         time.Time
}

// DailyData 每日資料
type DailyData struct {
	TimeStamp time.Time
	Date      string
	Items     []Item
}
