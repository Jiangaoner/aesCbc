package aesCbc

import (
	"testing"
)

const (
	CRYPT_KEY  = "1~$c31kjtR^@@c2#9&iy"
)

type testInfo struct {
	key      string //秘钥
	iv       string //iv
	encText  string //加密后文本
	origData string //原文
	encrData []byte //加密后内容
}

func TestAesCbc256(t *testing.T) {
	tests := []*testInfo{
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "704114497615264",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es0414497615272",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es704144222222297615406",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertysdddddduiopzxcvbnmgqo",
			origData: "df041449712619896",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwssaaertyuiopzxcvbnmgqo",
			origData: "fd041449768759744",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ds41449722659772",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ff4144977fe304674",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "vfr414497615388",
		},
		{
			key:      CRYPT_KEY,
			iv:       "1234567890qwertyu",
			origData: "bth14497615418",
		},
		{
			key:      CRYPT_KEY,
			iv:       "123456qo",
			origData: "4rb2auut300790842620672",
		},
	}

	for index, test := range tests {
		aesCipher := NewAesCipher([]byte(test.key), []byte(test.iv))
		encrData := aesCipher.Encrypt([]byte(test.origData))

		aesCipher = NewAesCipher([]byte(test.key), []byte(test.iv))
		origData := aesCipher.Decrypt(encrData)

		if string(origData) != test.origData {
			t.Error(index, " fail")
		}
	}
}
