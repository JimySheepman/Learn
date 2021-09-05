let total = 0, count = 1;
while(count <= 10){
    total  +=count;
    count +=1;
}
console.log(total);

function factorial(n){
    if (n == 0) {
        return 1;        
    }else{
        return factorial(n-1)*n;
    }
}
console.log(factorial(8));


console.log("Values:");
console.log("1. number = 13 or 9.81 or 2.99e8(2.998 × 108 = 299,800,000.)");
console.log("2. aritmethic =  (100+4)*11  or 100+4*11");
console.log("3. special number = Infinity - Infinity or 0/0");
console.log("4. string = `merlins` or "+ "merlins" +" or 'merlins' ");

console.log("Unary operators:");
console.log(typeof 4.5);
console.log(typeof "x");
console.log(-(10-2));

console.log("Boolean Values");
console.log(3>2);
console.log(3<2);
console.log("aardvark" < "zoroaster");
console.log(NaN == NaN);

console.log("Logical Operators");
console.log(true && false);
console.log(true || false);
console.log(1 + 1 == 2 && 10 * 10 > 50);
console.log(true ? 1 : 2);
console.log(false ? 1 : 2);


console.log("Automatic Type Conversion");
console.log(8 * null);
console.log("5" - 1);
console.log("5" + 1);
console.log("five" * 2);
console.log(false == 0);
console.log(null == undefined);
console.log(null == 0);


console.log("Short-Circuiting of Logical Operators");
console.log(null || "user");
console.log("Agnes" || "user");