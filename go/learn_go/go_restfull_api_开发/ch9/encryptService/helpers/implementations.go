package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// 加密服务实例  微服务接口的实现
// EncryptServiceInstance is the implementation of interface for micro service
type EncryptServiceInstance struct{}

// Implements AES encryption algorithm(Rijndael Algorithm)    执行AES（Rijndael算法的加密算法）
/* Initialization vector for the AES algorithm
   More details visit this link https://en.wikipedia.org/wiki/Advanced_Encryption_Standard */
var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// Encrypt encrypts the string with given key  加密用给定的密钥加密字符串
func (EncryptServiceInstance) Encrypt(_ context.Context, key string, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the encrypted string to original   解密将加密的字符串解密为原来的字符串。
func (EncryptServiceInstance) Decrypt(_ context.Context, key string, text string) (string, error) {
	if key == "" || text == "" {
		return "", errEmpty
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}

var errEmpty = errors.New("Sectt Key or Text should not be empty")
