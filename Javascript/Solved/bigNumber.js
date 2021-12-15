function solution(numbers) {
    var answer = '';
    
    answer = numbser.map(item => item + '').sort((a, b) => (b+a) - (a+b)).join('');

    
    return answer[0] === '0' ? '0' : answer;
}
