# Move Method

- **Sorun:** Bir metot, kendi sınıfından daha çok başka bir sınıfta kullanılır.

- **Çözüm:** Yöntemi en çok kullanan sınıfta yeni bir yöntem oluşturun, ardından kodu eski yöntemden oraya taşıyın. Orijinal yöntemin kodunu, diğer sınıftaki yeni yönteme referans haline getirin veya tamamen kaldırın.

## Neden Refactor

1. Bir yöntemi, yöntem tarafından kullanılan verilerin çoğunu içeren bir sınıfa taşımak istiyorsunuz. Bu, sınıfları dahili olarak daha tutarlı hale getirir .

2. Yöntemi çağıran sınıfın bulunduğu sınıfa olan bağımlılığını azaltmak veya ortadan kaldırmak için bir yöntemi taşımak istiyorsunuz. Çağıran sınıf, yöntemi taşımayı planladığınız sınıfa zaten bağımlıysa bu yararlı olabilir. Bu , sınıflar arasındaki bağımlılığı azaltır .

## Yeniden Düzenleme Nasıl Yapılır?

1. Eski yöntem tarafından kullanılan tüm özellikleri kendi sınıfında doğrulayın. Onları taşımak da iyi bir fikir olabilir. Kural olarak, bir özellik sadece söz konusu yöntemle kullanılıyorsa, özelliği kesinlikle ona taşımalısınız. Özellik başka yöntemlerle de kullanılıyorsa, bu yöntemleri de taşımanız gerekir. Bazen çok sayıda yöntemi taşımak, farklı sınıflarda aralarında ilişkiler kurmaktan çok daha kolaydır.

2. Yöntemin üst sınıflarda ve alt sınıflarda bildirilmediğinden emin olun. Bu durumda, verici sınıflar arasında bölünen bir yöntemin farklı işlevselliklerini sağlamak için alıcı sınıfta ya hareket etmekten kaçınmanız ya da bir tür polimorfizm uygulamanız gerekecektir.

3. Alıcı sınıfında yeni yöntemi bildirin. Yeni sınıfta metot için daha uygun olan yeni bir isim vermek isteyebilirsiniz.

4. Alıcı sınıfa nasıl başvuracağınıza karar verin. Halihazırda uygun bir nesne döndüren bir alanınız veya yönteminiz olabilir, ancak değilse, alıcı sınıfın nesnesini depolamak için yeni bir yöntem veya alan yazmanız gerekir.

5. Artık alıcı nesneye başvurmak için bir yolunuz ve sınıfında yeni bir yönteminiz var. Tüm bunlarla birlikte eski yöntemi yeni yönteme referans haline getirebilirsiniz.

6. Bir göz atın: eski yöntemi tamamen silebilir misiniz? Eğer öyleyse, eski yöntemi kullanan tüm yerlere yeni yönteme bir referans koyun.
