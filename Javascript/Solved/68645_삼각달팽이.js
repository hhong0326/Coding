function solution(n) {
    var answer = [];

    // 하나씩 증가하는 이차원배열
    var triangle = new Array(n).fill(0).map((item, i) => new Array(i+1).fill(0));
    
    var y = -1
    var x = 0
    var count = 1
    
    for (var i=0; i < n; i++) {
        for (var j = i; j < n; j++) {
            if (i % 3 === 0) {
                y +=1
            } else if (i % 3 === 1) {
                x += 1
            } else {
                y -= 1
                x -= 1
            }
            triangle[y][x] = count++
        }
    }

    for (var i = 0; i < n; i++) {
        answer = answer.concat(triangle[i])
    }
    
    return answer;
}

solution(4);

// n = 1, 1
// 1

// n = 2, 3(2, 1)
// 1
//2 3
// 1,2,3

// n = 3, 6(3, 2, 1)
//  1
// 2 6
//3 4 5
// 1,2,6,3,4,5

// n = 4, 10 (4, 3, 2, 1)
//   1,
//  2 9,
// 3 10 8
//4 5  6 7
// 1,2,9,3,10,8,4,5,6,7

// n = 5, 15 (5, 4, 3, 2, 1)