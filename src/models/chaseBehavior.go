package models

import (
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/contexts"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/structures"
	"github.com/LuisPalominoTrevilla/MultithreadedPacman/src/utils"
)

// BlinkyChaseBehavior according to the original PacMan game
type BlinkyChaseBehavior struct {
	ghost *Ghost
	ctx   *contexts.GameContext
}

// SwitchDirection by heading towards the player
func (b *BlinkyChaseBehavior) SwitchDirection() {
	b.ghost.turnTowards(b.ctx.MainPlayer.GetPosition(), false, true)
}

// PinkyChaseBehavior according to the original PacMan game
type PinkyChaseBehavior struct {
	ghost *Ghost
	ctx   *contexts.GameContext
}

// SwitchDirection by heading 3 steps in the direction of the player
func (p *PinkyChaseBehavior) SwitchDirection() {
	from := p.ctx.MainPlayer.GetPosition()
	direction := p.ctx.MainPlayer.GetDirection()
	cols, rows := p.ctx.Maze.Dimensions()
	toX := utils.Mod(from.X()+direction.X*3, cols)
	toY := utils.Mod(from.Y()+direction.Y*3, rows)
	target := structures.InitPosition(toX, toY)
	p.ghost.turnTowards(target, false, true)
}

// InkyChaseBehavior according to the original PacMan game
type InkyChaseBehavior struct {
	ghost *Ghost
	ctx   *contexts.GameContext
}

// SwitchDirection by heading 3 steps opposite to the direction of the player
func (i *InkyChaseBehavior) SwitchDirection() {
	from := i.ctx.MainPlayer.GetPosition()
	direction := i.ctx.MainPlayer.GetDirection()
	cols, rows := i.ctx.Maze.Dimensions()
	toX := utils.Mod(from.X()+direction.X*-3, cols)
	toY := utils.Mod(from.Y()+direction.Y*-3, rows)
	target := structures.InitPosition(toX, toY)
	i.ghost.turnTowards(target, false, true)
}

// ClydeChaseBehavior according to the original PacMan game
type ClydeChaseBehavior struct {
	ghost *Ghost
	ctx   *contexts.GameContext
}

// SwitchDirection by heading towards the player only if its distance is greater than 5
func (i *ClydeChaseBehavior) SwitchDirection() {
	pacmanPosition := i.ctx.MainPlayer.GetPosition()
	distance := i.ghost.position.DistanceTo(pacmanPosition)
	if distance < 3 {
		i.ghost.turnTowards(nil, false, true)
	} else {
		i.ghost.turnTowards(pacmanPosition, false, true)
	}
}
