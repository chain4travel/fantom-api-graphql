package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"time"
)

// Default values of configuration options
const (
	// this defines default application name
	defApplicationName = "Fantom GraphQL API Server (custom)"

	// defSelfAddress is a default address used as a placeholder
	// for actual API server identification.
	// Please make sure to configure your real key for your API server on the wild.
	defSelfAddress    = "0x0000000000000000000000000000000000000000"
	defSelfPrivateKey = ""

	// EmptyAddress defines an empty address
	EmptyAddress = "0x0000000000000000000000000000000000000000"

	// defServerBind holds default API server binding address
	defServerBind = "localhost:16761"

	// default set of timeouts for the server
	defReadTimeout     = 2
	defWriteTimeout    = 15
	defIdleTimeout     = 1
	defHeaderTimeout   = 1
	defResolverTimeout = 30

	// defServerDomain holds default API server domain address
	defServerDomain = "localhost:16761"

	// defLoggingLevel holds default Logging level
	// See `godoc.org/github.com/op/go-logging` for the full format specification
	// See `golang.org/pkg/time/` for time format specification
	defLoggingLevel = "INFO"

	// defLoggingFormat holds default format of the Logger output
	defLoggingFormat = "%{color}%{level:-8s} %{shortpkg}/%{shortfunc}%{color:reset}: %{message}"

	// defOperaUrl holds default opera connection string
	defOperaUrl = "~/.opera/opera.ipc"

	// defMongoUrl holds default MongoDB connection string
	defMongoUrl = "mongodb://localhost:27017"

	// defMongoDatabase holds the default name of the API persistent database
	defMongoDatabase = "fantom"

	// defCacheEvictionTime holds default time for in-memory eviction periods
	defCacheEvictionTime = 15 * time.Minute

	// defCacheMax size represents the default max size of the cache in MB
	defCacheMaxSize = 4096

	// defSolCompilerPath represents the default SOL compiler path
	defSolCompilerPath = "/usr/bin/solc"

	// defApiStateOrigin represents the default origin used for API state syncing
	defApiStateOrigin = "https://localhost"

	// defNetworkInitializerContract is the default address of the NetworkInitializer contract
	defNetworkInitializerContract = "0xd1005eed00000000000000000000000000000000"

	// defNodeDriverContract is the default address of the NodeDriver contract
	defNodeDriverContract = "0xd100a01e00000000000000000000000000000000"

	// defSfcContract is the default address of the SFC contract
	defSfcContract = "0xFC00FACE00000000000000000000000000000000"

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiFMintAddressProvider = "0x730e27f6c52d07b1a6ab39b639b617dc566c91af"

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiUniswapCore = EmptyAddress

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiUniswapRouter = EmptyAddress

	// defTokenLogoFilePath represents the default path to the tokens map file
	defTokenLogoFilePath = "tokens.json"

	// defBlockScanRescanDepth represents the amount of blocks re-scanned on server start
	defBlockScanRescanDepth = 200
)

// default list of API peers
var defApiPeers = []string{"https://localhost:16761/api"}

// defCorsAllowOrigins holds CORS default allowed origins.
var defCorsAllowOrigins = []string{"*"}

// default list of API peers
var defVotingSources = make([]string, 0)

// defERC20Logo defines default no-URL value for ERC20 logo list
var defERC20Logo = map[common.Address]string{
	common.HexToAddress(EmptyAddress): "https://repository.fantom.network/logos/erc20.svg",
}

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	// set simple details
	cfg.SetDefault(keyAppName, defApplicationName)
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyDomainAddress, defServerDomain)
	cfg.SetDefault(keySignatureAddress, defSelfAddress)
	cfg.SetDefault(keySignaturePrivateKey, defSelfPrivateKey)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyOperaUrl, defOperaUrl)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
	cfg.SetDefault(keyMongoDatabase, defMongoDatabase)
	cfg.SetDefault(keySolCompilerPath, defSolCompilerPath)
	cfg.SetDefault(keyApiPeers, defApiPeers)
	cfg.SetDefault(keyApiStateOrigin, defApiStateOrigin)
	cfg.SetDefault(keyErc20TokenMapFilePath, defTokenLogoFilePath)
	cfg.SetDefault(keyErc20Logos, defERC20Logo)

	// in-memory cache
	cfg.SetDefault(keyCacheEvictionTime, defCacheEvictionTime)
	cfg.SetDefault(keyCacheMaxSize, defCacheMaxSize)

	// server timeouts
	cfg.SetDefault(keyTimeoutRead, defReadTimeout)
	cfg.SetDefault(keyTimeoutWrite, defWriteTimeout)
	cfg.SetDefault(keyTimeoutHeader, defHeaderTimeout)
	cfg.SetDefault(keyTimeoutIdle, defIdleTimeout)
	cfg.SetDefault(keyTimeoutResolver, defResolverTimeout)

	// no voting sources by default
	cfg.SetDefault(keyVotingSources, defVotingSources)

	// cors
	cfg.SetDefault(keyCorsAllowOrigins, defCorsAllowOrigins)

	// staking configuration defaults
	cfg.SetDefault(keyStakingNetworkInitializerContract, defNetworkInitializerContract)
	cfg.SetDefault(keyStakingNodeDriverContract, defNodeDriverContract)
	cfg.SetDefault(keyStakingSfcContract, defSfcContract)
	cfg.SetDefault(keyStakingTokenizerContract, EmptyAddress)
	cfg.SetDefault(keyStakingERC20Token, EmptyAddress)

	// DeFi configuration
	cfg.SetDefault(keyDefiFMintAddressProvider, defDefiFMintAddressProvider)
	cfg.SetDefault(keyDefiUniswapCore, defDefiUniswapCore)
	cfg.SetDefault(keyDefiUniswapRouter, defDefiUniswapRouter)
}
