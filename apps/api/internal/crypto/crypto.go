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

	// HIGH H3 — keyring versionado pra rotação.
	// Mapa version → key. Versão "1" é alias do defaultKey por padrão.
	// Ciphertext criado a partir de agora carrega prefixo "v1:".
	// Ciphertext sem prefixo (legacy) usa defaultKey direto.
	keyring = map[string]string{}
	// currentVersion — versão usada em novos encrypts. Default "1".
	currentVersion = "1"
)

// RegisterKeyVersion adiciona uma versão de chave ao keyring (ex: "2" → newKey).
// Usar em rotação: registra v2, deploy, depois reescreve dados antigos.
func RegisterKeyVersion(version, key string) {
	keyring[version] = key
}

// SetCurrentVersion — define qual versão é usada em novos encrypts.
func SetCurrentVersion(version string) {
	currentVersion = version
}

// keyForVersion retorna a key (ou defaultKey se version="1" e não houver override).
func keyForVersion(version string) (string, bool) {
	if k, ok := keyring[version]; ok {
		return k, true
	}
	if version == "1" && defaultKey != "" {
		return defaultKey, true
	}
	return "", false
}

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

// EncryptWithDefaultKey criptografa usando a chave default (versão atual).
// Saída tem prefixo "v<version>:" — Decrypt detecta e usa a chave certa.
func EncryptWithDefaultKey(plaintext string) (string, error) {
	if defaultKey == "" {
		return "", errors.New("default encryption key not set - call crypto.SetDefaultKey() first")
	}
	key, ok := keyForVersion(currentVersion)
	if !ok {
		return "", errors.New("current encryption key version not registered")
	}
	enc, err := EncryptString(plaintext, key)
	if err != nil {
		return "", err
	}
	return "v" + currentVersion + ":" + enc, nil
}

// DecryptWithDefaultKey — descriptografa lidando com prefixo de versão.
// Sem prefixo: usa defaultKey (legacy). Com prefixo "vN:": usa keyring[N].
func DecryptWithDefaultKey(ciphertextBase64 string) (string, error) {
	if defaultKey == "" {
		return "", errors.New("default encryption key not set - call crypto.SetDefaultKey() first")
	}
	// Detect prefix vN:
	if len(ciphertextBase64) > 3 && ciphertextBase64[0] == 'v' {
		colon := -1
		for i := 1; i < len(ciphertextBase64) && i < 8; i++ {
			if ciphertextBase64[i] == ':' {
				colon = i
				break
			}
		}
		if colon > 1 {
			version := ciphertextBase64[1:colon]
			rest := ciphertextBase64[colon+1:]
			key, ok := keyForVersion(version)
			if !ok {
				return "", errors.New("encryption key version not found: " + version)
			}
			return DecryptString(rest, key)
		}
	}
	// Legacy: sem prefixo, usa default
	return DecryptString(ciphertextBase64, defaultKey)
}

// EncryptWithOAuthKey — usa chave dedicada pra OAuth tokens (H3).
// Chama site usa via wrapper em google_calendar_service. Caller passa a key
// explicitamente (não via global) pra deixar a separação clara.
func EncryptWithOAuthKey(plaintext, oauthKey string) (string, error) {
	if oauthKey == "" {
		return "", errors.New("oauth key is empty")
	}
	return EncryptString(plaintext, oauthKey)
}

// DecryptWithOAuthKey — par do EncryptWithOAuthKey.
func DecryptWithOAuthKey(ciphertextBase64, oauthKey string) (string, error) {
	if oauthKey == "" {
		return "", errors.New("oauth key is empty")
	}
	return DecryptString(ciphertextBase64, oauthKey)
}
