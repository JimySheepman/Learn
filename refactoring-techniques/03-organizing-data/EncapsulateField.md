# Encapsulate Field

- **Sorun:** Public alanınız var.

```Java
class Person {
  public String name;
}
```

- **Çözüm:** Alanı privete yapın ve bunun için erişim yöntemleri oluşturun.

```Java
class Person {
  private String name;

  public String getName() {
    return name;
  }
  public void setName(String arg) {
    name = arg;
  }
}
```

## Neden Refactor

Nesne yönelimli programlamanın temel direklerinden biri, nesne verilerini gizleme yeteneği olan Kapsülleme'dir . Aksi takdirde, tüm nesneler genel olacaktır ve diğer nesneler, herhangi bir kontrol ve denge olmaksızın nesnenizin verilerini alıp değiştirebilir! Veriler, bu verilerle ilişkili davranışlardan ayrılır, program bölümlerinin modülerliği tehlikeye girer ve bakım karmaşık hale gelir.

## Faydalar

- Bir bileşenin verileri ve davranışı birbiriyle yakından ilişkiliyse ve kodda aynı yerdeyse, bu bileşeni korumanız ve geliştirmeniz çok daha kolaydır.

- Nesne alanlarına erişimle ilgili karmaşık işlemleri de gerçekleştirebilirsiniz.

## Ne Zaman Kullanılmamalı

Bazı durumlarda, performans değerlendirmeleri nedeniyle kapsülleme tavsiye edilmez. Bu durumlar nadirdir ancak gerçekleştiğinde bu durum çok önemlidir.

Diyelim ki x ve y koordinatlarına sahip nesneleri içeren bir grafik düzenleyiciniz var. Bu alanların gelecekte değişmesi olası değildir. Ayrıca program, bu alanların bulunduğu birçok farklı nesneyi içerir. Bu nedenle, koordinat alanlarına doğrudan erişim, erişim yöntemlerinin çağrılmasıyla ele alınacak olan önemli CPU döngülerini kaydeder.

Bu olağandışı duruma bir örnek olarak, Java'da Point sınıfı vardır. Bu sınıfın tüm alanları geneldir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Alan için bir alıcı ve ayarlayıcı oluşturun.

2. Alanın tüm çağrılarını bulun. Alan değerinin alınmasını alıcı ile değiştirin ve yeni alan değerlerinin ayarını ayarlayıcı ile değiştirin.

3. Tüm alan çağrıları değiştirildikten sonra alanı özel yapın.

## Sonraki adımlar

Encapsulate Field , verileri ve bu verileri içeren davranışları birbirine yaklaştırmanın yalnızca ilk adımıdır. Erişim alanları için basit yöntemler oluşturduktan sonra bu yöntemlerin çağrıldığı yerleri tekrar kontrol etmelisiniz. Bu alanlardaki kodun erişim yöntemlerinde daha uygun görünmesi oldukça olasıdır.
