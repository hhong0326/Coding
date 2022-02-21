/**
 * @param {number[][]} isConnected
 * @return {number}
 */
 var findCircleNum = function(isConnected) {
    
    let visited = new Set();
    
    let provinces = 0;
    
    var dfs = function(node) {
        for(let i = 0; i < isConnected.length; i++){
            if(isConnected[node][i] === 1 && !visited.has(i)){
                visited.add(i);
                dfs(i);
            }
        }
    }
    
    for(let i=0; i<isConnected.length; i++){
        if(!visited.has(i)){
            
            dfs(i);
            provinces++;
        }
    }
    
    return provinces; 
};