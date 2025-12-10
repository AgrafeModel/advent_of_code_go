// ...existing code...
package utils

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenW = 1000
	screenH = 700
)

type PolygonVisualizer struct {
	Polygon *Polygon2DEdge
	Color   color.Color
}

type polygonGame struct {
	polys  []PolygonVisualizer
	points map[Position2D]struct{}

	offsetX float64
	offsetY float64
	scale   float64
}

func (g *polygonGame) Update() error {
	speed := 1.0 / (g.scale * 10)
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.offsetX += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.offsetX -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.offsetY += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.offsetY -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyEqual) || ebiten.IsKeyPressed(ebiten.KeyKPAdd) { // +
		g.scale *= 1.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyMinus) || ebiten.IsKeyPressed(ebiten.KeyKPSubtract) { // -
		g.scale /= 1.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		// reset view
		g.scale = 10.0
		g.offsetX = screenW / 2
		g.offsetY = screenH / 2
	}
	return nil
}

func (g *polygonGame) worldToScreen(p Position2D) (float64, float64) {
	x := g.offsetX + float64(p.X)*g.scale
	y := g.offsetY - float64(p.Y)*g.scale
	return x, y
}

func (g *polygonGame) Draw(screen *ebiten.Image) {
	// background
	screen.Fill(color.RGBA{20, 20, 30, 255})

	// draw grid axes (light)
	for x := -100; x <= 100; x++ {
		xs := float64(x)*g.scale + g.offsetX
		ebitenutil.DrawLine(screen, xs, 0, xs, screenH, color.RGBA{40, 40, 50, 255})
	}
	for y := -100; y <= 100; y++ {
		ys := g.offsetY - float64(y)*g.scale
		ebitenutil.DrawLine(screen, 0, ys, screenW, ys, color.RGBA{40, 40, 50, 255})
	}

	// draw points (red)
	pointSize := g.scale * 0.6
	if pointSize < 1 {
		pointSize = 1
	}
	for p := range g.points {
		x, y := g.worldToScreen(p)
		vector.DrawFilledRect(screen, float32(x-pointSize/2), float32(y-pointSize/2), float32(pointSize), float32(pointSize), color.RGBA{255, 0, 0, 255}, true)
	}

	// draw polygons
	for _, polyVis := range g.polys {
		poly := polyVis.Polygon
		col := polyVis.Color
		n := len(poly.edges)
		for i := 0; i < n; i++ {
			start := poly.edges[i].start
			end := poly.edges[i].end
			x1, y1 := g.worldToScreen(start)
			x2, y2 := g.worldToScreen(end)
			ebitenutil.DrawLine(screen, x1, y1, x2, y2, col)
		}
	}

}

func (g *polygonGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

// ShowPolygonEBit launches an Ebiten window showing the polygon and points with navigation.
// Call as: utils.ShowPolygonEBit(polygon, pointSet)
func ShowPolygonEBit(polys []PolygonVisualizer, points map[Position2D]struct{}) error {
	// compute bounding box of all content
	minX, minY := int(1e9), int(1e9)
	maxX, maxY := -int(1e9), -int(1e9)
	found := false
	for _, pv := range polys {
		if pv.Polygon == nil {
			continue
		}
		for _, v := range pv.Polygon.Vertices {
			if v.X < minX {
				minX = v.X
			}
			if v.Y < minY {
				minY = v.Y
			}
			if v.X > maxX {
				maxX = v.X
			}
			if v.Y > maxY {
				maxY = v.Y
			}
			found = true
		}
	}
	for p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		found = true
	}

	// default view
	scale := 10.0
	offsetX := float64(screenW) / 2
	offsetY := float64(screenH) / 2
	if found {
		wUnits := math.Max(1, float64(maxX-minX))
		hUnits := math.Max(1, float64(maxY-minY))
		// leave margin
		scale = math.Min(float64(screenW)*0.8/wUnits, float64(screenH)*0.8/hUnits)
		if scale <= 0 {
			scale = 1.0
		}
		centerX := (float64(minX) + float64(maxX)) / 2.0
		centerY := (float64(minY) + float64(maxY)) / 2.0
		offsetX = float64(screenW)/2.0 - centerX*scale
		offsetY = float64(screenH)/2.0 + centerY*scale
	}

	g := &polygonGame{
		polys:   polys,
		points:  points,
		scale:   scale,
		offsetX: offsetX,
		offsetY: offsetY,
	}
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Polygon viewer")
	if err := ebiten.RunGame(g); err != nil {
		log.Println("ebiten run error:", err)
		return err
	}
	return nil
}
