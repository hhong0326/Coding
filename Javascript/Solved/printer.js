// function solution(priorities, location) {
//     var answer = 0;
//     var count = 0;
    
//     while (priorities.length > 0) {
//         if(priorities[0] === Math.max(...priorities)) {
//             count++;
            
//             priorities.shift();
            
//             if(location === 0) {
//                 console.log(count)
//                 return count;
//             } 
//             location--;
//         } else {
//             if (location > 0) {
//                 location--;
//             } else {
//                 location = priorities.length-1;
//             }
//             var out = priorities.shift();
//             priorities.push(out);
//         }
//     }
    
// }



function solution(s) {
    var answer = [];
    
    var zeros = 0;
    var count = 0;
    var two = [];
    
    var a = 127;
    console.log(a.toString(2));
    
    while (s !== '1') {
        var two = s.split('');
        two = two.filter((ele) => {
            if (ele === '1') {
                return ele;
            } else {
                
                zeros++;
            }
        
        });

        s = parseInt(two.length, 2);
        console.log(s)
        count++;
    }
    
    console.log(count, zeros);
    
    return answer;
}

solution("110110101001");