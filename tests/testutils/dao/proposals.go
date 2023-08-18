package dao

import (
	"github.com/Seb369888/poolsea-go/dao"
	trustednodedao "github.com/Seb369888/poolsea-go/dao/trustednode"
	"github.com/Seb369888/poolsea-go/rocketpool"
	trustednodesettings "github.com/Seb369888/poolsea-go/settings/trustednode"
	rptypes "github.com/Seb369888/poolsea-go/types"

	"github.com/Seb369888/poolsea-go/tests/testutils/accounts"
	"github.com/Seb369888/poolsea-go/tests/testutils/evm"
)

// Pass and execute a proposal
func PassAndExecuteProposal(rp *rocketpool.RocketPool, proposalId uint64, trustedNodeAccounts []*accounts.Account) error {

	// Get proposal voting delay
	voteDelayTime, err := trustednodesettings.GetProposalVoteDelayTime(rp, nil)
	if err != nil {
		return err
	}

	// Increase time until proposal voting delay has passed
	if err := evm.IncreaseTime(int(voteDelayTime)); err != nil {
		return err
	}

	// Vote on proposal until passed
	for _, account := range trustedNodeAccounts {
		if state, err := dao.GetProposalState(rp, proposalId, nil); err != nil {
			return err
		} else if state == rptypes.Succeeded {
			break
		}
		if _, err := trustednodedao.VoteOnProposal(rp, proposalId, true, account.GetTransactor()); err != nil {
			return err
		}
	}

	// Execute proposal
	if _, err := trustednodedao.ExecuteProposal(rp, proposalId, trustedNodeAccounts[0].GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
