/*
new Date()
new Date(milisaniye)
new Date(tarih_metni)
new Date(yıl, ay, gün, saat, dakika, saniye, milisaniye)
getFullYear()
getMonth()
getDate()
getHours()
getMinutes()
getSeconds()
getMilliseconds()
getDay()
getTime()
getTimezoneOffset()


Tarih bileşeninin ayarlama

setFullYear(year [, month, date])
setMonth(month [, date])
setDate(date)
setHours(hour [, min, sec, ms])
setMinutes(min [, sec, ms])
setSeconds(sec [, ms])
setMilliseconds(ms)
setTime(milliseconds)


Date.now() 
new Date().getTime() 


Karakter dizisinden Date.parse ile tarih alma.
Date.parse(str) metodu karakterden tarih ayrıştırmaya yarar.

Metin formatı: YYYY-MM-DDTHH:mm:ss.sssZ şeklindedir, burada :

YYYY-MM-DD – tarih : yıl-ay-gün
"T" karakteri ayraç.
HH:mm:ss.sss – zaman: saat:dakika:saniye.sarise şeklindedir.
İsteğe bağlı olarak eklenen 'Z' +-hh:mm şeklinde UTC’ye göre zaman ayarlamaya yarar. Varsayılan Z değeri UTC+0 anlamına gelir.
Daha kısa YYYY-MM-DD veya YYYY-MM hatta YYYY gibi şeklinde bile olabilir.

Date.parse(str) çağrısı verilen formatta karakterleri alır ve timestamp( 1 Ocak 1970 UTC+0’dan itibaren geçen sarise ) olarak geri döner. Eğer format doğru değilse, NaN döner.


*/


let now = new Date();
alert( now ); // o anki tarih/saati gösterir.

// `0`  01.01.1970 UTC+0 demektir.
let Jan01_1970 = new Date(0);
alert( Jan01_1970 );

// Buna 24 saat eklemek için, get 02.01.1970 UTC+0
let Jan02_1970 = new Date(24 * 3600 * 1000);
alert( Jan02_1970 );

function diffSubtract(date1, date2) {
    return date2 - date1;
  }
  
  function diffGetTime(date1, date2) {
    return date2.getTime() - date1.getTime();
  }
  
  function bench(f) {
    let date1 = new Date(0);
    let date2 = new Date();
  
    let start = Date.now();
    for (let i = 0; i < 100000; i++) f(date1, date2);
    return Date.now() - start;
  }
  
  alert( 'Time of diffSubtract: ' + bench(diffSubtract) + 'ms' );
  alert( 'Time of diffGetTime: ' + bench(diffGetTime) + 'ms' );


  function diffSubtract(date1, date2) {
    return date2 - date1;
  }
  
  function diffGetTime(date1, date2) {
    return date2.getTime() - date1.getTime();
  }
  
  function bench(f) {
    let date1 = new Date(0);
    let date2 = new Date();
  
    let start = Date.now();
    for (let i = 0; i < 100000; i++) f(date1, date2);
    return Date.now() - start;
  }
  
  let time1 = 0;
  let time2 = 0;
  
  // Paketi 10 defa çalışacak şekilde ayarlayın
  for (let i = 0; i < 10; i++) {
    time1 += bench(diffSubtract);
    time2 += bench(diffGetTime);
  }
  
  alert( 'Total time for diffSubtract: ' + time1 );
  alert( 'Total time for diffGetTime: ' + time2 );


  alert(`Loading started ${performance.now()}ms ago`);
// Sonuç : Yüklemeye 4731.26000000001ms önce başladı
// .26 mikrosaniye (260 mikrosaniye)
// noktanın 3. basamağından sonraki değerler sapmadır fakat ilk 3 basamak doğrudur.