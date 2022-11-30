package repository

import (
	"context"
	"database/sql"
	"golang_rest_api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (repository CategoryRepositoryImpl) Save(ctx context.Context, tx sql.Tx, domain.Category) domain.Category {
	
}
