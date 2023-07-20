package tokens

import (
	"math/big"

	"github.com/RedDuck-Software/poolsea-go/deposit"
	"github.com/RedDuck-Software/poolsea-go/rocketpool"
	"github.com/RedDuck-Software/poolsea-go/tokens"

	"github.com/RedDuck-Software/poolsea-go/tests/testutils/accounts"
)

// Mint an amount of rETH to an account
func MintRETH(rp *rocketpool.RocketPool, toAccount *accounts.Account, amount *big.Int) error {

	// Get ETH value of amount
	ethValue, err := tokens.GetETHValueOfRETH(rp, amount, nil)
	if err != nil {
		return err
	}

	// Deposit from account to mint rETH
	opts := toAccount.GetTransactor()
	opts.Value = ethValue
	_, err = deposit.Deposit(rp, opts)
	return err

}
