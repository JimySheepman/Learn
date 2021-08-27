function standart_us(x, n) {
  let result = 1;
  for (let i = 0; i < n; i++) {
    result *= x;
  }
  return result;
}

console.log(standart_us(2, 3));

function recursion_us(x, n) {
  if (n == 1) {
    return x;
  } else {
    return x * recursion_us(x, n - 1);
  }
}

console.log(recursion_us(2, 3));



// let arr=[obj1,obj2,obj3];

let list = {
    deger:1,
    sonraki:{
        deger:2,
        sonraki:{
            deger:3,
            sonraki:{
                deger:4,
                sonraki:null
            }
        }
    }
}

console.log(list);

let list_2 = {value : 1};
list_2.next = {value: 2};
list_2.next.next = {value : 3};
list_2.next.next.next = {value : 4};

console.log(list_2);

let firma = {
    satis: [{
      adi: 'Ahmet',
      maasi: 1000
    }, {
      adi: 'Mehmet',
      salary: 150
    }],
  
    gelistirme: {
      siteler: [{
        adi: 'Mustafa',
        ucret: 200
      }, {
        adi: 'Mazlum',
        ucret: 50
      }],
  
      dahili: [{
        adi: 'Zafer',
        ucret: 1300
      }]
    }
  };
  
  // İşi yapan fonksiyon
  function maaslariTopla(firma) {
    if (Array.isArray(firma)) { // (1). durum
      return firma.reduce((onceki, suanki) => onceki + suanki.salary, 0); // diziyi topla
    } else { // (2.) durum
      let toplam = 0;
      for(let altDep of Object.values(altDep)) {
        sum += maaslariTopla(altDep); // özçağrı ile alt departmanların çağrılması, bunu sum ile topla.
      }
      return sum;
    }
  }
  
  alert(maaslariTopla(firma)); 
console.log(maaslariTopla(firma));
