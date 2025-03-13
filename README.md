# 🎮 Go-Tetris

A **modern Tetris game** built with **Go** using the **Ebiten game engine**.

## 🚀 Features
- ✅ Smooth Tetromino movement & rotation
- ✅ Wall Kick System (SRS)
- ✅ Hard Drop (SPACE Key)
- ✅ Scoring System
- ✅ Game Speed Increases Over Time
- 🚧 Upcoming: Hold & Swap (SHIFT Key),  & Leveling

---
## 📂 **Project Structure**
This project follows a **Go workspace** structure:
```
go-tetris-ws/
│── tetris/          # Core Tetris game logic (library)
│   ├── board.go
│   ├── config.go
│   ├── game.go
│   ├── input.go
│   ├── shapes.go
│   ├── tetromino.go
│   ├── ui.go
│   ├── go.mod
│── tetris-main/      # Main entry point
│   ├── go.mod
│   ├── main.go
│── go.work           # Go workspace file
│── README.md
│── LICENSE
│── .gitignore
```
---

## 🚀 **Installation & Running**
### **1️⃣ Clone the Repository**
```sh
$ git clone https://github.com/dynamicy/go-tetris.git
$ cd go-tetris/go-tetris-ws
```

### 2️⃣ Setup Go Workspace
```sh
$ go work use ./tetris ./tetris-main
```
### 3️⃣ Run the Game
```sh
$ cd tetris-main
$ go run .
```

## 📦 Using Go-Tetris as a Module
#### You can import the core Tetris game logic into other projects:
```sh
$ go get github.com/dynamicy/go-tetris-ws/tetris
```
#### Then, use it in your own Go project:
```go
import "github.com/dynamicy/go-tetris-ws/tetris"
```

## 🕹️ Controls
| Key     | Action                           |
|---------|----------------------------------|
| ← / →   | Move left / right               |
| ↓       | Soft Drop                        |
| Space   | Hard Drop                        |
| Z / X   | Rotate CCW / CW                  |
| Shift   | Hold Tetromino *(Upcoming Feature)* |

## 📢 How to Contribute
Want to improve Go-Tetris? Feel free to fork the repository, make changes, and submit a Pull Request!
