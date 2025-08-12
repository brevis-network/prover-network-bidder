package config

type Config struct {
	ChainID          uint64 `mapstructure:"submit_chain_id"`
	ChainRpc         string `mapstructure:"submit_chain_rpc"`
	BrevisMarketAddr string `mapstructure:"brevis_market_addr"`

	BidderEthAddr    string `mapstructure:"bidder_eth_addr"`
	BidderKeystore   string `mapstructure:"bidder_keystore"`
	BidderPassphrase string `mapstructure:"bidder_passphrase"`
}
