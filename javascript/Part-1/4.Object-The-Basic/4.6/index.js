let user = null;

alert(user?.address); // undefined
alert(user?.address.street); // undefined

let user_2 = null;
let x = 0;

user_2?.sayHi(x++); // no "sayHi", so the execution doesn't reach x++

alert(x); // 0, value not incremented
