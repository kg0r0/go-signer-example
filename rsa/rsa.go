package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func signData(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
}

func verifyData(data []byte, signature []byte, publicKey *rsa.PublicKey) error {
	hash := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemKey := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: privKeyBytes,
	})
	fmt.Printf("[*] Private key: \n%s\n", pemKey)

	data := []byte("hello world")
	signature, err := signData(data, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[*] Signature: \n%x\n", signature)

	publicKey := privateKey.PublicKey
	err = verifyData(data, signature, &publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Print("[*] Result: \nsignature is a valid signature of message from the public key.")
}
