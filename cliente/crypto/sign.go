package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/sha256"
	"encoding/pem"
	"encoding/base64" 

	utils "cliente/utils"
)

const WHO_SIGN = "crypto sign"

func SignMessage(privKeyString string, message string) string {
	payload := []byte(message)
	hash := sha256.Sum256(payload)
	privateKeyBytes := []byte(privKeyString)
	privateBlock, _ := pem.Decode(privateKeyBytes)

	privateKey, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	utils.FailOnError(err, "Erro ao parsear a chave privada")

	signatureRaw, err := rsa.SignPSS(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash[:], nil)
	utils.FailOnError(err, "Erro ao assinar mensagem")

	signature := base64.StdEncoding.EncodeToString(signatureRaw)

	utils.LogInfo("Assinatura gerada com sucesso", WHO_SIGN)

	return signature
}
