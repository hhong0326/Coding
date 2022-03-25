// 정규식 연습

function solution(new_id) {
    var answer = '';
    
    // 1
    new_id = new_id.toLowerCase();
    
    // 2

    new_id = new_id.replace(/[^\-_.|^a-z|^0-9]/g, "");
    new_id = new_id.replace(/[\^]/g, "");
    
    
    // 3
    new_id = new_id.replace(/\.{2,}/g, ".")
    
    // 4
    if (new_id[0] === '.') {
        new_id = new_id.slice(1);
    }
    if (new_id[new_id.length-1] === '.'){
        new_id = new_id.slice(0, new_id.length-1);
    }
    
    
    // 5
    if (new_id === "") {
        new_id += "a";
    }
    
    // 6
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
    while (new_id.length < 3) {
        new_id += new_id[new_id.length-1];
    }
    
    answer = new_id;
    return answer;
}