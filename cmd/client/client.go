package main

import (
	"fmt"
	"os"

	"github.com/tclemos/go-expert-desafio-client-server-api/config"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/infra/storage"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/services"
)

func main() {
	configPath := os.Args[1]
	cfg := config.MustLoadClientConfig(configPath)
	cotacaoService := services.NewCotacaoService(cfg.DolarProvider)
	cotacaoDolar, err := cotacaoService.GetCotacaoDolar()
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}
	cotacaoDolarStorage := storage.NewCotacaoDolarFileStorage(cfg)
	cotacaoDolarStorage.Save(*cotacaoDolar)
}
