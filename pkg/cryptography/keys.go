package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type RSAKeyPair struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func GenerateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
func GeneratePemKeyPair() (*RSAKeyPair, error) {
	privateKey, err := GenerateKeyPair(2048)
	if err != nil {
		return nil, fmt.Errorf("key generation error: %v", err)
	}

	privatePEM := EncodePrivateKeyToPEM(privateKey)

	publicPEM, err := EncodePublicKeyToPEM(&privateKey.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("public key encoding error: %v", err)
	}

	return &RSAKeyPair{
		PrivateKey: privatePEM,
		PublicKey:  publicPEM,
	}, nil
}
func EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) string {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}
	return string(pem.EncodeToMemory(pemBlock))
}

func EncodePublicKeyToPEM(publicKey *rsa.PublicKey) (string, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	}
	return string(pem.EncodeToMemory(pemBlock)), nil
}
