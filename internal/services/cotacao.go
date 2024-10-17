package services

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
	"github.com/tclemos/go-expert-desafio-client-server-api/pkg/entity"
)

type CotacaoService struct {
	cfg config.DolarProviderConfig
}

func NewCotacaoService(cfg config.DolarProviderConfig) *CotacaoService {
	return &CotacaoService{
		cfg: cfg,
	}
}

func (s *CotacaoService) GetCotacaoDolar() (*entity.Cotacao, error) {
	endpointURL, err := url.Parse(s.cfg.Endpoint)
	checkErr(err)

	readTimeout := time.Duration(s.cfg.ReadTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), readTimeout)
	defer cancel()

	req := &http.Request{
		URL:    endpointURL,
		Header: http.Header{},
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if errors.Is(err, context.DeadlineExceeded) {
		return nil, fmt.Errorf("o tempo para receber a cotação do dólar ultrapassou o tempo limite de %v, tente novamente", readTimeout.String())
	} else {
		checkErr(err)
	}

	body, err := io.ReadAll(res.Body)
	checkErr(err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("o servidor de cotação do dólar falhou ao retornar a cotação atual")
	}

	var cotacaoResponse dto.CotacaoDTO
	err = json.Unmarshal(body, &cotacaoResponse)
	checkErr(err)

	return &entity.Cotacao{
		Date: cotacaoResponse.Date,
		Bid:  cotacaoResponse.Bid,
	}, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
