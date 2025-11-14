let myName = prompt("Lütfen adınızı giriniz:");
document.getElementById("myName").innerHTML = myName;

var saatim = document.getElementById("saat");
var gunum = document.getElementById("gunn");

setInterval(function () {
  var tarih = new Date();
  var yil = tarih.getFullYear();
  var ay = tarih.getMonth() + 1;
  //var gun = tarih.getDate();
  var saat = tarih.getHours();
  var dk = tarih.getMinutes();
  var sn = tarih.getSeconds();
  var gun = [
    "Pazartesi",
    "Salı",
    "Çarşamba",
    "Perşembe",
    "Cuma",
    "Cumartesi",
    "Pazar",
  ][new Date().getDate()];
  saatim.textContent =
    ("0" + saat).substr(-2) +
    "." +
    ("0" + dk).substr(-2) +
    "." +
    ("0" + sn).substr(-2);
  gunum.textContent = gun;
}, 1000);
