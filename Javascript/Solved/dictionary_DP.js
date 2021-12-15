function solution(word) {
    var answer = 0;
    
    var dp = new Array(6).fill(0).map(el => new Set());
    dp[1].add('A');
    dp[1].add('E');
    dp[1].add('I');
    dp[1].add('O');
    dp[1].add('U');
    
    for (var i=2; i<=5; i++) {
        
        for (var j=1; j<i; j++) {
            for(var item1 of dp[j]){
                for(var item2 of dp[i-j]){
                
                    dp[i].add(item1 + item2);
                }
            }
        }
    }
    
    var dictionary = [];
    var dp2 = dp.map(item => [...item].sort()).forEach(item => dictionary.push(...item));
    
    dictionary.sort();
    
    answer = binarySearch(dictionary, word);
    
    return answer+1;
}

function binarySearch(dictionary, word) {
    
    var s = 0, e = dictionary.length -1;
    
    while (s <= e) {
        var mid = Math.floor((s+e) / 2);
        if(dictionary[mid] === word) {
            return mid;
        } else if((dictionary[mid] <= word) == 1) {
            s = mid + 1;
        } else {
            e = mid - 1;
        }
    }
    
    return -1;
}