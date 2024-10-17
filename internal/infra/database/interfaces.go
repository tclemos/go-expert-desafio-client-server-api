package database

import (
	"context"

	"github.com/tclemos/go-expert-desafio-client-server-api/pkg/entity"
)

type CotacoesDolarRepository interface {
	Add(context.Context, entity.Cotacao) error
	GetAll(context.Context) ([]entity.Cotacao, error)
}
