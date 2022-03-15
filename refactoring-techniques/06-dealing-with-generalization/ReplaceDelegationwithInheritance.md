# Replace Delegation with Inheritance

- **Sorun:** Bir sınıf, başka bir sınıfın tüm yöntemlerine yetki veren birçok basit yöntem içerir.

- **Çözüm:** Sınıfı, delegasyon yöntemlerini gereksiz kılan bir temsilci mirasçı yapın.

## Neden Refactor

Delegasyon, delegasyonun nasıl uygulanacağını değiştirmeye ve diğer sınıfları oraya yerleştirmeye izin verdiği için mirastan daha esnek bir yaklaşımdır. Bununla birlikte, eylemleri yalnızca bir sınıfa ve onun tüm genel yöntemlerine devrederseniz, delegasyon yararlı olmaktan çıkar.

Böyle bir durumda, delegasyonu kalıtımla değiştirirseniz, sınıfı çok sayıda delegasyon yönteminden arındırır ve her yeni temsilci sınıfı yöntemi için bunları oluşturma ihtiyacından kendinizi kurtarırsınız.

## Faydalar

- Kod uzunluğunu azaltır. Tüm bu yetkilendirme yöntemleri artık gerekli değildir.

Ne Zaman Kullanılmamalı

- Sınıf, temsilci sınıfının genel yöntemlerinin yalnızca bir kısmına yetkilendirme içeriyorsa bu tekniği kullanmayın. Bunu yaparak, Liskov ikame ilkesini ihlal etmiş olursunuz .

- Bu teknik, yalnızca sınıfın hala ebeveynleri yoksa kullanılabilir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Sınıfı, temsilci sınıfının bir alt sınıfı yapın.

2. Geçerli nesneyi, temsilci nesnesine başvuru içeren bir alana yerleştirin.

3. Basit yetkilendirme ile yöntemleri tek tek silin. Adları farklıysa, tüm yöntemlere tek bir ad vermek için Rename Method kullanın.

4. Temsilci alanına yapılan tüm başvuruları geçerli nesneye yapılan başvurularla değiştirin.

5. Temsilci alanını kaldırın.
