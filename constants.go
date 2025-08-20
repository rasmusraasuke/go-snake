package main

const GRID_SIZE = 15 // Tiles
const TILE_SIZE = 40 // Pixels

const backgroundSize = 225 // Pixels
const cherrySize = 360     // Pixels

const tileXScale = float64(TILE_SIZE) / float64(backgroundSize)
const tileYScale = float64(TILE_SIZE) / float64(backgroundSize)
const cherryXScale = float64(TILE_SIZE) / float64(cherrySize)
const cherryYScale = float64(TILE_SIZE) / float64(cherrySize)
