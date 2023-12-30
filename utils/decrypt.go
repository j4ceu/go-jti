package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"go-jti/config"
)

func GetAESDecrypted(encryptedString string) ([]byte, error) {
	config := config.LoadEncryptConfig()

	// Decode Base64 encrypted string to bytes
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return nil, nil
	}

	block, err := aes.NewCipher([]byte(config.Key))
	if err != nil {
		return nil, err
	}

	if len(encryptedData) < aes.BlockSize {
		return nil, fmt.Errorf("encrypted data is too short")
	}

	decrypted := make([]byte, len(encryptedData))
	mode := cipher.NewCBCDecrypter(block, []byte(config.IV))
	mode.CryptBlocks(decrypted, encryptedData)

	// Remove PKCS#7 padding
	padLength := int(decrypted[len(decrypted)-1])
	return decrypted[:len(decrypted)-padLength], nil
}
