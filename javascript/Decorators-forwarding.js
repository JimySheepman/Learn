function slow(x) {
    // burada baya yoğun işlemci gücüne ihtiyaş duyan işler yapılmaktadır.
    console.log(`${x} ile çağırıldı`);
    return x;
  }
  
  function cachingDecorator(func) {
    let cache = new Map();
  
    return function(x) {
      if (cache.has(x)) { // eğer sonuç map içerisinde ise
        return cache.get(x); // değeri gönder
      }
  
      let result = func(x); // aksi halde hesap yap
  
      cache.set(x, result); // sonra sonucu sakla
      return result;
    };
  }
  
  slow = cachingDecorator(slow);
  
  console.log( slow(1) ); // slow(1) saklandı
  console.log( "Tekrar: " + slow(1) ); // aynısı döndü
  
  console.log( slow(2) ); // slow(2) saklandı
  console.log( "Tekrar: " + slow(2) ); // bir önceki ile aynısı döndü.



  function sayHi() {
    console.log(this.name);
  }
  
  let user = { name: "John" };
  let admin = { name: "Admin" };
  
  // farklı objeler "this" objesi olarak gönderilebilir.
  sayHi.call( user ); // John
  sayHi.call( admin ); // Admin


  function say(phrase) {
    console.log(this.name + ': ' + phrase);
  }
  
  let users = { name: "John" };
  
  // user `this` olmakta ve `phrase` ilk argüman olmaktadır.
  say.call( users, "Hello" ); // John: Hello


  let worker = {
    someMethod() {
      return 1;
    },
  
    slow(x) {
        console.log(x + "ile çağırıldı");
      return x * this.someMethod(); // (*)
    }
  };
  
  function cachingDecorator(func) {
    let cache = new Map();
    return function(x) {
      if (cache.has(x)) {
        return cache.get(x);
      }
      let result = func.call(this, x); // "this" is passed correctly now
      cache.set(x, result);
      return result;
    };
  }
  
  worker.slow = cachingDecorator(worker.slow); // now make it caching
  
  console.log( worker.slow(2) ); // çalışır
  console.log( worker.slow(2) ); // orjinali değilde hafızadaki çalışır.


  function say(time, phrase) {
    console.log(`[${time}] ${this.name}: ${phrase}`);
  }
  
  let user_1 = { name: "John" };
  
  let messageData = ['10:00', 'Hello']; // time, phrase'e dönüşür.
  
  // this = user olur , messageData liste olarak (time,phrase) şeklinde gönderilir.
  say.apply(user_1, messageData); // [10:00] John: Hello (this=user)

  function hash() {
    alert( [].join.call(arguments) ); // 1,2
  }
  
  hash(1, 2);

  let worker_1 = {
    slow(min, max) {
      console.log(`${min},${max} ile çağırıldı`);
      return min + max;
    }
  };
  
