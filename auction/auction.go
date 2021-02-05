package auction

import (
    "fmt"
    "math/big"
    "sync"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"

    "github.com/rocket-pool/rocketpool-go/rocketpool"
)


// Get the total RPL balance of the auction contract
func GetTotalRPLBalance(rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    totalRplBalance := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, totalRplBalance, "getTotalRPLBalance"); err != nil {
        return nil, fmt.Errorf("Could not get auction contract total RPL balance: %w", err)
    }
    return *totalRplBalance, nil
}


// Get the allotted RPL balance of the auction contract
func GetAllottedRPLBalance(rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    allottedRplBalance := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, allottedRplBalance, "getAllottedRPLBalance"); err != nil {
        return nil, fmt.Errorf("Could not get auction contract allotted RPL balance: %w", err)
    }
    return *allottedRplBalance, nil
}


// Get the remaining RPL balance of the auction contract
func GetRemainingRPLBalance(rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    remainingRplBalance := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, remainingRplBalance, "getRemainingRPLBalance"); err != nil {
        return nil, fmt.Errorf("Could not get auction contract remaining RPL balance: %w", err)
    }
    return *remainingRplBalance, nil
}


// Get the number of lots for auction
func GetLotCount(rp *rocketpool.RocketPool, opts *bind.CallOpts) (uint64, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return 0, err
    }
    lotCount := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotCount, "getLotCount"); err != nil {
        return 0, fmt.Errorf("Could not get lot count: %w", err)
    }
    return (*lotCount).Uint64(), nil
}


// Lot details
func GetLotExists(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (bool, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return false, err
    }
    lotExists := new(bool)
    if err := rocketAuctionManager.Call(opts, lotExists, "getLotExists", big.NewInt(int64(lotIndex))); err != nil {
        return false, fmt.Errorf("Could not get lot %d exists status: %w", lotIndex, err)
    }
    return *lotExists, nil
}
func GetLotStartBlock(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (uint64, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return 0, err
    }
    lotStartBlock := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotStartBlock, "getLotStartBlock", big.NewInt(int64(lotIndex))); err != nil {
        return 0, fmt.Errorf("Could not get lot %d start block: %w", lotIndex, err)
    }
    return (*lotStartBlock).Uint64(), nil
}
func GetLotEndBlock(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (uint64, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return 0, err
    }
    lotEndBlock := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotEndBlock, "getLotEndBlock", big.NewInt(int64(lotIndex))); err != nil {
        return 0, fmt.Errorf("Could not get lot %d end block: %w", lotIndex, err)
    }
    return (*lotEndBlock).Uint64(), nil
}
func GetLotStartPrice(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotStartPrice := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotStartPrice, "getLotStartPrice", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d start price: %w", lotIndex, err)
    }
    return *lotStartPrice, nil
}
func GetLotReservePrice(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotReservePrice := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotReservePrice, "getLotReservePrice", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d reserve price: %w", lotIndex, err)
    }
    return *lotReservePrice, nil
}
func GetLotTotalRPLAmount(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotTotalRplAmount := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotTotalRplAmount, "getLotTotalRPLAmount", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d total RPL amount: %w", lotIndex, err)
    }
    return *lotTotalRplAmount, nil
}
func GetLotTotalBidAmount(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotTotalBidAmount := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotTotalBidAmount, "getLotTotalBidAmount", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d total ETH bid amount: %w", lotIndex, err)
    }
    return *lotTotalBidAmount, nil
}
func GetLotRPLRecovered(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (bool, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return false, err
    }
    lotRplRecovered := new(bool)
    if err := rocketAuctionManager.Call(opts, lotRplRecovered, "getLotRPLRecovered", big.NewInt(int64(lotIndex))); err != nil {
        return false, fmt.Errorf("Could not get lot %d RPL recovered status: %w", lotIndex, err)
    }
    return *lotRplRecovered, nil
}
func GetLotPriceByTotalBids(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotPriceByTotalBids := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotPriceByTotalBids, "getLotPriceByTotalBids", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d price by total bids: %w", lotIndex, err)
    }
    return *lotPriceByTotalBids, nil
}
func GetLotCurrentPrice(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotCurrentPrice := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotCurrentPrice, "getLotCurrentPrice", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d current price: %w", lotIndex, err)
    }
    return *lotCurrentPrice, nil
}
func GetLotClaimedRPLAmount(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotClaimedRplAmount := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotClaimedRplAmount, "getLotClaimedRPLAmount", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d claimed RPL amount: %w", lotIndex, err)
    }
    return *lotClaimedRplAmount, nil
}
func GetLotRemainingRPLAmount(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotRemainingRplAmount := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotRemainingRplAmount, "getLotRemainingRPLAmount", big.NewInt(int64(lotIndex))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d remaining RPL amount: %w", lotIndex, err)
    }
    return *lotRemainingRplAmount, nil
}
func GetLotIsCleared(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.CallOpts) (bool, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return false, err
    }
    lotIsCleared := new(bool)
    if err := rocketAuctionManager.Call(opts, lotIsCleared, "getLotIsCleared", big.NewInt(int64(lotIndex))); err != nil {
        return false, fmt.Errorf("Could not get lot %d cleared status: %w", lotIndex, err)
    }
    return *lotIsCleared, nil
}


// Get the ETH amount bid on a lot by an address
func GetLotAddressBidAmount(rp *rocketpool.RocketPool, lotIndex uint64, bidder common.Address, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lot := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lot, "getLotAddressBidAmount", big.NewInt(int64(lotIndex)), bidder); err != nil {
        return nil, fmt.Errorf("Could not get lot %d address ETH bid amount: %w", lotIndex, err)
    }
    return *lot, nil
}


// Get the price of a lot at a specific block
func GetLotPriceAtBlock(rp *rocketpool.RocketPool, lotIndex, blockNumber uint64, opts *bind.CallOpts) (*big.Int, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    lotPriceAtBlock := new(*big.Int)
    if err := rocketAuctionManager.Call(opts, lotPriceAtBlock, "getLotPriceAtBlock", big.NewInt(int64(lotIndex)), big.NewInt(int64(blockNumber))); err != nil {
        return nil, fmt.Errorf("Could not get lot %d price at block: %w", lotIndex, err)
    }
    return *lotPriceAtBlock, nil
}


// Create a new lot
func CreateLot(rp *rocketpool.RocketPool, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    txReceipt, err := rocketAuctionManager.Transact(opts, "createLot")
    if err != nil {
        return nil, fmt.Errorf("Could not create lot: %w", err)
    }
    return txReceipt, nil
}


// Place a bid on a lot
func PlaceBid(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    txReceipt, err := rocketAuctionManager.Transact(opts, "placeBid", big.NewInt(int64(lotIndex)))
    if err != nil {
        return nil, fmt.Errorf("Could not place bid on lot %d: %w", lotIndex, err)
    }
    return txReceipt, nil
}


// Claim RPL from a lot that was bid on
func ClaimBid(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    txReceipt, err := rocketAuctionManager.Transact(opts, "claimBid", big.NewInt(int64(lotIndex)))
    if err != nil {
        return nil, fmt.Errorf("Could not claim bid from lot %d: %w", lotIndex, err)
    }
    return txReceipt, nil
}


// Recover unclaimed RPL from a lot
func RecoverUnclaimedRPL(rp *rocketpool.RocketPool, lotIndex uint64, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketAuctionManager, err := getRocketAuctionManager(rp)
    if err != nil {
        return nil, err
    }
    txReceipt, err := rocketAuctionManager.Transact(opts, "recoverUnclaimedRPL", big.NewInt(int64(lotIndex)))
    if err != nil {
        return nil, fmt.Errorf("Could not recover unclaimed RPL from lot %d: %w", lotIndex, err)
    }
    return txReceipt, nil
}


// Get contracts
var rocketAuctionManagerLock sync.Mutex
func getRocketAuctionManager(rp *rocketpool.RocketPool) (*rocketpool.Contract, error) {
    rocketAuctionManagerLock.Lock()
    defer rocketAuctionManagerLock.Unlock()
    return rp.GetContract("rocketAuctionManager")
}
