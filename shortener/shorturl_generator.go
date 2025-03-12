package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// Hash function to hash original input
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// Binary-to-text encoding to produce final output
// * Why Base58?
//   - Remove confusing characters (0, O, I, l)
//   - Remove special characters like +, /, =
//   - ...
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortURL(originalURL string, userID string) string {
	// Generate SHA-256 hash
	urlHashBytes := sha256Of(originalURL + userID)

	// Convert hash bytes to uint64 via big.Int
	// - the 256-bit hash is too large to fit into standard integer types
	// - converting it to big.Int allows handling the full 256-bit value temporarily
	// - calling .Uint64() truncates the big integer to 64 bits (least significant bits)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// Convert uint64 to string then to byte slice, then encode
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	// Return the first 8 characters of the Base58-encoded string
	return finalString[:8]
}
