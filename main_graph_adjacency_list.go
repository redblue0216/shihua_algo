package main

import (
	"fmt"
	"shihua_algo/graph"
)

func main() {
	fmt.Println("===== 邻接表无向图 测试示例 =====")
	// 构建图边集合，顶点值 0~4
	// 图结构：
	//    0
	//   / \
	//  1   2
	//  |   |
	//  3---4
	edges := [][]graph.Vertex{
		{{Val: 0}, {Val: 1}},
		{{Val: 0}, {Val: 2}},
		{{Val: 1}, {Val: 3}},
		{{Val: 2}, {Val: 4}},
		{{Val: 3}, {Val: 4}},
	}
	// 初始化邻接表无向图
	g := graph.NewGraphAdjList(edges)

	// 1. 获取顶点总数 & 打印原始邻接表
	fmt.Println("\n1. 图基础信息与邻接表输出")
	fmt.Printf("顶点总数量 size = %d\n", g.Size())
	g.Print()

	// 2. 测试新增顶点、新增边
	fmt.Println("\n2. 新增顶点 5，新增边 0-5")
	g.AddVertex(graph.Vertex{Val: 5})
	g.AddEdge(graph.Vertex{Val: 0}, graph.Vertex{Val: 5})
	fmt.Printf("新增后顶点总数 size = %d\n", g.Size())
	g.Print()

	// 3. 删除边测试（删除 0-2）
	fmt.Println("\n3. 删除边 0 <-> 2")
	g.RemoveEdge(graph.Vertex{Val: 0}, graph.Vertex{Val: 2})
	g.Print()

	// 4. 删除顶点测试（删除顶点 4，自动清除所有关联边）
	fmt.Println("\n4. 删除顶点 4")
	g.RemoveVertex(graph.Vertex{Val: 4})
	fmt.Printf("删除后顶点总数 size = %d\n", g.Size())
	g.Print()

	// 5. BFS 广度优先遍历（从顶点0出发）
	fmt.Println("\n5. BFS 广度优先遍历（起点0）")
	bfsResult := graph.GraphBFS(g, graph.Vertex{Val: 0})
	fmt.Print("BFS 遍历序列: ")
	for _, v := range bfsResult {
		fmt.Printf("%d ", v.Val)
	}
	fmt.Println()

	// 6. DFS 深度优先遍历（从顶点0出发）
	fmt.Println("\n6. DFS 深度优先遍历（起点0）")
	dfsResult := graph.GraphDFS(g, graph.Vertex{Val: 0})
	fmt.Print("DFS 遍历序列: ")
	for _, v := range dfsResult {
		fmt.Printf("%d ", v.Val)
	}
	fmt.Println()

	// 演示另一张独立图测试
	fmt.Println("\n===== 新建另一张无向图测试 =====")
	edges2 := [][]graph.Vertex{
		{{Val: 10}, {Val: 20}},
		{{Val: 10}, {Val: 30}},
		{{Val: 20}, {Val: 30}},
		{{Val: 20}, {Val: 40}},
	}
	g2 := graph.NewGraphAdjList(edges2)
	fmt.Printf("第二张图顶点总数: %d\n", g2.Size())
	fmt.Println("邻接表：")
	g2.Print()

	fmt.Println("第二张图 BFS(起点10): ")
	bfs2 := graph.GraphBFS(g2, graph.Vertex{Val: 10})
	for _, v := range bfs2 {
		fmt.Printf("%d ", v.Val)
	}
	fmt.Println()

	fmt.Println("第二张图 DFS(起点10): ")
	dfs2 := graph.GraphDFS(g2, graph.Vertex{Val: 10})
	for _, v := range dfs2 {
		fmt.Printf("%d ", v.Val)
	}
	fmt.Println()
}
