package games_postgres_repository

import core_postgres_pool "github.com/Sklame132/rep/internal/core/repository/postgres/pool"

type GamesRepository struct {
	pool core_postgres_pool.Pool
}

func NewGamesRepository(pool core_postgres_pool.Pool) *GamesRepository {
	return &GamesRepository{
		pool: pool,
	}
}