# Chaos

## Intro

In all steps, you'll find two files:
 - `game.go` draw the game and call grid updates,
 - `grid.go` contains the logic to update the grid, this is the real game state.

> 💡 Those two files are partially implemented, and you need to fill in the blanks.

Each step can be run independently, using:

```bash
$ go run . 01
```

Where `01` is the step number.

## Goal

For this first step, we'll just try to draw random black and white pixels on the screen. Order will come later.

 - On each `Update()` call
   - Randomize the grid cells states (alive or dead, `true` or `false`)
   - Update the game pixels state according to the grid state
 - On each `Draw()` call
   - Draw the game pixels

Once all is done, you should see a random pattern of black and white pixels on the screen, which, if you're old enough, 
will remind you of the old TV static noise, AKA snow.

> 💡 Once again, keep an eye in the top left corner of the screen, where the FPS & TPS counters are displayed. They
> should still be around 60 FPS and 60 TPS.

If so, congratulations, you have successfully completed the first step,
and you're ready to move on to [the next one](./../02-order/README.md).