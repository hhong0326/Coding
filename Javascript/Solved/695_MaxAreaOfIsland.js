/**
 * @param {number[][]} grid
 * @return {number}
 */

// BFS

var maxAreaOfIsland = function(grid) {
    
    var max = 0;
    var count = 0;
    var pos = [[-1,0], [0,1], [1,0], [0,-1]];
    
    var bfs = function(q) {
        
        while(q.length) {
            var [i, j] = q.shift();
        
            if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length || grid[i][j] !== 1) {
                continue; // while, q
            }

            grid[i][j] = 2;

            count++;

            for (var k=0; k<pos.length; k++) { 
                q.push([i + pos[k][0], j+ pos[k][1]]);
            }

        }
        
        
    }
    
    for (var i=0; i<grid.length; i++) {
        for (var j=0; j<grid[0].length; j++) {
            if (grid[i][j] == 1) {
                bfs([[i, j]]);
                if (max < count) max = count;
                
                count = 0;
            }
        }
    }
        
    return max;
};