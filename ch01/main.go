package main

import (
	"fmt"
	"sync"
)

var version string
var revision string
var goversion string
var builddate string

// KeyValue type for KeyValue
type KeyValue struct {
	store map[string]string
	mu    sync.RWMutex
}

// NewKeyValue create KeyValue
func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)}
}

// Set set key-value
func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.store[key] = val
}

// Get get value by key
func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.store[key]
	return val, ok
}

func main() {
	kv := NewKeyValue()
	kv.Set("key", "Hello, world")
	value, ok := kv.Get("key")
	if ok {
		fmt.Println(value)
		fmt.Println()
		fmt.Println("version:   " + version)
		fmt.Println("revision:  " + revision)
		fmt.Println("goversion: " + goversion)
		fmt.Println("builddate: " + builddate)
	}

}
