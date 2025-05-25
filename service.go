package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

// Access the pointer to CatFactService and access or modify its context
func (s *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error) {
	res, err := http.Get(s.url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fact := &CatFact{} // Create new instance and get the address
	if err := json.NewDecoder(res.Body).Decode(fact); err != nil {
		log.Fatal(err)
	}

	return fact, nil
}
