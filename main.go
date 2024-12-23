package main
import (
    _ "image/png"
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
    Tiles []MapTile
}

type GameData struct {
    ScreenWidth int
	  ScreenHeight int
	  TileWidth int
    TileHeight int
}

// Constructors
func NewGameData() *GameData {
    gamedata := &GameData{
        ScreenWidth: 80,
		    ScreenHeight: 50,
		    TileWidth: 16,
		    TileHeight: 16,
    }
    return gamedata
}

func NewGame() *Game {
    game := &Game{
        Tiles: CreateTiles(),
    }
    return game
}

// Game object functions

// Update is called each "tic" of the clock
func (game *Game) Update() error {
    return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
    gd := NewGameData()
    //Draw the Map
    for x := 0; x < gd.ScreenWidth; x++ {
        for y := 0; y < gd.ScreenHeight; y++ {
            tile := game.Tiles[GetIndexFromXY(x, y)]
            op := &ebiten.DrawImageOptions{}
            op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
            screen.DrawImage(tile.Image, op)
	      }
    }
}

func (game *Game) Layout(width, height int) (int, int) {
    return 1280, 800
}

type MapTile struct {
    PixelX int
    PixelY int
    Blocked bool
    Image *ebiten.Image
}

//GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
//This coordinate is logical tiles, not pixels.
func GetIndexFromXY(x int, y int) int {
    gd := NewGameData()
    return (y * gd.ScreenWidth) + x
}

// Create border of walls on map
func CreateTiles() []MapTile {
    gd := NewGameData()
    tiles := make([]MapTile, 0)

    for x := 0; x < gd.ScreenWidth; x++ {
        for y := 0; y < gd.ScreenHeight; y++ {
            if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {
                wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")

                if err != nil {
                    log.Fatal(err)
                }
                tile := MapTile{
                    PixelX: x * gd.TileWidth,
                    PixelY: y * gd.TileHeight,
                    Blocked: true,
                    Image: wall,
                }
                tiles = append(tiles, tile)
            } else {
                floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
                if err != nil {
                    log.Fatal(err)
                }
                tile := MapTile{
                    PixelX: x * gd.TileWidth,
                    PixelY: y * gd.TileHeight,
                    Blocked: false,
                    Image: floor,
                }
                tiles = append(tiles, tile)
            }
        }
    }

    return tiles
}

func main() {
    game := NewGame()
    ebiten.SetWindowResizable(true)
    ebiten.SetWindowTitle("Go Rogue")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
