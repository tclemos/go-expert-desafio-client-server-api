package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tclemos/go-expert-desafio-client-server-api/config"
	"github.com/tclemos/go-expert-desafio-client-server-api/pkg/entity"
)

const fileContentTemplate = "Dólar: %v"

type CotacaoDolarFileStorage struct {
	cfg config.ClientConfig
}

func NewCotacaoDolarFileStorage(cfg config.ClientConfig) CotacaoDolarStorage {
	return &CotacaoDolarFileStorage{
		cfg: cfg,
	}
}

func (s CotacaoDolarFileStorage) Save(cotacao entity.Cotacao) {
	outputPath, err := filepath.Abs(s.cfg.Output)
	if err != nil {
		panic(err)
	}

	os.Remove(outputPath)
	fileContent := fmt.Sprintf(fileContentTemplate, cotacao.Bid)

	err = os.WriteFile(outputPath, []byte(fileContent), 0644)
	if err != nil {
		panic(err)
	}
}
