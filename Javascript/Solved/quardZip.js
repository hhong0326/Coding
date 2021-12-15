function solution(arr) {
    var answer = [0, 0];
    
    // 반으로 쪼개고 만들고 다시 아닌 부분을 위해 반으로 쪼개고..
    
    var size = arr.length;
    var s = new Set();
    var obj = {};
    var m = new Map();
    
    while(size != 1) {
        size/=2
        for (var i=0; i<arr.length; i+=size) {
            for (var j=0; j<arr[0].length; j+=size) {
                console.log(i + "___________" + j)
                dfs(i, j, size);
                
            }
            
        }
       
    }

    
    function dfs(s1, s2, len) {

        // console.log([...s])
        console.log(obj)
        console.log('' +  s1 +s2)
        if (len == 1 || s.has('' + s1 + s2)){
            // console.log("pass")
            return;
        }

        if (s1 <= obj[s1] && obj[s2] <= s2) {
            console.log("pass")
            return
        }
        var count = 0;
        
        for (var k=0; k<len; k++) {
            for (var l=0; l<len; l++) {
                if(arr[s1+k][s2+l] == 1) {
                    count++; // 1 else 0
                    // console.log(arr[i+k][j+l])
                }   
            }
        }

        console.log(count)
        if (count == len*len) {
            s.add(''  + s1 + s2);
            obj[s1] = s2;
            obj[s2] = s1;
            answer[1] += 1;
            console.log('압춥')
        } else if (count == 0){
            s.add('' + s1 + s2);
            obj[s1] = s2;
            obj[s2] = s1;
            console.log('압')
            answer[0] += 1;
        } else {
            console.log("Fdsfasdfa")
            // answer[1] += count;
            // answer[0] += len*len - count;
            // console.log(len/2*len/2 - count, count)
            // dfs(s1, s2, len/2);
        }
        if(len == 2) {
            answer[1] += count;
            answer[0] += len*len - count;
        }
    }
    
    return answer;
}

console.log(solution([[1,1,1,1,1,1,1,1],[0,1,1,1,1,1,1,1],[0,0,0,0,1,1,1,1],[0,1,0,0,1,1,1,1],[0,0,0,0,0,0,1,1],[0,0,0,0,0,0,0,1],[0,0,0,0,1,0,0,1],[0,0,0,0,1,1,1,1]]))