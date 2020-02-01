package assets

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	inventoryScale float32 = 0.5
)

const (
	turretCost  = 30
	factoryCost = 100
	bridgeCost  = 60
)

type Inventory struct {
	Player        uint
	MechanicParts float32
	Turret        uint
}

var PlayerInventory Inventory

var InventoryBlockTexture rl.Texture2D

func CreateInventory() {
	PlayerInventory = Inventory{
		Player:        PlayerID,
		MechanicParts: 130,
		Turret:        0,
	}

	InventoryBlockTexture = rl.LoadTexture("sprites/texturesquare.png")
}

func DrawInventory() {
	if PlayerInventory.MechanicParts < turretCost {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 150), float32(Background.Height - InventoryBlockTexture.Height)}, rl.Black)
	} else {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 150), float32(Background.Height - InventoryBlockTexture.Height)}, rl.White)
	}

	if PlayerInventory.MechanicParts < factoryCost {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 118), float32(Background.Height - InventoryBlockTexture.Height)}, rl.Black)
	} else {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 118), float32(Background.Height - InventoryBlockTexture.Height)}, rl.White)
	}

	if PlayerInventory.MechanicParts < bridgeCost {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 86), float32(Background.Height - InventoryBlockTexture.Height)}, rl.Black)
	} else {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 86), float32(Background.Height - InventoryBlockTexture.Height)}, rl.White)
	}
}
