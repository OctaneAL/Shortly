package util

const Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Base = 62
const Mod = int64(1e9 + 7)
const FixedLength = 10

func ToBase62(num int64) string {
	if num == 0 {
		return string(Alphabet[0])
	}

	var result []byte
	for num > 0 {
		remainder := num % Base
		result = append(result, Alphabet[remainder])
		num /= Base
	}

	return string(result)
}

func HashAndConvert(input string) string {
	var hash int64 = 0
	for _, char := range input {
		hash = (hash*31 + int64(char)) % Mod
	}

	shortCode := ToBase62(hash)

	if len(shortCode) > FixedLength {
		shortCode = shortCode[:FixedLength]
	}

	return shortCode
}
