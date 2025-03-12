module github.com/dynamicy/go-tetris-ws/tetris-main

go 1.24.1

require (
	github.com/dynamicy/go-tetris-ws/go-tetris v0.0.0
	github.com/hajimehoshi/ebiten/v2 v2.8.6
)

require (
	github.com/ebitengine/gomobile v0.0.0-20250209143333-6071a2a2351c // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/ebitengine/purego v0.8.2 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)

replace github.com/dynamicy/go-tetris-ws/go-tetris => ./../tetris
