package weLib

import (
	"sync"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

type Manager struct {
	players   []*Player
	playersMu sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		players: make([]*Player, 0),
	}
}

func (m *Manager) Player(p *player.Player) (*Player, bool) {
	m.playersMu.RLock()
	for _, pl := range m.players {
		if p == pl.p {
			m.playersMu.RUnlock()
			return pl, true
		}
	}
	m.playersMu.RUnlock()
	return nil, false
}

func (m *Manager) storePlayer(p *Player) {
	m.playersMu.Lock()
	defer m.playersMu.Unlock()
	m.players = append(m.players, p)
}
func (m *Manager) RemovePlayer(p *player.Player) {
	m.playersMu.RLock()
	for n, pl := range m.players {
		if p == pl.p {
			m.playersMu.RUnlock()
			m.playersMu.Lock()
			m.players = append(m.players[:n], m.players[n+1:]...)
			m.playersMu.Unlock()
		}
	}
	m.playersMu.RUnlock()
}

func (m *Manager) SetPos1(p *player.Player, pos mgl64.Vec3) {
	if pl, ok := m.Player(p); ok {
		pl.area.SetPos1(pos)
	} else {
		area := NewArea(pos, mgl64.Vec3{})
		pl := NewPlayer(p, area)
		m.storePlayer(pl)
	}
}
func (m *Manager) SetPos2(p *player.Player, pos mgl64.Vec3) {
	if pl, ok := m.Player(p); ok {
		pl.area.SetPos2(pos)
	} else {
		area := NewArea(mgl64.Vec3{}, pos)
		pl := NewPlayer(p, area)
		m.storePlayer(pl)
	}
}
