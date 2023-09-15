package ports

import "example.com/go_chantest/internal/core/domain"

type WebsiteRepository interface {
	GetAll() ([]domain.Website)
}

type WebsiteService interface {
	GetAll() ([]domain.Website)
}
