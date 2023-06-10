package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func isBepaSignatureValid(payload, secret, timestamp, sentSignature string) bool {
	payload = fmt.Sprintf("%s\n%s", timestamp, payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	byteSignature, err := hex.DecodeString(sentSignature)
	if err != nil {
		return false
	}
	return hmac.Equal(h.Sum(nil), byteSignature)
}

// TODO add a function to set required headers for testing
