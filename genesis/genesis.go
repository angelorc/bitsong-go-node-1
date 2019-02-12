package genesis

import (
	"encoding/base64"
	"encoding/json"
	"github.com/BitSongOfficial/bitsong-go-node/core/types"
	"github.com/BitSongOfficial/bitsong-go-node/helpers"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/types"
	"math/big"
	"time"
)

var (
	Network     = "bitsong-test-network-2"
	genesisTime = time.Date(2019, 2, 12, 10, 0, 0, 0, time.UTC)
)

func GetTestnetGenesis() (*tmtypes.GenesisDoc, error) {
	validatorsPubKeys := []string{
		"Sp0p+r4oUW74aXwX1N/YMu6ZDxEy5b/dCAsu7rwih6w=",
		"ReGPdo/CS/uhLRXiU5E6rtwUK8ttQLxhMEPEWUsaodw=",
		"GR5IvNYZXppTsOnfG/oGOOzrBYmkgkzmd/u0Uw5BVig=",
		"Yr7cZQzJC9dq2fJ5zRgUmtsUqmNxwVFcOpSZ9o3qF+g=",
	}
	validators := make([]tmtypes.GenesisValidator, len(validatorsPubKeys))

	for i, val := range validatorsPubKeys {
		validatorPubKeyBytes, _ := base64.StdEncoding.DecodeString(val)
		var validatorPubKey ed25519.PubKeyEd25519
		copy(validatorPubKey[:], validatorPubKeyBytes)

		validators[i] = tmtypes.GenesisValidator{
			PubKey: validatorPubKey,
			Power:  int64(100000000 / len(validatorsPubKeys)),
		}
	}

	appHash := [32]byte{}

	appState := AppState{
		FirstValidatorAddress: types.HexToAddress("Mx0fc2abbbac27d85d59f57b4ca9b1f001bdedb1a9"),
		InitialBalances: []Account{
			{
				Address: types.HexToAddress("Mx0fc2abbbac27d85d59f57b4ca9b1f001bdedb1a9"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx8ec0c88d843318686bcaa1a7f98e0a5f5339dee0"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx39ad3ab9c9cc23b5d3bb850cbcb5f0b03ffea2bf"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx9305fb4644220e515528ce62c9bcf55d0b5c8b4c"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx448fb95a04a20496c1658c159aa3d25344809627"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx6afaad5211a4ccfc06d6e33f9da43b6f5149fa53"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
		},
	}

	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return nil, err
	}

	genesis := tmtypes.GenesisDoc{
		ChainID:         Network,
		GenesisTime:     genesisTime,
		ConsensusParams: nil,
		Validators:      validators,
		AppHash:         appHash[:],
		AppState:        json.RawMessage(appStateJSON),
	}

	err = genesis.ValidateAndComplete()
	if err != nil {
		return nil, err
	}

	return &genesis, nil
}

func GetPreTestnetGenesis() (*tmtypes.GenesisDoc, error) {
	validatorsPubKeys := []string{
		"Sp0p+r4oUW74aXwX1N/YMu6ZDxEy5b/dCAsu7rwih6w=",
		"ReGPdo/CS/uhLRXiU5E6rtwUK8ttQLxhMEPEWUsaodw=",
		"GR5IvNYZXppTsOnfG/oGOOzrBYmkgkzmd/u0Uw5BVig=",
		"Yr7cZQzJC9dq2fJ5zRgUmtsUqmNxwVFcOpSZ9o3qF+g=",
	}
	validators := make([]tmtypes.GenesisValidator, len(validatorsPubKeys))

	for i, val := range validatorsPubKeys {
		validatorPubKeyBytes, _ := base64.StdEncoding.DecodeString(val)
		var validatorPubKey ed25519.PubKeyEd25519
		copy(validatorPubKey[:], validatorPubKeyBytes)

		validators[i] = tmtypes.GenesisValidator{
			PubKey: validatorPubKey,
			Power:  int64(100000000 / len(validatorsPubKeys)),
		}
	}

	appHash := [32]byte{}

	appState := AppState{
		FirstValidatorAddress: types.HexToAddress("Mx0fc2abbbac27d85d59f57b4ca9b1f001bdedb1a9"),
		InitialBalances: []Account{
			{
				Address: types.HexToAddress("Mx0fc2abbbac27d85d59f57b4ca9b1f001bdedb1a9"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx8ec0c88d843318686bcaa1a7f98e0a5f5339dee0"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx39ad3ab9c9cc23b5d3bb850cbcb5f0b03ffea2bf"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx9305fb4644220e515528ce62c9bcf55d0b5c8b4c"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx448fb95a04a20496c1658c159aa3d25344809627"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx6afaad5211a4ccfc06d6e33f9da43b6f5149fa53"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
		},
	}

	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return nil, err
	}

	genesis := tmtypes.GenesisDoc{
		ChainID:         Network + "-pre",
		GenesisTime:     genesisTime,
		ConsensusParams: nil,
		Validators:      validators,
		AppHash:         appHash[:],
		AppState:        json.RawMessage(appStateJSON),
	}

	err = genesis.ValidateAndComplete()
	if err != nil {
		return nil, err
	}

	return &genesis, nil
}

type AppState struct {
	FirstValidatorAddress types.Address `json:"first_validator_address"`
	InitialBalances       []Account     `json:"initial_balances"`
}

type Account struct {
	Address types.Address     `json:"address"`
	Balance map[string]string `json:"balance"`
}
