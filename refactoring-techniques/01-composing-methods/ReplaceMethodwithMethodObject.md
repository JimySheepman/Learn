# Replace Method with Method Object

- **Sorun:** Yerel değişkenlerin o kadar iç içe olduğu uzun bir yönteminiz var ki Extract Method uygulayamazsınız .

```Java
class Order {
  // ...
  public double price() {
    double primaryBasePrice;
    double secondaryBasePrice;
    double tertiaryBasePrice;
    // Perform long computation.
  }
}
```

- **Çözüm:** Yerel değişkenlerin sınıfın alanları olması için yöntemi ayrı bir sınıfa dönüştürün. Ardından, yöntemi aynı sınıf içinde birkaç yönteme bölebilirsiniz.

```Java
class Order {
  // ...
  public double price() {
    return new PriceCalculator(this).compute();
  }
}

class PriceCalculator {
  private double primaryBasePrice;
  private double secondaryBasePrice;
  private double tertiaryBasePrice;
  
  public PriceCalculator(Order order) {
    // Copy relevant information from the
    // order object.
  }
  
  public double compute() {
    // Perform long computation.
  }
}
```

## Neden Refactor

Bir yöntem çok uzun ve birbirinden izole edilmesi zor yerel değişkenlerin karışık kütleleri nedeniyle onu ayıramazsınız.

İlk adım, tüm yöntemi ayrı bir sınıfa ayırmak ve yerel değişkenlerini sınıfın alanlarına dönüştürmektir.

İlk olarak, bu, sorunun sınıf düzeyinde izole edilmesini sağlar. İkincisi, büyük ve hantal bir yöntemi, orijinal sınıfın amacına zaten uymayan daha küçük olanlara bölmenin yolunu açar.

## Faydalar

Uzun bir yöntemi kendi sınıfında yalıtmak, bir yöntemin boyut olarak balonlaşmasını durdurmaya izin verir. Bu aynı zamanda, orijinal sınıfı faydalı yöntemlerle kirletmeden, onu sınıf içinde alt yöntemlere ayırmaya da izin verir.

## Dezavantajları

Programın genel karmaşıklığını artıran başka bir sınıf eklenir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yeni bir sınıf oluşturun. Yeniden düzenlediğiniz yöntemin amacına göre adlandırın.

2. Yeni sınıfta, yöntemin daha önce bulunduğu sınıfın bir örneğine başvuruyu depolamak için özel bir alan oluşturun. Gerekirse orijinal sınıftan bazı gerekli verileri almak için kullanılabilir.

3. Yöntemin her yerel değişkeni için ayrı bir özel alan oluşturun.

4. Yöntemin tüm yerel değişkenlerinin değerlerini parametre olarak kabul eden ve ayrıca karşılık gelen özel alanları başlatan bir kurucu oluşturun.

5. Ana yöntemi bildirin ve yerel değişkenleri özel alanlarla değiştirerek orijinal yöntemin kodunu ona kopyalayın.

6. Bir yöntem nesnesi oluşturarak ve ana yöntemini çağırarak orijinal yöntemin gövdesini orijinal sınıfta değiştirin.
