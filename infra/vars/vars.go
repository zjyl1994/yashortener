package vars

import "gorm.io/gorm"

var (
	ListenAddr string
	BaseURL    string
	DBPath     string

	AdminUser string
	AdminPass string

	DB              *gorm.DB
	DebugMode       bool
	AnonymousCreate bool
)
