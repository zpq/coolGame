package lib

import (
	"errors"
)

type Skill struct {
	skillId   uint16
	skillName string
	skillType uint8
	cost      uint8
	needLevel uint8
}

//one method represent a concrete skill

func (this Skill) Health(obj *Player) (int, error) {
	if this.needLevel > obj.Level {
		return 0, errors.New("Invalid level")
	} else {
		return obj.Health + 100, nil
	}
}
