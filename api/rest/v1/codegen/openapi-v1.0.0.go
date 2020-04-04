// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AssetDetail defines model for AssetDetail.
type AssetDetail struct {
	Asset       *Asset  `json:"asset,omitempty"`
	DateCreated *int64  `json:"dateCreated,omitempty"`
	PriceRune   *string `json:"priceRune,omitempty"`
}

// BlockRewards defines model for BlockRewards.
type BlockRewards struct {
	BlockReward *string `json:"blockReward,omitempty"`
	BondReward  *string `json:"bondReward,omitempty"`
	StakeReward *string `json:"stakeReward,omitempty"`
}

// BondMetrics defines model for BondMetrics.
type BondMetrics struct {
	AverageActiveBond  *string `json:"averageActiveBond,omitempty"`
	AverageStandbyBond *string `json:"averageStandbyBond,omitempty"`
	MaximumActiveBond  *string `json:"maximumActiveBond,omitempty"`
	MaximumStandbyBond *string `json:"maximumStandbyBond,omitempty"`
	MedianActiveBond   *string `json:"medianActiveBond,omitempty"`
	MedianStandbyBond  *string `json:"medianStandbyBond,omitempty"`
	MinimumActiveBond  *string `json:"minimumActiveBond,omitempty"`
	MinimumStandbyBond *string `json:"minimumStandbyBond,omitempty"`
	TotalActiveBond    *string `json:"totalActiveBond,omitempty"`
	TotalStandbyBond   *string `json:"totalStandbyBond,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Error string `json:"error"`
}

// NetworkInfo defines model for NetworkInfo.
type NetworkInfo struct {
	ActiveBonds      *[]string     `json:"activeBonds,omitempty"`
	ActiveNodeCount  *int          `json:"activeNodeCount,omitempty"`
	BlockRewards     *BlockRewards `json:"blockRewards,omitempty"`
	BondMetrics      *BondMetrics  `json:"bondMetrics,omitempty"`
	BondingROI       *string       `json:"bondingROI,omitempty"`
	NextChurnHeight  *string       `json:"nextChurnHeight,omitempty"`
	PoolShareFactor  *string       `json:"poolShareFactor,omitempty"`
	StakingROI       *string       `json:"stakingROI,omitempty"`
	StandbyBonds     *[]string     `json:"standbyBonds,omitempty"`
	StandbyNodeCount *int          `json:"standbyNodeCount,omitempty"`
	TotalReserve     *string       `json:"totalReserve,omitempty"`
	TotalStaked      *string       `json:"totalStaked,omitempty"`
}

// PoolDetail defines model for PoolDetail.
type PoolDetail struct {
	Asset            *Asset  `json:"asset,omitempty"`
	AssetDepth       *string `json:"assetDepth,omitempty"`
	AssetROI         *string `json:"assetROI,omitempty"`
	AssetStakedTotal *string `json:"assetStakedTotal,omitempty"`
	BuyAssetCount    *string `json:"buyAssetCount,omitempty"`
	BuyFeeAverage    *string `json:"buyFeeAverage,omitempty"`
	BuyFeesTotal     *string `json:"buyFeesTotal,omitempty"`
	BuySlipAverage   *string `json:"buySlipAverage,omitempty"`
	BuyTxAverage     *string `json:"buyTxAverage,omitempty"`
	BuyVolume        *string `json:"buyVolume,omitempty"`
	PoolDepth        *string `json:"poolDepth,omitempty"`
	PoolFeeAverage   *string `json:"poolFeeAverage,omitempty"`
	PoolFeesTotal    *string `json:"poolFeesTotal,omitempty"`
	PoolROI          *string `json:"poolROI,omitempty"`
	PoolROI12        *string `json:"poolROI12,omitempty"`
	PoolSlipAverage  *string `json:"poolSlipAverage,omitempty"`
	PoolStakedTotal  *string `json:"poolStakedTotal,omitempty"`
	PoolTxAverage    *string `json:"poolTxAverage,omitempty"`
	PoolUnits        *string `json:"poolUnits,omitempty"`
	PoolVolume       *string `json:"poolVolume,omitempty"`
	PoolVolume24hr   *string `json:"poolVolume24hr,omitempty"`
	Price            *string `json:"price,omitempty"`
	RuneDepth        *string `json:"runeDepth,omitempty"`
	RuneROI          *string `json:"runeROI,omitempty"`
	RuneStakedTotal  *string `json:"runeStakedTotal,omitempty"`
	SellAssetCount   *string `json:"sellAssetCount,omitempty"`
	SellFeeAverage   *string `json:"sellFeeAverage,omitempty"`
	SellFeesTotal    *string `json:"sellFeesTotal,omitempty"`
	SellSlipAverage  *string `json:"sellSlipAverage,omitempty"`
	SellTxAverage    *string `json:"sellTxAverage,omitempty"`
	SellVolume       *string `json:"sellVolume,omitempty"`
	StakeTxCount     *string `json:"stakeTxCount,omitempty"`
	StakersCount     *string `json:"stakersCount,omitempty"`
	StakingTxCount   *string `json:"stakingTxCount,omitempty"`
	Status           *string `json:"status,omitempty"`
	SwappersCount    *string `json:"swappersCount,omitempty"`
	SwappingTxCount  *string `json:"swappingTxCount,omitempty"`
	WithdrawTxCount  *string `json:"withdrawTxCount,omitempty"`
}

// Stakers defines model for Stakers.
type Stakers string

// StakersAddressData defines model for StakersAddressData.
type StakersAddressData struct {
	PoolsArray  *[]Asset `json:"poolsArray,omitempty"`
	TotalEarned *string  `json:"totalEarned,omitempty"`
	TotalROI    *string  `json:"totalROI,omitempty"`
	TotalStaked *string  `json:"totalStaked,omitempty"`
}

// StakersAssetData defines model for StakersAssetData.
type StakersAssetData struct {
	Asset           *Asset  `json:"asset,omitempty"`
	AssetEarned     *string `json:"assetEarned,omitempty"`
	AssetROI        *string `json:"assetROI,omitempty"`
	AssetStaked     *string `json:"assetStaked,omitempty"`
	DateFirstStaked *int64  `json:"dateFirstStaked,omitempty"`
	PoolEarned      *string `json:"poolEarned,omitempty"`
	PoolROI         *string `json:"poolROI,omitempty"`
	PoolStaked      *string `json:"poolStaked,omitempty"`
	RuneEarned      *string `json:"runeEarned,omitempty"`
	RuneROI         *string `json:"runeROI,omitempty"`
	RuneStaked      *string `json:"runeStaked,omitempty"`
	StakeUnits      *string `json:"stakeUnits,omitempty"`
}

// StatsData defines model for StatsData.
type StatsData struct {
	DailyActiveUsers   *string `json:"dailyActiveUsers,omitempty"`
	DailyTx            *string `json:"dailyTx,omitempty"`
	MonthlyActiveUsers *string `json:"monthlyActiveUsers,omitempty"`
	MonthlyTx          *string `json:"monthlyTx,omitempty"`
	PoolCount          *string `json:"poolCount,omitempty"`
	TotalAssetBuys     *string `json:"totalAssetBuys,omitempty"`
	TotalAssetSells    *string `json:"totalAssetSells,omitempty"`
	TotalDepth         *string `json:"totalDepth,omitempty"`
	TotalEarned        *string `json:"totalEarned,omitempty"`
	TotalStakeTx       *string `json:"totalStakeTx,omitempty"`
	TotalStaked        *string `json:"totalStaked,omitempty"`
	TotalTx            *string `json:"totalTx,omitempty"`
	TotalUsers         *string `json:"totalUsers,omitempty"`
	TotalVolume        *string `json:"totalVolume,omitempty"`
	TotalVolume24hr    *string `json:"totalVolume24hr,omitempty"`
	TotalWithdrawTx    *string `json:"totalWithdrawTx,omitempty"`
}

// ThorchainEndpoint defines model for ThorchainEndpoint.
type ThorchainEndpoint struct {
	Address *string `json:"address,omitempty"`
	Chain   *string `json:"chain,omitempty"`
	PubKey  *string `json:"pub_key,omitempty"`
}

// ThorchainEndpoints defines model for ThorchainEndpoints.
type ThorchainEndpoints struct {
	Current *[]ThorchainEndpoint `json:"current,omitempty"`
}

// TxDetails defines model for TxDetails.
type TxDetails struct {
	Date    *int64  `json:"date,omitempty"`
	Events  *Event  `json:"events,omitempty"`
	Gas     *Gas    `json:"gas,omitempty"`
	Height  *string `json:"height,omitempty"`
	In      *Tx     `json:"in,omitempty"`
	Options *Option `json:"options,omitempty"`
	Out     *[]Tx   `json:"out,omitempty"`
	Pool    *Asset  `json:"pool,omitempty"`
	Status  *string `json:"status,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// Asset defines model for asset.
type Asset string

// Coin defines model for coin.
type Coin struct {
	Amount *string `json:"amount,omitempty"`
	Asset  *Asset  `json:"asset,omitempty"`
}

// Coins defines model for coins.
type Coins []Coin

// Event defines model for event.
type Event struct {
	Fee        *string `json:"fee,omitempty"`
	Slip       *string `json:"slip,omitempty"`
	StakeUnits *string `json:"stakeUnits,omitempty"`
}

// Gas defines model for gas.
type Gas struct {
	Amount *string `json:"amount,omitempty"`
	Asset  *Asset  `json:"asset,omitempty"`
}

// Option defines model for option.
type Option struct {
	Asymmetry           *string `json:"asymmetry,omitempty"`
	PriceTarget         *string `json:"priceTarget,omitempty"`
	WithdrawBasisPoints *string `json:"withdrawBasisPoints,omitempty"`
}

// Tx defines model for tx.
type Tx struct {
	Address *string `json:"address,omitempty"`
	Coins   *Coins  `json:"coins,omitempty"`
	Memo    *string `json:"memo,omitempty"`
	TxID    *string `json:"txID,omitempty"`
}

// AssetsDetailedResponse defines model for AssetsDetailedResponse.
type AssetsDetailedResponse []AssetDetail

// GeneralErrorResponse defines model for GeneralErrorResponse.
type GeneralErrorResponse Error

// NetworkResponse defines model for NetworkResponse.
type NetworkResponse NetworkInfo

// PoolsDetailedResponse defines model for PoolsDetailedResponse.
type PoolsDetailedResponse []PoolDetail

// PoolsResponse defines model for PoolsResponse.
type PoolsResponse []Asset

// StakersAddressDataResponse defines model for StakersAddressDataResponse.
type StakersAddressDataResponse StakersAddressData

// StakersAssetDataResponse defines model for StakersAssetDataResponse.
type StakersAssetDataResponse []StakersAssetData

// StakersResponse defines model for StakersResponse.
type StakersResponse []Stakers

// StatsResponse defines model for StatsResponse.
type StatsResponse StatsData

// ThorchainEndpointsResponse defines model for ThorchainEndpointsResponse.
type ThorchainEndpointsResponse ThorchainEndpoints

// TxsResponse defines model for TxsResponse.
type TxsResponse struct {
	Count *int64       `json:"count,omitempty"`
	Txs   *[]TxDetails `json:"txs,omitempty"`
}

// GetAssetInfoParams defines parameters for GetAssetInfo.
type GetAssetInfoParams struct {

	// One or more comma separated unique asset (CHAIN.SYMBOL)
	Asset string `json:"asset"`
}

// GetPoolsDataParams defines parameters for GetPoolsData.
type GetPoolsDataParams struct {

	// One or more comma separated unique asset (CHAIN.SYMBOL)
	Asset string `json:"asset"`
}

// GetStakersAddressAndAssetDataParams defines parameters for GetStakersAddressAndAssetData.
type GetStakersAddressAndAssetDataParams struct {

	// One or more comma separated unique asset (CHAIN.SYMBOL)
	Asset string `json:"asset"`
}

// GetTxDetailsParams defines parameters for GetTxDetails.
type GetTxDetailsParams struct {

	// Address of sender or recipient of any in/out tx in event
	Address *string `json:"address,omitempty"`

	// ID of any in/out tx in event
	Txid *string `json:"txid,omitempty"`

	// Any asset used in event (CHAIN.SYMBOL)
	Asset *string `json:"asset,omitempty"`

	// Requested type of events
	Type *string `json:"type,omitempty"`

	// pagination offset
	Offset int64 `json:"offset"`

	// pagination limit
	Limit int64 `json:"limit"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Asset Information// (GET /v1/assets)
	GetAssetInfo(ctx echo.Context, params GetAssetInfoParams) error
	// Get Documents// (GET /v1/doc)
	GetDocs(ctx echo.Context) error
	// Get Health// (GET /v1/health)
	GetHealth(ctx echo.Context) error
	// Get Network Data// (GET /v1/network)
	GetNetworkData(ctx echo.Context) error
	// Get Asset Pools// (GET /v1/pools)
	GetPools(ctx echo.Context) error
	// Get Pools Data// (GET /v1/pools/detail)
	GetPoolsData(ctx echo.Context, params GetPoolsDataParams) error
	// Get Stakers// (GET /v1/stakers)
	GetStakersData(ctx echo.Context) error
	// Get Staker Data// (GET /v1/stakers/{address})
	GetStakersAddressData(ctx echo.Context, address string) error
	// Get Staker Pool Data// (GET /v1/stakers/{address}/pools)
	GetStakersAddressAndAssetData(ctx echo.Context, address string, params GetStakersAddressAndAssetDataParams) error
	// Get Global Stats// (GET /v1/stats)
	GetStats(ctx echo.Context) error
	// Get Swagger// (GET /v1/swagger.json)
	GetSwagger(ctx echo.Context) error
	// Get the Proxied Pool Addresses// (GET /v1/thorchain/pool_addresses)
	GetThorchainProxiedEndpoints(ctx echo.Context) error
	// Get details of a tx by address, asset or tx-id// (GET /v1/txs)
	GetTxDetails(ctx echo.Context, params GetTxDetailsParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAssetInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetInfo(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAssetInfoParams
	// ------------- Required query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument asset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetInfo(ctx, params)
	return err
}

// GetDocs converts echo context to params.
func (w *ServerInterfaceWrapper) GetDocs(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDocs(ctx)
	return err
}

// GetHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHealth(ctx)
	return err
}

// GetNetworkData converts echo context to params.
func (w *ServerInterfaceWrapper) GetNetworkData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNetworkData(ctx)
	return err
}

// GetPools converts echo context to params.
func (w *ServerInterfaceWrapper) GetPools(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPools(ctx)
	return err
}

// GetPoolsData converts echo context to params.
func (w *ServerInterfaceWrapper) GetPoolsData(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPoolsDataParams
	// ------------- Required query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument asset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPoolsData(ctx, params)
	return err
}

// GetStakersData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersData(ctx)
	return err
}

// GetStakersAddressData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersAddressData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersAddressData(ctx, address)
	return err
}

// GetStakersAddressAndAssetData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersAddressAndAssetData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetStakersAddressAndAssetDataParams
	// ------------- Required query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument asset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersAddressAndAssetData(ctx, address, params)
	return err
}

// GetStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetStats(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStats(ctx)
	return err
}

// GetSwagger converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwagger(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSwagger(ctx)
	return err
}

// GetThorchainProxiedEndpoints converts echo context to params.
func (w *ServerInterfaceWrapper) GetThorchainProxiedEndpoints(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetThorchainProxiedEndpoints(ctx)
	return err
}

// GetTxDetails converts echo context to params.
func (w *ServerInterfaceWrapper) GetTxDetails(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTxDetailsParams
	// ------------- Optional query parameter "address" -------------
	if paramValue := ctx.QueryParam("address"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "address", ctx.QueryParams(), &params.Address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// ------------- Optional query parameter "txid" -------------
	if paramValue := ctx.QueryParam("txid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "txid", ctx.QueryParams(), &params.Txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	// ------------- Optional query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// ------------- Optional query parameter "type" -------------
	if paramValue := ctx.QueryParam("type"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "type", ctx.QueryParams(), &params.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
	}

	// ------------- Required query parameter "offset" -------------
	if paramValue := ctx.QueryParam("offset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument offset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Required query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument limit is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTxDetails(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v1/assets", wrapper.GetAssetInfo)
	router.GET("/v1/doc", wrapper.GetDocs)
	router.GET("/v1/health", wrapper.GetHealth)
	router.GET("/v1/network", wrapper.GetNetworkData)
	router.GET("/v1/pools", wrapper.GetPools)
	router.GET("/v1/pools/detail", wrapper.GetPoolsData)
	router.GET("/v1/stakers", wrapper.GetStakersData)
	router.GET("/v1/stakers/:address", wrapper.GetStakersAddressData)
	router.GET("/v1/stakers/:address/pools", wrapper.GetStakersAddressAndAssetData)
	router.GET("/v1/stats", wrapper.GetStats)
	router.GET("/v1/swagger.json", wrapper.GetSwagger)
	router.GET("/v1/thorchain/pool_addresses", wrapper.GetThorchainProxiedEndpoints)
	router.GET("/v1/txs", wrapper.GetTxDetails)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

