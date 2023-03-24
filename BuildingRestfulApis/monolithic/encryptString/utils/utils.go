package utils

import (
	"crypto/aes"
	"encoding/base64"
)
import "crypto/cipher"

// Implements AES encryption algorithm (Rijndael Algo)
var initVector = []byte{35, 46, 23, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// creates a cipher block a given int key, the we pass the block to a cipher block encrutor func
// This encryptor rakes the block and initialization vector.
// Then XORKeyStream creates a cipher on the cipher block. it fills the ciphertext.
// Then we do a base64 encoding to generate the protected string
func EncryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// here we decode the base64 encoding,
// create a cipher block with the key. and then reverse the process
// XORKeyStream with go from cipher text to plain text.
func DecryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}