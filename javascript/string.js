let tek = "tek-tırnak";
let cift = "çift-tırnak";
let us_isareti = `üs işareti`;

function toplam(a, b) {
  return a + b;
}

alert(`1 + 2 = ${toplam(1, 2)}.`);

let davetliListesi = `davetliler:
* ihsan
* cemal
* murtaza
`;

alert(davetliListesi);

let davetliListesi = "Davetliler:\n * İhsan\n * Cemal\n * Muzaffer";

alert(davetliListesi);

alert("Merhaba\nDünya");

alert(`Merhaba Dünya`);

alert("\u00A9");
alert("\u{20331}");
alert("\u{1F60D}");

alert("N'aber canım - Tavşan !");

alert(`\\`);

alert(`Naber\n`.length);

let str = `Selam`;

// ilk karakter
alert(str[0]); // S
alert(str.charAt(0)); // S

// son karakter
alert(str[str.length - 1]); // m

let str = `Selam`;

alert(str[1000]); // undefined
alert(str.charAt(1000)); // '' (boş karakter)

for (let karakter of "Selam") {
  alert(karakter); // S,e,l,a,m (karakter önce "S", sonra "e", sonra "a" vs)
}

let str = "Selam";

str[0] = "s"; // hata
alert(str[0]); // çalışmaz, değişiklik olmaz


let str = 'Selam';

str = str[0] + 'ELAM' ;  // karakter dizisini tamamen değiştir.

alert( str ); // SELAM

alert( 'Arayüz'.toUpperCase() ); // ARAYÜZ
alert( 'Arayüz'.toLowerCase() ); // arayüz

alert( 'Arayüz'[0].toLowerCase() ); // 'a'

let str = "N`aber Canım - Tavşan";

alert( str.indexOf("N'aber") ); // 0, çünkü N`aber başlangıçta
alert( str.indexOf("n'aber") ); // -1, bulunamadı, arama büyük/küçük harf duyarlıdır.

alert( str.indexOf("Tavşan") ); // 15, "Tavşan" 15. pozisyonda bulunmaktadır.

let str = "N`aber Canım - Tavşan";

alert( str.indexOf('an', 9) ) // 19

let str = 'Bir berber bir berbere gel birader beraber bir berber dükkanı açalım demiş';

let hedef = 'bir';

let poz = 0;
while (true) {
  let bulunanPoz = str.indexOf(hedef, poz);
  if (bulunanPoz == -1) break;

  alert( `Bulunan poz: ${bulunanPoz}` );
  poz = bulunanPoz + 1; // bir sonraki pozisyondan aramaya devam et.
}

let str = 'Bir berber bir berbere gel birader beraber bir berber dükkanı açalım demiş';
let hedef = "bir";


let poz = -1;
while ((poz = str.indexOf(hedef, poz + 1)) != -1) {
  alert( poz );
}

let str = "Bin berber bir berbere gel birader beraber bir berber dükkanı açalım demiş";

if (str.lastIndexOf("Bin")) {
    alert("Buldum!"); // çalışmaz!
}

let str = "Bin berber bir berbere gel birader beraber bir berber dükkanı açalım demiş";

if (str.indexOf("Bin") != -1) {
    alert("Buldum"); // Şimdi oldu!
}


alert( ~2 ); // -3,  -(2+1) demektir.
alert( ~1 ); // -2,  -(1+1) demektir.
alert( ~0 ); // -1,  -(0+1) demektir.
alert( ~-1 ); // 0,  -(-1+1) demektir.

let str = "Bin berber bir berbere gel birader beraber bir berber dükkanı açalım demiş";

if (~str.indexOf("Bin")) {
  alert( 'Buldum!' ); // Çalıştı
}

alert( "Bin berber bir berbere gel birader beraber bir berber dükkanı açalım demiş".includes("Bin") ); // true

alert( "Merhaba".includes("Güle Güle") ); // false


alert( "birader".includes("ir") ); // true
alert( "birader".includes("ir", 3) ); // false, 3. pozisyondan sonra `ir` bulunmamaktadır.

alert( "birader".startsWith("bir") ); // true, "birader" "bir" ile başlar.
alert( "birader".endsWith("er") );   // true, "birader" "er" ile biter.

let str = "stringify";
alert( str.slice(0,5) ); // 'strin',  0 ile 5 arasındaki alt karakter dizisi (5 dahil değil)
alert( str.slice(0,1) ); // 's', 0 ile 1, fakat 1 dahil değil, yani sadece 0'ıncı karakter.


let str = "stringify";
alert( str.slice(2) ); // ringify, ikinci pozisyondan sonuna kadar.

let str = "stringify";

// sağdan 4. pozisyondan başla ve yine sağdan 1. pozisyona kadar al.
alert( str.slice(-4, -1) ); // gif

let str = "stringify";

// alt karakter dizisi için aynıdır.
alert( str.substring(2, 6) ); // "ring"
alert( str.substring(6, 2) ); // "ring"

// ...fakat slice için farklıdır:
alert( str.slice(2, 6) ); // "ring" (aynı)
alert( str.slice(6, 2) ); // "" (boş karakter)

let str = "stringify";
alert( str.substr(2, 4) ); // ring, 2. pozisyondan 4 karakter al.

let str = "stringify";
alert( str.substr(-4, 2) ); // gi, 4. pozisyondan 2 karakter al.

alert( 'a' > 'Z' ); // true

alert( 'Österreich' > 'Zealand' ); // true

// Büyük küçük harflerde farklı kodlar döndürülür.
alert( "z".codePointAt(0) ); // 122
alert( "Z".codePointAt(0) ); // 90
alert( String.fromCodePoint(90) ); // Z
// 90 hexa decimal sistemde 5a ya denk gelmektedir.
alert( '\u005a' ); // Z
let str = '';

for (let i = 65; i <= 220; i++) {
  str += String.fromCodePoint(i);
}
alert( str );
// ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
// ¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜ
alert( 'Österreich'.localeCompare('Zealand') ); // -1
alert( '𝒳'.length ); // 2, MATHEMATICAL SCRIPT CAPITAL X
alert( '😂'.length ); // 2, FACE WITH TEARS OF JOY
alert( '𩷶'.length ); // 2, a rare chinese hieroglyph
alert( '𝒳'[0] ); // garip semboller...
alert( '𝒳'[1] ); // ...her biri ikilinin parçaları
// charCodeAt çiftlere uygun değildir, bundan dolayı sadece kodlar verilir.
alert( '𝒳'.charCodeAt(0).toString(16) ); // d835, 0xd800 ile 0xdbff arasında
alert( '𝒳'.charCodeAt(1).toString(16) ); // dcb3, 0xdc00 ile 0xdfff arasında
alert( 'S\u0307' ); // Ṡ
alert( 'S\u0307\u0323' ); // Ṩ
alert( 'S\u0307\u0323' ); // Ṩ, S + üst nokta + alt nokta
alert( 'S\u0323\u0307' ); // Ṩ, S + alt nokta + üst nokta
alert( 'S\u0307\u0323' == 'S\u0323\u0307' ); // false
alert( "S\u0307\u0323".normalize() == "S\u0323\u0307".normalize() ); // true
alert( "S\u0307\u0323".normalize().length ); // 1
alert( "S\u0307\u0323".normalize() == "\u1e68" ); // true

/*
Özet
3 tip tırnak bulunmaktadır. “`” işareti ile birkaç satırdan oluşan karakter dizisi yazmak mümkündür
JavaScript’te karakterler UTF-16 ile kodlanmıştır.
\n gibi özel karakterler veya \u.. ile unicode kullanılabilir.
Karakteri almak için: [] kullanılır.
Alt karakter kümesi almak için slice veya substring kullanılır.
Küçük/büyük harf değişimi için: toLowerCase/toUpperCase.
Alt karakter dizisi aramak için : indexOf veya includes/startsWith/endsWith kullanılabilir.
Karakterleri dile göre karşılaştırmak için localceCompare kullanılabilir. Diğer türlü karakterler kodlarına göre karşılaştırılırlar.
Bunun yanında karakter dizileri için daha başka yardımcı metodlar bulunmaktadır:

str.trim() – başlangıç ve bitişteki boşlukları siler.
str.repeat(n) – str'yi istendiği kadar tekrar eder…
… Daha fazlası için manual adresine bakabilirsiniz.
Karakter dizileri bunun yanında arama/değiştirme veya regular expression için metodlar barındırmaktadır. Fakat bu konular ayrı bölümleri hak etmektedir. Bu konulara ilerleyen bölümlerde dönülecektir.

*/