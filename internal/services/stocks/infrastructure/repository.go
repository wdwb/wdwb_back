package infrastructure

import (
	"github.com/changchanghwang/wdwb_back/internal/services/stocks/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepository interface {
	FindAll(db *gorm.DB) ([]*domain.Stock, error)
	FindByCusips(db *gorm.DB, cusips []string) ([]*domain.Stock, error)
	FindOneOrFail(db *gorm.DB, id uuid.UUID) (*domain.Stock, error)
	Save(db *gorm.DB, stocks []*domain.Stock) error
}
