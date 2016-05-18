package lib

import (
// "math"
)

const (
	WEAPON = 0
	ARMOR  = 1
)

type Player struct {
	Name         string
	Exp          int
	Level        uint8
	Job          uint8
	Strength     int
	Agility      int
	Intelligence int
	MaxHealth    int
	Health       int
	Mana         int
	MaxMana      int
	Attack       int
	Defend       int
	Dodge        int
	Updated      bool //is finished increase attributes after level up
	// Skill     []Skill
	// Equipment []Equipment
}

func NewPlayer(name string) *Player {

	strength, agility, intelligence := 20, 15, 5
	maxHealth := PerStrToHealth * strength
	maxMana := PerIntToMana * intelligence

	p := &Player{
		Name:         name,
		Exp:          0,
		Level:        0,
		Strength:     strength,
		Agility:      agility,
		Intelligence: intelligence,
		MaxHealth:    maxHealth,
		Health:       maxHealth,
		MaxMana:      maxMana,
		Mana:         maxMana,
		Attack:       0,
		Defend:       0,
		Dodge:        0,
	}
	p.Attack = p.getBaseAttact()
	p.Defend = p.getBaseDefend()
	return p
}

func (this *Player) RefreshPlayer() {
	this.Health = this.MaxHealth
	this.Mana = this.MaxMana
}

func (this *Player) LevelUp() {

}

func (this *Player) getBaseAttact() int {
	return this.Strength * PerStrToAttack
}

func (this *Player) getBaseDefend() int {
	return this.Strength * PerStrToDefend

}
