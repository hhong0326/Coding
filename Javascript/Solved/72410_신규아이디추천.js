// 정규식 연습

function solution(new_id) {
    var answer = '';
    
    // 1
    new_id = new_id.toLowerCase();
    
    // 2

    // 해당 문자 제외한 나머지 문자 지우기
    new_id = new_id.replace(/[^\-_.|^a-z|^0-9]/g, "");

    // 해당 문자만 지우기 and 특수문자 앞에는 \ 들어가기
    new_id = new_id.replace(/[\^]/g, "");
    
    // /^[ㄱ-ㅎ|가-힣|a-z|A-Z|0-9|]+$/
    
    // 3
    // 2회 이상 연속된 중복 문자 하나로 만들기
    new_id = new_id.replace(/\.{2,}/g, ".")
    
    // 4
    // 앞 뒤 문자 slice 얕은 복사
    if (new_id[0] === '.') {
        new_id = new_id.slice(1);
    }
    if (new_id[new_id.length-1] === '.'){
        new_id = new_id.slice(0, new_id.length-1);
    }
    
    
    // 5
    // 문자열에 + 
    if (new_id === "") {
        new_id += "a";
    }
    
    // 6
    // slice(a, b) b에 해당하는 값은 미포함!
    if (new_id.length > 15) {
        new_id = new_id.slice(0, 15);
        if (new_id[0] === '.') {
            new_id = new_id.slice(1);
        }
        if (new_id[new_id.length-1] === '.'){
            new_id = new_id.slice(0, new_id.length-1);
        }
    }
    
    // 7
    // 3 미만이면 마지막 문자를 복사해서 더하기
    while (new_id.length < 3) {
        new_id += new_id[new_id.length-1];
    }
    
    answer = new_id;
    return answer;
}