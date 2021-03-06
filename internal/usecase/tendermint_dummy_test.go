package usecase

import (
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"gitlab.com/thorchain/midgard/pkg/clients/thorchain"
)

var _ thorchain.Tendermint = (*TendermintDummy)(nil)

// TendermintDummy is test purpose implementation of Tendermint.
type TendermintDummy struct{}

func (t *TendermintDummy) BlockchainInfo(_, _ int64) (*coretypes.ResultBlockchainInfo, error) {
	return nil, ErrNotImplemented
}

func (t *TendermintDummy) BlockResults(_ *int64) (*coretypes.ResultBlockResults, error) {
	return nil, ErrNotImplemented
}

func (t *TendermintDummy) Send() ([]interface{}, error) {
	return nil, nil
}
