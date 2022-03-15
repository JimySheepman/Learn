# Replace Conditional with Polymorphism

- **Sorun:** Nesne türüne veya özelliklerine bağlı olarak çeşitli eylemler gerçekleştiren bir koşula sahipsiniz.

```Java
class Bird {
  // ...
  double getSpeed() {
    switch (type) {
      case EUROPEAN:
        return getBaseSpeed();
      case AFRICAN:
        return getBaseSpeed() - getLoadFactor() * numberOfCoconuts;
      case NORWEGIAN_BLUE:
        return (isNailed) ? 0 : getBaseSpeed(voltage);
    }
    throw new RuntimeException("Should be unreachable");
  }
}
```

- **Çözüm:** Koşulun dallarıyla eşleşen alt sınıflar oluşturun. Onlarda, paylaşılan bir yöntem oluşturun ve kodu, koşulun ilgili dalından ona taşıyın. Ardından koşulluyu ilgili yöntem çağrısıyla değiştirin. Sonuç, nesne sınıfına bağlı olarak polimorfizm yoluyla uygun uygulamanın elde edilmesidir.

```Java
abstract class Bird {
  // ...
  abstract double getSpeed();
}

class European extends Bird {
  double getSpeed() {
    return getBaseSpeed();
  }
}
class African extends Bird {
  double getSpeed() {
    return getBaseSpeed() - getLoadFactor() * numberOfCoconuts;
  }
}
class NorwegianBlue extends Bird {
  double getSpeed() {
    return (isNailed) ? 0 : getBaseSpeed(voltage);
  }
}

// Somewhere in client code
speed = bird.getSpeed();
```

## Neden Refactor

Bu yeniden düzenleme tekniği, kodunuz aşağıdakilere göre değişen çeşitli görevleri gerçekleştiren operatörler içeriyorsa yardımcı olabilir:

- Uyguladığı nesnenin veya arabirimin sınıfı

- Bir nesnenin alanının değeri

- Bir nesnenin yöntemlerinden birinin çağrılmasının sonucu

Yeni bir nesne özelliği veya türü belirirse, tüm benzer koşullularda kod aramanız ve eklemeniz gerekir. Bu nedenle, bir nesnenin yöntemlerinin tümüne dağılmış birden çok koşul varsa, bu tekniğin yararı katlanır.

## Faydalar

- Bu teknik, Söyle-Sor-Sor ilkesine bağlı kalır: bir nesneye durumu hakkında soru sormak ve ardından buna dayalı eylemler gerçekleştirmek yerine, nesneye ne yapması gerektiğini söylemek ve nasıl yapacağına kendisinin karar vermesine izin vermek çok daha kolaydır. bunu yapmak için.

- Yinelenen kodu kaldırır. Neredeyse aynı olan birçok koşuldan kurtulursunuz.

- Yeni bir yürütme varyantı eklemeniz gerekiyorsa, tek yapmanız gereken mevcut koda ( Open/Closed Principle ) dokunmadan yeni bir alt sınıf eklemektir.

## Yeniden Düzenleme Nasıl Yapılır?

### Refactor'a hazırlanıyor

Bu yeniden düzenleme tekniği için, alternatif davranışları içerecek hazır bir sınıf hiyerarşisine sahip olmalısınız. Böyle bir hiyerarşiniz yoksa, bir tane oluşturun. Diğer teknikler bunun gerçekleşmesine yardımcı olacaktır:

- Replace Type Code with Subclasses ile değiştirin . Belirli bir nesne özelliğinin tüm değerleri için alt sınıflar oluşturulacaktır. Bu yaklaşım basittir ancak daha az esnektir çünkü nesnenin diğer özellikleri için alt sınıflar oluşturamazsınız.

- Replace Type Code with State/Strategy ile değiştirin . Belirli bir nesne özelliği için bir sınıf ayrılacak ve özelliğin her değeri için ondan alt sınıflar oluşturulacaktır. Geçerli sınıf, bu tür nesnelere referanslar içerecek ve onlara yürütme yetkisi verecek.

Aşağıdaki adımlar, hiyerarşiyi zaten oluşturduğunuzu varsayar.

## Yeniden Düzenleme Adımları

1. Koşul, diğer eylemleri de gerçekleştiren bir yöntemdeyse, Extract Method gerçekleştirin .

2. Her hiyerarşi alt sınıfı için, koşulluyu içeren yöntemi yeniden tanımlayın ve karşılık gelen koşullu dalın kodunu o konuma kopyalayın.

3. Bu dalı koşulludan silin.

4. Koşul boşalana kadar değiştirmeyi tekrarlayın. Ardından koşulluyu silin ve yöntem özetini bildirin.
