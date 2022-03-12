# Self Encapsulate Field

- **Sorun:** Bir sınıf içindeki özel alanlara doğrudan erişimi kullanırsınız.

```Java
class Range {
  private int low, high;
  boolean includes(int arg) {
    return arg >= low && arg <= high;
  }
}
```

- **Çözüm:** Alan için bir alıcı ve ayarlayıcı oluşturun ve alana erişmek için yalnızca bunları kullanın.

```Java
class Range {
  private int low, high;
  boolean includes(int arg) {
    return arg >= getLow() && arg <= getHigh();
  }
  int getLow() {
    return low;
  }
  int getHigh() {
    return high;
  }
}
```

## Neden Refactor

Bazen bir sınıf içindeki özel bir alana doğrudan erişmek yeterince esnek değildir. İlk sorgu yapıldığında bir alan değeri başlatabilmek veya atandığında alanın yeni değerleri üzerinde belirli işlemler yapabilmek veya belki tüm bunları alt sınıflarda çeşitli şekillerde yapabilmek istiyorsunuz.

## Faydalar

- Alanlara dolaylı erişim , bir alana erişim yöntemleri (alıcılar ve ayarlayıcılar) aracılığıyla işlem yapıldığında gerçekleşir. Bu yaklaşım, alanlara doğrudan erişimden çok daha esnektir .

  - İlk olarak, sahadaki veriler ayarlandığında veya alındığında karmaşık işlemler gerçekleştirebilirsiniz. Alan değerlerinin tembel başlatma ve doğrulaması, alan alıcıları ve ayarlayıcıları içinde kolayca uygulanır.

  - İkincisi ve daha da önemlisi, alt sınıflardaki alıcıları ve ayarlayıcıları yeniden tanımlayabilirsiniz.

-Bir alan için hiç ayarlayıcı uygulamama seçeneğiniz vardır. Alan değeri yalnızca yapıcıda belirtilecektir, böylece alanı tüm nesne ömrü boyunca değiştirilemez hale getirir.

## Dezavantajları

Alanlara doğrudan erişim kullanıldığında, esneklik azalmasına rağmen kod daha basit ve daha prezentabl görünür.

## Yeniden Düzenleme Nasıl Yapılır?

1. Alan için bir alıcı (ve isteğe bağlı ayarlayıcı) oluşturun. protected ya da public olmalılar.

2. Alanın tüm doğrudan çağrılarını bulun ve bunları alıcı ve ayarlayıcı çağrılarıyla değiştirin.
