// bfs

var canReach = function(arr, start) {
    
    let q = [];
    q.push(start);
    visited = new Array(arr).fill(false);
    
    while (q.length > 0) {
        
        let s = q.shift();
        visited[s] = true;
        
        if (arr[s] === 0) {
            return true;
        }
        
        let l = s - arr[s];
        let r = s + arr[s];
        if (l >= 0 && !visited[l]) {
            q.push(l);
        }
        if (r < arr.length && !visited[r]) {
            q.push(r);
        }
    }
    
    return false;
};
