package util



import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb string.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Gera dados aleatórios para as accounts

func RandomOwner() string {
	return RandomString(6)
}

func RandomAmmount() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currecies := ["BRL", "USD", "EUR"]
	n := len(currecies)
	return currecies[rand.Intn(n)]
}