package model

import "time"

// Tweet はツイートのモデル
type Tweet struct {
	ID        int64
	Text      string `form:"text"`
	CreatedAt time.Time
	UpdatedAT time.Time
}
