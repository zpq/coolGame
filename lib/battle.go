package lib

import (
	"math"
)

type Battle struct {
	// boss       *Boss
	// player     []*Player
	// monster    []*monster
	// turn       uint8
	// difficulty uint8
	// round      uint16
}

func (this *Battle) getDodge(p *Player, b *Boss) int {
	if p.Agility == b.Agility {
		return 0
	}
	if p.Agility > b.Agility {
		return int(math.Ceil((float64((float64(p.Agility) - b.Agility) / p.Agility)) * 100))
	}
	if p.Agility < b.Agility {
		return int(math.Ceil((float64((float64(b.Agility) - p.Agility) / b.Agility)) * 100))
	}
	return 0
}
