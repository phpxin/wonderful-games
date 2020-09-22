package core

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"wonderful-games/guessing-idiom/tools/log"
)

const (
	Size = 8
)

type StageItem struct {
	Idiom string `json:"idiom"`
	Mean  string `json:"mean"`
	Pos   string `json:"pos"`
	Vis   string `json:"vis"`
}

type Position struct {
	X int32
	Y int32
}

type WordPos struct {
	Word string
	Pos  *Position
}

type Game struct {
	Matrix   [Size][Size]*WordPos
	Idiom    string
	IdiomPos map[string][]*WordPos
}

func NewGame(idiom string) *Game {
	return &Game{
		Idiom:    idiom,
		IdiomPos: make(map[string][]*WordPos),
	}
}

func (game *Game) GetResult() []*StageItem {
	stage := make([]*StageItem, 0)

	hiddenIndexes := make(map[*Position]int)

	for _,wordPoses := range game.IdiomPos {
		invisible := int(rand.Uint32()%4)
		for i,wordPos := range wordPoses {
			if i==invisible {
				hiddenIndexes[wordPos.Pos] = 0
			}
		}
	}

	for idiomName,wordPoses := range game.IdiomPos {
		stageItem := new(StageItem)
		stageItem.Idiom = idiomName
		idiomInfo,_ := Idioms[idiomName]
		stageItem.Mean = idiomInfo.Expl

		var buffer bytes.Buffer

		log.Debug("", "name %s", idiomName)

		visables := make([]string, 4)
		for i,wordPos := range wordPoses {
			log.Debug("", "pos %+v", wordPos)
			if buffer.Len()>0 {
				buffer.WriteByte(';')
			}

			buffer.WriteString(strconv.Itoa(int(wordPos.Pos.X))) // x,y
			buffer.WriteByte(',') // x,y
			buffer.WriteString(strconv.Itoa(int(wordPos.Pos.Y))) // x,y

			if _,ok := hiddenIndexes[wordPos.Pos]; ok {
				visables[i] = "0"
			}else{
				visables[i] = "1"
			}
		}

		stageItem.Pos = buffer.String()

		stageItem.Vis = strings.Join(visables, "")
		stage = append(stage, stageItem)
	}

	return stage
}

func (game *Game) PrintMatrix() {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if game.Matrix[i][j] == nil {
				fmt.Print("口")
				continue
			}
			fmt.Print(game.Matrix[i][j].Word)
		}
		fmt.Println()
	}
}

func (game *Game) PrintPos() {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			var x int32 = -100
			var y int32 = -100
			var w string = ""

			if game.Matrix[j][i] != nil {
				x = game.Matrix[j][i].Pos.X
				y = game.Matrix[j][i].Pos.Y
				w = game.Matrix[j][i].Word
			}

			fmt.Println(i, j, x, y, w)
		}
	}
}

func (game *Game) FindResult(direction bool) { // true -x false |y
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("panic info")
			fmt.Println(x)
			fmt.Println("matrix")
			game.PrintMatrix()
			fmt.Println("positions")
			game.PrintPos()
			fmt.Println("stack tracker")
			panic(x)
		}
	}()

	words := strings.Split(game.Idiom, "")
	wps := make([]*WordPos, 0)
	for k, w := range words {
		wp := &WordPos{
			Word: w,
			Pos: &Position{
				X: 0,
				Y: 0,
			},
		}
		if direction {
			wp.Pos.X = int32(k)
			game.Matrix[0][k] = wp
		} else {
			wp.Pos.Y = int32(k)
			game.Matrix[k][0] = wp
		}

		wps = append(wps, wp)
	}



	game.IdiomPos[game.Idiom] = wps

	for k, wp := range wps {
		if game.dfs(int32(k), wp, !direction, 1) {
			break
		}
	}
}

func (game *Game) movingMatrix(x, y int) bool {
	if x > 0 {
		// right
		// check
		for i := 0; i < Size; i++ {
			for m := 1; m <= x; m++ {
				if game.Matrix[i][Size-m] != nil {
					// The grid has element
					return false
				}
			}
		}

		// move
		for i := 0; i < Size; i++ {
			for j := Size - 1; j >= x; j-- {

				if game.Matrix[i][j-x] != nil {
					game.Matrix[i][j-x].Pos.X = int32(j)
					game.Matrix[i][j-x].Pos.Y = int32(i)
				}
				game.Matrix[i][j] = game.Matrix[i][j-x]
				game.Matrix[i][j-x] = nil
			}
		}
	} else if x < 0 {
		x = int(math.Abs(float64(x)))
		// left
		// check
		for i := 0; i < Size; i++ {
			for m := 0; m < x; m++ {
				if game.Matrix[i][m] != nil {
					// The grid has element
					return false
				}
			}
		}

		// move
		for i := 0; i < Size; i++ {
			for j := x; j < Size; j++ {

				if game.Matrix[i][j] != nil {
					game.Matrix[i][j].Pos.X = int32(j - x)
					game.Matrix[i][j].Pos.Y = int32(i)
				}
				game.Matrix[i][j-x] = game.Matrix[i][j]
				game.Matrix[i][j] = nil
			}
		}
	}

	if y > 0 {
		// down
		// check
		for i := 0; i < Size; i++ {
			for m := 1; m <= y; m++ {
				if game.Matrix[Size-m][i] != nil {
					// The grid has element
					return false
				}
			}
		}

		// move
		for i := 0; i < Size; i++ {
			for j := Size - 1; j >= y; j-- {

				if game.Matrix[j-y][i] != nil {
					game.Matrix[j-y][i].Pos.X = int32(i)
					game.Matrix[j-y][i].Pos.Y = int32(j)
				}

				game.Matrix[j][i] = game.Matrix[j-y][i]
				if game.Matrix[j][i] != nil {
				}

				game.Matrix[j-y][i] = nil
			}
		}
	} else if y < 0 {
		y = int(math.Abs(float64(y)))
		// up
		// check
		for i := 0; i < Size; i++ {
			for m := 0; m < y; m++ {
				if game.Matrix[m][i] != nil {
					// The grid has element
					return false
				}
			}
		}

		// move
		for i := 0; i < Size; i++ {
			for j := y; j < Size; j++ {

				//tmp := game.Matrix[j][i]
				if game.Matrix[j][i] != nil {
					game.Matrix[j][i].Pos.X = int32(i)
					game.Matrix[j][i].Pos.Y = int32(j - y)
				}
				game.Matrix[j-y][i] = game.Matrix[j][i]
				game.Matrix[j][i] = nil
			}
		}
	}

	return true
}

// depth first search
func (game *Game) dfs(wpIndex int32, wordPos *WordPos, direction bool, counter int32) bool {
	if len(game.IdiomPos) >= 8 {
		return true
	}
	if counter >= 20 {
		return true
	}
	associatedIdioms := InvertedIndex[wordPos.Word]
	for _, ai := range associatedIdioms {
		_, ok := game.IdiomPos[ai]
		if ok {
			continue
		}
		aiws := strings.Split(ai, "")
		var posIndex int32 = 0
		for i, aiw := range aiws {
			if aiw == wordPos.Word {
				posIndex = int32(i) // confirm the index of the associated word of the idiom
				break
			}
		}
		wps := make([]*WordPos, 0)
		if direction {
			// -
			// does it cover other word that exists in the matrix but not exists in the idiom
			var coverOtherIdiom bool = false
			var xOutOfBound int32 = 0
			for i, aiw := range aiws {
				aiw_y := wordPos.Pos.Y
				aiw_x := wordPos.Pos.X - posIndex + int32(i)
				if aiw_x < 0 {
					if aiw_x < xOutOfBound {
						xOutOfBound = aiw_x
					}
					continue // out of bound
				}
				if aiw_x >= Size {
					if aiw_x-(Size-1) > xOutOfBound {
						xOutOfBound = aiw_x - (Size - 1)
					}
					continue
				}
				if game.Matrix[aiw_y][aiw_x] != nil && game.Matrix[aiw_y][aiw_x].Word != aiw {
					coverOtherIdiom = true
					break // it cover other word that exists in the matrix but not exists in the idiom
				}
			}

			if coverOtherIdiom {
				continue
			}

			// does it make matrix out of bound
			if xOutOfBound != 0 {
				if !game.movingMatrix(int(xOutOfBound)*-1, 0) {
					continue // continue to next idiom
				}
			}

			// insert word to slot
			for i, aiw := range aiws {
				aiw_y := wordPos.Pos.Y
				aiw_x := wordPos.Pos.X - posIndex + int32(i)
				wp := &WordPos{
					Word: aiw,
					Pos: &Position{
						X: aiw_x,
						Y: aiw_y,
					},
				}
				if game.Matrix[aiw_y][aiw_x] != nil && game.Matrix[aiw_y][aiw_x].Word == wp.Word {
					// 已存在，不要覆盖它，保留原 WordPos 指针，因为回溯的时候如果被修改，下面的程序就会由于指针被修改，找不到准确的值
					wps = append(wps, game.Matrix[aiw_y][aiw_x])
					continue
				}

				game.Matrix[aiw_y][aiw_x] = wp
				wps = append(wps, wp)

			}



		} else {
			// |
			// does it cover other word that exists in the matrix but not exists in the idiom
			var coverOtherIdiom bool = false
			var yOutOfBound int32 = 0
			for i, aiw := range aiws {
				aiw_x := wordPos.Pos.X
				aiw_y := wordPos.Pos.Y - posIndex + int32(i)
				if aiw_y < 0 {
					if aiw_y < yOutOfBound {
						yOutOfBound = aiw_y
					}
					continue // out of bound
				}
				if aiw_y >= Size {
					if aiw_y-(Size-1) > yOutOfBound {
						yOutOfBound = aiw_y - (Size - 1)
					}
					continue
				}
				if game.Matrix[aiw_y][aiw_x] != nil && game.Matrix[aiw_y][aiw_x].Word != aiw {
					coverOtherIdiom = true
					break // it cover other word that exists in the matrix but not exists in the idiom
				}

			}

			if coverOtherIdiom {
				continue
			}

			// does it make matrix out of bound
			if yOutOfBound != 0 {
				// @todo moving matrix
				if !game.movingMatrix(0, int(yOutOfBound)*-1) {
					continue // continue to next idiom
				}
			}

			// insert word to slot
			for i, aiw := range aiws {
				aiw_x := wordPos.Pos.X
				aiw_y := wordPos.Pos.Y - posIndex + int32(i)
				wp := &WordPos{
					Word: aiw,
					Pos: &Position{
						X: aiw_x,
						Y: aiw_y,
					},
				}

				if game.Matrix[aiw_y][aiw_x] != nil && game.Matrix[aiw_y][aiw_x].Word == wp.Word {
					// 已存在，不要覆盖它，保留原 WordPos 指针，因为回溯的时候如果被修改，下面的程序就会由于指针被修改，找不到准确的值
					wps = append(wps, game.Matrix[aiw_y][aiw_x])
					continue
				}
				game.Matrix[aiw_y][aiw_x] = wp
				wps = append(wps, wp)
			}

		}
		//break
		log.Debug("", "%s pos len %d", ai, len(wps))
		game.IdiomPos[ai] = wps

		for k, wp := range wps {
			if game.dfs(int32(k), wp, !direction, counter+1) {
				break
			}
		}
	}
	return true
}
