# Optimization

You can now try to run this step code

```bash
$ go run . 03
```

You should be able to make two observations:
 1. The window is much wider, and there is a lot more cells.
 2. The FPS and TPS metrics should be lowed (depending on your machine).
 3. In the top right corner of the window, you should see the number of CPU (or core) available on your machine.

This code run much slower than the previous ones because we went from a `500x500` grid to a `1800x900`, or from
`250 000` cells to `1 620 000` cells, about `6.5` times more.

There is a lot of know ways to optimize the Game of Life implementation.
And one of them is to use a one dimensional array, which is already done here.

But, as we are trying to improve our Golang skills, we will use a Golang feature to achieve a good optimization.

## Goal

Make the game run at 60 FPS and 60 TPS (more or less) on a `1800x900` grid.

The number of CPU displayed in the top right corner of the window is a good hint on what to do.