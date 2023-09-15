package repositories

import (
	"fmt"
	"log"

	"example.com/go_chantest/internal/core/domain"
	"github.com/BurntSushi/toml"
)

type memkvs struct {
	Websites []domain.Website
}

func LoadMemKVS(cfg string) *memkvs {
	var websites memkvs
	_, err := toml.Decode(cfg, &websites)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("we found websites %v\n", websites.Websites)
	
	return &websites
}

func (repo *memkvs) GetAll() ([]domain.Website) {
	return repo.Websites
}

