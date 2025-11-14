# Form Template Method

- **Sorun:** Alt sınıflarınız, aynı sırada benzer adımları içeren algoritmaları uygular.

- **Çözüm:** Algoritma yapısını ve aynı adımları bir üst sınıfa taşıyın ve farklı adımların alt sınıflarda uygulanmasını bırakın.

## Neden Refactor

Alt sınıflar paralel olarak, bazen farklı kişiler tarafından geliştirilir, bu da her değişikliğin tüm alt sınıflarda yapılması gerektiğinden, kod tekrarına, hatalara ve kod bakımında zorluklara yol açar.

## Faydalar

- Kod çoğaltma, her zaman basit kopyala/yapıştır durumlarına atıfta bulunmaz. Sayıları sıralamak için bir yönteminiz ve yalnızca öğelerin karşılaştırılmasıyla farklılaştırılan nesne koleksiyonlarını sıralamak için bir yönteminiz olduğunda olduğu gibi, çoğaltma genellikle daha yüksek bir düzeyde gerçekleşir. Bir şablon yöntemi oluşturmak, paylaşılan algoritma adımlarını bir üst sınıfta birleştirerek ve yalnızca alt sınıflardaki farklılıkları bırakarak bu çoğaltmayı ortadan kaldırır.

- Bir şablon yöntemi oluşturmak , eylem halindeki Açık/Kapalı İlkenin bir örneğidir . Yeni bir algoritma sürümü göründüğünde, yalnızca yeni bir alt sınıf oluşturmanız gerekir; mevcut kodda herhangi bir değişiklik gerekli değildir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Alt sınıflardaki algoritmaları, ayrı yöntemlerde açıklanan bileşenlerine ayırın. Extract Method bu konuda yardımcı olabilir.

2. Tüm alt sınıflar için aynı olan elde edilen yöntemler, Pull Up Method aracılığıyla bir üst sınıfa taşınabilir .

3. Benzer olmayan yöntemlere, Rename Method aracılığıyla tutarlı adlar verilebilir .

4. Pull Up Method kullanarak benzer olmayan yöntemlerin imzalarını soyut olanlar olarak bir üst sınıfa taşıyın . Uygulamalarını alt sınıflarda bırakın.

5. Ve son olarak, algoritmanın ana yöntemini üst sınıfa çıkarın. Şimdi hem gerçek hem de soyut olarak üst sınıfta açıklanan yöntem adımlarıyla çalışmalıdır.
