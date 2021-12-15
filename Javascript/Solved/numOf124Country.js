function solution(n) {
    var answer = '';
    
    if (n==1) {
        return '1'
    } else if (n==2) {
        return '2'
    } else if (n==3) {
        return '4'
    }
    
    var arr = [];
    
    while (n >= 3) {
        var b = n % 3;
        if (b == 0) {
            
            arr.push('4');
            n = Math.floor(n/3)-1;
            
        } else {
            arr.push(b + '');
            n = Math.floor(n/3);
        }
       
    } 
    if (n != 0) {
        arr.push(n);
    }
    
    answer = arr.reverse().join('');

    return answer;
}
