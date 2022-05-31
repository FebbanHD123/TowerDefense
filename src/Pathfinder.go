package main

type Pathfinder struct {
	level      Level
	openList   []Node
	closedList []Node
}

type Path struct {
	nodes []Node
}

type Node struct {
	gScore              float64
	fScore              float64
	location            Location
	predecessorLocation Location
	hasCameFrom         bool
}

func CreatePathFinder(level Level) Pathfinder {
	return Pathfinder{
		level: level,
	}
}

func (p *Pathfinder) FindPath(fromLocation, toLocation Location) (Path, bool) {
	p.openList = make([]Node, 0)
	p.closedList = make([]Node, 0)
	goalNode := p.GetNode(int(toLocation.x), int(toLocation.y))
	startNode := p.GetNode(int(fromLocation.x), int(fromLocation.y))
	p.openList = append(p.openList, startNode)
	for len(p.openList) > 0 {
		currentNode := p.RemoveMinFromOpenList()
		if currentNode.equals(goalNode) {
			return p.recustructPath(currentNode), true
		}
		p.closedList = append(p.closedList, currentNode)
		p.expandNode(&goalNode, &currentNode)
	}
	return Path{}, false

}

func (p *Pathfinder) expandNode(goalNode, node *Node) {
	for _, neighbor := range node.getNeighbors(*p) {
		if p.IsNodeClosed(neighbor) {
			continue
		}
		tentativG := node.gScore + 1
		if tentativG > neighbor.gScore && neighbor.gScore >= 0 {
			continue
		}
		neighbor.SetCameFrom(*node)
		neighbor.gScore = tentativG
		neighbor.fScore = tentativG + neighbor.GetHScore(*goalNode)
		if !p.IsNodeOpen(neighbor) {
			p.openList = append(p.openList, neighbor)
		}
	}
}

func (p *Pathfinder) recustructPath(currentNode Node) Path {
	var nodes []Node
	var current = currentNode
	nodes = append(nodes, current)
	for {
		current = current.GetCameFrom(*p)

		nodes = append(nodes[:1], nodes[0:]...)
		nodes[0] = current
		if !current.hasCameFrom {
			break
		}
	}
	return Path{nodes: nodes}
}

func (n Node) equals(node Node) bool {
	return n.location.x == node.location.x && n.location.y == node.location.y
}

func (n Node) GetHScore(node Node) float64 {
	distance := node.location.Distance(n.location)
	return distance
}

func (n Node) getNeighbors(pathfinder Pathfinder) []Node {

	var neighbors []Node
	for dx := 0; dx < 3; dx++ {
		for dy := 0; dy < 3; dy++ {
			x := int(n.location.x) + 1 - dx
			y := int(n.location.y) + 1 - dy
			if x >= width || y >= height || x < 0 || y < 0 {
				continue
			}
			node := pathfinder.GetNode(x, y)
			if !pathfinder.IsNodeObstacle(node) && !pathfinder.IsNodeClosed(node) {
				neighbors = append(neighbors, node)
			}
		}
	}

	return neighbors
}

func (p *Pathfinder) GetNode(x, y int) Node {
	node := Node{
		gScore:              -1,
		fScore:              -1,
		location:            CreateLocation(uint16(x), uint16(y)),
		predecessorLocation: CreateLocation(0, 0),
		hasCameFrom:         false,
	}
	for _, n := range p.openList {
		if n.equals(node) {
			return n
		}
	}
	for _, n := range p.closedList {
		if n.equals(node) {
			return n
		}
	}
	return node
}

func (p *Pathfinder) IsNodeClosed(node Node) bool {
	for _, n := range p.closedList {
		if n.equals(node) {
			return true
		}
	}
	return false
}

func (p *Pathfinder) IsNodeOpen(node Node) bool {
	for _, n := range p.openList {
		if n.equals(node) {
			return true
		}
	}
	return false
}

func (p *Pathfinder) IsNodeObstacle(node Node) bool {
	return p.level.GetRegionAtLocation(node.location).Type != RTYPE_PATH
}

func (p *Pathfinder) RemoveMinFromOpenList() Node {
	min := p.openList[0]
	var index int
	for i, node := range p.openList {
		if node.fScore >= 0 && node.fScore < min.fScore {
			min = node
			index = i
		}
	}
	p.openList = append(p.openList[:index], p.openList[index+1:]...)
	return min
}

func (n *Node) SetCameFrom(node Node) {
	n.predecessorLocation = node.location
	n.hasCameFrom = true
}

func (n Node) GetCameFrom(pathfinder Pathfinder) Node {
	return pathfinder.GetNode(int(n.predecessorLocation.x), int(n.predecessorLocation.y))
}
