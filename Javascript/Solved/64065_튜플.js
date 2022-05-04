function solution(s) {
    var answer = [];
    
    // 분할하기
    // {} 지우고, ,로 split
    
    s = s.slice(2, s.length-2);
    console.log(s)
    var arr = s.split("},{")
    var newArr = arr.map(item => {
        return item.split(",")
    })
    
    newArr.sort((a, b) => {
        return a.length - b.length
    })

    console.log(newArr)

    answer.push(newArr[0][0])

    // newArr[0]
    for (var i=1; i<newArr.length; i++) {
        for (var j=0; j<newArr[i].length; j++) {

            if (!answer.includes(newArr[i][j])) {
                answer.push(newArr[i][j])
            }
        }
        
    }



    // 표현하는 튜플 값을 배열에 담아 리턴
    
    
    return answer.map(item=> item*1);
}

solution("{{1,2,3},{2,1},{1,2,4,3},{2}}");