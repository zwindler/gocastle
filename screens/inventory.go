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
	inventoryContainer := container.NewVBox()

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
			// Find the index of the selected item in the player's inventory
			itemIndex := -1
			for i, item := range player.Inventory {
				if item.Name == selected {
					itemIndex = i
					break
				}
			}

			// Check if the item is equipped
			if itemIndex >= 0 {
				if player.Inventory[itemIndex].Equipped {
					// If equipped, un-equip the item
					err := player.UnequipItem(itemIndex)
					if err != nil {
						dialog.ShowError(err, window)
					}
				} else {
					// If not equipped, equip the item
					err := player.EquipItem(itemIndex)
					if err != nil {
						dialog.ShowError(err, window)
					}
				}

				// Refresh the inventory display after equipping/un-equipping the item
				//ShowInventoryScreen(window)
			}
		})

		// Set the selected item in the dropdown list to the first item in the category (if available)
		if len(itemsInCategory) > 0 {
			itemDropdown.SetSelected(itemsInCategory[0])
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
		ShowGameScreen(window)
	})

	// Create the content container to hold the inventory items and the back button
	content := container.NewVBox(
		inventoryContainer,
		backButton,
	)

	window.SetContent(content)
}
