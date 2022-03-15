# Push Down Method

- **Sorun:** Davranış, yalnızca bir (veya birkaç) alt sınıf tarafından kullanılan bir üst sınıfta mı uygulanıyor?

- **Çözüm:** Bu davranışı alt sınıflara taşıyın.

## Neden Refactor

İlk başta belirli bir yöntemin tüm sınıflar için evrensel olması gerekiyordu, ancak gerçekte yalnızca bir alt sınıfta kullanılıyor. Bu durum, planlanan özelliklerin gerçekleşmemesi durumunda ortaya çıkabilir.

Bu tür durumlar, yalnızca bir alt sınıfta kullanılan bir yöntem bırakarak, bir sınıf hiyerarşisinden işlevselliğin kısmen çıkarılmasından (veya kaldırılmasından) sonra da ortaya çıkabilir.

Bir metoda birden fazla alt sınıf tarafından ihtiyaç duyulduğunu, ancak hepsinin gerekli olmadığını görüyorsanız, bir ara alt sınıf oluşturup metodu ona taşımak faydalı olabilir. Bu, bir yöntemin tüm alt sınıflara itilmesinden kaynaklanacak kod tekrarından kaçınmayı sağlar.

## Faydalar

Sınıf bütünlüğünü geliştirir. Bir yöntem, onu görmeyi beklediğiniz yerde bulunur.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yöntemi bir alt sınıfta bildirin ve kodunu üst sınıftan kopyalayın.

2. Yöntemi üst sınıftan kaldırın.

3. Yöntemin kullanıldığı tüm yerleri bulun ve gerekli alt sınıftan çağrıldığını doğrulayın.
