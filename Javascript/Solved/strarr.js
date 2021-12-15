

//문자열 잘라서 배열에 넣기
var str1 = '123'
str1.split('') // [1,2,3]
var str2 = '1-2-3'
str2.split('-') //[1,2,3]
//문자열이 ~로만 이루어진걸로 배열로 리턴
var str = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
var regexp = /[A-E]/gi;
var matches_array = str.match(regexp);

console.log(matches_array);
// ['A', 'B', 'C', 'D', 'E', 'a', 'b', 'c', 'd', 'e']


//문자열 자르기

//slice는 새로운 string을 생성
var str2 = '012345678'
str2.slice(0,8) //01234567

str2.slice(1,7) //123456

//substr은 string의 일부를 리턴

str2.substr(1,3) //1번인덱스부터 3개
str2.substr(1) //1번 인덱스부터 쭉

//substring은 시작인덱스와 마지막인덱스 전까지 리턴

str2.substring(1,3) //1번인덱스부터 3번인덱스 전까지 갯수는 3-1
str2.substring(1) //1번 인덱스부터 쭉 str.substr(1)이랑 동일

//다 소문자로
str2.toLowerCase()
//다 대문자로
str2.toUpperCase()




//배열

//fill .fill(무엇으로,언제부터,언제전까지) 채울꺼냐

[1,2,3].fill(4) //[4,4,4]

[1,2,3].fill(4,1) //[1,4,4]

[1,2,3].fill(4,1,2) //[1,4,3]


//join  배열을 문자열로

var ele = ['a','b','c']

ele.join() //'a,b,c'
ele.join('') //'abc'
ele.join('-') //'a-b-c'

