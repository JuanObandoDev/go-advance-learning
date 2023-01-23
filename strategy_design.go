package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

func NewPasswordProtector(user, passwordName string, hashAlgorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user,
		passwordName,
		hashAlgorithm,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}
type MD5 struct{}

func (SHA) Hash(p *PasswordProtector) {
	h := sha1.New()
	h.Write([]byte(p.passwordName))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing password %s for user %s with SHA1: %s\n", p.passwordName, p.user, sha1Hash)
}

func (MD5) Hash(p *PasswordProtector) {
	h := md5.New()
	h.Write([]byte(p.passwordName))
	md5Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing password %s for user %s with MD5: %s\n", p.passwordName, p.user, md5Hash)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	password := NewPasswordProtector("John", "Facebook", sha)
	password.Hash()
	password.SetHashAlgorithm(md5)
	password.Hash()
}
