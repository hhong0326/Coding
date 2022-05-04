function solution(gems) {
    var answer = [];

    // 보석 종류를 알아야 함
    var s = new Set();
    gems.forEach(gem => {
        s.add(gem);
    })


    if (s.size == 1) {
        return [1,1];
    }

    // 종류가 다 선택되는 최소의 구간

    var l = 0;
    var ll = s.size;

    var start = 0, end = 0;
    var period = [];
           
    // 0 4 / 1 5 / 2 6 ...
    // 0 5 / 1 6 / 2 7 ...
    
    while (ll < gems.length) {
        for (var i=0; i<gems.length-ll; i++) {
            var copySet = new Set(JSON.parse(
                JSON.stringify(Array.from(s))
              ));
            for (var j=i; j<=ll+i; j++) {
                // 체크

                console.log(copySet)
                if (copySet.has(gems[j])) {
                    copySet.delete(gems[j])
                } 
                if (copySet.size == 0) {
                    
                    answer = [i+1, j+1]
                    ll = gems.length;
                    break
                }
            }
        }

        ll++
    }

//     while (l < gems.length-s.size) {

//         var copySet = new Set(s)

//         for (var i=l; i<gems.length; i++) {
//             if (copySet.has(gems[i])) {
//                 copySet.delete(gems[i])
//             } 

//             if (copySet.size == 0) {
//                 start = l
//                 end = i
//                 period.push([start+1, end+1])
//                 break
//             }
//         }
//         l++
//     }


//     var min = gems.length;
//     var minIdx = -1;
//     for (var i=0; i<period.length; i++) {
//         var p = period[i][1] - period[i][0];
//         if (min > p) {
//             min = p;
//             minIdx = i
//         }
//     }

//     answer = period[minIdx];
    


    return answer;
}

// 다른 사람 풀이 Map, Set 활용
function others_solution(gems) {
    const gemVarietyCounts = new Set(gems).size;
    const gemMap = new Map();
    const gemLengths = [];
    gems.forEach((gem,i) => {
        gemMap.delete(gem);
        gemMap.set(gem,i);
        if(gemMap.size === gemVarietyCounts) {
            gemLengths.push([gemMap.values().next().value + 1, i + 1 ]);
        }
    });

    gemLengths.sort((a, b) => {
        if ((a[1] - a[0]) === (b[1] - b[0])) {
            return a[1] - b[1];
        }
        return (a[1] - a[0]) - (b[1] - b[0]);
    });

    return gemLengths[0];
}

solution(["DIA", "RUBY", "RUBY", "DIA", "DIA", "EMERALD", "SAPPHIRE", "DIA"])