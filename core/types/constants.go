package types

func GetBaseCoin() CoinSymbol {
	return getBaseCoin(2)
}

func getBaseCoin(chainId int) CoinSymbol {
	var coin CoinSymbol

	switch chainId {
	case 1:
		copy(coin[:], []byte("BTSG"))
	case 2:
		copy(coin[:], []byte("BTST"))
	}

	coin[4] = byte(0)

	return coin
}
