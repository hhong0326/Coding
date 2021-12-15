function solution(dartResult) {
    var answer = 0;

    // 0 - 10 3번
    // 
    // # * 이 있으면 3글자 아니면 두글자

    
    dartResult = dartResult.replace(/([*#SDT])/g,"$1 ").split(" ");
    dartResult.pop();

    var dart = [];
    for (var i=0; i<dartResult.length; i++) {
        if (dartResult[i] === '*' || dartResult[i] === '#') {
            continue;
        } 
        if (dartResult[i+1] === '*' || dartResult[i+1] === '#') {
            dart.push(dartResult[i] + dartResult[i+1]);
        } else {
            if (i === dartResult.length - 2) {
                dart.push(dartResult[i], dartResult[i+1]);
            } else {

                dart.push(dartResult[i]);
            }
        }
    }
    var result2 = dart.reduce((pre, now) => {
        if(now.length !== 2) { // 3
            var n = zegop(Number(now[0]), now[1]);
            console.log(n, starshap(pre, n, now[2]))
            return starshap(pre, n, now[2]);
        } else {
            console.log(zegop(Number(now[0]), now[1]))
            return pre + zegop(Number(now[0]), now[1]);
        }
    }, 0);
    
    
    return result2;
}
                                
function zegop(num, letter) {
        
    switch(letter) {
        case 'S':
            return Math.pow(num, 1);
        case 'D':
            return Math.pow(num, 2);
        case 'T':
            return Math.pow(num, 3);
    }
    
    return 0;
}
                       
function starshap(preNum, nowNum, letter) {
        
    console.log(preNum, nowNum)
    switch(letter) {
        case '*':
            return preNum*2 + nowNum*2;
        case '#':
            return preNum;
    }
}

solution('1S2D*3T');