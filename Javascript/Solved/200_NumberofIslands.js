/**
 * @param {character[][]} grid
 * @return {number}
 */
 var numIslands = function(grid) {
    
    const pos = [[-1,0],[0,1],[1,0],[0,-1]];
    
    const bfs = function(q) {
        while(q.length > 0) {
            let [i, j] = q.shift();
            
            if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length || grid[i][j] === '0') {
                continue;
            }
            
            grid[i][j] = '0';
            
            for (let k=0; k<pos.length; k++) {
                q.push([i + pos[k][0], j + pos[k][1]]);
            }
            
        }
    }
    
//     const dfs = function(i, j) {
//         if (i < 0 || j < 0 || i == grid.length || j == grid[0].length || grid[i][j] === '0') {
//             return;
//         }
        
//         if (grid[i][j] === '1') {
            
//             grid[i][j] = '0';
//             dfs(i-1,j);
//             dfs(i,j-1);
//             dfs(i+1,j);
//             dfs(i,j+1);
//         }
//     }
    
    let count = 0;
    
    for (let i=0; i<grid.length; i++ ) {
        for (let j=0; j<grid[0].length; j++ ) {
            if (grid[i][j] === '1') {
                
                count++;
                // dfs(i, j);
                bfs([[i, j]]);
            }
        }
    }
    
    return count;
    

// DFS using stack
//     var counter = 0;
    
//     for (var i = 0; i < grid.length; i++) {
//         for (var j = 0; j < grid[0].length; j++) {
//             var stack = [];
//             if (grid[i][j] === '1') {
//                 counter ++;
//                 grid[i][j] = 0;
//                 stack.push([i, j]);
//             }
            
//             while (stack.length !== 0) {
//                 var [row, col] = stack.pop();
                
//                 if (grid[row][col+1] === '1') {
//                     grid[row][col+1] = '0'
//                     stack.push([row, col+1])
//                 }
//                 if (grid[row+1] && grid[row+1][col] === '1') {
//                     grid[row+1][col] = '0'
//                     stack.push([row+1, col])
//                 }
//                 if (grid[row][col-1] === '1') {
//                     grid[row][col-1] = '0'
//                     stack.push([row, col-1])
//                 }
//                 if (grid[row-1] && grid[row-1][col] === '1') {
//                     grid[row-1][col] = '0'
//                     stack.push([row-1, col])
//                 }
//             }
//         }
//     }
    
//     return counter;
};