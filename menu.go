package main

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/colornames"
)

func CreateMenu(main *Main) *widget.Container {
	font := DefaultFont()

	singlePlayer := widget.NewButton(
		widget.ButtonOpts.TextLabel("Single Player"),
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
				main.StartSinglePlayer("Test", ARROWS)
			},
		),
	)

	oneVone := widget.NewButton(
		widget.ButtonOpts.TextLabel("1 VS 1"),
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
				main.StartTwoPlayers("Test1", "Test2", WASD, ARROWS)
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

	menuContainer.AddChild(singlePlayer)
	menuContainer.AddChild(oneVone)

	root.AddChild(menuContainer)
	return root
}
