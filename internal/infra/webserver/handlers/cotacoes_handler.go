package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tclemos/go-expert-desafio-client-server-api/config"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/dto"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/infra/database"
)

type CotacoesDolarHandler struct {
	db database.CotacoesDolarRepository
}

func NewCotacoesDolarHandler(cfg config.DolarProviderConfig, db database.CotacoesDolarRepository) *CotacoesDolarHandler {
	return &CotacoesDolarHandler{
		db: db,
	}
}

func (h *CotacoesDolarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cotacoesResponse := []dto.CotacaoDTO{}

	cotacoes, err := h.db.GetAll(r.Context())
	if handleError(w, err) {
		return
	}

	for _, cotacao := range cotacoes {
		cotacoesResponse = append(cotacoesResponse, dto.CotacaoDTO{
			Date: cotacao.Date,
			Bid:  cotacao.Bid,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	cotacoesResponseJson, err := json.Marshal(cotacoesResponse)
	if handleError(w, err) {
		return
	}

	_, err = w.Write(cotacoesResponseJson)
	if handleError(w, err) {
		return
	}
}
