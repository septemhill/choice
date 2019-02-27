package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type ConnectorType int32

type Connector struct {
	g *Grid
	//ct        ConnectorType
	ctorp     *Connector
	connected bool
}

func (c *Connector) connect(ctorp *Connector) error {
	//if c.ct == ctorp.ct {
	//Pointer to each other
	c.ctorp = ctorp
	ctorp.ctorp = c

	//Update status
	c.connected = true
	ctorp.connected = true

	//Update grid connector space
	c.g.Space--
	ctorp.g.Space--

	return nil
	//}

	//return errors.New("connector type not match")
}

type Grid struct {
	Name       string
	Connectors [GRID_DIMENSION]*Connector
	Space      int
}

//const (
//	CONN_TYPE_NONE ConnectorType = iota
//	CONN_TYPE_1
//	CONN_TYPE_2
//	CONN_TYPE_3
//	CONN_TYPE_4
//	CONN_TYPE_5
//	CONN_TYPE_6
//	CONN_TYPE_MAX
//)

const (
	GRID_DIMENSION = 4
)

func (g *Grid) findFirstFreeConnector() *Connector {
	for i := 0; i < GRID_DIMENSION; i++ {
		if !g.Connectors[i].connected {
			return g.Connectors[i]
		}
	}

	return nil
}

func NewGrid(name string) *Grid {
	rand.Seed(time.Now().UnixNano())
	grid := &Grid{
		Name:  name,
		Space: GRID_DIMENSION,
	}

	grid.Connectors = [GRID_DIMENSION]*Connector{
		&Connector{g: grid /*ct: ConnectorType(rand.Int31n(int32(CONN_TYPE_MAX)))*/},
		&Connector{g: grid /*ct: ConnectorType(rand.Int31n(int32(CONN_TYPE_MAX)))*/},
		&Connector{g: grid /*ct: ConnectorType(rand.Int31n(int32(CONN_TYPE_MAX)))*/},
		&Connector{g: grid /*ct: ConnectorType(rand.Int31n(int32(CONN_TYPE_MAX)))*/},
	}

	return grid
}

func traverse(g *Grid, ctorp *Connector, m map[*Grid]struct{}) {
	_, ok := m[g]
	if g != nil && !ok {
		fmt.Println(g.Name, g.Space)
		m[g] = struct{}{}
		for i := 0; i < GRID_DIMENSION; i++ {
			if g.Connectors[i] != ctorp && g.Connectors[i].connected {
				grid := g.Connectors[i].ctorp.g
				inConn := g.Connectors[i].ctorp
				traverse(grid, inConn, m)
			}
		}
	}
}

func Traverse(g *Grid, ctorp *Connector) {
	m := make(map[*Grid]struct{})
	traverse(g, ctorp, m)
}

func createTreePath(gridCount int, in, out, noSpace *[]*Grid) *Grid {
	var root *Grid
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		*out = append(*out, NewGrid(strconv.FormatInt(int64(i), 10)))
	}

	for i := 0; i < gridCount; i++ {
		inMapLen := len(*in)
		outMapLen := len(*out)

		if inMapLen == 0 {
			ornd := rand.Intn(outMapLen)
			grid := (*out)[ornd]
			root = grid
			*out = append((*out)[:ornd], (*out)[ornd+1:]...)
			*in = append(*in, grid)
		} else {
			irnd, ornd := rand.Intn(inMapLen), rand.Intn(outMapLen)
			igrid := (*in)[irnd]
			ogrid := (*out)[ornd]

			for i := 0; i < len(igrid.Connectors); i++ {
				if !igrid.Connectors[i].connected {
					igrid.Connectors[i].connect(ogrid.Connectors[i])
					break
				}
			}

			if igrid.Space == 0 {
				*noSpace = append(*noSpace, igrid)
				*in = append((*in)[:irnd], (*in)[irnd+1:]...)
			}

			*out = append((*out)[:ornd], (*out)[ornd+1:]...)
			*in = append(*in, ogrid)
		}
	}

	return root
}

func makeTreeCycle(gridCount int, in, noSpace *[]*Grid) {
	rand.Seed(time.Now().UnixNano())

	freeConnectorCount := (gridCount * 4) - (gridCount-1)*2
	connectPairCount := freeConnectorCount * 25 / 100

	for i := 0; i < connectPairCount; i++ {
		inLen := len(*in)
		rand.Shuffle(inLen, func(ia, ja int) {
			(*in)[ia], (*in)[ja] = (*in)[ja], (*in)[ia]
		})

		firstGrid := (*in)[0]
		secondGrid := (*in)[1]

		firstGrid.findFirstFreeConnector().connect(secondGrid.findFirstFreeConnector())

		if secondGrid.Space == 0 {
			*noSpace = append(*noSpace, secondGrid)
			*in = append((*in)[:1], (*in)[2:]...)
		}

		if firstGrid.Space == 0 {
			*noSpace = append(*noSpace, firstGrid)
			*in = append((*in)[1:])
		}
	}
}

func makeExitPoint(root *Grid, in, noSpace []*Grid) *Grid {
	rand.Seed(time.Now().UnixNano())
	grids := append(in, noSpace...)

	for {
		rand.Shuffle(len(grids), func(i, j int) {
			grids[i], grids[j] = grids[j], grids[i]
		})

		if root != grids[len(grids)-1] {
			return grids[len(grids)-1]
		}
	}
}

func CreateRandomMap() *RandomMap {
	var root, exit *Grid

	outMapGrids := make([]*Grid, 0)
	inMapGrids := make([]*Grid, 0)
	noSpaceGrids := make([]*Grid, 0)

	root = createTreePath(100, &inMapGrids, &outMapGrids, &noSpaceGrids)
	makeTreeCycle(100, &inMapGrids, &noSpaceGrids)

	exit = makeExitPoint(root, inMapGrids, noSpaceGrids)

	return &RandomMap{root: root, exit: exit}
}

type RandomMap struct {
	root        *Grid
	exit        *Grid
	currPostion *Grid
	team        *Team
}

func (m *RandomMap) Enter(t *Team) {
	m.team = t
	m.currPostion = m.root
}

func (m *RandomMap) Walk() {
	type wayoption struct {
		option     int
		optionName string
		grid       *Grid
	}

	fmt.Printf("root: %s, exit: %s\n", m.root.Name, m.exit.Name)
	for {
		if m.currPostion == m.exit {
			fmt.Println("wowowowowowowowowowowowowowowow, u DONE")
			break
		}

		options := make([]wayoption, 0)
		idx := 1

		for i := 0; i < len(m.currPostion.Connectors); i++ {
			if m.currPostion.Connectors[i].connected {
				options = append(options, wayoption{option: idx, optionName: m.currPostion.Connectors[i].ctorp.g.Name, grid: m.currPostion.Connectors[i].ctorp.g})
				idx++
			}
		}

		for i := 0; i < len(options); i++ {
			fmt.Printf("[%d]. %s\n", options[i].option, options[i].optionName)
		}

		choice := readUserChoice()

		if (choice > 0) && (choice <= len(options)) {
			m.currPostion = options[choice-1].grid
		}
	}
}
