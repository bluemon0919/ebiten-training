package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600

	BoardWidth  = 300
	BoardHeight = 300

	KeyWidth  = 50
	KeyHeight = 50
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
	keyOnColor      = color.RGBA{0x00, 0x00, 0x00, 0xff}
	keyOffColor     = color.RGBA{0x80, 0x80, 0xff, 0xff}
)

type myGame struct {
	boardImage *ebiten.Image
	pointImage *ebiten.Image
	keys       map[int]myKey
}

// NewGame generates a new Game object.
func NewGame() (*myGame, error) {
	g := &myGame{
		keys: make(map[int]myKey, 4),
	}
	g.keyAppend()
	return g, nil
}

// keyAppend キーの配置を設定します
func (g *myGame) keyAppend() {
	const margin = 10 // boardに対するマージン
	bw, bh := BoardWidth, BoardHeight
	kw, kh := KeyWidth, KeyHeight
	key := myKey{
		state: false,
		w:     KeyWidth,
		h:     KeyHeight,
	}

	// keyImageはboardImage上の座標系に所属するため、
	// boardImage上の(0,0)-(bw,bh)の範囲で指定する
	key.x = (bw - kw) / 2
	key.y = margin
	key.key = ebiten.KeyUp
	g.keys[int(ebiten.KeyUp)] = key

	key.x = (bw - kw) / 2
	key.y = bh - kh - margin
	key.key = ebiten.KeyDown
	g.keys[int(ebiten.KeyDown)] = key

	key.x = margin
	key.y = (bh - kh) / 2
	key.key = ebiten.KeyLeft
	g.keys[int(ebiten.KeyLeft)] = key

	key.x = bw - kw - margin
	key.y = (bh - kh) / 2
	key.key = ebiten.KeyRight
	g.keys[int(ebiten.KeyRight)] = key
}

func (g *myGame) Update() error {
	// キー入力などはUpdateに実装する
	for i, k := range g.keys {
		err := k.Update()
		if err != nil {
			return err
		}
		g.keys[i] = k
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

	for _, k := range g.keys {
		k.Draw(g.boardImage)
	}

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := BoardWidth, BoardHeight
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}

func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

type myKey struct {
	key        ebiten.Key // キーの種類
	x, y, w, h int
	state      bool // キーの状態ON(1)/OFF(0)
}

func (k *myKey) Update() error {
	if ebiten.IsKeyPressed(k.key) {
		k.state = true
	} else if k.state {
		k.state = false
	} else {
	}
	return nil
}

func (k *myKey) Draw(boardImage *ebiten.Image) {
	keyImage := ebiten.NewImage(k.w, k.h)
	if k.state {
		keyImage.Fill(keyOnColor)
	} else {
		keyImage.Fill(keyOffColor)
	}

	op := &ebiten.DrawImageOptions{}
	x, y := k.x, k.y
	op.GeoM.Translate(float64(x), float64(y))

	boardImage.DrawImage(keyImage, op)
}
