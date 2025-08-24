package main

import (
	"strconv"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/colornames"
)

func CreateGameOver(main *Main, gameOverText string, scores map[string]int) *widget.Container {
	font := DefaultFont()

	endText := widget.NewText(
		widget.TextOpts.Text(gameOverText, &font, colornames.Forestgreen),
		widget.TextOpts.Position(
			widget.TextPositionCenter,
			widget.TextPositionCenter,
		),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
	)

	var scoreString string
	for player, score := range scores {
		scoreString += player + ": " + strconv.Itoa(score) + "   "
	}

	scoresText := widget.NewText(
		widget.TextOpts.Text(scoreString, &font, colornames.Forestgreen),
		widget.TextOpts.Position(
			widget.TextPositionCenter,
			widget.TextPositionCenter,
		),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
	)

	backButton := widget.NewButton(
		widget.ButtonOpts.TextLabel("Back To Menu"),
		widget.ButtonOpts.TextFace(&font),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:    colornames.Gainsboro,
			Hover:   colornames.Gainsboro,
			Pressed: colornames.Gainsboro,
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewNineSliceColor(colornames.Forestgreen),
			Hover:   image.NewNineSliceColor(Mix(colornames.Forestgreen, colornames.Mediumseagreen, 0.4)),
			Pressed: image.NewNineSliceColor(Mix(colornames.Forestgreen, colornames.Black, 0.4)),
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  false,
			}),
			widget.WidgetOpts.MinSize(180, 48),
		),
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				main.state = MAIN_MENU
				main.ui.Container = CreateMenu(main)
			},
		),
	)

	menuContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(
				widget.DirectionVertical,
			),
			widget.RowLayoutOpts.Spacing(25),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(50, 50),
		),
	)

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Darkolivegreen),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	menuContainer.AddChild(scoresText)
	menuContainer.AddChild(endText)
	menuContainer.AddChild(backButton)

	root.AddChild(menuContainer)
	return root
}
