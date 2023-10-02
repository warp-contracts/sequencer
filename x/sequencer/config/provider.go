package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/privval"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

const (
	CONFIG_DIR                 = "config"
	LAST_ARWEAVE_BLOCK_FILE    = "last_arweave_block.json"
	LAST_SORT_KEYS_FILE        = "last_sort_keys.json"
	VALIDATOR_PRIVATE_KEY_FILE = "priv_validator_key.json"
)

// Configuration provider for the sequencer
type ConfigProvider struct {
	homeDir string
	logger  log.Logger
}

func NewConfigProvider(homeDir string, logger log.Logger) *ConfigProvider {
	return &ConfigProvider{homeDir, logger}
}

func (cp ConfigProvider) LastArweaveBlock() (*types.LastArweaveBlock, error) {
	filePath := cp.configFilePath(LAST_ARWEAVE_BLOCK_FILE)
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var blockInfo types.ArweaveBlockInfo
	err = json.Unmarshal(jsonFile, &blockInfo)
	if err != nil {
		return nil, err
	}

	cp.logger.
		With("height", blockInfo.Height).
		With("hash", blockInfo.Hash).
		With("timestamp", blockInfo.Timestamp).
		With("file", filePath).
		Info("Last Arweave block loaded from the config file")

	return &types.LastArweaveBlock{
		ArweaveBlock:         &blockInfo,
		SequencerBlockHeight: 0,
	}, nil
}

func (cp ConfigProvider) LastSortKeys() []types.LastSortKey {
	filePath := cp.configFilePath(LAST_SORT_KEYS_FILE)
	var keys []types.LastSortKey

	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		cp.logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to retrieve last sort keys from the config file")
		return keys
	}

	err = json.Unmarshal(jsonFile, &keys)
	if err != nil {
		cp.logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to unmarshal last sort keys from the config file")
		return keys
	}

	cp.logger.
		With("number of keys", len(keys)).
		With("file", filePath).
		Info("Last sort keys loaded from the config file")
	return keys
}

func (cp ConfigProvider) ValidatorPrivateKey() crypto.PrivKey {
	filePath := cp.configFilePath(VALIDATOR_PRIVATE_KEY_FILE)
	keyFile := privval.LoadFilePVEmptyState(filePath, "")

	cp.logger.
		With("address", keyFile.Key.Address).
		With("public key", keyFile.Key.PubKey).
		With("file", filePath).
		Info("Validator private key loaded from the config file")
	return keyFile.Key.PrivKey
}

func (cp ConfigProvider) configFilePath(fileName string) string {
	return filepath.Join(cp.homeDir, CONFIG_DIR, fileName)
}
