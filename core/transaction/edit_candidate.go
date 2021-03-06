package transaction

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/BitSongOfficial/bitsong-go-node/core/code"
	"github.com/BitSongOfficial/bitsong-go-node/core/commissions"
	"github.com/BitSongOfficial/bitsong-go-node/core/state"
	"github.com/BitSongOfficial/bitsong-go-node/core/types"
	"github.com/BitSongOfficial/bitsong-go-node/formula"
	"github.com/tendermint/tendermint/libs/common"
	"math/big"
)

type EditCandidateData struct {
	PubKey        types.Pubkey  `json:"pub_key"`
	RewardAddress types.Address `json:"reward_address"`
	OwnerAddress  types.Address `json:"owner_address"`
}

func (data EditCandidateData) TotalSpend(tx *Transaction, context *state.StateDB) (TotalSpends, []Conversion, *big.Int, *Response) {
	panic("implement me")
}

func (data EditCandidateData) BasicCheck(tx *Transaction, context *state.StateDB) *Response {
	if data.PubKey == nil {
		return &Response{
			Code: code.DecodeError,
			Log:  "Incorrect tx data"}
	}

	if !context.CoinExists(tx.GasCoin) {
		return &Response{
			Code: code.CoinNotExists,
			Log:  fmt.Sprintf("Coin %s not exists", tx.GasCoin)}
	}

	if !context.CandidateExists(data.PubKey) {
		return &Response{
			Code: code.CandidateNotFound,
			Log:  fmt.Sprintf("Candidate with such public key (%s) not found", data.PubKey.String())}
	}

	candidate := context.GetStateCandidate(data.PubKey)
	sender, _ := tx.Sender()
	if !bytes.Equal(candidate.OwnerAddress.Bytes(), sender.Bytes()) {
		return &Response{
			Code: code.IsNotOwnerOfCandidate,
			Log:  fmt.Sprintf("Sender is not an owner of a candidate")}
	}

	return nil
}

func (data EditCandidateData) String() string {
	return fmt.Sprintf("EDIT CANDIDATE pubkey: %x",
		data.PubKey)
}

func (data EditCandidateData) Gas() int64 {
	return commissions.EditCandidate
}

func (data EditCandidateData) Run(tx *Transaction, context *state.StateDB, isCheck bool, rewardPool *big.Int, currentBlock int64) Response {
	sender, _ := tx.Sender()

	response := data.BasicCheck(tx, context)
	if response != nil {
		return *response
	}

	commissionInBaseCoin := tx.CommissionInBaseCoin()
	commission := big.NewInt(0).Set(commissionInBaseCoin)

	if !tx.GasCoin.IsBaseCoin() {
		coin := context.GetStateCoin(tx.GasCoin)

		if coin.ReserveBalance().Cmp(commissionInBaseCoin) < 0 {
			return Response{
				Code: code.CoinReserveNotSufficient,
				Log:  fmt.Sprintf("Coin reserve balance is not sufficient for transaction. Has: %s, required %s", coin.ReserveBalance().String(), commissionInBaseCoin.String())}
		}

		commission = formula.CalculateSaleAmount(coin.Volume(), coin.ReserveBalance(), coin.Data().Crr, commissionInBaseCoin)
	}

	if context.GetBalance(sender, tx.GasCoin).Cmp(commission) < 0 {
		return Response{
			Code: code.InsufficientFunds,
			Log:  fmt.Sprintf("Insufficient funds for sender account: %s. Wanted %s %s", sender.String(), commission, tx.GasCoin)}
	}

	if !isCheck {
		rewardPool.Add(rewardPool, commissionInBaseCoin)

		context.SubCoinReserve(tx.GasCoin, commissionInBaseCoin)
		context.SubCoinVolume(tx.GasCoin, commission)

		context.SubBalance(sender, tx.GasCoin, commission)
		context.EditCandidate(data.PubKey, data.RewardAddress, data.OwnerAddress)
		context.SetNonce(sender, tx.Nonce)
	}

	tags := common.KVPairs{
		common.KVPair{Key: []byte("tx.type"), Value: []byte(hex.EncodeToString([]byte{byte(TypeEditCandidate)}))},
		common.KVPair{Key: []byte("tx.from"), Value: []byte(hex.EncodeToString(sender[:]))},
	}

	return Response{
		Code:      code.OK,
		GasUsed:   tx.Gas(),
		GasWanted: tx.Gas(),
		Tags:      tags,
	}
}
