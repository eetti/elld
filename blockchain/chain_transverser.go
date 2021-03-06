package blockchain

import (
	"github.com/ellcrys/elld/types"
	"github.com/ellcrys/elld/types/core"
)

// NewChainTraverser creates a new ChainTransverser instance
func (b *Blockchain) NewChainTraverser() *ChainTraverser {
	return &ChainTraverser{
		bChain: b,
	}
}

// ChainTraverseFunc is the function type that
// runs a query against a given chain
type ChainTraverseFunc func(chain types.Chainer) (bool, error)

// ChainTraverser allows a user to run a query function
// against a chain of chains. If no result is found in
// the start or initial chain, the parent chain is passed
// to the query function till we reach a chain with no parent.
type ChainTraverser struct {
	chain  types.Chainer
	bChain *Blockchain
}

// Start sets the start chain
func (t *ChainTraverser) Start(chain types.Chainer) *ChainTraverser {
	t.chain = chain
	return t
}

// Query begins a chain traversal session. If false is returned,
// the chain's parent is searched next and so on, until a chain
// with no parent is encountered. If true is returned, the query
// ends. If error is returned, the query ends with the error from
// qf returned.
func (t *ChainTraverser) Query(qf ChainTraverseFunc) error {
	for {
		found, err := qf(t.chain)
		if err != nil {
			return err
		}
		if found {
			return nil
		}

		// Get the chain info of the current chain, if it has a
		// parent chain, look it up to get the parent chain info,
		// then create new chain instance based on the parent chain
		// and set as the next chain.
		if ci := t.chain.GetInfo(); ci.GetParentChainID() != "" {
			parentChainInfo, err := t.bChain.findChainInfo(ci.GetParentChainID())
			if err != nil {
				if err != core.ErrChainNotFound {
					return err
				}
			}
			if parentChainInfo == nil {
				break
			}
			t.chain = t.bChain.NewChainFromChainInfo(parentChainInfo)
		} else {
			break
		}
	}
	return nil
}
