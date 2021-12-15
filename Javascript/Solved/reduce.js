
//reduce

//배열.reduce((누적값, 현잿값, 인덱스, 요소) => { return 결과 }, 초깃값);
var initialValue = 0;
var sum = [{x: 1}, {x:2}, {x:3}].reduce(function (acc, cur) {
    return acc + cur.x;
},initialValue)

console.log(sum) // logs 6

var flattened = [[0, 1], [2, 3], [4, 5]].reduce(
    function(accumulator, currentValue) {
      return accumulator.concat(currentValue);
    },
    []
  );
  // flattened is [0, 1, 2, 3, 4, 5]


var names = ['Alice', 'Bob', 'Tiff', 'Bruce', 'Alice'];

var countedNames = names.reduce(function (allNames, name) { 
  if (name in allNames) {
    allNames[name]++;
  }
  else {
    allNames[name] = 1;
  }
  return allNames;
}, {});
// countedNames is:
// { 'Alice': 2, 'Bob': 1, 'Tiff': 1, 'Bruce': 1 }

const oneTwoThree = [1, 2, 3];

result = oneTwoThree.reduce((acc, cur) => {
    acc.push(cur % 2 ? '홀수' : '짝수');
    return acc;
  }, []);
  result; // ['홀수', '짝수', '홀수']









