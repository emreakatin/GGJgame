package assets

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	inventoryScale float32 = 0.5
)

const (
	TurretCost  = 30
	FactoryCost = 100
)

type Inventory struct {
	Player        uint
	MechanicParts float32
	Turret        uint
}

var PlayerInventory Inventory
var Player2Inventory Inventory
var Player3Inventory Inventory

var InventoryBlockTexture rl.Texture2D

func CreateInventory() {
	PlayerInventory = Inventory{
		Player:        PlayerID,
		MechanicParts: 130000,
		Turret:        0,
	}

	Player2Inventory = Inventory{
		Player:        Player2ID,
		MechanicParts: 130000,
		Turret:        0,
	}

	Player3Inventory = Inventory{
		Player:        Player3ID,
		MechanicParts: 130000,
		Turret:        0,
	}

	InventoryBlockTexture = rl.LoadTexture("sprites/texturesquare.png")
}

func DrawInventory() {
	rl.DrawText(strconv.Itoa(int(PlayerInventory.MechanicParts)), Background.Width-195, Background.Height-26, 23, rl.White)

	if PlayerInventory.MechanicParts < TurretCost {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 150), float32(Background.Height - InventoryBlockTexture.Height)}, rl.Black)
	} else {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 150), float32(Background.Height - InventoryBlockTexture.Height)}, rl.White)
	}

	if PlayerInventory.MechanicParts < FactoryCost {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 118), float32(Background.Height - InventoryBlockTexture.Height)}, rl.Black)
	} else {
		rl.DrawTextureRec(InventoryBlockTexture, rl.Rectangle{0, 0, float32(InventoryBlockTexture.Width), float32(InventoryBlockTexture.Height)}, rl.Vector2{float32(Background.Width - 118), float32(Background.Height - InventoryBlockTexture.Height)}, rl.White)
	}
}
