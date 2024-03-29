package framework

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type Crypto struct {
}

// Not so secret key!
var key = getKey()

func getKey() []byte {
	// Open file
	file, err := os.Open("/home/per/.softteam_key")
	if err != nil {
		log.Fatal(err)
	}

	// Read file
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Close file
	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	return b[:len(b)-1]
}

// https://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64

// Encrypt : Encrypts a string
func (c Crypto) Encrypt(plainText string) (string, error) {
	text := []byte(plainText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt : Decrypts an encrypted string
func (c Crypto) Decrypt(encryptedString string) (string, error) {
	text, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}
	return string(data), nil
}
