package currency

import (
	"time"

	"github.com/otter-trade/coin-exchange-api/currency/coinmarketcap"
)

// Config holds all the information needed for currency related manipulation
type Config struct {
	ForexProviders                AllFXSettings `json:"forexProviders"`
	CryptocurrencyProvider        Provider      `json:"cryptocurrencyProvider"`
	CurrencyPairFormat            *PairFormat   `json:"currencyPairFormat"`
	FiatDisplayCurrency           Code          `json:"fiatDisplayCurrency"`
	CurrencyFileUpdateDuration    time.Duration `json:"currencyFileUpdateDuration"`
	ForeignExchangeUpdateDuration time.Duration `json:"foreignExchangeUpdateDuration"`
}

// Provider defines coinmarketcap tools
type Provider struct {
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled"`
	Verbose     bool   `json:"verbose"`
	APIKey      string `json:"apiKey"`
	AccountPlan string `json:"accountPlan"`
}

// BotOverrides defines a bot overriding factor for quick running currency
// subsystems
type BotOverrides struct {
	Coinmarketcap     bool
	CurrencyConverter bool
	CurrencyLayer     bool
	ExchangeRates     bool
	Fixer             bool
	OpenExchangeRates bool
	ExchangeRateHost  bool
}

// CoinmarketcapSettings refers to settings
type CoinmarketcapSettings coinmarketcap.Settings

// SystemsSettings defines incoming system settings
type SystemsSettings struct {
	Coinmarketcap     coinmarketcap.Settings
	Currencyconverter FXSettings
	Currencylayer     FXSettings
	Fixer             FXSettings
	Openexchangerates FXSettings
}

// AllFXSettings defines all the foreign exchange settings
type AllFXSettings []FXSettings

// FXSettings defines foreign exchange requester settings
type FXSettings struct {
	Name            string `json:"name"`
	Enabled         bool   `json:"enabled"`
	Verbose         bool   `json:"verbose"`
	APIKey          string `json:"apiKey"`
	APIKeyLvl       int    `json:"apiKeyLvl"`
	PrimaryProvider bool   `json:"primaryProvider"`
}

// File defines a full currency file generated by the currency storage
// analysis system
type File struct {
	LastMainUpdate interface{} `json:"lastMainUpdate"`
	Cryptocurrency []*Item     `json:"cryptocurrencies"`
	FiatCurrency   []*Item     `json:"fiatCurrencies"`
	UnsetCurrency  []*Item     `json:"unsetCurrencies"`
	Contracts      []*Item     `json:"contracts"`
	Token          []*Item     `json:"tokens"`
	Stable         []*Item     `json:"stableCurrencies"`
}

// Const here are packaged defined delimiters
const (
	UnderscoreDelimiter   = "_"
	DashDelimiter         = "-"
	ForwardSlashDelimiter = "/"
	ColonDelimiter        = ":"
)

// delimiters is a delimiter list
var delimiters = []string{
	DashDelimiter,
	UnderscoreDelimiter,
	ForwardSlashDelimiter,
	ColonDelimiter,
}
