package models

// Ayah هو النموذج الخاص بالآيات
type Quran_texts struct {
	ID   int    `gorm:"column:index;primaryKey"`
	Sura int    `gorm:"column:sura"`
	Aya  int    `gorm:"column:aya"`
	Text string `gorm:"column:text"`
}

type User struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}
