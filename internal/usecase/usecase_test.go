package usecase

import (
	"errors"
	"math"
	"sync"
	"testing"
	"time"

	"gitlab.com/thorchain/midgard/internal/common"
	"gitlab.com/thorchain/midgard/internal/models"
	. "gopkg.in/check.v1"
)

var _ = Suite(&UsecaseSuite{})

type UsecaseSuite struct {
	dummyStore     *StoreDummy
	dummyThorchain *ThorchainDummy
}

func (s *UsecaseSuite) SetUpSuite(c *C) {
	s.dummyStore = &StoreDummy{}
	s.dummyThorchain = &ThorchainDummy{}
}

func Test(t *testing.T) {
	TestingT(t)
}

type TestScanningThorchain struct {
	ThorchainDummy
	chains []common.Chain
	err    error
	mu     sync.Mutex
}

func (t *TestScanningThorchain) GetChains() ([]common.Chain, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.chains, t.err
}

func (t *TestScanningThorchain) setChains(chains []common.Chain) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.chains = chains
}

func (s *UsecaseSuite) TestScanningUpdateScanners(c *C) {
	client := &TestScanningThorchain{
		chains: []common.Chain{
			common.BNBChain,
			common.BTCChain,
		},
	}
	conf := &Config{
		// We don't want to test thorchain.Scanner
		ScanInterval:           time.Hour,
		ScannersUpdateInterval: time.Second,
	}
	uc, err := NewUsecase(client, s.dummyStore, conf)
	c.Assert(err, IsNil)

	err = uc.StartScanner()
	c.Assert(err, IsNil)

	time.Sleep(conf.ScannersUpdateInterval + time.Second)
	uc.multiScanner.mu.Lock()
	for _, chain := range client.chains {
		_, ok := uc.multiScanner.scanners[chain]
		c.Assert(ok, Equals, true)
	}
	uc.multiScanner.mu.Unlock()

	newChains := []common.Chain{
		common.BNBChain,
		common.BTCChain,
		common.ETHChain,
	}
	client.setChains(newChains)

	time.Sleep(conf.ScannersUpdateInterval + time.Second)
	uc.multiScanner.mu.Lock()
	for _, chain := range client.chains {
		_, ok := uc.multiScanner.scanners[chain]
		c.Assert(ok, Equals, true)
	}
	uc.multiScanner.mu.Unlock()
}

func (s *UsecaseSuite) TestScanningRestart(c *C) {
	client := &TestScanningThorchain{
		chains: []common.Chain{},
	}
	conf := &Config{
		// We don't want to test thorchain.Scanner
		ScanInterval:           time.Hour,
		ScannersUpdateInterval: time.Second,
	}
	uc, err := NewUsecase(client, s.dummyStore, conf)
	c.Assert(err, IsNil)

	// Scanner should be able to restart.
	err = uc.StartScanner()
	c.Assert(err, IsNil)
	err = uc.StopScanner()
	c.Assert(err, IsNil)
	err = uc.StartScanner()
	c.Assert(err, IsNil)
	err = uc.StopScanner()
	c.Assert(err, IsNil)
}

func (s *UsecaseSuite) TestScanningFaultTolerant(c *C) {
	client := &TestScanningThorchain{
		chains: []common.Chain{
			common.BNBChain,
		},
		err: errors.New("could not fetch chains"),
	}
	conf := &Config{
		// We don't want to test thorchain.Scanner
		ScanInterval:           time.Hour,
		ScannersUpdateInterval: time.Second,
	}
	uc, err := NewUsecase(client, s.dummyStore, conf)
	c.Assert(err, IsNil)

	// Scanner should be able to restart.
	err = uc.StartScanner()
	c.Assert(err, IsNil)

	// Scanner should not be terminated in case of any error.
	time.Sleep(conf.ScannersUpdateInterval + time.Second)
}

type TestGetTxDetailsStore struct {
	StoreDummy
	address   common.Address
	txID      common.TxID
	asset     common.Asset
	eventType string
	offset    int64
	limit     int64
	txDetails []models.TxDetails
	count     int64
	err       error
}

func (s *TestGetTxDetailsStore) GetTxDetails(address common.Address, txID common.TxID, asset common.Asset, eventType string, offset, limit int64) ([]models.TxDetails, int64, error) {
	s.address = address
	s.txID = txID
	s.asset = asset
	s.eventType = eventType
	s.offset = offset
	s.limit = limit
	return s.txDetails, s.count, s.err
}

func (s *UsecaseSuite) TestGetTxDetails(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetTxDetailsStore{
		txDetails: []models.TxDetails{
			{
				Pool:   common.BNBAsset,
				Type:   "stake",
				Status: "Success",
				In: models.TxData{
					Address: "bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38",
					Coin: []common.Coin{
						{
							Asset:  common.RuneB1AAsset,
							Amount: 100,
						},
						{
							Asset:  common.BNBAsset,
							Amount: 10,
						},
					},
					Memo: "stake:BNB.BNB",
					TxID: "2F624637DE179665BA3322B864DB9F30001FD37B4E0D22A0B6ECE6A5B078DAB4",
				},
				Out:     nil,
				Gas:     models.TxGas{},
				Options: models.Options{},
				Events: models.Events{
					StakeUnits: 100,
					Slip:       0,
					Fee:        0,
				},
				Date:   uint64(time.Now().Unix()),
				Height: 1,
			},
			{
				Pool: common.Asset{
					Chain:  "BNB",
					Symbol: "TOML-4BC",
					Ticker: "TOML",
				},
				Type:   "stake",
				Status: "Success",
				In: models.TxData{
					Address: "bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38",
					Coin: []common.Coin{
						{
							Asset:  common.RuneB1AAsset,
							Amount: 100,
						},
						{
							Asset: common.Asset{
								Chain:  "BNB",
								Symbol: "TOML-4BC",
								Ticker: "TOML",
							},
							Amount: 10,
						},
					},
					Memo: "stake:TOML",
					TxID: "E7A0395D6A013F37606B86FDDF17BB3B358217C2452B3F5C153E9A7D00FDA998",
				},
				Out:     nil,
				Gas:     models.TxGas{},
				Options: models.Options{},
				Events: models.Events{
					StakeUnits: 100,
					Slip:       0,
					Fee:        0,
				},
				Date:   uint64(time.Now().Unix()),
				Height: 2,
			},
		},
		count: 10,
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	address, _ := common.NewAddress("bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38")
	txID, _ := common.NewTxID("E7A0395D6A013F37606B86FDDF17BB3B358217C2452B3F5C153E9A7D00FDA998")
	asset, _ := common.NewAsset("BNB.TOML-4BC")
	eventType := "stake"
	page := models.NewPage(0, 2)
	details, count, err := uc.GetTxDetails(address, txID, asset, eventType, page)
	c.Assert(err, IsNil)
	c.Assert(details, DeepEquals, store.txDetails)
	c.Assert(count, Equals, store.count)
	c.Assert(store.address, Equals, address)
	c.Assert(store.txID, Equals, txID)
	c.Assert(store.asset, Equals, asset)
	c.Assert(store.eventType, Equals, eventType)
	c.Assert(store.offset, Equals, page.Offset)
	c.Assert(store.limit, Equals, page.Limit)

	store = &TestGetTxDetailsStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, _, err = uc.GetTxDetails(address, txID, asset, eventType, page)
	c.Assert(err, NotNil)
}

type TestGetPoolsStore struct {
	StoreDummy
	pools []common.Asset
	err   error
}

func (s *TestGetPoolsStore) GetPools() ([]common.Asset, error) {
	return s.pools, s.err
}

func (s *UsecaseSuite) TestGetPools(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetPoolsStore{
		pools: []common.Asset{
			common.BNBAsset,
			{
				Chain:  "BNB",
				Symbol: "TOML-4BC",
				Ticker: "TOML",
			},
		},
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	pools, err := uc.GetPools()
	c.Assert(err, IsNil)
	c.Assert(pools, DeepEquals, store.pools)

	store = &TestGetPoolsStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetPools()
	c.Assert(err, NotNil)
}

type TestGetAssetDetailsStore struct {
	StoreDummy
	pool        common.Asset
	priceInRune float64
	dateCreated uint64
	err         error
}

func (s *TestGetAssetDetailsStore) GetPool(asset common.Asset) (common.Asset, error) {
	return s.pool, s.err
}

func (s *TestGetAssetDetailsStore) GetPriceInRune(asset common.Asset) (float64, error) {
	return s.priceInRune, s.err
}

func (s *TestGetAssetDetailsStore) GetDateCreated(asset common.Asset) (uint64, error) {
	return s.dateCreated, s.err
}

func (s *UsecaseSuite) TestGetAssetDetails(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetAssetDetailsStore{
		pool: common.Asset{
			Chain:  "BNB",
			Symbol: "TOML-4BC",
			Ticker: "TOML",
		},
		priceInRune: 1.5,
		dateCreated: uint64(time.Now().Unix()),
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	details, err := uc.GetAssetDetails(store.pool)
	c.Assert(err, IsNil)
	c.Assert(details, DeepEquals, &models.AssetDetails{
		PriceInRune: store.priceInRune,
		DateCreated: int64(store.dateCreated),
	})

	store = &TestGetAssetDetailsStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetAssetDetails(store.pool)
	c.Assert(err, NotNil)
}

type TestGetStatsStore struct {
	StoreDummy
	dailyActiveUsers   uint64
	monthlyActiveUsers uint64
	totalUsers         uint64
	dailyTx            uint64
	monthlyTx          uint64
	totalTx            uint64
	totalVolume24hr    uint64
	totalVolume        uint64
	totalStaked        uint64
	totalDepth         uint64
	totalEarned        uint64
	poolCount          uint64
	totalAssetBuys     uint64
	totalAssetSells    uint64
	totalStakeTx       uint64
	totalWithdrawTx    uint64
	err                error
}

func (s *TestGetStatsStore) GetStatsData() (models.StatsData, error) {
	stats := models.StatsData{
		DailyActiveUsers:   s.dailyActiveUsers,
		MonthlyActiveUsers: s.monthlyActiveUsers,
		TotalUsers:         s.totalUsers,
		DailyTx:            s.dailyTx,
		MonthlyTx:          s.monthlyTx,
		TotalTx:            s.totalTx,
		TotalVolume24hr:    s.totalVolume24hr,
		TotalVolume:        s.totalVolume,
		TotalStaked:        s.totalStaked,
		TotalDepth:         s.totalDepth,
		TotalEarned:        s.totalEarned,
		PoolCount:          s.poolCount,
		TotalAssetBuys:     s.totalAssetBuys,
		TotalAssetSells:    s.totalAssetSells,
		TotalStakeTx:       s.totalStakeTx,
		TotalWithdrawTx:    s.totalWithdrawTx,
	}
	return stats, s.err
}

func (s *UsecaseSuite) TestGetStats(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetStatsStore{
		dailyActiveUsers:   2,
		monthlyActiveUsers: 10,
		totalUsers:         20,
		dailyTx:            100,
		monthlyTx:          200,
		totalTx:            500,
		totalVolume24hr:    10000,
		totalVolume:        50000,
		totalStaked:        30000,
		totalDepth:         35000,
		totalEarned:        5000,
		poolCount:          3,
		totalAssetBuys:     50,
		totalAssetSells:    60,
		totalStakeTx:       15,
		totalWithdrawTx:    5,
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	stats, err := uc.GetStats()
	c.Assert(err, IsNil)
	c.Assert(stats, DeepEquals, &models.StatsData{
		DailyActiveUsers:   store.dailyActiveUsers,
		MonthlyActiveUsers: store.monthlyActiveUsers,
		TotalUsers:         store.totalUsers,
		DailyTx:            store.dailyTx,
		MonthlyTx:          store.monthlyTx,
		TotalTx:            store.totalTx,
		TotalVolume24hr:    store.totalVolume24hr,
		TotalVolume:        store.totalVolume,
		TotalStaked:        store.totalStaked,
		TotalDepth:         store.totalDepth,
		TotalEarned:        store.totalEarned,
		PoolCount:          store.poolCount,
		TotalAssetBuys:     store.totalAssetBuys,
		TotalAssetSells:    store.totalAssetSells,
		TotalStakeTx:       store.totalStakeTx,
		TotalWithdrawTx:    store.totalWithdrawTx,
	})

	store = &TestGetStatsStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetStats()
	c.Assert(err, NotNil)
}

type TestGetPoolDetailsStore struct {
	StoreDummy
	status           string
	asset            common.Asset
	assetDepth       uint64
	assetROI         float64
	assetStakedTotal uint64
	buyAssetCount    uint64
	buyFeeAverage    float64
	buyFeesTotal     uint64
	buySlipAverage   float64
	buyTxAverage     float64
	buyVolume        uint64
	poolDepth        uint64
	poolFeeAverage   float64
	poolFeesTotal    uint64
	poolROI          float64
	poolROI12        float64
	poolSlipAverage  float64
	poolStakedTotal  uint64
	poolTxAverage    float64
	poolUnits        uint64
	poolVolume       uint64
	poolVolume24hr   uint64
	price            float64
	runeDepth        uint64
	runeROI          float64
	runeStakedTotal  uint64
	sellAssetCount   uint64
	sellFeeAverage   float64
	sellFeesTotal    uint64
	sellSlipAverage  float64
	sellTxAverage    float64
	sellVolume       uint64
	stakeTxCount     uint64
	stakersCount     uint64
	stakingTxCount   uint64
	swappersCount    uint64
	swappingTxCount  uint64
	withdrawTxCount  uint64
	err              error
}

func (s *TestGetPoolDetailsStore) GetPoolData(asset common.Asset) (models.PoolData, error) {
	data := models.PoolData{
		Status:           s.status,
		Asset:            s.asset,
		AssetDepth:       s.assetDepth,
		AssetROI:         s.assetROI,
		AssetStakedTotal: s.assetStakedTotal,
		BuyAssetCount:    s.buyAssetCount,
		BuyFeeAverage:    s.buyFeeAverage,
		BuyFeesTotal:     s.buyFeesTotal,
		BuySlipAverage:   s.buySlipAverage,
		BuyTxAverage:     s.buyTxAverage,
		BuyVolume:        s.buyVolume,
		PoolDepth:        s.poolDepth,
		PoolFeeAverage:   s.poolFeeAverage,
		PoolFeesTotal:    s.poolFeesTotal,
		PoolROI:          s.poolROI,
		PoolROI12:        s.poolROI12,
		PoolSlipAverage:  s.poolSlipAverage,
		PoolStakedTotal:  s.poolStakedTotal,
		PoolTxAverage:    s.poolTxAverage,
		PoolUnits:        s.poolUnits,
		PoolVolume:       s.poolVolume,
		PoolVolume24hr:   s.poolVolume24hr,
		Price:            s.price,
		RuneDepth:        s.runeDepth,
		RuneROI:          s.runeROI,
		RuneStakedTotal:  s.runeStakedTotal,
		SellAssetCount:   s.sellAssetCount,
		SellFeeAverage:   s.sellFeeAverage,
		SellFeesTotal:    s.sellFeesTotal,
		SellSlipAverage:  s.sellSlipAverage,
		SellTxAverage:    s.sellTxAverage,
		SellVolume:       s.sellVolume,
		StakeTxCount:     s.stakeTxCount,
		StakersCount:     s.stakersCount,
		StakingTxCount:   s.stakingTxCount,
		SwappersCount:    s.swappersCount,
		SwappingTxCount:  s.swappingTxCount,
		WithdrawTxCount:  s.withdrawTxCount,
	}
	return data, s.err
}

func (s *UsecaseSuite) TestGetPoolDetails(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetPoolDetailsStore{
		status: models.Enabled.String(),
		asset: common.Asset{
			Chain:  "BNB",
			Symbol: "TOML-4BC",
			Ticker: "TOML",
		},
		assetDepth:       50000000010,
		assetROI:         0.1791847095714499,
		assetStakedTotal: 50000000010,
		buyAssetCount:    2,
		buyFeeAverage:    3730.5,
		buyFeesTotal:     7461,
		buySlipAverage:   0.12300000339746475,
		buyTxAverage:     0.0000149245672606,
		buyVolume:        140331491,
		poolDepth:        4698999994,
		poolFeeAverage:   0.0000000003961796,
		poolFeesTotal:    14939056,
		poolROI:          1.89970001,
		poolROI12:        1.89970001,
		poolSlipAverage:  0.06151196360588074,
		poolStakedTotal:  4341978343,
		poolTxAverage:    59503608,
		poolUnits:        25025000100,
		poolVolume:       357021653,
		poolVolume24hr:   140331492,
		price:            0.0010000019997999997,
		runeDepth:        2349499997,
		runeROI:          3.80000002,
		runeStakedTotal:  2349500000,
		sellAssetCount:   3,
		sellFeeAverage:   7463556,
		sellFeesTotal:    14927112,
		sellSlipAverage:  0.12302392721176147,
		sellTxAverage:    119007217,
		sellVolume:       357021653,
		stakeTxCount:     1,
		stakersCount:     1,
		stakingTxCount:   1,
		swappersCount:    3,
		swappingTxCount:  3,
		withdrawTxCount:  1,
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	asset, _ := common.NewAsset("BNB.TOML-4BC")
	stats, err := uc.GetPoolDetails(asset)
	c.Assert(err, IsNil)
	c.Assert(stats, DeepEquals, &models.PoolData{
		Status:           store.status,
		Asset:            store.asset,
		AssetDepth:       store.assetDepth,
		AssetROI:         store.assetROI,
		AssetStakedTotal: store.assetStakedTotal,
		BuyAssetCount:    store.buyAssetCount,
		BuyFeeAverage:    store.buyFeeAverage,
		BuyFeesTotal:     store.buyFeesTotal,
		BuySlipAverage:   store.buySlipAverage,
		BuyTxAverage:     store.buyTxAverage,
		BuyVolume:        store.buyVolume,
		PoolDepth:        store.poolDepth,
		PoolFeeAverage:   store.poolFeeAverage,
		PoolFeesTotal:    store.poolFeesTotal,
		PoolROI:          store.poolROI,
		PoolROI12:        store.poolROI12,
		PoolSlipAverage:  store.poolSlipAverage,
		PoolStakedTotal:  store.poolStakedTotal,
		PoolTxAverage:    store.poolTxAverage,
		PoolUnits:        store.poolUnits,
		PoolVolume:       store.poolVolume,
		PoolVolume24hr:   store.poolVolume24hr,
		Price:            store.price,
		RuneDepth:        store.runeDepth,
		RuneROI:          store.runeROI,
		RuneStakedTotal:  store.runeStakedTotal,
		SellAssetCount:   store.sellAssetCount,
		SellFeeAverage:   store.sellFeeAverage,
		SellFeesTotal:    store.sellFeesTotal,
		SellSlipAverage:  store.sellSlipAverage,
		SellTxAverage:    store.sellTxAverage,
		SellVolume:       store.sellVolume,
		StakeTxCount:     store.stakeTxCount,
		StakersCount:     store.stakersCount,
		StakingTxCount:   store.stakingTxCount,
		SwappersCount:    store.swappersCount,
		SwappingTxCount:  store.swappingTxCount,
		WithdrawTxCount:  store.withdrawTxCount,
	})

	store = &TestGetPoolDetailsStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetPoolDetails(asset)
	c.Assert(err, NotNil)
}

type TestGetStakersStore struct {
	StoreDummy
	stakers []common.Address
	err     error
}

func (s *TestGetStakersStore) GetStakerAddresses() ([]common.Address, error) {
	return s.stakers, s.err
}

func (s *UsecaseSuite) TestGetStakers(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetStakersStore{
		stakers: []common.Address{
			common.Address("bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38"),
			common.Address("bnb1llvmhawaxxjchwmfmj8fjzftvwz4jpdhapp5hr"),
			common.Address("bnb1u3xts5zh9zuywdjlfmcph7pzyv4f9t4e95jmdq"),
		},
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	stakers, err := uc.GetStakers()
	c.Assert(err, IsNil)
	c.Assert(stakers, DeepEquals, store.stakers)

	store = &TestGetStakersStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetStakers()
	c.Assert(err, NotNil)
}

type TestGetStakerDetailsStore struct {
	StoreDummy
	pools       []common.Asset
	totalEarned int64
	totalROI    float64
	totalStaked int64
	err         error
}

func (s *TestGetStakerDetailsStore) GetStakerAddressDetails(_ common.Address) (models.StakerAddressDetails, error) {
	details := models.StakerAddressDetails{
		PoolsDetails: s.pools,
		TotalEarned:  s.totalEarned,
		TotalROI:     s.totalROI,
		TotalStaked:  s.totalStaked,
	}
	return details, s.err
}

func (s *UsecaseSuite) TestGetStakerDetails(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetStakerDetailsStore{
		pools: []common.Asset{
			{
				Chain:  "BNB",
				Symbol: "TOML-4BC",
				Ticker: "TOML",
			},
			{
				Chain:  "BNB",
				Symbol: "BNB",
				Ticker: "BNB",
			},
		},
		totalEarned: 20,
		totalROI:    1.002,
		totalStaked: 10000,
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	address, _ := common.NewAddress("bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38")
	stats, err := uc.GetStakerDetails(address)
	c.Assert(err, IsNil)
	c.Assert(stats, DeepEquals, &models.StakerAddressDetails{
		PoolsDetails: store.pools,
		TotalEarned:  store.totalEarned,
		TotalROI:     store.totalROI,
		TotalStaked:  store.totalStaked,
	})

	store = &TestGetStakerDetailsStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetStakerDetails(address)
	c.Assert(err, NotNil)
}

type TestGetStakerAssetDetailsStore struct {
	StoreDummy
	asset           common.Asset
	stakeUnits      uint64
	runeStaked      int64
	assetStaked     int64
	poolStaked      int64
	runeEarned      int64
	assetEarned     int64
	poolEarned      int64
	runeROI         float64
	assetROI        float64
	poolROI         float64
	dateFirstStaked uint64
	err             error
}

func (s *TestGetStakerAssetDetailsStore) GetStakersAddressAndAssetDetails(_ common.Address, _ common.Asset) (models.StakerAddressAndAssetDetails, error) {
	details := models.StakerAddressAndAssetDetails{
		Asset:           s.asset,
		StakeUnits:      s.stakeUnits,
		RuneStaked:      s.runeStaked,
		AssetStaked:     s.assetStaked,
		PoolStaked:      s.poolStaked,
		RuneEarned:      s.runeEarned,
		AssetEarned:     s.assetEarned,
		PoolEarned:      s.poolEarned,
		RuneROI:         s.runeROI,
		AssetROI:        s.assetROI,
		PoolROI:         s.poolROI,
		DateFirstStaked: s.dateFirstStaked,
	}
	return details, s.err
}

func (s *UsecaseSuite) TestGetStakerAssetDetails(c *C) {
	client := &ThorchainDummy{}
	store := &TestGetStakerAssetDetailsStore{
		asset: common.Asset{
			Chain:  "BNB",
			Symbol: "TOML-4BC",
			Ticker: "TOML",
		},
		stakeUnits:      100,
		runeStaked:      10000,
		assetStaked:     20000,
		poolStaked:      15000,
		runeEarned:      200,
		assetEarned:     100,
		poolEarned:      250,
		runeROI:         1.005,
		assetROI:        1.02,
		poolROI:         0.0166666666666667,
		dateFirstStaked: uint64(time.Now().Unix()),
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	asset, _ := common.NewAsset("BNB.TOML-4BC")
	address, _ := common.NewAddress("bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38")
	stats, err := uc.GetStakerAssetDetails(address, asset)
	c.Assert(err, IsNil)
	c.Assert(stats, DeepEquals, &models.StakerAddressAndAssetDetails{
		Asset:           store.asset,
		StakeUnits:      store.stakeUnits,
		RuneStaked:      store.runeStaked,
		AssetStaked:     store.assetStaked,
		PoolStaked:      store.poolStaked,
		RuneEarned:      store.runeEarned,
		AssetEarned:     store.assetEarned,
		PoolEarned:      store.poolEarned,
		RuneROI:         store.runeROI,
		AssetROI:        store.assetROI,
		PoolROI:         store.poolROI,
		DateFirstStaked: store.dateFirstStaked,
	})

	store = &TestGetStakerAssetDetailsStore{
		err: errors.New("could not fetch requested data"),
	}
	uc, err = NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetStakerAssetDetails(address, asset)
	c.Assert(err, NotNil)
}

type TestGetNetworkInfoStore struct {
	StoreDummy
	totalDepth uint64
	err        error
}

func (s *TestGetNetworkInfoStore) GetTotalDepth() (uint64, error) {
	return s.totalDepth, s.err
}

type TestGetNetworkInfoThorchain struct {
	ThorchainDummy
	bondMetrics      models.BondMetrics
	activeBonds      []uint64
	standbyBonds     []uint64
	activeNodeCount  int
	standbyNodeCount int
	totalReserve     uint64
	poolShareFactor  float64
	blockReward      models.BlockRewards
	bondingROI       float64
	stakingROI       float64
	nextChurnHeight  int64
	err              error
}

func (s *TestGetNetworkInfoThorchain) GetNetworkInfo(totalDepth uint64) (models.NetworkInfo, error) {
	info := models.NetworkInfo{
		BondMetrics:      s.bondMetrics,
		ActiveBonds:      s.activeBonds,
		StandbyBonds:     s.standbyBonds,
		TotalStaked:      totalDepth,
		ActiveNodeCount:  s.activeNodeCount,
		StandbyNodeCount: s.standbyNodeCount,
		TotalReserve:     s.totalReserve,
		PoolShareFactor:  s.poolShareFactor,
		BlockReward:      s.blockReward,
		BondingROI:       s.bondingROI,
		StakingROI:       s.stakingROI,
		NextChurnHeight:  s.nextChurnHeight,
	}
	return info, s.err
}

func (s *UsecaseSuite) TestGetNetworkInfo(c *C) {
	client := &TestGetNetworkInfoThorchain{
		bondMetrics: models.BondMetrics{
			AverageActiveBond:  50000000,
			AverageStandbyBond: 4000000,
			MaximumActiveBond:  100000000,
			MaximumStandbyBond: 1000000,
			MedianActiveBond:   10000000,
			MedianStandbyBond:  20000,
			MinimumActiveBond:  2500,
			MinimumStandbyBond: 100,
			TotalActiveBond:    100000000,
			TotalStandbyBond:   50000000,
		},
		activeBonds:      []uint64{2500, 10000},
		standbyBonds:     []uint64{100, 1000},
		activeNodeCount:  2,
		standbyNodeCount: 2,
		totalReserve:     17514301182,
		poolShareFactor:  33150836.453073956,
		blockReward: models.BlockRewards{
			BlockReward: 462.8123726851852,
			BondReward:  -15342616812.533314,
			StakeReward: 15342617275.345686,
		},
		bondingROI:      math.Inf(-1),
		stakingROI:      173904.5062587643,
		nextChurnHeight: 539690,
	}
	store := &TestGetNetworkInfoStore{
		totalDepth: 556448810677,
	}
	uc, err := NewUsecase(client, store, &Config{})
	c.Assert(err, IsNil)

	stats, err := uc.GetNetworkInfo()
	c.Assert(err, IsNil)
	c.Assert(stats, DeepEquals, &models.NetworkInfo{
		BondMetrics:      client.bondMetrics,
		ActiveBonds:      client.activeBonds,
		StandbyBonds:     client.standbyBonds,
		TotalStaked:      store.totalDepth,
		ActiveNodeCount:  client.activeNodeCount,
		StandbyNodeCount: client.standbyNodeCount,
		TotalReserve:     client.totalReserve,
		PoolShareFactor:  client.poolShareFactor,
		BlockReward:      client.blockReward,
		BondingROI:       client.bondingROI,
		StakingROI:       client.stakingROI,
		NextChurnHeight:  client.nextChurnHeight,
	})

	uc, err = NewUsecase(client, s.dummyStore, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetNetworkInfo()
	c.Assert(err, NotNil)

	uc, err = NewUsecase(s.dummyThorchain, store, &Config{})
	c.Assert(err, IsNil)

	_, err = uc.GetNetworkInfo()
	c.Assert(err, NotNil)
}
