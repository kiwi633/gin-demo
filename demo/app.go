package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func encrypt(key, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return hex.EncodeToString(ciphertext), nil
}

func decrypt(key, ciphertext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	decodedCiphertext, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	iv := decodedCiphertext[:aes.BlockSize]
	decodedCiphertext = decodedCiphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(decodedCiphertext))
	cfb.XORKeyStream(plaintext, decodedCiphertext)
	return string(plaintext), nil
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
func SumNumbers[K comparable, V string](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
func main() {
	key128 := "0123456789abcdef"                 // 128-bit key
	key256 := "0123456789abcdef0123456789abcdef" // 256-bit key

	plaintext := "Hello, AES!"

	// 128-bit AES encryption
	ciphertext128, err := encrypt(key128, plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("128-bit AES ciphertext:", ciphertext128)

	decrypted128, err := decrypt(key128, ciphertext128)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("128-bit AES decrypted plaintext:", decrypted128)

	// 256-bit AES encryption
	ciphertext256, err := encrypt(key256, plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("256-bit AES ciphertext:", ciphertext256)

	decrypted256, err := decrypt(key256, ciphertext256)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("256-bit AES decrypted plaintext:", decrypted256)
}
