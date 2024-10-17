package sqlite

import (
	"context"
	"database/sql"

	"github.com/tclemos/go-expert-desafio-client-server-api/internal/infra/database"
	"github.com/tclemos/go-expert-desafio-client-server-api/pkg/entity"
)

type CotacaoDolarRepository struct {
	db *sql.DB
}

func NewCotacaoDolarRepository(db *sql.DB) database.CotacoesDolarRepository {
	return &CotacaoDolarRepository{
		db: db,
	}
}

func (r *CotacaoDolarRepository) Add(ctx context.Context, cotacao entity.Cotacao) error {
	insertStudentSQL := `INSERT INTO cotacoes(date, bid) VALUES (?, ?)`
	statement, err := r.db.Prepare(insertStudentSQL)
	if err != nil {
		return err
	}
	_, err = statement.ExecContext(ctx, cotacao.Date, cotacao.Bid)
	if err != nil {
		return err
	}
	return nil
}

func (r *CotacaoDolarRepository) GetAll(ctx context.Context) ([]entity.Cotacao, error) {
	cotacoes := []entity.Cotacao{}
	row, err := r.db.QueryContext(ctx, "SELECT * FROM cotacoes")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var cotacao entity.Cotacao
		row.Scan(&cotacao.Date, &cotacao.Bid)
		cotacoes = append(cotacoes, cotacao)
	}
	return cotacoes, nil
}
