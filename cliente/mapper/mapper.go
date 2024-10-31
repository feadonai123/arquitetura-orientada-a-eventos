package mapper

import (
	"fmt"
	"strings"
	"errors"

	crypto "cliente/crypto"
	utils "cliente/utils"
)

const WHO_VERIFY = "mapper"

func MessageToWord(content string, privateKeys string) string {
	contentHash := crypto.ToHash(content)
	signature := crypto.SignMessage(privateKeys, contentHash)
	utils.LogInfo("Mensagem pronta para envio", WHO_VERIFY)
	return fmt.Sprintf("%s.%s", contentHash, signature)
}

func WordToMessage(word string, publicKeys string) (string, error) {
	parts := strings.Split(word, ".")
	messageHash, signature := parts[0], parts[1]

	if !crypto.VerifySignature(publicKeys, messageHash, signature) {
		return "", errors.New("Assinatura inválida")
	}

	utils.LogInfo("Mensagem recebida e verificada", WHO_VERIFY)
	return crypto.FromHash(messageHash), nil
}