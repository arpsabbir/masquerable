package client

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func encrypt(iv []byte, key []byte, plaintext []byte) []byte {
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}

// MakeRandomField makes the random value that can pass the check at server side
func MakeRandomField(sta *State) []byte {
	h := sha256.New()
	t := int(sta.Now().Unix()) / (12 * 60 * 60)
	h.Write([]byte(fmt.Sprintf("%v", t) + sta.Key))
	goal := h.Sum(nil)[0:16]
	iv := make([]byte, 16)
	io.ReadFull(rand.Reader, iv)
	rest := encrypt(iv, sta.AESKey, goal)
	ret := make([]byte, 32)
	copy(ret, iv)
	copy(ret[16:], rest)
	return ret
}
