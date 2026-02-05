package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

var (
	ErrInvalidKey        = errors.New("encryption key must be 32 bytes")
	ErrInvalidCiphertext = errors.New("ciphertext too short or invalid")

	// defaultKey é a chave padrão usada pelas funções wrapper
	// Deve ser inicializada no main.go com SetDefaultKey()
	defaultKey string
)

// DeriveKey deriva uma chave de 32 bytes a partir de uma string usando SHA-256
func DeriveKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

// Encrypt criptografa dados usando AES-256-GCM
// Retorna os dados criptografados em Base64
func Encrypt(plaintext []byte, keyString string) (string, error) {
	key := DeriveKey(keyString)

	// Criar cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Criar GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Criar nonce aleatório
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Criptografar (nonce é prefixado ao ciphertext)
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// Retornar como Base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt descriptografa dados AES-256-GCM de Base64
func Decrypt(ciphertextBase64 string, keyString string) ([]byte, error) {
	key := DeriveKey(keyString)

	// Decodificar Base64
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return nil, err
	}

	// Criar cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Criar GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Verificar tamanho mínimo
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, ErrInvalidCiphertext
	}

	// Extrair nonce e ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Descriptografar
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// EncryptString é um helper para criptografar strings
func EncryptString(plaintext string, keyString string) (string, error) {
	return Encrypt([]byte(plaintext), keyString)
}

// DecryptString é um helper para descriptografar para string
func DecryptString(ciphertextBase64 string, keyString string) (string, error) {
	plaintext, err := Decrypt(ciphertextBase64, keyString)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// Init inicializa o pacote de criptografia com a chave padrão
// Deve ser chamada no main.go com a ENCRYPTION_KEY do config
func Init(key string) error {
	if key == "" {
		return errors.New("encryption key cannot be empty")
	}
	defaultKey = key
	return nil
}

// SetDefaultKey define a chave de criptografia padrão
// Deve ser chamada no main.go com a ENCRYPTION_KEY do config
func SetDefaultKey(key string) {
	defaultKey = key
}

// GetDefaultKey retorna a chave padrão (para testes)
func GetDefaultKey() string {
	return defaultKey
}

// EncryptWithDefaultKey criptografa usando a chave padrão
// Compatível com código legado que não passa a key
func EncryptWithDefaultKey(plaintext string) (string, error) {
	if defaultKey == "" {
		return "", errors.New("default encryption key not set - call crypto.SetDefaultKey() first")
	}
	return EncryptString(plaintext, defaultKey)
}

// DecryptWithDefaultKey descriptografa usando a chave padrão
// Compatível com código legado que não passa a key
func DecryptWithDefaultKey(ciphertextBase64 string) (string, error) {
	if defaultKey == "" {
		return "", errors.New("default encryption key not set - call crypto.SetDefaultKey() first")
	}
	return DecryptString(ciphertextBase64, defaultKey)
}
