# Replace Inheritance with Delegation

- **Sorun:** Üst sınıfının yöntemlerinin yalnızca bir bölümünü kullanan bir alt sınıfınız var (veya üst sınıf verilerini devralmak mümkün değil).

- **Çözüm:** Bir alan oluşturun ve içine bir üst sınıf nesnesi koyun, yöntemleri üst sınıf nesnesine devredin ve mirastan kurtulun.

## Neden Refactor

Kalıtımın kompozisyonla değiştirilmesi, aşağıdaki durumlarda sınıf tasarımını önemli ölçüde iyileştirebilir:

- Alt sınıfınız Liskov ikame ilkesini ihlal ediyor , yani kalıtım yalnızca ortak kodu birleştirmek için uygulandıysa, ancak alt sınıf üst sınıfın bir uzantısı olduğu için değil.

- Alt sınıf, üst sınıfın yöntemlerinin yalnızca bir bölümünü kullanır. Bu durumda, birisinin çağırmaması gereken bir üst sınıf yöntemini çağırması yalnızca bir zaman meselesidir.

Özünde, bu yeniden düzenleme tekniği her iki sınıfı da böler ve üst sınıfı, alt sınıfın ebeveyni değil, yardımcısı yapar. Tüm üst sınıf yöntemlerini devralmak yerine, alt sınıf, üst sınıf nesnesinin yöntemlerine delegasyon yapmak için yalnızca gerekli yöntemlere sahip olacaktır.

## Faydalar

- Bir sınıf, üst sınıftan miras alınan gereksiz yöntemler içermez.

- Temsilci alanına çeşitli uygulamalara sahip çeşitli nesneler konabilir. Aslında Strateji tasarım modelini elde edersiniz.

## Dezavantajları

- Birçok basit yetkilendirme yöntemi yazmanız gerekir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Üst sınıfı tutmak için alt sınıfta bir alan oluşturun. İlk aşamada, mevcut nesneyi içine yerleştirin.

2. Alt sınıf yöntemlerini, yerine üst sınıf nesnesini kullanacak şekilde değiştirin this.

3. İstemci kodunda çağrılan üst sınıftan devralınan yöntemler için alt sınıfta basit temsilci yöntemleri oluşturun.

4. Alt sınıftan devralma bildirimini kaldırın.

5. Yeni bir nesne oluşturarak eski üst sınıfın depolandığı alanın başlatma kodunu değiştirin.
