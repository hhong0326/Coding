function solution(user_id, banned_id) {
    var answer = 0;


    // *을 제외하고 가능한 모든 id 찾기
    // banned의 갯수만큼 제거하는 경우의 수를 구해야함

    var obj = {};
    var sameMap = new Map();
    banned_id.forEach(bid => {
        
        sameMap.has(bid) ? sameMap.set(bid, sameMap.get(bid) + 1) : sameMap.set(bid, 1)
            
        for (var i=0; i<user_id.length; i++) {
            
            var flag = true;
            if (user_id[i].length == bid.length) {
                for (var j=0; j<bid.length; j++) {
                    if (bid[j] != "*" && bid[j] != user_id[i][j]) {
                        flag = false
                        break
                    }
                }

                if (flag) {
                    if (obj[user_id[i]]) {
                        obj[user_id[i]].push(bid)
                    } else {
                        obj[user_id[i]] = [bid]
                    }
                }
            }
        }
    });

    console.log(obj)
    console.log(sameMap)

    var len = banned_id.length;
    

    // 


    // 모든 경우의 수
    
    return answer;
}

solution(["frodo", "fradi", "crodo", "abc123", "frodoc"], ["fr*d*", "*rodo", "******", "******"]);