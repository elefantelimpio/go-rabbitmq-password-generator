package pwdgenerator

import (
	"crypto/sha256"
	"fmt"
	mathrand "math/rand"

	"encoding/base64"
	"encoding/hex"
)

type pwdService struct{}

//NewPwdService NewPwdService
func NewPwdService() PwdIService {
	return pwdService{}
}

//GetSHA256Hash Get SHA256 password hash
func (srv pwdService) GetSHA256Hash(pwd string) (string, error) {

	i := getRandomUInt32()
	salt := fmt.Sprintf("%X", i)

	pwdHash, err := getPwdSHA256Hash(salt, pwd)
	if err != nil {
		return "", err
	}

	base4Result, err := encodeToBase64(salt + pwdHash)
	if err != nil {
		return "", err
	}

	return base4Result, nil
}

func getRandomUInt32() uint32 {
	return mathrand.Uint32()
}

func getPwdSHA256Hash(salt string, pwd string) (string, error) {
	pwdHexString := hex.EncodeToString([]byte(pwd))

	bytess, err := hex.DecodeString(salt + pwdHexString)
	if err != nil {
		return "", err
	}
	ha := sha256.New()
	ha.Write(bytess)

	firstHash := fmt.Sprintf("%x", ha.Sum(nil))
	return firstHash, nil
}

func encodeToBase64(toEncode string) (string, error) {
	bytess, err := hex.DecodeString(toEncode)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(bytess), nil
}
