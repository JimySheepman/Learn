let kullanici = {
    adi: "Ahmet",
    yasi: 30
  };
  
  //  değerler üzerinden döngü
  for(let deger of Object.values(kullanici)) {
    alert(deger); // Ahmet,sonrasında 30
  }


  let fiyatlar = {
    muz: 1,
    portakal: 2,
    et: 4,
  };
  
  let ikiKatiFiyatlar = {};
  for(let [product, price] of Object.entries(fiyatlar)) {
    ikiKatiFiyatlar[product] = price * 2;
  }
  
  alert(ikiKatiFiyatlar.et); // 8