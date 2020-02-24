package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"

	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var decrypted string
var privateKey, publicKey []byte

func init() {
	var err error
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		fmt.Printf("读取公钥错误: %s\n", err.Error())
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		fmt.Printf("读取私钥错误: %s\n", err.Error())
		os.Exit(-1)
	}
}

func main() {
	/*var bits int = 2048
	if err := GenRsaKey(bits); err != nil {
		fmt.Printf("秘钥文件生成失败: %s\n", err.Error())
	}
	fmt.Printf("秘钥文件生成成功: %s\n", err.Error())
	*/

	initData := []byte("abcdefghijklmnopq")
	/*data, err := RsaEncrypt(initData)
	if err != nil {
		panic(err)
	}
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))*/

	data, err := RsaSign(initData)
	if err != nil {
		panic(err)
	}
	err = RsaVerySign(initData, data)
	if err != nil {
		panic(err)
	}

}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, origData)
}

// 签名
func RsaSign(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	sha := sha1.New()
	_, err = sha.Write(origData)
	if err != nil {
		return nil, err
	}
	hashed := sha.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hashed[:])
}

// 验签
func RsaVerySign(origData, signData []byte) error {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	sha := sha1.New()
	_, err = sha.Write(origData)
	if err != nil {
		return err
	}
	hashed := sha.Sum(nil)

	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed, signData)
}
