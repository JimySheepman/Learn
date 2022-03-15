# Collapse Hierarchy

- **Sorun:** Bir alt sınıfın pratik olarak üst sınıfıyla aynı olduğu bir sınıf hiyerarşiniz var.

- **Çözüm:** Alt sınıfı ve üst sınıfı birleştirin.

## Neden Refactor

Programınız zamanla büyüdü ve bir alt sınıf ile üst sınıf neredeyse aynı hale geldi. Bir alt sınıftan bir özellik kaldırıldı, bir yöntem üst sınıfa taşındı... ve şimdi iki benzer sınıfınız var.

## Faydalar

- Program karmaşıklığı azalır. Daha az sınıf, kafanızda düz tutacak daha az şey ve gelecekteki kod değişiklikleri sırasında endişelenecek daha az kırılabilir hareketli parça anlamına gelir.

- Yöntemler bir sınıfta erken tanımlandığında, kodunuzda gezinmek daha kolaydır. Belirli bir yöntemi bulmak için tüm hiyerarşiyi taramanız gerekmez.

## Ne Zaman Kullanılmamalı

- Yeniden düzenlediğiniz sınıf hiyerarşisinin birden fazla alt sınıfı var mı? Eğer öyleyse, yeniden düzenleme tamamlandıktan sonra kalan alt sınıflar, hiyerarşinin çöktüğü sınıfın mirasçıları olmalıdır.

- Ancak bunun Liskov ikame ilkesinin ihlaline yol açabileceğini unutmayın . Örneğin, programınız şehir ulaşım ağlarını öykünüyorsa ve yanlışlıkla Transportüst sınıfı Caralt sınıfa daraltırsanız, Planesınıf Car. Hata!

## Yeniden Düzenleme Nasıl Yapılır?

1. Hangi sınıfın kaldırılmasının daha kolay olduğunu seçin: üst sınıf veya alt sınıfı.

2. Alt sınıftan kurtulmaya karar verirseniz, Pull Up Field ve Pull Up Yöntemini kullanın . Üst sınıfı ortadan kaldırmayı seçerseniz, Push Down Field ve Push Down Method'a gidin .

3. Silmekte olduğunuz sınıfın tüm kullanımlarını, alanların ve yöntemlerin taşınacağı sınıfla değiştirin. Genellikle bu, sınıf oluşturma kodu, değişken ve parametre yazma ve kod yorumlarında belgeler olacaktır.

4. Boş sınıfı silin.
