package helpers

import (
	"strings"

	"github.com/openlyinc/pointy"
	"github.com/pkg/errors"
	"gitlab.com/thorchain/midgard/internal/models"

	api "gitlab.com/thorchain/midgard/api/rest/v1/codegen"
	"gitlab.com/thorchain/midgard/internal/common"
)

func ConvertAssetForAPI(asset common.Asset) *api.Asset {
	assetString := api.Asset(asset.String())
	return &assetString
}

func ConvertEventDataForAPI(events models.Events) *api.Event {
	return &api.Event{
		Fee:        pointy.Int64(int64(events.Fee)),
		Slip:       pointy.Float64(events.Slip),
		StakeUnits: pointy.Int64(int64(events.StakeUnits)),
	}
}

func ConvertGasForAPI(gas models.TxGas) *api.Gas {
	if gas.Amount == 0 {
		return nil
	}

	a, _ := common.NewAsset(gas.Asset.Symbol.String())
	asset := ConvertAssetForAPI(a)

	return &api.Gas{
		Amount: pointy.Int64(int64(gas.Amount)),
		Asset:  asset,
	}
}

func ConvertCoinForAPI(coin common.Coin) *api.Coin {
	a, _ := common.NewAsset(coin.Asset.Symbol.String())
	asset := ConvertAssetForAPI(a)

	return &api.Coin{
		Amount: pointy.Int64(coin.Amount),
		Asset:  asset,
	}
}

func ConvertCoinsForAPI(coins common.Coins) api.Coins {
	var c []api.Coin
	for _, coin := range coins {
		converted := ConvertCoinForAPI(coin)
		c = append(c, *converted)
	}

	return c
}

func ConvertTxForAPI(tx models.TxData) *api.Tx {
	if tx.Address == "" {
		return nil
	}

	coins := ConvertCoinsForAPI(tx.Coin)
	return &api.Tx{
		Address: pointy.String(tx.Address),
		Coins:   &coins,
		Memo:    pointy.String(tx.Memo),
		TxID:    pointy.String(tx.TxID),
	}
}

func ConvertTxsForAPI(txs []models.TxData) *[]api.Tx {
	apiTxs := make([]api.Tx, 0, len(txs))
	for _, tx := range txs {
		apiTxs = append(apiTxs, *ConvertTxForAPI(tx))
	}

	return &apiTxs
}

func ConvertOptionsForAPI(options models.Options) *api.Option {
	return &api.Option{
		Asymmetry:           pointy.Float64(options.Asymmetry),
		PriceTarget:         pointy.Int64(int64(options.PriceTarget)),
		WithdrawBasisPoints: pointy.Int64(int64(options.WithdrawBasisPoints)),
	}
}

func PrepareTxDataResponseForAPI(txData []models.TxDetails) api.TxDetailedResponse {
	var response api.TxDetailedResponse
	for _, d := range txData {
		txD := api.TxDetails{
			Date:    pointy.Int64(int64(d.Date)),
			Events:  ConvertEventDataForAPI(d.Events),
			Gas:     ConvertGasForAPI(d.Gas),
			Height:  pointy.Int64(int64(d.Height)),
			In:      ConvertTxForAPI(d.In),
			Options: ConvertOptionsForAPI(d.Options),
			Out:     ConvertTxsForAPI(d.Out),
			Pool:    ConvertAssetForAPI(d.Pool),
			Status:  pointy.String(d.Status),
			Type:    pointy.String(d.Type),
		}
		response = append(response, txD)
	}

	return response
}

// ParseAssets parses comma separated assets from string.
func ParseAssets(str string) (asts []common.Asset, err error) {
	parts := strings.Split(str, ",")

	asts = make([]common.Asset, len(parts))
	for i, part := range parts {
		asts[i], err = common.NewAsset(part)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid asset '%s'", part)
		}
	}
	return asts, nil
}
