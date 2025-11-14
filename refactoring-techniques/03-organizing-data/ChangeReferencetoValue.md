# Change Reference to Value

- **Sorun:** Yaşam döngüsünü yönetmek için çok küçük ve nadiren değiştirilen bir referans nesneniz var.

- **Çözüm:**  Bir değer nesnesine dönüştürün.

## Neden Refactor

Bir referanstan bir değere geçmek için ilham, referansla çalışmanın zorluğundan gelebilir. Referanslar sizin tarafınızdan yönetim gerektirir:

- Her zaman gerekli nesneyi depolamadan talep etmeyi gerektirirler.

- Bellekteki referanslarla çalışmak sakıncalı olabilir.

- Referanslarla çalışmak, dağıtılmış ve paralel sistemlerde değerlere kıyasla özellikle zordur.

Değerler, yaşamları boyunca durumları değişebilecek nesneler yerine değiştirilemez nesnelere sahip olmayı tercih ediyorsanız özellikle yararlıdır.

## Faydalar

- Nesnelerin önemli bir özelliği de değiştirilemez olmalarıdır. Bir nesne değeri döndüren her sorgu için aynı sonuç alınmalıdır. Bu doğruysa, aynı şeyi temsil eden çok sayıda nesne varsa hiçbir sorun çıkmaz.

- Değerlerin uygulanması çok daha kolaydır.

## Dezavantajları

Bir değer değişebilirse, herhangi bir nesnenin değişip değişmediğini, aynı varlığı temsil eden diğer tüm nesnelerdeki değerlerin güncellendiğinden emin olun. Bu o kadar külfetli ki, bu amaç için bir referans oluşturmak daha kolay.

## Yeniden Düzenleme Nasıl Yapılır?

1. Nesneyi değiştirilemez hale getirin. Nesne, durumunu ve verilerini değiştiren herhangi bir ayarlayıcıya veya başka yönteme sahip olmamalıdır ( Remove Setting Method burada yardımcı olabilir). Verilerin bir değer nesnesinin alanlarına atanması gereken tek yer bir yapıcıdır.

2. İki değeri karşılaştırabilmek için bir karşılaştırma yöntemi oluşturun.

3. Fabrika yöntemini silip silemeyeceğinizi ve nesne oluşturucuyu herkese açık yapıp yapamayacağınızı kontrol edin.
