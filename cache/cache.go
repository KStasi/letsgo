package cache

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// SCache - type for cashe
type SCache struct {
	Item    interface{}
	Expires time.Time
}

// Cache - type for cashe
type Cache map[string]*SCache

// ICache - interface for Cache abstraction
type ICache interface {
	Add(key interface{}) error
	Get(key interface{}) (interface{}, error)
	Set(key interface{}) error
	Delete(key interface{}) (interface{}, error)
}

// Get - request value by key
func (cache Cache) Get(key interface{}) (interface{}, error) {

	hash, err := hashKey(key)

	if err != nil {
		fmt.Printf("fn=get key=%q status=error error=%q", hash, err)
		return nil, err
	}
	value := cache[hash]
	if time.Now().After(value.Expires) {
		err := errors.New("key not exists")
		fmt.Printf("fn=get key=%q status=error error=%q", hash, err)
		return nil, err
	}
	return value, nil
}

// Delete - remove cache value
func (cache Cache) Delete(key interface{}) (interface{}, error) {

	hash, err := hashKey(key)

	if err != nil {
		fmt.Printf("fn=get key=%q status=error error=%q", hash, err)
		return nil, err
	}
	value := cache[hash]
	if time.Now().After(value.Expires) {
		err := errors.New("key not exists")
		fmt.Printf("fn=delete key=%q status=error error=%q", hash, err)
	}
	cache[hash] = &SCache{}
	return value, nil
}

// Add - store the value by the new key
func (cache Cache) Add(key, value interface{}, ttl time.Duration) error {
	hash, err := hashKey(key)

	if err != nil {
		fmt.Printf("fn=add key=%q status=error error=%q", hash, err)
		return err
	}

	prevValue := cache[hash]
	if time.Now().Before(prevValue.Expires) {
		err := errors.New("key exists")
		fmt.Printf("fn=get key=%q status=error error=%q", hash, err)
		return err
	}

	cache[hash] = &SCache{
		Item:    value,
		Expires: time.Now().Add(ttl),
	}
	return nil
}

// Set - store the value by the key even if the key pair exists
func (cache Cache) Set(key, value interface{}, ttl time.Duration) error {

	hash, err := hashKey(key)

	if err != nil {
		fmt.Printf("fn=set key=%q status=error error=%q", hash, err)
		return err
	}

	cache[hash] = &SCache{
		Item:    value,
		Expires: time.Now().Add(ttl),
	}
	return nil
}

func hashKey(key interface{}) (string, error) {
	data, err := json.Marshal(key)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha256.Sum256(data))[0:32], nil
}
