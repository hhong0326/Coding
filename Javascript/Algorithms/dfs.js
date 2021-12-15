function solution(n, computers) {
    var answer = 0;
    var s = [];
    s.push(0);
    
    while (s.length !== 0) {
        
        var pc = s.pop();
         
        for (var i=0; i<computers[pc].length; i++) {
            if (i !== pc && computers[pc][i] === 1) {

                s.push(i);
                console.log(i)
            }
        }
        
        
    }
    return answer;
}

solution(3, [[1, 1, 0], [1, 1, 0], [0, 0, 1]]);