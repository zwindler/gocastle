package maps

type TileInfo struct {
	Path       string
	IsWalkable bool
}

var (
	TilesTypes = []TileInfo{
		{Path: "./static/grass.png", IsWalkable: true},
		{Path: "./static/building_corner_left_up.png", IsWalkable: false},
		{Path: "./static/building_corner_right_up.png", IsWalkable: false},
		{Path: "./static/building_corner_right_down.png", IsWalkable: false},
		{Path: "./static/building_corner_left_down.png", IsWalkable: false},
		{Path: "./static/building_left.png", IsWalkable: false},
		{Path: "./static/building_down.png", IsWalkable: false},
		{Path: "./static/building_right.png", IsWalkable: false},
		{Path: "./static/building_up.png", IsWalkable: false},
		{Path: "./static/opaque.png", IsWalkable: false},
	}
)
