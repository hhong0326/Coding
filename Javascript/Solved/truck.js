function solution(bridge_length, weight, truck_weights) {
    var answer = 0;
    var bridge = new Array(bridge_length);
    
    console.log(bridge);
    var t = truck_weights.shift();
    bridge.push(t);
    
    var sum = t;

    while(bridge.length !== 0) {
        answer++;
        if (bridge.length !== 0) {
            sum = bridge.reduce((pre, value) => pre + value);
        }
        bridge.shift();
        
        
        
        if (truck_weights[0] + sum  <= weight) { // 무게 확인
                t = truck_weights.shift();
                bridge.push(t);
                
        } else {
            bridge.push(0);
        }  
    }
    
    
    return answer;
}
solution(2, 10, [7,4,5,6])