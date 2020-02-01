package assets

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	inventoryScale float32 = 0.5
)

type Inventory struct {
	Player        uint
	MechanicParts uint
	Turret        uint
}

var PlayerInventory Inventory

var InventoryBlockTexture rl.Texture2D

func CreateInventory() {
	PlayerInventory = Inventory{
		Player:        PlayerID,
		MechanicParts: 0,
		Turret:        0,
	}

	InventoryBlockTexture = rl.LoadTexture("sprites/texturesquare.png")
}

func DrawInventory() {
	rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 150), float32(Background.Height - InventoryBlockTexture.Height)}, rl.White)
}
