package storage

import "github.com/tclemos/go-expert-desafio-client-server-api/pkg/entity"

type CotacaoDolarStorage interface {
	Save(entity.Cotacao)
}
