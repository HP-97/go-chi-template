package service

import (
	"example.com/go_chantest/internal/core/domain"
	"example.com/go_chantest/internal/core/ports"
)

type service struct {
	websiteRepository ports.WebsiteRepository
}

func New(websiteRepository ports.WebsiteRepository) *service {
	return &service{
		websiteRepository: websiteRepository,
	}
}

func (srv *service) GetAll() []domain.Website {
	return srv.websiteRepository.GetAll()
}
