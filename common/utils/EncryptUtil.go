package utils

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/rsa"
	"encoding/hex"
	"io"
)

// Md5Encrypt MD5 加密
func Md5Encrypt(str string) string {
	encrypt := md5.New()
	encrypt.Write([]byte(str))
	 return hex.EncodeToString(encrypt.Sum(nil))
}

// RSAGeneratePriKey RSA 生成秘钥
func RSAGeneratePriKey(reader io.Reader,bits int) (*rsa.PrivateKey,error) {
	return rsa.GenerateKey(reader, bits)
}

// RSAGeneratePubKey RSA 生成公钥
func RSAGeneratePubKey(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	return &privateKey.PublicKey
}

// RSADecrypt RSA解密
func RSADecrypt(reader io.Reader,key *rsa.PrivateKey,context []byte) ([]byte,error) {
	return rsa.DecryptPKCS1v15(reader, key, context)
}


// RSAEncrypt RSA加密
func RSAEncrypt(reader io.Reader,key *rsa.PublicKey,context []byte) ([]byte,error) {
	return rsa.EncryptPKCS1v15(reader, key, context)
}

// AESEncrypt AES 加密
func AESEncrypt(context,key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err !=nil{
		panic(err)
	}
	newContext := make([]byte, len(context))
	cipher.Encrypt(context,newContext)
	return newContext
}

// AESDecrypt AES 解密
func AESDecrypt(context,key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err !=nil{
		panic(err)
	}
	newContext := make([]byte, len(context))
	cipher.Decrypt(context,newContext)
	return newContext
}

