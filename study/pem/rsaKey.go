package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenRsaKey(bits int) error {
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Printf("私钥生成失败: %s\n", err.Error())
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	//pkey := pem.EncodeToMemory(block)
	//fmt.Printf("%s\n", string(pkey))

	file, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("创建私钥文件失败: %s\n", err.Error())
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		fmt.Printf("写入私钥失败: %s\n", err.Error())
		return err
	}

	//生成公钥
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Printf("公钥生成失败: %s\n", err.Error())
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	//pkey = pem.EncodeToMemory(block)
	//fmt.Printf("%s\n", string(pkey))

	file, err = os.Create("public.pem")
	if err != nil {
		fmt.Printf("创建公钥文件失败: %s\n", err.Error())
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		fmt.Printf("写入公钥文件失败: %s\n", err.Error())
		return err
	}
	return nil
}

func main() {
	GenRsaKey(1024)
}
