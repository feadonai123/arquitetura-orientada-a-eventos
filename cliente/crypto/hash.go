package crypto

import (
	"encoding/base64"

	utils "cliente/utils"
)

const WHO_HASH = "crypto hash"

func ToHash(mesage string) string {
	utils.LogInfo("Gerando hash", WHO_HASH)
	return base64.StdEncoding.EncodeToString([]byte(mesage))
}

func FromHash(hash string) string {
	decoded, err := base64.StdEncoding.DecodeString(hash)
	utils.FailOnError(err, "Erro ao decodificar a hash")
	utils.LogInfo("Hash decodificado com sucesso", WHO_HASH)
	return string(decoded)
}
