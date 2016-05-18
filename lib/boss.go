package lib

type Boss struct {
	Name         string
	Exp          int
	Level        uint8
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
}

func NewBoss() *Boss {
	boss := &Boss{
		Name:         "Hades",
		Exp:          100,
		Level:        10,
		Strength:     30,
		Agility:      20,
		Intelligence: 5,
		MaxHealth:    700,
		Health:       700,
		Mana:         200,
		MaxMana:      200,
		Attack:       50,
		Defend:       10,
		Dodge:        0,
	}
	return boss
}
