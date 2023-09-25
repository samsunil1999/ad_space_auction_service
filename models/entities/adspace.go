package entities

import (
	"time"
)

type AdSpaces struct {
	ID             int64     `gorm:"column:adspace_id;primaryKey;auto_increment;not null"`
	Uuid           string    `gorm:"column:uuid;not null;unique;type:varchar(28)"`
	Name           string    `gorm:"column:name;not null;type:varchar(50)"`
	Description    string    `gorm:"column:description;not null;type:varchar(150)"`
	BasePrice      float32   `gorm:"column:base_price;not null;type:float"`
	SoldPrice      float32   `gorm:"column:sold_price;default:null;type:float" json:"sold_price,omitempty"`
	BidderId       string    `gorm:"column:bidder_id;default:null;type:varchar(25)" json:"bidder_id,omitempty"`
	AuctionEndTime time.Time `gorm:"column:auction_end_time;default:null"`
	ExpiredAt      time.Time `gorm:"column:expired_at;not null"`
	Status         string    `gorm:"column:status;not null;type:varchar(25)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
