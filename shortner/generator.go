package shortner

import "hash/fnv"

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// encodeBase62 encodes a number to a base62 string
func EncodeBase62(num uint64) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		result = string(base62Chars[num%62]) + result
		num /= 62
	}
	return result
}

// hashURL returns a hash of the URL using FNV-1a
func HashURL(url string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(url))
	return h.Sum64()
}

// shortenURL uses FNV-1a + base62
func ShortenURL(longURL string) string {
	hash := HashURL(longURL)
	return EncodeBase62(hash)
}
