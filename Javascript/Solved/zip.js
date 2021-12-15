function solution(msg) {
    var answer = [];
    
    var index = {
        'A': 1,
        'B': 2,
        'C': 3,
        'D': 4,
        'E': 5,
        'F': 6,
        'G': 7,
        'H': 8,
        'I': 9,
        'J': 10,
        'K': 11,
        'L': 12,
        'M': 13,
        'M': 14,
        'O': 15,
        'P': 16,
        'Q': 17,
        'R': 18,
        'S': 19,
        'T': 20,
        'U': 21,
        'V': 22,
        'W': 23,
        'X': 24,
        'Y': 25,
        'Z': 26
    };

    if (msg.length === 1) {
        return [index[msg]];
    }
    
    var count = 27
    
    var i = 0;
    var l = msg.length;
    while(i < l) {
        var cur = '';
        for(var j=i; j<l; j++) {
            const target = msg.substring(i, j+1);
            if(index[target]) {
                cur = target;
            } else {
                index[target] = ++count;
                break;
            }
        }
        answer.push(index[cur]);
        i += cur.length;
    }
    
    
    return answer;
}

solution('KAKAO');