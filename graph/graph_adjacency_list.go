/*
这是一个邻接表图相关模块
整合：无向邻接表图、BFS、DFS、内置工具函数，无外部pkg依赖
*/
package graph


import (
	"fmt"
	"strconv"
	"strings"
)


/*
基本数据结构 -- 邻接表图
技术实现 -- 链表
*/
// Vertex 顶点结构体，替代外部 pkg 导入
type Vertex struct {
	Val int
}

// DeleteSliceElms 切片删除指定元素工具函数，内置替代pkg工具
func DeleteSliceElms[T comparable](slice []T, target T) []T {
	newSlice := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != target {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// 基于邻接表实现的无向图类（首字母大写导出）
type GraphAdjList struct {
	// 邻接表，key：顶点，value：该顶点的所有邻接顶点
	adjList map[Vertex][]Vertex
}

// 构造函数，大写导出
func NewGraphAdjList(edges [][]Vertex) *GraphAdjList {
	g := &GraphAdjList{
		adjList: make(map[Vertex][]Vertex),
	}
	// 添加所有顶点和边
	for _, edge := range edges {
		g.AddVertex(edge[0])
		g.AddVertex(edge[1])
		g.AddEdge(edge[0], edge[1])
	}
	return g
}

// 获取顶点数量
func (g *GraphAdjList) Size() int {
	return len(g.adjList)
}

// 添加边
func (g *GraphAdjList) AddEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error: vertex not exist or self-loop edge invalid")
	}
	// 添加无向双向边
	g.adjList[vet1] = append(g.adjList[vet1], vet2)
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
}

// 删除边
func (g *GraphAdjList) RemoveEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error: vertex not exist or self-loop edge invalid")
	}
	g.adjList[vet1] = DeleteSliceElms(g.adjList[vet1], vet2)
	g.adjList[vet2] = DeleteSliceElms(g.adjList[vet2], vet1)
}

// 添加顶点
func (g *GraphAdjList) AddVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	g.adjList[vet] = make([]Vertex, 0)
}

// 删除顶点
func (g *GraphAdjList) RemoveVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if !ok {
		panic("error: vertex not exist")
	}
	// 删除顶点key
	delete(g.adjList, vet)
	// 所有邻接列表移除该顶点
	for v, list := range g.adjList {
		g.adjList[v] = DeleteSliceElms(list, vet)
	}
}

// 打印邻接表
func (g *GraphAdjList) Print() {
	var builder strings.Builder
	fmt.Printf("邻接表 = \n")
	for k, v := range g.adjList {
		builder.WriteString("\t\t" + strconv.Itoa(k.Val) + ": ")
		for _, vet := range v {
			builder.WriteString(strconv.Itoa(vet.Val) + " ")
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}

// ====================== BFS 图遍历 ======================
// GraphBFS 广度优先遍历，返回遍历顶点序列
func GraphBFS(g *GraphAdjList, start Vertex) []Vertex {
	// 已访问集合
	visited := make(map[Vertex]bool)
	// 队列
	queue := []Vertex{start}
	visited[start] = true
	// 遍历结果
	res := make([]Vertex, 0)

	for len(queue) > 0 {
		// 队首出队
		vet := queue[0]
		queue = queue[1:]
		res = append(res, vet)

		// 遍历邻接顶点
		for _, adjVet := range g.adjList[vet] {
			if !visited[adjVet] {
				visited[adjVet] = true
				queue = append(queue, adjVet)
			}
		}
	}
	return res
}

// ====================== DFS 图遍历 ======================
// GraphDFS 深度优先遍历入口
func GraphDFS(g *GraphAdjList, start Vertex) []Vertex {
	visited := make(map[Vertex]bool)
	res := make([]Vertex, 0)
	var dfs func(cur Vertex)
	dfs = func(cur Vertex) {
		visited[cur] = true
		res = append(res, cur)
		// 递归访问邻接点
		for _, adjVet := range g.adjList[cur] {
			if !visited[adjVet] {
				dfs(adjVet)
			}
		}
	}
	dfs(start)
	return res
}
