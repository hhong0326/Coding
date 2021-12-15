function solution(str1, str2) {
    var answer = 0;
    
    var str1Arr = [];
    var str2Arr = [];
    
    var str1Set = new Set();
    var str2Set = new Set();
    
    for (var i=0; i<str1.length-1; i++) {
        str1Arr.push(str1[i]+str1[i+1]);
        str1Set.add(str1[i]+str1[i+1]);
    }
    
    for (var i=0; i<str2.length-1; i++) {
        str2Arr.push(str2[i]+str2[i+1]);
        str2Set.add(str2[i]+str2[i+1]);
    }
    
    var union = new Set([...str1Set, ...str2Set]);

    console.log(str1Set)
    
    var intersection = new Set(
    [...str1Set].filter(item => str2Set.has(item)));
    console.log(intersection);

    
    // console.log(str1Arr, str2Arr)
    
    var regex = /^[a-z|A-Z]+$/;
    
    return answer;
}