package gencryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func NewPrivateKey(privateKey *rsa.PrivateKey, pemName ...string) string {
	var name string
	if len(pemName) > 0 {
		name = pemName[0] + ".pem"
	} else {
		name = "private_key.pem"
	}
	// 将私钥转换为 PEM 格式
	if privateKey != nil {
		privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privBlock := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privBytes,
		}
		err := ioutil.WriteFile(name, pem.EncodeToMemory(privBlock), 0600)
		if err != nil {
			panic(err)
		}
	}
	return name
}

func ReadPrivateKey(str string, ciphertext []byte) []byte {
	// 读取私钥
	privateKeyBytes, err := ioutil.ReadFile(str)
	if err != nil {
		panic(err)
	}
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	privateKeyInterface, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	// 使用私钥解密数据
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyInterface, ciphertext)
	if err != nil {
		panic(err)
	}
	return plaintext
}

func NewPublicKey(privateKey *rsa.PrivateKey, pemName ...string) string {
	var name string
	if len(pemName) > 0 {
		name = pemName[0] + ".pem"
	} else {
		name = "public_key.pem"
	}
	// 将公钥转换为 PEM 格式
	if privateKey != nil {
		pubBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
		if err != nil {
			panic(err)
		}
		pubBlock := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubBytes,
		}
		err = ioutil.WriteFile(name, pem.EncodeToMemory(pubBlock), 0644)
		if err != nil {
			panic(err)
		}
	}
	return name
}

func ReadPublicKey(str, data string) []byte {
	// 读取公钥
	publicKeyBytes, err := ioutil.ReadFile(str)
	if err != nil {
		panic(err)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	// 使用公钥加密数据
	message := []byte(data)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

func GenerateKey() *rsa.PrivateKey {
	// 生成 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	return privateKey
}
