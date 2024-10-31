package crypto

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"crypto/sha256"
	"encoding/pem"
	"encoding/base64" 

	utils "cliente/utils"
)

const WHO_VERIFY = "crypto verify"

func VerifySignature(pubKeyString string, message string, signature string) bool {
	payload := []byte(message)
	hash := sha256.Sum256(payload)
	publicKeyBytes := []byte(pubKeyString)
	publicBlock, _ := pem.Decode(publicKeyBytes)

	if publicBlock == nil || publicBlock.Type != "PUBLIC KEY" {
		utils.LogInfo("Chave pública inválida", WHO_VERIFY)
		return false
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		utils.LogInfo("Erro ao parsear a chave pública", WHO_VERIFY)
		return false
	}

	signatureDecoded, err := base64.StdEncoding.DecodeString(string(signature))
	if err != nil {
		utils.LogInfo("Erro ao decodificar a assinatura", WHO_VERIFY)
		return false
	}

	err = rsa.VerifyPSS(publicKey.(*rsa.PublicKey), crypto.SHA256, hash[:], signatureDecoded, nil)
	if err != nil {
		return false
	}

	utils.LogInfo("Assinatura verificada com sucesso", WHO_VERIFY)
	return true
}