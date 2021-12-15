function solution(n, edge) {
    var answer = 0;
    
    //bfs
    
    var q = [1];
    var visited = [];
    var depth = 0;
    
    while(q.length !== 0) {
        
        var l = q.length;
        
        for (var i=0; i<l; i++) {
            var line = q.shift();

            if(visited.includes(line)) {
                continue;
            }
            visited.push(line);
            
            edge.forEach(item => {
                if (item[0] === line) {
                    q.push(item[1]);
                } else if (item[1] === line) {
                    q.push(item[0]);
                }
            });
            depth++;
        }
        
        console.log(q)
        console.log(visited)
        console.log(depth)
    }
    
    return answer;
}

solution(6, [[3, 6], [4, 3], [3, 2], [1, 3], [1, 2], [2, 4], [5, 2]]);