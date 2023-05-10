package main

import (
 	"fmt"
 	"os"

 	tea "github.com/charmbracelet/bubbletea"
	dgf "github.com/snowel/dgf"
)


type BoardEditor struct {
	Board []byte
	Record []int
	Mouse tea.MouseEvent
	Curse int //intersection on board mouse cursor
}

var Starpoints = []int{61, 175, 289, 295, 181, 67, 73, 187, 301}

// For a set of coordinates, give the DGF intersectoins ID
// Negative value means it's not a valid one.
func TermCoordToIntersect(x int, y int, xOff int, yOff int) int {
  yCoor := y - yOff + 1

  xTemp := x - xOff + 1
  var xCoor int
  if xTemp % 2 != 0 {
		xCoor = (xTemp + 1 ) / 2
  } else {
		return -1
  }

  return xCoor + (yCoor - 1) * 19
}

func (m BoardEditor) PlaceStone(inter int) {

 	dgf.ApplyMove(m.Board, inter)
	//m.Board[inter+dgf.BOff] = dgf.GetPlayer(m.Board)
	//dgf.SwitchPlayer(m.Board)
 } 

func InitialBoard(board string) BoardEditor {
	if board == "" {
		return BoardEditor{
		  Board: []byte{1, 0, 5, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}} 
	}
  return BoardEditor{ Board: []byte{12, 0, 0, 0, 0, 0, 2, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} } 
}

func (m BoardEditor) Init() tea.Cmd{
	return nil
}

func (m BoardEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {

  case tea.KeyMsg:

	 switch msg.String() {
	 case "ctrl+c", "q":
		return m, tea.Quit
	 }

  case tea.MouseMsg:
	 m.Mouse = tea.MouseEvent(msg)
	 inter := TermCoordToIntersect(m.Mouse.X, m.Mouse.Y, 0, 0)

	 if inter > 0 {
		m.Curse = inter
		if m.Mouse.Type == tea.MouseRelease {
		  m.PlaceStone(inter)
		}
	 } else {
		m.Curse = 0
	 }
  }

  return m, nil
}

func (m BoardEditor) View() string {
    var s string

    for index, v := range m.Board {
				if index <= dgf.BOff {continue}
				i := index - dgf.BOff
				if i == m.Curse && v == 0 {
					  switch dgf.GetPlayer(m.Board){
						case 1:
						 	s += fmt.Sprintf("󰸳 ")
						case 2:
							s += fmt.Sprintf("󰸶 ")}
					  if i % 19 == 0 {
						 s += fmt.Sprintf("\n")
					  }
					  continue
				}	
				if v == 0 {
					if dgf.Contains(i, Starpoints){
					  s += fmt.Sprintf("󰧞 ")
					} else {
					  s += fmt.Sprintf("󰧟 ")
					}
				}
				if v == 1 {
					  s += fmt.Sprintf(" ")
				}
				if v == 2 {
					  s += fmt.Sprintf(" ")
				}

				if i % 19 == 0 {
					s += fmt.Sprintf("\n")
				 }
		}

	var colIcon string
	switch dgf.GetPlayer(m.Board){
	case 1:
		colIcon = ""
	case 2:
		colIcon = ""
	  }
	 s += fmt.Sprintf("\nPress q to quit.\n Player is %s", colIcon)

    // Send the UI for rendering
    return s
}

func main() {
	 
	if len(os.Args) > 1 {
var baseBoard = []byte{0b00001001, 0, 5, 19,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
						 }
	dgf.NfPrintBoard(baseBoard)
	var cmd string


	fmt.Scanln(&cmd)
	fmt.Println(cmd)
	for cmd != "q" {
		switch cmd {
		case "m":
			var (
				col int
			 	row int
			 )
			 fmt.Println("Please input a move using japanese coordiantes. i.e. uper left startpoint is '3 3' ")
			fmt.Scanf("%d %d", &col, &row)
			move := dgf.JCoordToMove(col, row, 19)
			fmt.Println("You've player the move: ", move, col, row)
			dgf.ApplyMove(baseBoard, move)
			dgf.NfPrintBoard(baseBoard)
		}
	//fmt.Scanln(&cmd)
	}

	fmt.Println("Bye for now...")

	} else {
  p := tea.NewProgram(InitialBoard(""), tea.WithAltScreen(), tea.WithMouseAllMotion()) 
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }

  }
}

