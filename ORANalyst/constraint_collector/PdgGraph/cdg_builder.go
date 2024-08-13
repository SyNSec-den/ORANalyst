package PdgGraph

import (
	"fmt"
	"sort"
	"strconv"
)

type Node struct {
	dom   domInfo
	Index int
	Preds []*Node
	Succs []*Node
}

func NewNode(label int) *Node {
	return &Node{
		Index: label,
		dom:   domInfo{},
	}
}

func (a *Node) isDescendant(b *Node) bool {
	idxB := b.Index
	if a.dom.ancestorBitVector[idxB+1] == 1 {
		return true
	}
	return false
}

func (g *PostDomGraph) getNodeWithIdx(idx int) *Node {
	n, ok := g.BlocksMap[idx]
	if !ok {
		return nil
	}
	return n
}

func (a *Node) isControlDependent(b *Node) bool {
	idxB := b.Index
	if a.dom.controlDepBitVector[idxB+1] == 1 {
		return true
	}
	return false
}

type PostDomGraph struct {
	BlocksMap map[int]*Node
	Blocks    []*Node
	Edges     []*PostDomEdge
}

type PostDomEdge struct {
	from *Node
	to   *Node
}

func NewPostDomGraph(idxStart int, idxEnd int, reverse bool) *PostDomGraph {
	//from idxStart to idxEnd - 1
	blocksMap := make(map[int]*Node)
	var blocks []*Node
	if reverse == false {
		for i := idxStart; i < idxEnd; i++ {
			blocksMap[i] = NewNode(i)
			blocks = append(blocks, blocksMap[i])
		}
	} else {
		for i := idxEnd - 1; i >= idxStart; i-- {
			blocksMap[i] = NewNode(i)
			blocks = append(blocks, blocksMap[i])
		}
	}

	f := &PostDomGraph{
		BlocksMap: blocksMap,
		Blocks:    blocks,
	}

	for _, b := range f.Blocks {
		b.dom = domInfo{}
		b.dom.idom = b
	}

	return f
}

func (g *PostDomGraph) AddEdge(from int, to int) {
	_, ok := g.BlocksMap[from]
	if !ok {
		g.BlocksMap[from] = NewNode(from)
		g.Blocks = append(g.Blocks, g.BlocksMap[from])
	}

	_, ok = g.BlocksMap[to]
	if !ok {
		g.BlocksMap[to] = NewNode(to)
		g.Blocks = append(g.Blocks, g.BlocksMap[to])
	}

	g.BlocksMap[from].Succs = append(g.BlocksMap[from].Succs, g.BlocksMap[to])
	g.BlocksMap[to].Preds = append(g.BlocksMap[to].Preds, g.BlocksMap[from])

	g.Edges = append(g.Edges, &PostDomEdge{
		from: g.BlocksMap[from],
		to:   g.BlocksMap[to],
	})

}

func (a *Node) Idom() *Node { return a.dom.idom }

func (a *Node) Dominees() []*Node { return a.dom.children }

func (a *Node) Dominates(c *Node) bool {
	return a.dom.pre <= c.dom.pre && c.dom.post <= a.dom.post
}

type byDomPreorder []*Node

func (a byDomPreorder) Len() int           { return len(a) }
func (a byDomPreorder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDomPreorder) Less(i, j int) bool { return a[i].dom.pre < a[j].dom.pre }

// DomPreorder returns a new slice containing the blocks of f in
// dominator tree preorder.
func (f *PostDomGraph) DomPreorder() []*Node {
	n := len(f.Blocks)
	order := make(byDomPreorder, n)
	copy(order, f.Blocks)
	sort.Sort(order)
	return order
}

// domInfo contains a BasicBlock's dominance information.
type domInfo struct {
	idom      *Node   // immediate dominator (parent in domtree)
	children  []*Node // nodes immediately dominated by this one
	pre, post int32   // pre- and post-order numbering within domtree

	ancestors         []*Node //all ancestors of node in domtree, EXCLUDING itself
	ancestorBitVector []int

	controlDepBlocks    []*Node
	controlDepBitVector []int
}

// ltState holds the working state for Lengauer-Tarjan algorithm
// (during which domInfo.pre is repurposed for CFG DFS preorder number).
type ltState struct {
	// Each slice is indexed by b.Index.
	sdom     []*Node // b's semidominator
	parent   []*Node // b's parent in DFS traversal of CFG
	ancestor []*Node // b's ancestor with least sdom
}

// dfs implements the depth-first search part of the LT algorithm.
func (lt *ltState) dfs(v *Node, i int32, preorder []*Node) int32 {
	preorder[i] = v
	v.dom.pre = i // For now: DFS preorder of spanning tree of CFG
	i++
	lt.sdom[v.Index] = v
	lt.link(nil, v)
	for _, w := range v.Succs {
		if lt.sdom[w.Index] == nil {
			lt.parent[w.Index] = v
			i = lt.dfs(w, i, preorder)
		}
	}
	return i
}

// eval implements the EVAL part of the LT algorithm.
func (lt *ltState) eval(v *Node) *Node {
	// TODO(adonovan): opt: do path compression per simple LT.
	u := v
	for ; lt.ancestor[v.Index] != nil; v = lt.ancestor[v.Index] {
		if lt.sdom[v.Index].dom.pre < lt.sdom[u.Index].dom.pre {
			u = v
		}
	}
	return u
}

// link implements the LINK part of the LT algorithm.
func (lt *ltState) link(v, w *Node) {
	lt.ancestor[w.Index] = v
}

// buildDomTree computes the dominator tree of f using the LT algorithm.
// Precondition: all blocks are reachable (e.g. optimizeBlocks has been run).
func buildDomTree(f *PostDomGraph) bool {
	// The step numbers refer to the original LT paper; the
	// reordering is due to Georgiadis.

	// Clear any previous domInfo.
	// fmt.Printf("D")

	for _, b := range f.Blocks {
		b.dom = domInfo{}
		b.dom.idom = b
	}

	// fmt.Printf("E")

	n := len(f.Blocks)
	// Allocate space for 5 contiguous [n]*Node arrays:
	// sdom, parent, ancestor, preorder, buckets.
	space := make([]*Node, 5*n)
	lt := ltState{
		sdom:     space[0:n],
		parent:   space[n : 2*n],
		ancestor: space[2*n : 3*n],
	}

	// fmt.Printf("F")

	// Step 1.  Number vertices by depth-first preorder.
	preorder := space[3*n : 4*n]
	root := f.Blocks[0]
	lt.dfs(root, 0, preorder)

	buckets := space[4*n : 5*n]
	copy(buckets, preorder)
	//fmt.Printf("here")
	// In reverse preorder...
	//for i := int32(n) - 1; i >= 0; i-- {
	//	w := preorder[i]
	//	if w != nil {
	//		fmt.Printf("Preorder %d is %d\n", i, w.Index)
	//	} else {
	//		fmt.Printf("Preorder %d is nil\n", i)
	//	}
	//}
	for i := int32(n) - 1; i > 0; i-- {
		w := preorder[i]
		if w == nil {
			return true
			continue // TODO: why this can be nil?
		}
		// Step 3. Implicitly define the immediate dominator of each node.
		for v := buckets[i]; v != w; v = buckets[v.dom.pre] {
			u := lt.eval(v)
			if lt.sdom[u.Index].dom.pre < i {
				v.dom.idom = u
			} else {
				v.dom.idom = w
			}
		}

		// Step 2. Compute the semidominators of all nodes.
		lt.sdom[w.Index] = lt.parent[w.Index]
		for _, v := range w.Preds {
			u := lt.eval(v)
			if lt.sdom[u.Index] == nil {
				fmt.Printf("lt.sdom[u.Index] is nil %d\n", u.Index)
				continue
			}
			if lt.sdom[w.Index] == nil {
				fmt.Printf("lt.sdom[w.Index] is nil %d\n", w.Index)
				continue
			}
			if lt.sdom[u.Index].dom.pre < lt.sdom[w.Index].dom.pre {
				lt.sdom[w.Index] = lt.sdom[u.Index]
			}
		}

		lt.link(lt.parent[w.Index], w)

		if lt.parent[w.Index] == lt.sdom[w.Index] {
			w.dom.idom = lt.parent[w.Index]
		} else {
			buckets[i] = buckets[lt.sdom[w.Index].dom.pre]
			buckets[lt.sdom[w.Index].dom.pre] = w
		}
	}

	// The final 'Step 3' is now outside the loop.
	for v := buckets[0]; v != root; v = buckets[v.dom.pre] {
		v.dom.idom = root
	}

	// Step 4. Explicitly define the immediate dominator of each
	// node, in preorder.
	for _, w := range preorder[1:] {
		if w == root {
			w.dom.idom = nil
		} else {
			if w.dom.idom != lt.sdom[w.Index] {
				w.dom.idom = w.dom.idom.dom.idom
			}
			// Calculate Children relation as inverse of Idom.
			w.dom.idom.dom.children = append(w.dom.idom.dom.children, w)
		}
	}

	numberDomTree(root, 0, 0)

	// printDomTreeDot(os.Stderr, f)        // debugging
	// printDomTreeText(os.Stderr, root, 0) // debugging
	return false
}

// numberDomTree sets the pre- and post-order numbers of a depth-first
// traversal of the dominator tree rooted at v.  These are used to
// answer dominance queries in constant time.
func numberDomTree(v *Node, pre, post int32) (int32, int32) {
	v.dom.pre = pre
	pre++
	for _, child := range v.dom.children {
		pre, post = numberDomTree(child, pre, post)
	}
	v.dom.post = post
	post++
	return pre, post
}

func populateAncestors(graph *PostDomGraph) {

	for _, node := range graph.BlocksMap {
		space := make([]int, len(graph.BlocksMap)+1)
		currNode := node
		for currNode.dom.idom != nil {
			if currNode.dom.idom.Index == currNode.Index {
				break
			}
			node.dom.ancestors = append(node.dom.ancestors, currNode.dom.idom)
			space[currNode.dom.idom.Index+1] = 1
			currNode = currNode.dom.idom
		}
		node.dom.ancestorBitVector = space
	}

}

func BuildControlDepTree(graph *PostDomGraph) {
	var S []*PostDomEdge
	for _, edge := range graph.Edges {
		A, B := edge.from, edge.to
		if B.isDescendant(A) == false {
			S = append(S, edge)
		}
	}

	for _, edge := range S {
		A, B := edge.from, edge.to
		lca := A
		if B.isDescendant(A) == false {
			lca = A.dom.idom
		}

		// fmt.Printf("%d %d %d\n", A.Index, B.Index, lca.Index)

		if lca == A {
			currNode := B
			for currNode != lca {
				//fmt.Printf("%d is control dependent on %d\n", currNode.Index, A.Index)

				currNode.dom.controlDepBlocks = append(currNode.dom.controlDepBlocks, A)
				currNode = currNode.dom.idom
			}

			//fmt.Printf("%d is control dependent on %d\n", currNode.Index, A.Index)
			currNode.dom.controlDepBlocks = append(currNode.dom.controlDepBlocks, A)

		} else {
			currNode := B
			for currNode != lca {
				//fmt.Printf("%d is control dependent on %d\n", currNode.Index, A.Index)
				currNode.dom.controlDepBlocks = append(currNode.dom.controlDepBlocks, A)
				currNode = currNode.dom.idom
			}
		}
	}

	controlDepMatrix := make([][]int, len(graph.BlocksMap)+1)
	for i := 0; i < len(graph.BlocksMap)+1; i++ {
		controlDepMatrix[i] = make([]int, len(graph.BlocksMap)+1)
	}

	ancestorMatrix := make([][]int, len(graph.BlocksMap)+1)
	for i := 0; i < len(graph.BlocksMap)+1; i++ {
		ancestorMatrix[i] = make([]int, len(graph.BlocksMap)+1)
	}

	for _, n := range graph.BlocksMap {
		for _, controldepNode := range n.dom.controlDepBlocks {
			controlDepMatrix[n.Index+1][controldepNode.Index+1] = 1
		}
	}

	for _, n := range graph.BlocksMap {
		for _, ancestorNode := range n.dom.ancestors {
			ancestorMatrix[n.Index+1][ancestorNode.Index+1] = 1
		}
	}

	for i := 0; i < len(controlDepMatrix); i++ {
		for j := 0; j < len(controlDepMatrix); j++ {
			for k := 0; k < len(controlDepMatrix); k++ {
				if controlDepMatrix[i][j] == 1 ||
					(controlDepMatrix[i][k] == 1 && controlDepMatrix[k][j] == 1 && ancestorMatrix[i][j] == 0) {
					if controlDepMatrix[i][j] == 0 {
						//fmt.Printf("New : %d is control dependent on %d\n", i-1, j-1)
						//fmt.Printf("%v\n", graph.getNodeWithIdx(i-1))
						graph.getNodeWithIdx(i - 1).dom.controlDepBlocks = append(
							graph.getNodeWithIdx(i-1).dom.controlDepBlocks, graph.getNodeWithIdx(j-1))

					}
					controlDepMatrix[i][j] = 1
				}
			}
		}
	}

	for i := 0; i < len(controlDepMatrix); i++ {
		if graph.getNodeWithIdx(i-1) != nil {
			for j := 0; j < len(controlDepMatrix); j++ {
				graph.getNodeWithIdx(i - 1).dom.controlDepBitVector =
					append(graph.getNodeWithIdx(i-1).dom.controlDepBitVector, controlDepMatrix[i][j])
			}
		}
	}

}

func (cfgGraph *CFGGraph) BuildCDGFromCFG() {
	postDomGraphForCfg := NewPostDomGraph(-1, len(cfgGraph.Nodes)-2, false)

	//for k, _ := range cfgGraph.Nodes {
	//	fmt.Printf("%s | ", k)
	//}
	//fmt.Printf("\n")

	for _, n := range cfgGraph.Nodes {
		//fmt.Printf("Node %s\n", n.Name)
		for _, succN := range n.SuccessorIds {
			idxN, err := strconv.Atoi(n.Id)
			if err != nil {
				return
			}
			idxSuccN, err := strconv.Atoi(succN)
			if err != nil {
				return
			}
			//fmt.Printf("Adding Edge %d %d\n", idxN, idxSuccN)
			postDomGraphForCfg.AddEdge(idxN, idxSuccN)
		}
	}
	//fmt.Printf("\n\n")

	reverseG := NewPostDomGraph(0, len(postDomGraphForCfg.BlocksMap), false)
	mapReverseGraph := make(map[int]int)
	mapGraph := make(map[int]int)
	for i := 0; i < len(postDomGraphForCfg.BlocksMap); i++ {
		mapReverseGraph[len(postDomGraphForCfg.BlocksMap)-2-i] = i
		mapGraph[i] = len(postDomGraphForCfg.BlocksMap) - 2 - i
	}

	for _, edge := range postDomGraphForCfg.Edges {
		reverseG.AddEdge(mapReverseGraph[edge.to.Index], mapReverseGraph[edge.from.Index])
	}

	if buildDomTree(reverseG) {
		return
	}

	for _, node := range postDomGraphForCfg.BlocksMap {
		reverseGraphNode := reverseG.BlocksMap[mapReverseGraph[node.Index]]
		node.dom.idom = postDomGraphForCfg.BlocksMap[mapGraph[reverseGraphNode.dom.idom.Index]]
	}

	populateAncestors(postDomGraphForCfg)
	BuildControlDepTree(postDomGraphForCfg)

	for _, n := range postDomGraphForCfg.BlocksMap {
		for j := 0; j < len(n.dom.controlDepBitVector); j++ {
			dependent := cfgGraph.getNodeById(n.Index)
			controller := cfgGraph.getNodeById(j - 1)
			if dependent != nil && controller != nil && n.dom.controlDepBitVector[j] == 1 {
				dependent.ControlDepBlocks[controller] = "-1"
				controller.ControllerOfBlocks = append(controller.ControllerOfBlocks, dependent)
			}
		}
	}

	//cfgGraph.CalculateControllerFlow()
}

func ProcessBFS(block *Block) {
	fmt.Printf("Processing Block %s\n", block.FullName())
	for _, n := range block.Function.ContainsBlock {
		n.BFSVisitController = false
	}

	if len(block.ControllerOfBlocks) != 0 {
		if len(block.Successors) == 1 {
			if block.LastInstr != nil && block.LastInstr.SSAType == "Jump" {
				for _, dep := range block.ControllerOfBlocks {
					dep.ControlDepBlocks[block] = "2"
				}
			}

		} else if len(block.Successors) == 2 {
			if block.LastInstr != nil && block.LastInstr.SSAType == "If" {
				var TrueBFSQueue []*Block
				TrueBFSQueue = append(TrueBFSQueue, block)

				for len(TrueBFSQueue) != 0 {
					currNode := TrueBFSQueue[0]
					fmt.Printf("%s ", currNode.Id)
					TrueBFSQueue = TrueBFSQueue[1:]

					if !currNode.BFSVisitController {
						currNode.BFSVisitController = true

						_, ok := currNode.ControlDepBlocks[block]
						if ok {
							currNode.ControlDepBlocks[block] = "1"
						}

						if currNode == block {
							TrueBFSQueue = append(TrueBFSQueue, block.TrueSuccessor)
						} else {
							for _, succ := range currNode.Successors {
								TrueBFSQueue = append(TrueBFSQueue, succ)
							}
						}

					}
				}
				fmt.Printf("\n")

				for _, dep := range block.ControllerOfBlocks {
					if dep.ControlDepBlocks[block] == "-1" {
						dep.ControlDepBlocks[block] = "0"
					}
				}

			}
		}
	}
}

func (cfgGraph *CFGGraph) CalculateControllerFlow() {
	for _, n := range cfgGraph.Nodes {
		ProcessBFS(n)
	}

	fmt.Printf("\n\n\n")
	for _, n := range cfgGraph.Nodes {
		fmt.Printf("Block %s\n", n.FullName())
		for k, v := range n.ControlDepBlocks {
			fmt.Printf("%s is Dependent on %s %s\n", n.FullName(), k.FullName(), v)
		}
	}
}

//
//func ComputeDomTree() {
//
//	g := NewPostDomGraph(-1, 5, false)
//
//	// 0 1
//	// 1 2
//	// 1 3
//	// 2 4
//	// 2 5
//	// 3 5
//	// 3 7
//	// 4 6
//	// 5 6
//	// 6 7
//
//	g.AddEdge(-1, 0)
//	g.AddEdge(0, 1)
//	g.AddEdge(1, 2)
//	g.AddEdge(1, 3)
//	g.AddEdge(2, 4)
//	g.AddEdge(2, 5)
//	g.AddEdge(3, 5)
//	g.AddEdge(3, 7)
//	g.AddEdge(4, 6)
//	g.AddEdge(5, 6)
//	g.AddEdge(6, 7)
//	g.AddEdge(7, 8)
//
//	//g.AddEdge(-1, 0)
//	//g.AddEdge(0, 1)
//	//g.AddEdge(0, 2)
//	//g.AddEdge(1, 3)
//	//g.AddEdge(1, 4)
//	//g.AddEdge(2, 5)
//	//g.AddEdge(3, 5)
//	//g.AddEdge(4, 5)
//
//	fmt.Printf("%d\n", len(g.BlocksMap))
//	reverseG := NewPostDomGraph(0, len(g.BlocksMap), false)
//	mapReverseGraph := make(map[int]int)
//	mapGraph := make(map[int]int)
//	for i := 0; i < len(g.BlocksMap); i++ {
//		mapReverseGraph[len(g.BlocksMap)-2-i] = i
//		mapGraph[i] = len(g.BlocksMap) - 2 - i
//	}
//
//	//for k, v := range mapReverseGraph {
//	//	fmt.Printf("Mapreversegraph %d : %d\n", k, v)
//	//}
//	//
//	//for k, v := range mapGraph {
//	//	fmt.Printf("Mapgraph %d : %d\n", k, v)
//	//}
//
//	for _, edge := range g.Edges {
//		//fmt.Printf("Adding edge %d %d\n", mapReverseGraph[edge.to.Index], mapReverseGraph[edge.from.Index])
//		reverseG.AddEdge(mapReverseGraph[edge.to.Index], mapReverseGraph[edge.from.Index])
//	}
//
//	buildDomTree(reverseG)
//
//	//for _, node := range reverseG.BlocksMap {
//	//	fmt.Printf("Idom of %d in RG is %d\n", node.Index, node.dom.idom.Index)
//	//}
//
//	for _, node := range g.BlocksMap {
//		reverseGraphNode := reverseG.BlocksMap[mapReverseGraph[node.Index]]
//		node.dom.idom = g.BlocksMap[mapGraph[reverseGraphNode.dom.idom.Index]]
//	}
//
//	populateAncestors(g)
//	BuildControlDepTree(g)
//
//	for _, n := range g.BlocksMap {
//		if n.dom.idom != nil {
//			fmt.Printf("Idom of %d is %d\n", n.Index, n.dom.idom.Index)
//		} else {
//			fmt.Printf("No Idom of %d", n.Index)
//		}
//	}
//	//for _, n := range g.BlocksMap {
//	//	if n.dom.idom != nil {
//	//		fmt.Printf("Ancestors of %d : ", n.Index)
//	//		for _, n := range n.dom.ancestors {
//	//			fmt.Printf("%d ", n.Index)
//	//		}
//	//		fmt.Printf("\n")
//	//		fmt.Printf("%v \n", n.dom.ancestorBitVector)
//	//	} else {
//	//		fmt.Printf("No dom of %d\n", n.Index)
//	//	}
//	//}
//}

// 0 1
// 1 2
// 1 3
// 2 4
// 2 5
// 3 5
// 3 7
// 4 6
// 5 6
// 6 7
