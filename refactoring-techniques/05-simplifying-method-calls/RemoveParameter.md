# Remove Parameter

- **Sorun:** Bir yöntemin gövdesinde bir parametre kullanılmaz.

- **Çözüm:** Kullanılmayan parametreyi kaldırın.

## Neden Refactor

Bir yöntem çağrısındaki her parametre, programcıyı bu parametrede hangi bilgilerin bulunduğunu bulmaya zorlar. Ve eğer bir parametre metot gövdesinde tamamen kullanılmıyorsa, bu "kazıma" boşunadır.

Ve her durumda, ek parametreler çalıştırılması gereken ekstra kodlardır.

Bazen geleceğe yönelik bir gözle parametreler ekleriz, parametrenin gerekli olabileceği yöntemde değişiklikler öngörürüz. Yine de deneyimler, yalnızca gerçekten ihtiyaç duyulduğunda bir parametre eklemenin daha iyi olduğunu gösteriyor. Ne de olsa, beklenen değişiklikler genellikle aynen kalır - tahmin edilir.

## Faydalar

Bir yöntem, yalnızca gerçekten ihtiyaç duyduğu parametreleri içerir.

## Ne Zaman Kullanılmamalı

Yöntem, alt sınıflarda veya bir üst sınıfta farklı şekillerde uygulanıyorsa ve parametreniz bu uygulamalarda kullanılıyorsa, parametreyi olduğu gibi bırakın.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yöntemin bir üst sınıfta mı yoksa alt sınıfta mı tanımlandığını görün. Eğer öyleyse, parametre orada kullanılıyor mu? Parametre bu uygulamalardan birinde kullanılıyorsa, bu yeniden düzenleme tekniğine devam edin.

2. Bir sonraki adım, yeniden düzenleme işlemi sırasında programı çalışır durumda tutmak için önemlidir. Eski yöntemi kopyalayarak yeni bir yöntem oluşturun ve ilgili parametreyi buradan silin. Eski yöntemin kodunu yenisine yapılan çağrıyla değiştirin.

3. Eski yönteme yapılan tüm referansları bulun ve bunları yeni yönteme yapılan referanslarla değiştirin.

4. Eski yöntemi silin. Eski yöntem bir genel arabirimin parçasıysa bu adımı gerçekleştirmeyin. Bu durumda, eski yöntemi kullanım dışı olarak işaretleyin.
