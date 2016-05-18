package lib

import (
	"math"
)

type Battle struct {
	Boss    *Boss
	Players []*Player
	// monster    []*monster
	Turn       uint8 // determine whether boss or player attack
	Difficulty uint8
	Round      uint16
}

func NewBattle(boss *Boss, players ...*Player) *Battle {
	battle := &Battle{
		Boss:       boss,
		Players:    players,
		Turn:       0, // 0 represent the status of player will attack boss | 1 represent the status of boss will attack player
		Difficulty: 1,
		Round:      0, // current round
	}
	return battle
}

func (this *Battle) getDodge(p *Player, b *Boss) int {
	if p.Agility == b.Agility {
		return 0
	}
	if p.Agility > b.Agility {
		return int(math.Ceil(float64(p.Agility-b.Agility)/float64(p.Agility)) * 100)
	}
	if p.Agility < b.Agility {
		return int(math.Ceil(float64(b.Agility-p.Agility)/float64(b.Agility)) * 100)
	}
	return 0
}

func (this *Battle) AttackPtoB(p *Player, b *Boss) {

}
