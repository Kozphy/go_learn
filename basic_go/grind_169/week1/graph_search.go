package week1

// TODO: to complete
// func dfs_practice(graph *map[string][]int, vertices_seg []int, target int) bool {
// S := &StackVertex{}
// S.Push()

// for _, vertex := range vertices_seg {
// 	if vertex == target {
// 		return true
// 	}
// 	if !S.IsEmpty() {
// 		str_vertex := strconv.Itoa(vertex)
// 		if !S.vertrices[S.length-1].visited {
// 			dfs_practice(graph, (*graph)[str_vertex], target)
// 		}
// 	}
// }

// 	return false
// }

// TODO: to complete
// short path problem
// func bfs_practice(graph map[string][]int, target int) bool {
// 	Q := &QueueVertex{}
// 	Q.enqueue(graph["0"])
// fmt.Println()
// fmt.Printf("%+v\n", Q.front())
// Q.dequeue()
// fmt.Println(Q.front())
// var searched []string
// for Q.length != 0 {
// 	vertex := Q.front()
// 	Q.dequeue()
// 	for _, searchValue := range searched {
// 		sv, err := strconv.Atoi(searchValue)
// 		if err != nil {
// 			log.Fatalf("convert string to value failed: %v\n", sv)
// 		}
// 		if vertex.value != sv {
// 			if vertex.value == target {
// 				fmt.Println("finded")
// 				return true
// 			} else {
// 				Q.enqueue()
// 				searched = append(searched)
// 			}
// 		}
// 	}
// }
// return false
// }

// func Execute_graph_search_practice() {
// 	graph := map[string][]int{
// 		"0": {1, 2},
// 		"1": {0, 3, 4},
// 		"2": {0},
// 		"3": {1},
// 		"4": {2, 3},
// 	}
// 	bfs_practice(graph, 0)
// 	// dfs_practice(&graph, graph["0"], 0)
// }
