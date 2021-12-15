/**
 * @param {number[][]} rooms
 * @return {boolean}
 */
var canVisitAllRooms = function(rooms) {
    
    // bfs
    let q = [0];
    let visited = new Set();
    
    while (q.length > 0) {
        
        let r = q.shift();
        
        if(visited.has(r)) continue;
        visited.add(r);
        
        rooms[r].forEach(item => {
            q.push(item);
        });
        
        
    }
    
    return rooms.length === visited.size;
    
    // dfs
    
//     let s = [0];
//     let visited = new Set();
    
//     while(s.length) {
        
//         let r = s.pop();
        
//         if(visited.has(r)) continue;
//         visited.add(r);
//         rooms[r].forEach(item => {
//             s.push(item);
//         });
//     }
    
//     return rooms.length === visited.size;
};
