package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Can't test encrypt, iv is different each run, so encrypted value is different each time
// func TestCrypto_Encrypt(t *testing.T) {
// 	fw := NewFramework()
//
// 	encrypted, err := fw.Crypto.Encrypt("Per Hultqvist")
// 	assert.Nil(t, err)<
// 	assert.Equal(t, "17efd7c819007d79c9a1c287163b77ea2549a7255dc913b93c3d3a9e4ca975b7c35829ef", encrypted)
// }

func TestCrypto_Decrypt(t *testing.T) {
	fw := NewFramework()

	decrypted, err := fw.Crypto.Decrypt("17efd7c819007d79c9a1c287163b77ea2549a7255dc913b93c3d3a9e4ca975b7c35829ef")
	assert.Nil(t, err)
	assert.Equal(t, "Per Hultqvist", decrypted)
}
