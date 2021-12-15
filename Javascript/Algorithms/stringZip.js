function solution(s) {
    var answer = s.length;

    var idx = 0;
    var count = 1;
    
    for(var i=1;i<=s.length/2;i++) {
        var str = "";
        var tmp = s.substring(0,i);
        
        for(idx=i; idx<=s.length; idx+=i) {
            if(tmp === s.substring(idx, idx+i)){
                count++;
            } else {
                if(count === 1) {
                    str += tmp;
                } else {
                    str += count + tmp;
                }
                    count = 1;
                    tmp = s.substring(idx, idx+i);
            }
        }
        
        if(count === 1) {
            str += tmp; 
        } else {
            str += count + tmp;
        }
        
        answer = Math.min(answer, str.length);
    }
    return answer;
}