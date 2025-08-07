package crypto

import (
	"crypto/ecdh"
	"encoding/base64"
)

func ParsePublicKeyFromBase64(base64Str string) (*ecdh.PublicKey, error) {
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	curve := ecdh.P521()
	pubKey, err := curve.NewPublicKey(data)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}
