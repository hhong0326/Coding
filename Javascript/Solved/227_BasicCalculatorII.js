/**
 * @param {string} s
 * @return {number}
 */

// Failed, Someone's Solution

 var calculate = function (s) {
    var num = 0;
    var stack = [];
    var sign = "+";
    
    var signNum = function (sign, num) {
      if (sign == "*") {
        stack.push(stack.pop() * num);
      } else if (sign == "/") {
        stack.push(stack.pop() / num | 0);
      } else if (sign == "-") {
        stack.push(-num);
      } else if (sign = "+") {
        stack.push(+num);
      }
    }
    
    for (var i=0; i<s.length; i++) {
      if (s[i] >= '0' && s[i] <= '9') {
        num = num * 10 + Number(s[i]);
      } else if (s[i] != " ") {
        signNum(sign, num);
        sign = s[i];
        num = 0;
      }
    }
    signNum(sign, num);
    return stack.reduce((a, b) => a + b);
    
  };