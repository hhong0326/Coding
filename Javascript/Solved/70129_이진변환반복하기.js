function solution(s) {
    var answer = [];
    
    var binaryCnt = 0;

    var count = 0;
    while (s != "1") {
    
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
        binaryCnt++;
    }

    
    return [binaryCnt, count];
}

console.log(solution("01110"))