// screens/inventory.go

package screens

import (
	"gocastle/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ShowInventoryScreen is the main function of the inventory screen
func ShowInventoryScreen(window fyne.Window) {
	// Load the background image
	backgroundImage := canvas.NewImageFromFile("static/inventory.png")
	backgroundImage.FillMode = canvas.ImageFillContain

	// Create a container to hold the dropdown lists for each category
	inventoryContainerLeft := container.NewVBox()
	inventoryContainerRight := container.NewVBox()

	// Iterate over each object category and display the dropdown list for the items in that category
	for index, category := range model.CategoryList {
		// Create a header label for the category
		categoryLabel := widget.NewLabel(category.Name)

		// Find the items in the player's inventory that belong to the current category
		itemsInCategory := make([]string, 0)
		for _, item := range player.Inventory {
			if item.Category == category.Name {
				itemsInCategory = append(itemsInCategory, item.Name)
			}
		}

		// Create a dropdown list to display the items in the category
		itemDropdown := widget.NewSelect(itemsInCategory, func(selected string) {
			for i, item := range player.Inventory {
				if item.Name == selected {
					player.EquipItem(i)
				} else if player.Inventory[i].Equipped {
					// If another item was equipped, un-equip it
					err := player.UnequipItem(i)
					if err != nil {
						dialog.ShowError(err, window)
					}
				}
			}
		})

		for _, item := range player.Inventory {
			if item.Equipped {
				itemDropdown.SetSelected(item.Name)
				break
			}
		}

		// Create a container to hold the category label and the dropdown list
		categoryContainer := container.NewVBox(
			categoryLabel,
			itemDropdown,
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
		player.RefreshStats(false)
		ShowGameScreen(window)
	})

	// Create the content container to hold the inventory items and the back button
	inventoryContainer := container.NewBorder(nil, nil, inventoryContainerLeft, inventoryContainerRight, backgroundImage)

	floorFakeContainer := widget.NewLabel("The floor once I implement it")
	floorContainer := container.NewVBox(floorFakeContainer)
	floorScroll := container.NewVScroll(floorContainer)
	rightColumn := container.NewBorder(nil, backButton, nil, nil, floorScroll)
	content := container.NewBorder(nil, nil, nil, rightColumn, inventoryContainer)
	window.SetContent(content)
}
