package main

//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//var Reset = "\033[0m"
//var Red = "\033[31m"
//var Green = "\033[32m"
//var Yellow = "\033[33m"
//var Blue = "\033[34m"
//var Purple = "\033[35m"
//var Cyan = "\033[36m"
//var Gray = "\033[37m"
//var White = "\033[97m"
//

//
//type World [][]bool
//
//func (w World) Display() {
//	for _, row := range w {
//		for _, cell := range row {
//			switch {
//			case cell:
//				//fmt.Printf(greenSquare)
//				fmt.Printf("\033[43m  ") // yellow
//			default:
//				fmt.Printf("\033[44m  ") // blue
//				//fmt.Printf(brownSquare)
//			}
//		}
//		fmt.Printf("\n")
//	}
//}
//
//func (w World) Seed() {
//	for _, row := range w {
//		for i := range row {
//			if rand.Intn(4) == 1 {
//				row[i] = true
//			}
//		}
//	}
//}
//
////func (w World) TestSeed() {
////	for i, row := range w {
////		for j := range row {
////			switch {
////			case i == 1 && j == 1:
////				row[j] = true
////			case i == 2 && j == 2:
////				row[j] = true
////			case i == 1 && j == 2:
////				row[j] = true
////				//case i == 1 && j == 0:
////				//	row[j] = true
////			}
////		}
////	}
////}
//
//func (w World) Alive(x, y int) bool {
//	y = (height + y) % height
//	x = (width + x) % width
//	return w[y][x]
//}
//
//func (w World) Neighbors(x, y int) int {
//	var neighbors int
//
//	for i := y - 1; i <= y+1; i++ {
//		for j := x - 1; j <= x+1; j++ {
//			if i == y && j == x {
//				continue
//			}
//			if w.Alive(j, i) {
//				neighbors++
//			}
//		}
//	}
//	return neighbors
//}
//
//func (w World) Next(x, y int) bool {
//	n := w.Neighbors(x, y)
//	alive := w.Alive(x, y)
//	if n < 4 && n > 1 && alive {
//		return true
//	} else if n == 3 && !alive {
//		return true
//	} else {
//		return false
//	}
//}
//
//func Step(a, b World) {
//	for i := 0; i < height; i++ {
//		for j := 0; j < width; j++ {
//			b[i][j] = a.Next(j, i)
//		}
//	}
//}
//
//func MakeWorld() World {
//	w := make(World, height)
//	for i := range w {
//		w[i] = make([]bool, width)
//	}
//	return w
//}
//
//func main() {
//	fmt.Println(ansiEscapeSeq)
//	rand.Seed(time.Now().UTC().UnixNano())
//	newWorld := MakeWorld()
//	nextWorld := MakeWorld()
//	newWorld.Seed()
//	for {
//		newWorld.Display()
//		Step(newWorld, nextWorld)
//		newWorld, nextWorld = nextWorld, newWorld
//		time.Sleep(sleepIteration * time.Millisecond)
//		fmt.Println(ansiEscapeSeq)
//		//fmt.Println("пурум пум пум")
//		fmt.Printf("\033[0m")
//	}
//	//fmt.Printf("ansiEscapeSeq: '%v'\n", ansiEscapeSeq)
//
//	//fmt.Printf("brownSquare: '%v'\n", brownSquare)
//	//
//	//fmt.Printf("greenSquare: '%v'\n", greenSquare)
//
//	//color.Cyan("Prints text in cyan.")
//	//fmt.Printf("test: '%v'\n", "\033[35m]  ")
//	//fmt.Println("\033[42m  ")
//	//fmt.Println("\033[44m  ") // blue
//	//fmt.Println("\033[43m  ") // yellow
//}
//
////package main
////
////import (
////	"fmt"
////	"math/rand"
////	"time"
////)
////
////const (
////	width  = 80
////	height = 15
////)
////
////// Universe является двухмерным полем клеток.
////type Universe [][]bool
////
////// NewUniverse возвращает пустую вселенную.
////func NewUniverse() Universe {
////	u := make(Universe, height)
////	for i := range u {
////		u[i] = make([]bool, width)
////	}
////	return u
////}
////
////// Seed заполняет вселенную случайными живыми клетками.
////func (u Universe) Seed() {
////	for i := 0; i < (width * height / 4); i++ {
////		u.Set(rand.Intn(width), rand.Intn(height), true)
////	}
////}
////
////// Set устанавливает состояние конкретной клетки.
////func (u Universe) Set(x, y int, b bool) {
////	u[y][x] = b
////}
////
////// Alive сообщает, является ли клетка живой.
////// Если координаты за пределами вселенной, возвращаемся к началу.
////func (u Universe) Alive(x, y int) bool {
////	x = (x + width) % width
////	y = (y + height) % height
////	return u[y][x]
////}
////
////// Neighbors подсчитывает прилегающие живые клетки.
////func (u Universe) Neighbors(x, y int) int {
////	n := 0
////	for v := -1; v <= 1; v++ {
////		for h := -1; h <= 1; h++ {
////			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
////				n++
////			}
////		}
////	}
////	return n
////}
////
////// Next возвращает состояние определенной клетки на следующем шаге.
////func (u Universe) Next(x, y int) bool {
////	n := u.Neighbors(x, y)
////	return n == 3 || n == 2 && u.Alive(x, y)
////}
////
////// String возвращает вселенную как строку
////func (u Universe) String() string {
////	var b byte
////	buf := make([]byte, 0, (width+1)*height)
////
////	for y := 0; y < height; y++ {
////		for x := 0; x < width; x++ {
////			b = ' '
////			if u[y][x] {
////				b = '*'
////			}
////			buf = append(buf, b)
////		}
////		buf = append(buf, '\n')
////	}
////
////	return string(buf)
////}
////
////// Show очищает экран и возвращает вселенную.
////func (u Universe) Show() {
////	fmt.Print("\x0c", u.String())
////}
////
////// Step обновляет состояние следующей вселенной (b) из
////// текущей вселенной (a).
////func Step(a, b Universe) {
////	for y := 0; y < height; y++ {
////		for x := 0; x < width; x++ {
////			b.Set(x, y, a.Next(x, y))
////		}
////	}
////}
////
////func main() {
////	a, b := NewUniverse(), NewUniverse()
////	a.Seed()
////
////	for i := 0; i < 1000; i++ {
////		Step(a, b)
////		a.Show()
////		time.Sleep(time.Second / 30)
////		a, b = b, a // Swap universes
////	}
////}
