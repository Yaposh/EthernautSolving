package client

import (
	"Ethernaut/internal/store"
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// https://github.com/trufflesuite/ganache/tree/master#options
var defaultGanacheChainID = big.NewInt(1337)

type Client struct {
	EthConn                     *ethclient.Client
	Wallets                     store.Wallets
	WalletAddressBySerialNumber map[int]string
}

func New(rawURL string) (*Client, error) {
	conn, err := ethclient.Dial(rawURL)
	if err != nil {
		return nil, fmt.Errorf("new client: %s", err)
	}

	wallets := store.LoadWallets()
	cli := &Client{
		EthConn:                     conn,
		Wallets:                     wallets,
		WalletAddressBySerialNumber: wallets.WalletsAddressesBySerial(),
	}

	return cli, nil
}

func (c *Client) SendEther(from, to string, value *big.Int, gazLimit uint64) error {
	nonce, err := c.EthConn.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		panic(err)
	}

	toAddressHex := common.HexToAddress(to)

	gazPrice, err := c.EthConn.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Value:    value,
		Nonce:    nonce,
		GasPrice: gazPrice,
		Gas:      gazLimit,
		To:       &toAddressHex,
		Data:     data,
	})

	privateKey, err := crypto.HexToECDSA(c.Wallets.PrivateKeys[from])
	if err != nil {
		return fmt.Errorf("hex to ecdsa private key: %s", err)
	}

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return fmt.Errorf("signed tx: %s", err)
	}

	if err := c.EthConn.SendTransaction(context.Background(), signedTx); err != nil {
		return fmt.Errorf("send tx: %s", err)
	}

	fmt.Printf("Sent %s from %q wei to %q\n", value, from, to)
	return nil
}

func (c *Client) TransactOptions(walletNumber int, value int64, gazLimit uint64) (*bind.TransactOpts, error) {
	address := c.WalletAddressBySerialNumber[walletNumber]
	nonce, err := c.EthConn.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return nil, fmt.Errorf("pending nonce: %s", err)
	}

	pKey := c.Wallets.PrivateKeys[address]
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa private key: %s", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, defaultGanacheChainID)
	if err != nil {
		return nil, fmt.Errorf("new keyed transactor: %s", err)
	}

	gazPrice, err := c.EthConn.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("suggest gaz price: %s", err)
	}

	auth.Value = big.NewInt(value)
	auth.From = common.HexToAddress(address)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = gazLimit
	auth.GasPrice = gazPrice

	return auth, nil
}

func (c *Client) GetWalletBalance(walletNumber int) *big.Float {
	address := common.HexToAddress(c.WalletAddressBySerialNumber[walletNumber])

	result, err := c.EthConn.BalanceAt(context.Background(), address, nil)
	if err != nil {
		panic("get account balance: " + err.Error())
	}

	fromBalance, _ := new(big.Float).SetString(result.String())
	return new(big.Float).Quo(fromBalance, big.NewFloat(math.Pow10(18)))
}

func (c *Client) GetBalance(address string) *big.Float {
	result, err := c.EthConn.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		panic("get account balance: " + err.Error())
	}

	fromBalance, _ := new(big.Float).SetString(result.String())
	return new(big.Float).Quo(fromBalance, big.NewFloat(math.Pow10(18)))
}

func (c *Client) PrintAllBalances() {
	fmt.Println("--------------------------------------")
	for i, address := range c.Wallets.ToSortSlice() {
		result, err := c.EthConn.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			panic("get account balance: " + err.Error())
		}

		fromBalance, _ := new(big.Float).SetString(result.String())
		etherValue := new(big.Float).Quo(fromBalance, big.NewFloat(math.Pow10(18)))

		fmt.Printf("%d) %s (%.4f)\n", i, address, etherValue)
	}
	fmt.Println("--------------------------------------")
}
