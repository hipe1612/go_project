package model

type Names struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Age         int    `json:"age"  gorm:"column:age;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (Names) TableName() string {
	return "names"
}

type NamesCreation struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Age         int    `json:"age"  gorm:"column:age;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (NamesCreation) TableName() string {
	return Names{}.TableName()
}

type NamesUpdate struct {
	Name        *string `json:"name" gorm:"column:name;"`
	Age         *int    `json:"age"  gorm:"column:age;"`
	Description *string `json:"description" gorm:"column:description;"`
}

func (NamesUpdate) TableName() string {
	return Names{}.TableName()
}
