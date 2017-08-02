package game

type Tile int

type TileType int

const (
	LiverTile TileType = iota
	LungTile
	HeartTile
	StomachTile
	BrainTile
)

func (t Tile) Place(x, y int, typ TileType) {

}
