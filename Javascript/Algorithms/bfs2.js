//BFS

const GRAPH = {
    A: ["B", "C"],
    B: ["A", "D"],
    C: ["A", "G", "H", "I"],
    D: ["B", "E", "F"],
    E: ["D"],
    F: ["D"],
    G: ["C"],
    H: ["C"],
    I: ["C", "J"],
    J: ["I"]
  };
  
  const bfs = (GRAPH, startNode) => {
    let visited = []; // 탐색을 마친 노드들
    let needVisit = []; // 탐색해야할 노드들
  
    needVisit.push(startNode); // 노드 탐색 시작
  
    while (needVisit.length !== 0) { // 탐색해야할 노드가 남아있다면
      const node = needVisit.shift(); // queue이기 때문에 선입선출, shift()를 사용한다.
      if (!visited.includes(node)) { // 해당 노드가 탐색된 적 없다면
        visited.push(node); 
        needVisit = [...needVisit, ...GRAPH[node]];
      }
    }
    return visited;
  };
  
  console.log(bfs(GRAPH, "A"));
  // ["A", "B", "C", "D", "G", "H", "I", "E", "F", "J"] 



  

  //DFS

  const graph = {
    A: ["B", "C"],
    B: ["A", "D"],
    C: ["A", "G", "H", "I"],
    D: ["B", "E", "F"],
    E: ["D"],
    F: ["D"],
    G: ["C"],
    H: ["C"],
    I: ["C", "J"],
    J: ["I"],
  };
  
  // (graph, 시작 정점)
  const dfs = (graph, startNode) => {
    let needVisitStack = []; // 탐색을 해야 할 노드들
    let visitedQueue = []; // 탐색을 마친 노드들
  
    needVisitStack.push(startNode);
  
    // 탐색을 해야 할 노드가 남아 있다면
    while (needVisitStack.length !== 0) {
      const node = needVisitStack.pop();
      if (!visitedQueue.includes(node)) {
        visitedQueue.push(node);
        needVisitStack = [...needVisitStack, ...graph[node]];
      }
    }
  
    return visitedQueue;
  };
  
  console.log(dfs(graph, "A"));
  
  // ["A", "C", "I", "J", "H", "G", "B", "D", "F", "E"]