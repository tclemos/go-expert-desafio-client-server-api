package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tclemos/go-expert-desafio-client-server-api/config"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/infra/database/sqlite"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/infra/webserver/handlers"
)

func main() {
	configPath := os.Args[1]
	cfg := config.MustLoadServerConfig(configPath)
	db := sqlite.MustOpenConn(cfg.DB)
	defer db.Close()

	fmt.Printf("DB connected: %v\n", cfg.DB.Path)
	cotacaoDolarRepository := sqlite.NewCotacaoDolarRepository(db)

	mux := http.NewServeMux()

	cotacaoDolarHandler := handlers.NewCotacaoDolarHandler(cfg.DolarProvider, cotacaoDolarRepository)
	mux.Handle("/cotacao", cotacaoDolarHandler)

	cotacoesDolarHandler := handlers.NewCotacoesDolarHandler(cfg.DolarProvider, cotacaoDolarRepository)
	mux.Handle("/cotacoes", cotacoesDolarHandler)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	fmt.Printf("http server listen: http://%v\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}
