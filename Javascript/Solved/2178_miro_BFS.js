function solution() {
    
    let maps =[[1,0,1,1,1,1], [1,0,1,0,1,0],[1,0,1,0,1,1],[1,1,1,0,1,1]];

    const N = 4;
    const M = 6;
    const dx = [0,0,1,-1]
    const dy = [1,-1,0,0];

    let visited = [[0, 0, 1]];
    while(visited.length){
        let [n,m,c] = visited.shift();
        maps[n][m] = 0;
        console.log(maps[n][m]);
        
        for(let i=0; i<4;i++){
            const x = n+dx[i]
            const y = m+dy[i]
            if(x>=0 && x<N && y>=0 && y<M){
                if(maps[x][y] === 1){
                    visited.push([x,y,c+1])
                }
            }
        } 
    }
}

solution();