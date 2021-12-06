let billion = 1000000000;
let billion_2 = 1_000_000_000;
let billion_3 = 1e9;
console.log(billion, billion_2, billion_3);
console.log(7.32e9);

console.log(1e3 === 1 * 1000);
console.log(1.23e6 === 1.23 * 1000000);

let ms = 0.000001;
let ms_2 = 1e-6;
console.log(ms, ms_2);

console.log(1e-3 === 1 / 1000);
console.log(1.23e-6 === 1.23 / 1000000);

let num = 255;
console.log(num.toString(16));
console.log(num.toString(2));

console.log((123456).toString(36));

let number = 1.23456;
console.log(Math.round(number * 100) / 100);

let number_2 = 12.34;
console.log(number_2.toFixed(1));

let number_3 = 12.34;
console.log(number_3.toFixed(5));

console.log(1e500);
console.log(0.1 + 0.2 == 0.3);
console.log(0.1 + 0.2);

console.log((0.1).toFixed(20));

console.log((0.1 + 0.2).toFixed(2));
console.log(+(0.1 + 0.2).toFixed(2));

console.log((0.1 * 10 + 0.2 * 10) / 10);
console.log((0.28 * 100 + 0.14 * 100) / 100);

console.log(9999999999999999);

console.log(isNaN(NaN));
console.log(isNaN("str"));
console.log(NaN === NaN);
console.log(isFinite("15"));
console.log(isFinite("str"));
console.log(isFinite(Infinity));

console.log(+"100px");
console.log(parseInt("100px"));
console.log(parseFloat("12.5em"));
console.log(parseInt("12.3"));
console.log(parseFloat("12.3.4"));

console.log(Math.random());
console.log(Math.max(3, 5, -10, 123, 1));
console.log(Math.min(3, 5, -10, 123, 1));
console.log(Math.pow(2, 8));
