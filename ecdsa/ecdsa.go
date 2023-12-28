package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func signData(data []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(data)
	return ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
}

func verifyData(data []byte, publicKey *ecdsa.PublicKey, signature []byte) bool {
	hash := sha256.Sum256(data)
	return ecdsa.VerifyASN1(publicKey, hash[:], signature)
}

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	privKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	pemKey := pem.EncodeToMemory(&pem.Block{
		Type: "EC PRIVATE KEY", Bytes: privKeyBytes,
	})
	fmt.Printf("[*] Private key:\n%s\n", pemKey)

	data := []byte("hello world")
	signature, err := signData(data, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[*] Signature:\n%x\n", signature)

	publicKey := privateKey.PublicKey
	res := verifyData(data, &publicKey, signature)
	fmt.Printf("[*] Result:\n%v\n", res)
}
