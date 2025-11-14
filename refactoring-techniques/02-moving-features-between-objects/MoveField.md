# Move Field

- **Sorun:** Bir alan, kendi sınıfından çok başka bir sınıfta kullanılır.

- **Çözüm:** Yeni bir sınıfta bir alan oluşturun ve eski alanın tüm kullanıcılarını ona yönlendirin.

## Neden Refactor

Genellikle alanlar,  Extract Class tekniğinin bir parçası olarak taşınır . Alanı hangi sınıfta bırakacağınıza karar vermek zor olabilir. Temel kuralımız şudur: Bir alanı, onu kullanan yöntemlerle aynı yere (veya bu yöntemlerin çoğunun olduğu yere) koyun.

Bu kural, bir alanın basitçe yanlış yere yerleştirildiği diğer durumlarda yardımcı olacaktır.

## Yeniden Düzenleme Nasıl Yapılır?

1. Alan public ise, alanı private yapıp genel erişim yöntemleri sağlarsanız yeniden düzenleme çok daha kolay olacaktır (bunun için Encapsulate Field kullanabilirsiniz ).

2. Alıcı sınıfında erişim yöntemleriyle aynı alanı oluşturun.

3. Alıcı sınıfa nasıl başvuracağınıza karar verin. Uygun nesneyi döndüren bir alanınız veya yönteminiz zaten olabilir; değilse, alıcı sınıfın nesnesini depolamak için yeni bir yöntem veya alan yazmanız gerekecektir.

4. Eski alana yapılan tüm başvuruları, alıcı sınıfındaki yöntemlere uygun çağrılarla değiştirin. Alan özel değilse, üst sınıf ve alt sınıflarda bununla ilgilenin.

5. Orijinal sınıftaki alanı silin.
