package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/GYJoker/jgt/config"
	"github.com/GYJoker/jgt/files"
	"github.com/golang-jwt/jwt"
	"io/fs"
	"os"
	"sync"
)

type Config struct {
	PwdSalt    string `json:"pwd_salt" yaml:"pwd_salt"`
	RsaPubFile string `json:"rsa_pub_file" yaml:"rsa_pub_file"`
	RsaPriFile string `json:"rsa_pri_file" yaml:"rsa_pri_file"`
}

type RSA struct {
	PrivateKey    *rsa.PrivateKey
	PrivateKeyStr string
	PublicKey     *rsa.PublicKey
	PublicKeyStr  string
	config        *config.EncryptConfig
}

var (
	r    *RSA
	once sync.Once
)

func GetRSA() *RSA {
	once.Do(func() {
		r = _initRsaKey(config.GetEncryptConfig())
	})
	return r
}

func _initRsaKey(config *config.EncryptConfig) *RSA {
	myRsa := &RSA{config: config}

	// 判断文件夹是否存在
	files.CheckAndCreatePath(config.RsaPriFile)

	if files.FileExists(config.RsaPriFile) && files.FileExists(config.RsaPubFile) {
		/// readKey
		rsaPrivateKeyByte, _ := os.ReadFile(config.RsaPriFile)
		myRsa.PrivateKeyStr = string(rsaPrivateKeyByte)
		myRsa.PrivateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(rsaPrivateKeyByte)
		rsaPublicKeyByte, _ := os.ReadFile(config.RsaPubFile)
		myRsa.PublicKeyStr = string(rsaPublicKeyByte)
		myRsa.PublicKey, _ = jwt.ParseRSAPublicKeyFromPEM(rsaPublicKeyByte)
		fmt.Println("[crypto] _initServerKey pass!")
		return myRsa
	}
	files.RemoveFile(config.RsaPriFile)
	files.RemoveFile(config.RsaPubFile)
	myRsa.genRsaKey()
	return _initRsaKey(config)
}

func (r *RSA) genRsaKey() {
	var prKey, puKey []byte
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	fmt.Println(err)
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prKey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	fmt.Println(err)
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	puKey = pem.EncodeToMemory(block)

	err = os.WriteFile(r.config.RsaPriFile, prKey, fs.ModePerm)
	fmt.Println(err)
	err = os.WriteFile(r.config.RsaPubFile, puKey, fs.ModePerm)
	fmt.Println(err)
}
