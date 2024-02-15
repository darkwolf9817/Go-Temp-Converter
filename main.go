package Go_Temp_Converter

import (
	"fmt"
	"strconv"
),
	"strconv",

	"fyne.io/fyne/v2",
	"fyne.io/fyne/v2/container",
	"fyne.io/fyne/v2/widget"
	)

func main() {

}

app := fyne.NewApp("Temperature Converter")
window := app.NewWindow("Convert!")

// Define UI Elements
unitDropdown := widget.NewSelect([]string{"Celsius", "Fahrenheit", "Kelvin"}, func(s string) {
	outLabel.SetText("")
})

inEntry := widget.NewEntry()

convertBtn := widget.NewButton("Convert", func() {
	value, err := strconv.ParseFloat(inEntry.Text(), 64)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}
	fromUnit := unitDropdown.Selected
	toUnit := 
})
