package usecase

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngobrut/beli-mang-api/config"
	"github.com/ngobrut/beli-mang-api/infra/aws"
	"github.com/ngobrut/beli-mang-api/internal/repository"
)

type Usecase struct {
	cnf  config.Config
	db   *pgxpool.Pool
	repo repository.IFaceRepository
	aws  aws.IFaceAWS
}

func New(cnf config.Config, db *pgxpool.Pool, repo repository.IFaceRepository, aws aws.IFaceAWS) IFaceUsecase {
	return &Usecase{
		cnf:  cnf,
		db:   db,
		repo: repo,
		aws:  aws,
	}
}
