package weLib

import (
	"github.com/df-plus/area"
	"github.com/go-gl/mathgl/mgl64"
)

type Area struct {
	pos1, pos2 mgl64.Vec3
	area.Vec3
}

func NewArea(pos1, pos2 mgl64.Vec3) *Area {
	return &Area{
		pos1: pos1,
		pos2: pos2,
		Vec3: area.NewVec3(pos1, pos2),
	}
}
func (a *Area) SetPos1(pos mgl64.Vec3) {
	a.pos1 = pos
	a.Vec3 = area.NewVec3(pos, a.pos2)
}
func (a *Area) SetPos2(pos mgl64.Vec3) {
	a.pos2 = pos
	a.Vec3 = area.NewVec3(a.pos1, pos)
}

func (a *Area) AllVec3() (v []mgl64.Vec3) {
	for x := a.Min().X(); x < a.Max().X()+1; x++ {
		for y := a.Min().Y(); y < a.Max().Y()+1; y++ {
			for z := a.Min().Z(); z < a.Max().Z()+1; z++ {
				v = append(v, mgl64.Vec3{x, y, z})
			}
		}
	}
	return v
}
