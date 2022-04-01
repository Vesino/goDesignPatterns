package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type HashAlgorithm interface {
	Hash(p *passwordProtector)
}

type passwordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func NewPasswordProtector(user string, passName string, algorithm HashAlgorithm) *passwordProtector {
	return &passwordProtector{
		user:          user,
		passwordName:  passName,
		hashAlgorithm: algorithm,
	}
}

func (p *passwordProtector) setHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *passwordProtector) Protect() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}
type SHA256 struct{}
type MD5 struct{}

func (SHA) Hash(p *passwordProtector) {
	h := sha1.New()
	h.Write([]byte(p.passwordName))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using SHA %s\n", sha256Hash)
}

func (SHA256) Hash(p *passwordProtector) {
	h := sha256.New()
	h.Write([]byte(p.passwordName))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using SHA256 %s\n", sha256Hash)
}

func (MD5) Hash(p *passwordProtector) {
	h := md5.New()
	h.Write([]byte(p.passwordName))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using MD5 %s\n", sha256Hash)
}

func main() {
	sha := SHA{}
	passProtector := NewPasswordProtector("Vesino", "MysuperPasswordPerrona", sha)
	passProtector.Protect()

	md5 := MD5{}
	passProtector2 := NewPasswordProtector("Vesino", "OtherSuperSecretePasswordPerrona", md5)
	passProtector2.Protect()

	sHa256 := SHA256{}
	passProtector3 := NewPasswordProtector("Vesino", "OtherSuperSecretePasswordPerrona", sHa256)
	passProtector3.Protect()
}
