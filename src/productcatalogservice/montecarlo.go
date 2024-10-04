package main

import (
	"math/rand"
	"time"
)

var (
	currentTicks = 0
)

const (
	screenWidth      = 640
	screenHeight     = 640
	noisePerTick     = 8000
	ticksBeforeStart = 120
)

func NewCalc() *Calc {
	return &Calc{
		pointsToCalculate: int32(rand.Intn(15000000) + 10000000),
	}
}

// Calc is our main game object.
type Calc struct {
	totalPoints       int32
	pointsInCircle    int32
	estimatedPi       float64
	finished          bool
	pointsToCalculate int32
}

// GenerateNoise generates new noise in
func (g *Calc) GenerateNoise() float64 {
	i := 0
	for i < noisePerTick {
		// We're done.
		if g.totalPoints == g.pointsToCalculate {
			g.finished = true
			break
		}
		// Generate random x/y coordinates
		x := rand.Intn(screenWidth)
		y := rand.Intn(screenHeight)
		g.totalPoints++

		// In my case, x-center, y-center and radius are the same.
		if withinCircle(x, y, screenWidth/2, screenHeight/2, screenWidth/2) {
			g.pointsInCircle++
		}
		g.estimatedPi = 4 * float64(g.pointsInCircle) / float64(g.totalPoints)
		i++
	}
	return g.estimatedPi
}

// withinCircle is a very simple Pythagoras implementation.
// See also: https://stackoverflow.com/questions/481144/equation-for-testing-if-a-point-is-inside-a-circle
func withinCircle(x, y, centerX, centerY, radius int) bool {
	return (x-centerX)*(x-centerX)+(y-centerY)*(y-centerY) < (radius * radius)
}

// CalculatePi takes a long time to estimate Pi
func (c *Calc) CalculatePi() float64 {
	ticker := time.NewTicker(5 * time.Millisecond)
	for range ticker.C {
		result := c.GenerateNoise()
		if c.finished {
			return result
		}
	}
	return 0
}