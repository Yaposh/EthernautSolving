package store

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
)

type Wallets struct {
	Addresses   map[string]Wallet `json:"addresses"`
	PrivateKeys map[string]string `json:"private_keys"`
}

type Wallet struct {
	SecretKey Key `json:"secretKey"`
	PublicKey Key `json:"publicKey"`
	Address   string
}

type Key struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type Account struct {
	Nonce     string `json:"nonce"`
	Balance   string `json:"balance"`
	StateRoot string `json:"stateRoot"`
	CodeHash  string `json:"codeHash"`
}

// LoadWallets return map with Key - address, value - privateKey
func LoadWallets() Wallets {
	file, err := os.Open("internal/store/keys.json")
	if err != nil {
		panic(err)
	}

	var wallets Wallets
	if err := json.NewDecoder(file).Decode(&wallets); err != nil {
		panic(err)
	}

	return wallets
}

func (ws Wallets) WalletsAddressesBySerial() map[int]string {
	sl := ws.ToSortSlice()

	result := make(map[int]string, len(ws.Addresses))
	for i, address := range sl {
		result[i] = address
	}

	return result
}

func (ws Wallets) ToSortSlice() []string {
	result := make([]string, 0, len(ws.Addresses))

	for _, addr := range ws.Addresses {
		result = append(result, addr.Address)
	}

	sort.Slice(result, func(i, j int) bool {
		if strings.Compare(result[i], result[j]) == 1 {
			return false
		}

		return true
	})

	return result
}
