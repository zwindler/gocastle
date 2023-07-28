// screens/inventory.go

package screens

import (
	"gocastle/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ShowInventoryScreen is the main function of the inventory screen
func ShowInventoryScreen(window fyne.Window) {
	// Create a container to hold the dropdown lists for each category
	inventoryContainer := container.NewAdaptiveGrid(2)

	// Iterate over each object category and display the dropdown list for the items in that category
	for _, category := range model.CategoryList {
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

		// Add the category container to the main inventory container
		inventoryContainer.Add(categoryContainer)
	}

	// Create a "Back" button to return to the main menu
	backButton := widget.NewButton("Back", func() {
		// gear may have changed, reset all secondary stats
		player.RefreshStats(false)
		ShowGameScreen(window)
	})

	// Create the content container to hold the inventory items and the back button
	content := container.NewVBox(
		inventoryContainer,
		backButton,
	)

	window.SetContent(content)
}
