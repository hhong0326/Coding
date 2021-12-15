// function solution(N, number) {
//     const cache = new Array(9).fill(0).map(el => new Set());
//     for(let i=1; i<9; i++){
//         cache[i].add('1'.repeat(i) * N);
//         for (let j=1; j < i; j++) {
//             for(const arg1 of cache[j]){
//                 for(const arg2 of cache[i-j]){
//                     cache[i].add(arg1+arg2);
//                     cache[i].add(arg1-arg2);
//                     cache[i].add(arg1*arg2);
//                     cache[i].add(arg1/arg2>>0);
//                 }
//             }
//         }
//         if(cache[i].has(number)) return i; 
//     }
//     return -1;
// }

function solution(N, number) {
    var answer = -1;
    
    var dp = new Array(9);
    dp[0] = new Set();
    dp[1] = new Set();
    // dp[2] = new Set();
    
    dp[0].add(0);
    dp[1].add(N);
//     dp[2].add(Number(String(N)+String(N))); //NN dp[2] = dp[1]*11 + dp[2]
//     dp[2].add(N+N);
//     dp[2].add(N-N);
//     dp[2].add(N*N);
//     dp[2].add(N/N);
    
//     for (var i=1; i<=2; i++) {
//         if (dp[i].has(number)) {
//             return i;
//         }
//     }
    

    for (var i=1; i<9; i++) {
        dp[i] = new Set();
        // console.log('1'.repeat(i)*N)
        dp[i].add('1'.repeat(i)*N);  
        for (var j = 1; j < i; j++) {
            
            for (var item1 of dp[j]) {
              for (var item2 of dp[i-j]) {
                dp[i].add(item1+item2)
                dp[i].add(item1-item2);
                dp[i].add(item1*item2);
                dp[i].add(item1/item2 >> 0); // arg1/arg2>>0 비트연산자
              }
            }
          }
          if (dp[i].has(number)) {
            return i;
          }
        // for (var ii of dp[i-1].entries()) {
        //     // console.log(ii)
        //     dp[i].add(Number(String(ii[1])+String(N))); // Number(String(N)+String(N))
        //     // console.log(Number(String(ii)+String(N)));
        //     dp[i].add(ii[1]+N);
        //     dp[i].add(ii[1]-N);
        //     dp[i].add(ii[1]*N);
        //     dp[i].add(ii[1]/N);
        // }
        
    }    
    
    return answer;
}

solution(5, 12);