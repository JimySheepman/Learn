# Pull Up Constructor Body

- **Sorun:** Alt sınıflarınız, çoğunlukla aynı koda sahip yapıcılara sahiptir.

```Java
class Manager extends Employee {
  public Manager(String name, String id, int grade) {
    this.name = name;
    this.id = id;
    this.grade = grade;
  }
  // ...
}
```

- **Çözüm:** Bir üst sınıf oluşturucu oluşturun ve alt sınıflarda aynı olan kodu ona taşıyın. Alt sınıf yapıcılarında üst sınıf yapıcısını çağırın.

```Java
class Manager extends Employee {
  public Manager(String name, String id, int grade) {
    super(name, id);
    this.grade = grade;
  }
  // ...
}
```

## Neden Refactor

Bu yeniden düzenleme tekniğinin Pull Up Yönteminden farkı nedir?

1. Java'da, alt sınıflar bir yapıcıyı devralamaz, bu nedenle alt sınıf yapıcısına Pull Up Yöntemini uygulayamaz ve tüm yapıcı kodunu üst sınıfa çıkardıktan sonra silemezsiniz. Üst sınıfta bir kurucu yaratmaya ek olarak, alt sınıflarda üst sınıf kurucusuna basit yetkilendirme ile kurucuların olması gerekir.

2. C++ ve Java'da (üst sınıf kurucusunu açıkça çağırmadıysanız) üst sınıf kurucusu otomatik olarak alt sınıf kurucusundan önce çağrılır, bu da ortak kodu yalnızca alt sınıf kurucularının başlangıcından itibaren taşımayı gerekli kılar (kazandığınız için) t bir alt sınıf oluşturucusunda rastgele bir yerden üst sınıf yapıcısını arayamazsınız).

3. Çoğu programlama dilinde, bir alt sınıf oluşturucu, üst sınıfın parametrelerinden farklı kendi parametre listesine sahip olabilir. Bu nedenle, yalnızca gerçekten ihtiyaç duyduğu parametrelerle bir üst sınıf oluşturucu oluşturmalısınız.

## Yeniden Düzenleme Nasıl Yapılır?

1. Bir üst sınıfta bir kurucu oluşturun.

2. Ortak kodu, her bir alt sınıfın kurucusunun başlangıcından üst sınıf kurucusuna çıkarın. Bunu yapmadan önce, mümkün olduğu kadar çok ortak kodu yapıcının başına taşımaya çalışın.

3. Üst sınıf kurucu çağrısını, alt sınıf kurucularının ilk satırına yerleştirin.
