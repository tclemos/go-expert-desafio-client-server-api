package dto

import "time"

type CotacaoDTO struct {
	Date time.Time `json:"date"`
	Bid  string    `json:"bid"`
}
