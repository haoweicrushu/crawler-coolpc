package models

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
}
