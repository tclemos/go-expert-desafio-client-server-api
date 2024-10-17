package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tclemos/go-expert-desafio-client-server-api/config"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/dto"
	"github.com/tclemos/go-expert-desafio-client-server-api/internal/infra/database"
	"github.com/tclemos/go-expert-desafio-client-server-api/pkg/entity"
)

type CotacaoDolarHandler struct {
	db  database.CotacoesDolarRepository
	cfg config.DolarProviderConfig
}

func NewCotacaoDolarHandler(cfg config.DolarProviderConfig, db database.CotacoesDolarRepository) *CotacaoDolarHandler {
	return &CotacaoDolarHandler{
		db:  db,
		cfg: cfg,
	}
}

func (h *CotacaoDolarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpointURL, err := url.Parse(h.cfg.Endpoint)
	if handleError(w, err) {
		return
	}

	readTimeout := time.Duration(h.cfg.ReadTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), readTimeout)
	defer cancel()

	req := &http.Request{
		URL:    endpointURL,
		Header: http.Header{},
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("ERROR: o tempo para receber a cotação do dólar ultrapassou o tempo limite de %v\n", readTimeout.String())
	}
	if handleError(w, err) {
		return
	}

	body, err := io.ReadAll(res.Body)
	if handleError(w, err) {
		return
	}
	defer res.Body.Close()

	var dolarProviderResponse map[string]any
	err = json.Unmarshal(body, &dolarProviderResponse)
	if handleError(w, err) {
		return
	}

	bid := dolarProviderResponse["USDBRL"].(map[string]any)["bid"].(string)

	writeTimeout := time.Duration(h.cfg.WriteTimeout)
	ctx, cancel = context.WithTimeout(r.Context(), writeTimeout)
	defer cancel()

	err = h.db.Add(ctx, entity.Cotacao{Date: time.Now(), Bid: bid})
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("ERROR: o tempo para escrever a cotação do dólar no banco de dados ultrapassou o tempo limite de %v\n", writeTimeout.String())
	}
	if handleError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	cotacaoResponse := dto.CotacaoDTO{Bid: bid}
	cotacaoResponseJson, err := json.Marshal(cotacaoResponse)
	if handleError(w, err) {
		return
	}

	_, err = w.Write(cotacaoResponseJson)
	if handleError(w, err) {
		return
	}
}
