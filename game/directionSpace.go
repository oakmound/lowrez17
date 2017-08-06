package game

import (
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/physics"
)

type DirectionSpace struct {
	*collision.Space
	physics.ForceVector
}

func (ds *DirectionSpace) Init() event.CID {
	return event.NextID(ds)
}

func NewDirectionSpace(s *collision.Space, v physics.ForceVector) *DirectionSpace {
	ds := &DirectionSpace{
		Space:       s,
		ForceVector: v,
	}
	s.CID = ds.Init()
	return ds
}
