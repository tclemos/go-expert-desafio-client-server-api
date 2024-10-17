package entity

import "time"

type Cotacao struct {
	Date time.Time
	Bid  string
}
