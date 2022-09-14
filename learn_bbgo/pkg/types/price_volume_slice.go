package types

import "github.com/zixsa/learn_bbgo/pkg/fixedpoint"

type PriceVolume struct {
	Price, Volume fixedpoint.Value
}

type PriceVolumeSlice []PriceVolume
