let mil = 1000000000;
let milyar = 1e9;

alert(7.3e9);

1e3 = 1 * 1000;
1.23e6 = 1.23 * 1000000;

let ms = 0.000001;
let ms = 1e-6;

1e-3 = 1 / 1000;
1.23e-6 = 1.23 / 100000;

alert(0xff);
alert(0xFF);
let a = 0b11111111;
let b = 0o377;


alert(a == b);
let sayi = 255;

alert(sayi.toString(16));
alert(sayi.toString(2));
alert(123456..toString(36));

let num = 1.23456;

alert( Math.floor(num * 100) / 100 );

let num = 12.34;
alert( num.toFixed(1) );

let num = 12.36;
alert( num.toFixed(1) );

let num = 12.34;
alert( num.toFixed(5) );

alert( 1e500 );

alert(0.1 + 0.2 == 0.3);

alert(0.1 + 0.2);

alert( 0.1.toFixed(20) );

let toplam = 0.1 + 0.2;
alert( toplam.toFixed(2));

let toplam = 0.1 + 0.2;
alert( +toplam.toFixed(2) );

alert( (0.1 * 10 + 0.2 * 10) / 10 );

alert( 9999999999999999 );

alert( isNaN(NaN) ); 
alert( isNaN("str") );

alert( NaN === NaN );

alert( isFinite("15") ); 
alert( isFinite("str") ); 
alert( isFinite(Infinity) );


let num = +prompt("Bir sayı giriniz", '');


alert( isFinite(num) );


alert( +"100px" );


alert( parseInt('100px') ); 
alert( parseFloat('12.5em') );

alert( parseInt('12.3') ); 
alert( parseFloat('12.3.4') ); 


alert( parseInt('a123') );


alert( parseInt('0xff', 16) );

alert( parseInt('ff', 16) ); 
alert( parseInt('2n9c', 36) );


alert( Math.random() );
alert( Math.random() );
alert( Math.random() );


alert( Math.max(3, 5, -10, 0, 1) ); 
alert( Math.min(1, 2) ); 


alert( Math.pow(2, 10) );