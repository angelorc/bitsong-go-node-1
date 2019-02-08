package eventsdb

import "github.com/BitSongOfficial/go-amino"

func RegisterAminoEvents(codec *amino.Codec) {
	codec.RegisterInterface((*Event)(nil), nil)
	codec.RegisterConcrete(RewardEvent{},
		"bitsong/RewardEvent", nil)
	codec.RegisterConcrete(SlashEvent{},
		"bitsong/SlashEvent", nil)
	codec.RegisterConcrete(UnbondEvent{},
		"bitsong/UnbondEvent", nil)
}
