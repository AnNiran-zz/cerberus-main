package cryptography

import (
	"bufio"
	"fmt"
	"os"
)

func EncryptDocument(filename, rsaPath, rsaFile string) ([]byte, []byte, error) {

	key, _, err := createCipherKey()

	if err != nil {
		return nil, nil, err
	}

	// symmetric encryption
	encryptedData, err := aesEncrypt(filename, key)

	if err != nil {
		return nil, nil, err
	}

	// copy provided rsa pair to specified location
	err = getRsaKeyFromUpload(rsaPath, rsaFile)

	if err != nil {
		return nil, nil, err
	}

	// asymmetric encryption
	cipherKey, err := encryptWithPublicKey(key, rsaPath)

	if err != nil {
		return nil, nil, err
	}

	return []byte(encryptedData), cipherKey, nil
}

func DecryptDocument(data, encryptedCipherKey []byte, rsaPath string) ([]byte, error) {

	// decrypt cipherKey with private key
	cipherKey, err := decryptWitPrivateKey(encryptedCipherKey, rsaPath)

	if err != nil {
		return nil, err
	}

	// what if the rsa key pair is lost after the encryption?
	// symmetric decryption
	dataAsBytes, err := aesDecrypt(data, cipherKey)

	return dataAsBytes, nil
}

func readImage(filename string) []byte {

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fileSize int64 = fileInfo.Size()
	fileBytes := make([]byte, fileSize)

	// read the file into bytes
	buffer := bufio.NewReader(file)

	_, err = buffer.Read(fileBytes)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return fileBytes
}
