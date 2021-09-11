package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600

	BoardWidth  = 300
	BoardHeight = 300
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)  // window size
	ebiten.SetWindowTitle("DrawBoard (Ebiten Demo)") // window title
	// これだけだと、即終了してしまう

	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(game); err != nil { // game main loop start
		log.Fatal(err)
	}
}

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	boardColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
	pointColor      = color.RGBA{0xff, 0x80, 0x80, 0xff}
)

type myGame struct {
	boardImage *ebiten.Image
	pointImage *ebiten.Image
	click      myClick
}

type myClick struct {
	state bool
	myPoint
}

type myPoint struct {
	x, y int
}

var points []myPoint // クリックポイントを記録する
var initDraw = false

// NewGame generates a new Game object.
func NewGame() (*myGame, error) {
	g := &myGame{}
	return g, nil
}

func (g *myGame) Update() error {
	// マウス操作などはUpdateに実装する
	if !g.click.state {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // マウスの左クリック
			g.click.state = true
			g.click.x, g.click.y = ebiten.CursorPosition()
			fmt.Println("mouse click ", g.click.x, g.click.y)
			points = append(points, myPoint{
				x: g.click.x,
				y: g.click.y,
			})
		}
	} else {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // マウスの左クリック
			g.click.state = false
		}
	}
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor) // 背景を白塗りする

	if g.boardImage == nil {
		w, h := BoardWidth, BoardHeight
		g.boardImage = ebiten.NewImage(w, h)
		g.boardImage.Fill(boardColor) // ボードを塗りつぶす
	}
	if g.pointImage == nil {
		w, h := 10, 10
		g.pointImage = ebiten.NewImage(w, h)
		g.pointImage.Fill(pointColor) // クリックポイントのオブジェクトを塗りつぶす
	}

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := BoardWidth, BoardHeight
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)

	for _, p := range points {
		op := &ebiten.DrawImageOptions{}
		x := p.x - 10/2
		y := p.y - 10/2
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(g.pointImage, op)
	}
}

func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
