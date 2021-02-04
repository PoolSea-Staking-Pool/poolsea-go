package network

import (
    "fmt"
    "math/big"
    "sync"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/core/types"

    "github.com/rocket-pool/rocketpool-go/rocketpool"
)


// Get the block number which network prices are current for
func GetPricesBlock(rp *rocketpool.RocketPool, opts *bind.CallOpts) (uint64, error) {
    rocketNetworkPrices, err := getRocketNetworkPrices(rp)
    if err != nil {
        return 0, err
    }
    pricesBlock := new(*big.Int)
    if err := rocketNetworkPrices.Call(opts, pricesBlock, "getPricesBlock"); err != nil {
        return 0, fmt.Errorf("Could not get network prices block: %w", err)
    }
    return (*pricesBlock).Uint64(), nil
}


// Get the current network RPL price in ETH
func GetRPLPrice(rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
    rocketNetworkPrices, err := getRocketNetworkPrices(rp)
    if err != nil {
        return nil, err
    }
    rplPrice := new(*big.Int)
    if err := rocketNetworkPrices.Call(opts, rplPrice, "getRPLPrice"); err != nil {
        return nil, fmt.Errorf("Could not get network RPL price: %w", err)
    }
    return *rplPrice, nil
}


// Submit network prices for an epoch
func SubmitPrices(rp *rocketpool.RocketPool, block uint64, rplPrice *big.Int, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketNetworkPrices, err := getRocketNetworkPrices(rp)
    if err != nil {
        return nil, err
    }
    txReceipt, err := rocketNetworkPrices.Transact(opts, "submitPrices", big.NewInt(int64(block)), rplPrice)
    if err != nil {
        return nil, fmt.Errorf("Could not submit network prices: %w", err)
    }
    return txReceipt, nil
}


// Get contracts
var rocketNetworkPricesLock sync.Mutex
func getRocketNetworkPrices(rp *rocketpool.RocketPool) (*rocketpool.Contract, error) {
    rocketNetworkPricesLock.Lock()
    defer rocketNetworkPricesLock.Unlock()
    return rp.GetContract("rocketNetworkPrices")
}
