package storageredis

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
)

func (rd *RedisStorage) encrypt(bytes []byte) ([]byte, error) {
	// No key? No encrypt
	if len(rd.Options.AESKey) == 0 {
		return bytes, nil
	}

	c, err := aes.NewCipher(rd.Options.GetAESKeyByte())
	if err != nil {
		return nil, fmt.Errorf("unable to create AES cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, fmt.Errorf("unable to create GCM cipher: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, fmt.Errorf("unable to generate nonce: %v", err)
	}

	return gcm.Seal(nonce, nonce, bytes, nil), nil
}

// EncryptStorageData encrypt storage data, so it won't be plain data
func (rd *RedisStorage) EncryptStorageData(data *StorageData) ([]byte, error) {
	// JSON marshal, then encrypt if key is there
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal: %v", err)
	}

	// Prefix with simple prefix and then encrypt
	bytes = append([]byte(rd.Options.ValuePrefix), bytes...)
	return rd.encrypt(bytes)
}

func (rd *RedisStorage) decrypt(bytes []byte) ([]byte, error) {
	// No key? No decrypt
	if len(rd.Options.AESKey) == 0 {
		return bytes, nil
	}
	if len(bytes) < aes.BlockSize {
		return nil, fmt.Errorf("invalid contents")
	}

	block, err := aes.NewCipher(rd.Options.GetAESKeyByte())
	if err != nil {
		return nil, fmt.Errorf("unable to create AES cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("unable to create GCM cipher: %v", err)
	}

	out, err := gcm.Open(nil, bytes[:gcm.NonceSize()], bytes[gcm.NonceSize():], nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failure: %v", err)
	}

	return out, nil
}

// DecryptStorageData decrypt storage data, so we can read it
func (rd *RedisStorage) DecryptStorageData(bytes []byte) (*StorageData, error) {
	// We have to decrypt if there is an AES key and then JSON unmarshal
	bytes, err := rd.decrypt(bytes)
	if err != nil {
		return nil, err
	}

	// Simple sanity check of the beginning of the byte array just to check
	if len(bytes) < len(rd.Options.ValuePrefix) || string(bytes[:len(rd.Options.ValuePrefix)]) != rd.Options.ValuePrefix {
		return nil, fmt.Errorf("invalid data format")
	}

	// Now just json unmarshal
	data := &StorageData{}
	if err := json.Unmarshal(bytes[len(rd.Options.ValuePrefix):], data); err != nil {
		return nil, fmt.Errorf("unable to unmarshal result: %v", err)
	}
	return data, nil
}
