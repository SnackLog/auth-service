package crypto

import argonhashutils "github.com/LightJack05/argon-hash-utils"

func HashPassword(password string, salt []byte) string {
	return argonhashutils.HashPassword(password, 16, 2, 1, salt, 128).ToString()
}
