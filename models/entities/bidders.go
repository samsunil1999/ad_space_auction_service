package entities

import (
	"database/sql"
	"time"
)

type Bidders struct {
	ID          int64  `gorm:"column:bidder_id;primaryKey;auto_increment;not null"`
	Uuid        string `gorm:"column:uuid;not null;unique;type:varchar(28)"`
	Name        string `gorm:"column:name;not null;type:varchar(50)"`
	Email       string `gorm:"column:email;not null;type:varchar(50)"`
	PhoneNumber string `gorm:"phone_number;not null;type:varchar(10)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
