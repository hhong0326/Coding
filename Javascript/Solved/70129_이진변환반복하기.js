function solution(s) {
    var answer = [];
    
    while (s != "1") {
        var count = 0;
        var newS = "";
        for (var i=0; i<s.length; i++) {
            if (s[i] != "0") {
                newS += s[i];
            } else {
                count++;
            }
        }
        
        var l = newS.length;
        var v = l.toString(2);
        s = "" + v
        answer.push(count);
    }
    
    return answer;
}

console.log(solution("0111010"))