function solution(user_id, banned_id) {
    var answer = 0;


    // *을 제외하고 가능한 모든 id 찾기
    // banned의 갯수만큼 제거하는 경우의 수를 구해야함

    var obj2 = {};
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

                    if (obj2[bid]) {
                        if (!obj2[bid].includes(user_id[i])) {
                            obj2[bid].push(user_id[i])
                        }
                    } else {
                        obj2[bid] = [user_id[i]]
                    }
                }
            }
        }
    });
    console.log(obj2)
    console.log(sameMap)

    // 모든 경우의 수
    var s = new Set();

    function backtrack(obj, banned_id, arr = [], idx) {
        if (idx == banned_id.length) {
            // sort
            arr.sort();
            var str = arr.join(".");
            return s.add(str);
        } else {
            for (var i=0; i<obj[banned_id[idx]].length; i++) {
                if(!arr.includes(obj[banned_id[idx]][i])) {
                    backtrack(obj2, banned_id, [...arr, obj[banned_id[idx]][i]], idx+1)
                }
            }
        }
    }
    
    backtrack(obj2, banned_id, [], 0);

    return s.size;
}


solution(["frodo", "fradi", "crodo", "abc123", "frodoc"], ["fr*d*", "*rodo", "******", "******"]);