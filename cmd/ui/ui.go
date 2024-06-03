package ui

import (
	"log/slog"
	"typrfr/cmd/processor"

	"github.com/rivo/tview"
)

func Render(game *processor.Game, app *tview.Application) {

	go func() {
		for gameState := range game.State {
			switch gameState {
			case processor.NOT_STARTED:
				button := tview.NewButton("Hit enter to start the game").SetSelectedFunc(func() {
					game.StartGame()
				})

				if err := app.SetRoot(button, false).EnableMouse(true).Run(); err != nil {

					slog.Error("some error")
					panic(err)
				}

				break

			case processor.IN_PROGRESS:
				textView := tview.NewTextView().SetText(game.Sentence)
				container := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(textView, 0, 1, false)

				if err := app.SetRoot(container, true).EnableMouse(true).Run(); err != nil {
					panic(err)
				}
				break
			case processor.FINISHED:
				slog.Info("Game ended", "state", game.HasFinished())
				break
			default:
			}

		}

	}()
}
