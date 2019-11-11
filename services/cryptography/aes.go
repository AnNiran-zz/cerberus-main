package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"io"
)

func createHash(key []byte) string {
	hasher := md5.New()
	hasher.Write(key)
	return hex.EncodeToString(hasher.Sum(nil))
}

func aesEncrypt(filename string, key []byte) ([]byte, error) {

	fileBytes := readImage(filename)

	block, err := aes.NewCipher(key) // encrypt the key

	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(fileBytes)

	cipherText := make([]byte, aes.BlockSize+len(b))

	// obtain the first 16 bytes in a slice
	initializationVector := cipherText[:aes.BlockSize]

	// write the first 16 bytes to fill the initialization vector
	if _, err := io.ReadFull(rand.Reader, initializationVector); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// return encrypted stream
	cfb := cipher.NewCFBEncrypter(block, initializationVector)

	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(b))

	// return cipher text that is going to be saved in ipfs
	return cipherText, nil
}

func aesDecrypt(text, cipherKey []byte) ([]byte, error) {

	// AES cipher -> create
	block, err := aes.NewCipher(cipherKey)

	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, err
	}

	// obtain initialization vector
	initializationVector := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, initializationVector)

	cfb.XORKeyStream(text, text)

	data, err := base64.StdEncoding.DecodeString(string(text))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	return data, nil
}

