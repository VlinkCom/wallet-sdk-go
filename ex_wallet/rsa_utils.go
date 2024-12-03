package ex_wallet

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// 初始化RSA密钥对
func InitKey() (string, string, error) {
	// 生成RSA密钥对
	privKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", err
	}

	// 公钥
	pubKeyBytes := x509.MarshalPKCS1PublicKey(&privKey.PublicKey)
	pubKey := base64.StdEncoding.EncodeToString(pubKeyBytes)

	// 私钥
	privKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
	privKeyStr := base64.StdEncoding.EncodeToString(privKeyBytes)

	return pubKey, privKeyStr, nil
}

func Sign(data []byte, priKey string) (string, error) {
	// Decode the Base64-encoded private key
	priKeyBytes, err := base64.StdEncoding.DecodeString(priKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %v", err)
	}

	// Parse the PKCS#8 private key
	priv, err := x509.ParsePKCS8PrivateKey(priKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// Assert that the private key is of type *rsa.PrivateKey
	privateKey, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("private key is not an RSA private key")
	}

	// Create an MD5 hash of the data
	hash := md5.New()
	_, err = hash.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to hash data: %v", err)
	}
	digest := hash.Sum(nil)

	// Sign the hash using the RSA private key
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.MD5, digest)
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// Return the signature in Base64 encoding
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Function to verify the signature using RSA public key (X.509 format)
func Verify(data []byte, sign []byte, pubKey string) (bool, error) {
	// Decode the Base64 encoded public key
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return false, fmt.Errorf("failed to decode public key: %v", err)
	}

	// Parse the X.509 encoded public key
	pub, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Assert that the public key is of type *rsa.PublicKey
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("public key is not an RSA public key")
	}

	// Create a SHA-256 hash of the data
	hash := sha256.New()
	_, err = hash.Write(data)
	if err != nil {
		return false, fmt.Errorf("failed to hash data: %v", err)
	}
	digest := hash.Sum(nil)

	// Verify the signature using the RSA public key
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest, sign)
	if err != nil {
		return false, fmt.Errorf("verification failed: %v", err)
	}

	// Return true if verification succeeds
	return true, nil
}
