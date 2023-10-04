// screens/inventory.go

package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/pkg/embedimages"
	"github.com/zwindler/gocastle/pkg/game"
	"github.com/zwindler/gocastle/pkg/newtheme"
)

var (
	// TODO clean this.
	totalWeightValueLabel    *canvas.Text
	equippedWeightValueLabel *canvas.Text
	goldAmountValueLabel     *canvas.Text

	itemDropdowns []dropdown
)

type dropdown struct {
	category model.Category
	widget   *widget.Select
}

// ShowInventoryScreen is the main function of the inventory screen.
func ShowInventoryScreen(window fyne.Window) {
	// Load the background image
	img, _ := embedimages.GetImageFromEmbed("static/inventory.png")
	backgroundImage := canvas.NewImageFromImage(img)
	backgroundImage.FillMode = canvas.ImageFillContain

	// Create a container to hold the dropdown lists for each category
	inventoryContainerLeft := container.NewVBox()
	inventoryContainerRight := container.NewVBox()
	inventoryStatsArea := createInventoryStatsArea()

	// Iterate over each object category and display the dropdown list for the items in that category
	for index, category := range model.CategoryList {
		// Create a header label for the category
		categoryLabel := widget.NewLabel(category.Name)

		// Find the items in the player's inventory that belong to the current category
		itemsInCategory := make([]string, 0)
		for _, item := range game.Player.Inventory {
			if item.Category == category.Name {
				itemsInCategory = append(itemsInCategory, item.Name)
			}
		}

		// Create a dropdown list to display the items in the category
		itemDropdown := dropdown{
			category: category,
			widget: widget.NewSelect(itemsInCategory, func(selected string) {
				for _, item := range game.Player.Inventory {
					if item.Name == selected {
						if err := game.Player.EquipItem(item); err != nil {
							fmt.Printf("Error equipping item: %s\n", err)
						}
					}
				}
				updateInventoryStatsArea()
			}),
		}

		// add the pointer of this widget in the list of items
		itemDropdowns = append(itemDropdowns, itemDropdown)

		for _, item := range game.Player.Inventory {
			if item.Equipped {
				itemDropdown.widget.SetSelected(item.Name)
				break
			}
		}

		// Create a container to hold the category label and the dropdown list
		categoryContainer := container.NewVBox(
			categoryLabel,
			itemDropdown.widget,
		)

		// Add the category container to the left of right inventory container
		if index < len(model.CategoryList)/2 {
			inventoryContainerLeft.Add(categoryContainer)
		} else {
			inventoryContainerRight.Add(categoryContainer)
		}
	}

	// Create a "Back" button to return to the main menu
	backButton := widget.NewButton("Back", func() {
		// gear may have changed, reset all secondary stats
		game.Player.RefreshStats(false)
		ShowGameScreen(window)
	})

	// Create the content container to hold the inventory items and the back button
	inventoryContainer := container.NewBorder(nil, nil, inventoryContainerLeft, inventoryContainerRight, backgroundImage)

	floorContainer := displayFloorItems()
	floorScroll := container.NewVScroll(floorContainer)
	bottomRightColumn := container.NewVBox(inventoryStatsArea, backButton)
	rightColumn := container.NewBorder(nil, bottomRightColumn, nil, nil, floorScroll)
	content := container.NewBorder(nil, nil, nil, rightColumn, inventoryContainer)
	window.SetContent(content)
}

// createInventoryStatsArea creates the stats area containing inventory weight and gold amount.
func createInventoryStatsArea() fyne.CanvasObject {
	totalWeightValueLabel = canvas.NewText("", newtheme.TextColor)
	equippedWeightValueLabel = canvas.NewText("", newtheme.TextColor)
	goldAmountValueLabel = canvas.NewText("", newtheme.TextColor)

	labels := container.NewVBox(
		canvas.NewText("Inventory Weight:", newtheme.TextColor),
		canvas.NewText("Equipped Items Weight:", newtheme.TextColor),
		canvas.NewText("Gold amount:", newtheme.TextColor))
	values := container.NewVBox(
		totalWeightValueLabel,
		equippedWeightValueLabel,
		goldAmountValueLabel)

	updateInventoryStatsArea()

	return container.NewHBox(labels, values)
}

// updateInventoryStatsArea refreshes the values in InventoryStatsArea.
func updateInventoryStatsArea() {
	totalWeightValueLabel.Text = fmt.Sprintf("%.3f kg", float32(game.Player.InventoryWeight)/1000)
	totalWeightValueLabel.Refresh()

	equippedWeightValueLabel.Text = fmt.Sprintf("%.3f kg", float32(game.Player.EquippedWeight)/1000)
	equippedWeightValueLabel.Refresh()

	goldAmountValueLabel.Text = fmt.Sprintf("%d gold pieces", game.Player.CurrentGold)
	goldAmountValueLabel.Refresh()
}

func displayFloorItems() (floorVBox *fyne.Container) {
	floorVBox = container.NewVBox()
	for _, item := range game.CurrentMap.ObjectList {
		if item.Coord.X == game.Player.Avatar.Coord.X && item.Coord.Y == game.Player.Avatar.Coord.Y {
			currentItemContainer := container.NewVBox()
			nameLabel := widget.NewLabel(item.Name)

			// Create a copy of the item variable for use in the closure
			itemCopy := item

			// Use a closure to pass arguments to the takeButton's callback function
			takeButton := widget.NewButton("Take", func() {
				takeItemFromFloor(itemCopy, floorVBox, currentItemContainer)
			})

			detailsButton := widget.NewButton("Details", func() {
				// TODO display object statistics
			})

			currentItemContainer.Add(nameLabel)
			currentItemContainer.Add(container.NewGridWithColumns(2, takeButton, detailsButton))
			floorVBox.Add(currentItemContainer)
		}
	}
	return floorVBox
}

func RefreshDropdownContent(categoryName, newItem string) {
	for _, dropdown := range itemDropdowns {
		if dropdown.category.Name == categoryName {
			dropdown.widget.Options = append(dropdown.widget.Options, newItem)
			dropdown.widget.Refresh()
			return
		}
	}
}

func takeItemFromFloor(item *model.Object, floorVBox, currentItemContainer *fyne.Container) {
	game.Player.AddObjectToInventory(item, false)

	// Remove object from currentMap ObjectList
	// TODO check error
	_ = game.CurrentMap.FindObjectToRemove(item)

	// Remove current container as well from floor container
	floorVBox.Remove(currentItemContainer)

	// Refresh items in inventory
	RefreshDropdownContent(item.Category, item.Name)
	updateInventoryStatsArea()
}
