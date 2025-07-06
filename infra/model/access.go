package model

type Access struct {
	ID        uint `gorm:"primarykey"`
	LinkID    uint
	CreateAt  int64
	UserAgent string
	IP        string
}
