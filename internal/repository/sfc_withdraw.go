/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
)

// WithdrawRequests extracts a list of partial withdraw requests for the given address.
func (p *proxy) WithdrawRequests(addr *common.Address, stakerID *hexutil.Big, cursor *string, count int32) (*types.WithdrawRequestList, error) {
	// get all the requests for the given delegator address
	if stakerID == nil {
		// log the action and pull the list for all vals
		p.log.Debugf("loading withdraw requests of %s to any validator", addr.String())
		return p.db.Withdrawals(cursor, count, &bson.D{{"addr", addr.String()}})
	}

	// log the action and pull the list for specific address and val
	p.log.Debugf("loading withdraw requests of %s to #%d", addr.String(), stakerID.ToInt().Uint64())
	return p.db.Withdrawals(cursor, count, &bson.D{{"addr", addr.String()}, {"to", stakerID.String()}})
}