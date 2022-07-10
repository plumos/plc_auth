package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
)

const pk = "2ec325478c3904ee0c199b69c14a27635d4476c7e02ff3b0bbf978fa90075a096bd73336c8668d4ad679a57510f83f3312c28b71d42ca0a3da1e97cbba689dd0794c745cf8277975879a7d63f94e3e406984baaa626ed8a6b0e82a2f2a4263993bd5af97f10728967c292cc9e4793b3d245fe136ad928557ca1eb1391015678d2cd6474a69408e7370371d2110d538a979c99342a5ea2c7ef3b7b20d16b19740f5c1d1e7e858166a4024b593ad15cc764f66e0694ac223733fafc9a83962f03603d745d1c66d5ac4e65644168d637343275786da0abedb1d570c844500698f9481ad44df6de8345a8f822f86e3e5c6ffa21cfc286bc2eb22fb5ed1606c6c74fd"

func RsaEncrypt(str string) []byte {

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		nil,
		[]byte("str"), //需要加密的字符串
		nil)
	if err != nil {
		return nil
	}
	return encryptedBytes
}

func PubKey() rsa.PublicKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	//生成公钥
	publicKey := privateKey.PublicKey
	return publicKey
}

func ReadKey() rsa.PublicKey {
	key, _ := hex.DecodeString(pk)
	publicKey, _ := x509.ParsePKIXPublicKey(key)
	return publicKey.(rsa.PublicKey)
}
