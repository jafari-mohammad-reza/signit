package main

import "signit/pkg/cryptography"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateKeyPair() (*cryptography.RSAKeyPair, error) {
	pair, err := cryptography.GeneratePemKeyPair()
	if err != nil {
		return nil, err
	}
	return pair, nil
}
