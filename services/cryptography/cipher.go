package cryptography

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

func createCipherKey() ([]byte, []byte, error) {

	key := make([]byte, 32)

	_, err := rand.Read(key)

	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, err
	}

	keyHash, err := bcrypt.GenerateFromPassword(key, 12)

	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, err
	}

	return key, keyHash, nil
}

func SaveCipherKey(cipherKey []byte, path, version string) error {

	cipherKeyPath := filepath.Join(path, "cipher", version)

	err := os.MkdirAll(cipherKeyPath, 0755)

	if err != nil {
		return err
	}

	file, err := os.Create(cipherKeyPath + "/cipher")

	if err != nil {
		return err
	}

	_, err = file.Write(cipherKey)

	if err != nil {
		return err
	}

	err = file.Close()

	if err != nil {
		return err
	}

	return nil
}

func DeleteCipherKeyFile(path, version string) error {

	cipherKeyPath := filepath.Join(path, "cipher", version)

	_, err := os.Stat(path)

	if os.IsNotExist(err) {

		if err != nil {
			return errors.New("Path to cipher key location does not exist")
		}
	}

	_, err = os.Stat(cipherKeyPath)

	if os.IsNotExist(err) {

		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	directory, err := ioutil.ReadDir(cipherKeyPath)

	if err != nil {
		return err
	}

	for _, item := range directory {
		os.RemoveAll(filepath.Join([]string{cipherKeyPath, item.Name()}...))
	}

	os.Remove(cipherKeyPath)

	return nil
}

func ReadCipherKey(cipherPath string) ([]byte, error) {

	filename := cipherPath + "/cipher"

	cipherKey, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return cipherKey, nil
}
