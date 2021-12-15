function solution(maps) {
    var answer = 0;
    
    var n = maps.length;
    var m = maps[0].length;
    
    var dis = new Array(n);
    for (var i=0; i<n; i++) {
        dis[i] = new Array(m);
        for (var j=0; j<m; j++) {
            dis[i][j] = 1;
        }
    }
    
    var dx = [0, 0, -1, 1]; // 방향키
    var dy = [-1, 1, 0, 0];
    
    var i = 0;
    var j = 0;
    var ni = 0;
    var nj = 0;
    var q = [];
    
    q.push([0, 0]);
    maps[0][0] = 2;

    while (q.length !== 0) {
        var a = q.shift();
        i = a[0];
        j = a[1];
        
        for (var k=0; k<4; k++) {
            
            ni = i + dx[k];
            nj = j + dy[k];
            if (ni >= 0 && nj >= 0 && ni < maps.length && nj < maps[0].length && maps[ni][nj] === 1) {
                q.push([ni, nj]); // BFS 
                maps[ni][nj] = 2; // 지난 것임을 확인하기 위해 2를 넣는다
                dis[ni][nj] = dis[i][j]+1; // 이전꺼 값에 1을 더한 다음 거리를 구한다
            } 
        }
            
    }
    
    return dis[n-1][m-1] === 1 ? -1 : dis[n-1][m-1];
}

solution([[1,0,1,1,1],[1,0,1,0,1],[1,0,1,1,1],[1,1,1,0,1],[0,0,0,0,1]]);