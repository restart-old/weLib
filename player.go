package worldEdit

import (
	"github.com/df-mc/dragonfly/server/player"
)

type Player struct {
	p    *player.Player
	area *Area
}

func NewPlayer(p *player.Player, area *Area) *Player {
	return &Player{
		p:    p,
		area: area,
	}
}
