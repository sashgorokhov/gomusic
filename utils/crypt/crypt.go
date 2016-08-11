package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
	"strings"
)

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func get_key() []byte {
	client_id := strconv.Itoa(CLIENT_ID)
	return []byte(strings.Repeat(client_id, 16/len(client_id)+1)[:16])
}

func EncryptToken(access_token string) (string, error) {
	b, err := encrypt(get_key(), []byte(access_token))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func DecryptToken(encrypted string) (string, error) {
	b, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	d, err := decrypt(get_key(), b)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
