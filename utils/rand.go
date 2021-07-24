package utils

import (
	"math/rand"
	"time"
)

const str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ~!@#$%^&*()_+,./<>?;\""

// RandomString 随机字符串
func RandomString(length int) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomDigitString 随机数字字符串
func RandomDigitString(length int) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(10)])
	}
	return string(result)
}

// RandomLowerString 随机小写字符串
func RandomLowerString(length int) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(26)+10])
	}
	return string(result)
}

// RandomUpperString 随机大写字符串
func RandomUpperString(length int) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(26)+36])
	}
	return string(result)
}

// RandomAlphaString 随机字母字符串
func RandomAlphaString(length int) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(52)+10])
	}
	return string(result)
}
