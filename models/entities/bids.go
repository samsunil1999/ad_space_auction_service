package entities

import (
	"database/sql"
	"time"
)

type Bids struct {
	ID        int64    `gorm:"column:bid_id;primaryKey;auto_increment;not null"`
	Uuid      string   `gorm:"column:uuid;not null;unique;type:varchar(28)"`
	AdSpaceId string   `gorm:"column:ad_space_id;not null;type:varchar(28)"`
	Adspace   AdSpaces `gorm:"foreignKey:AdSpaceId;references:Uuid"`
	BidderId  string   `gorm:"column:bidder_id;not null;type:varchar(28)"`
	Bidder    Bidders  `gorm:"foreignKey:BidderId;references:Uuid"`
	BidAmount float32  `gorm:"column:bid_amount;not null;type:float"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
