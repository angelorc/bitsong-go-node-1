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
	Network     = "bitsong-test-network-1"
	genesisTime = time.Date(2019, 2, 4, 12, 0, 0, 0, time.UTC)
)

func GetTestnetGenesis() (*tmtypes.GenesisDoc, error) {
	validatorsPubKeys := []string{
		"SAI2eUjhwLzvmq7pgzdS6MgA2IpXNVdGuiXK+Vzt0gE=",
		//"c42kG6ant9abcpSvoVi4nFobQQy/DCRDyFxf4krR3Rw=",
		//"bxbB/yGm+5RqrtD0wfzKJyty/ZBJiPkdOIMoK4rjG6I=",
		//"nhPy9UaN14KzFkRPvWZZXhPbp9e9Pvob7NULQgRfWMY=",
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
		FirstValidatorAddress: types.HexToAddress("Mx448fb95a04a20496c1658c159aa3d25344809627"),
		InitialBalances: []Account{
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
		"tSlRQl0lF1BPdnIVynepvj4M14j9ckQ9qbF0/GhqN/A=",
		//"c42kG6ant9abcpSvoVi4nFobQQy/DCRDyFxf4krR3Rw=",
		//"bxbB/yGm+5RqrtD0wfzKJyty/ZBJiPkdOIMoK4rjG6I=",
		//"nhPy9UaN14KzFkRPvWZZXhPbp9e9Pvob7NULQgRfWMY=",
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
		FirstValidatorAddress: types.HexToAddress("Mxee81347211c72524338f9680072af90744333146"),
		InitialBalances: []Account{
			{
				Address: types.HexToAddress("Mxee81347211c72524338f9680072af90744333146"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx184ac726059e43643e67290666f7b3195093f870"),
				Balance: map[string]string{
					"BTST": helpers.BipToPip(big.NewInt(100000000)).String(),
				},
			},
			{
				Address: types.HexToAddress("Mx4e828501f3a5325d8f7fad4c5bc9db8da1938afe"),
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
