# Parameterize Method

- **Sorun:** Birden çok yöntem, yalnızca iç değerleri, sayıları veya işlemlerinde farklı olan benzer eylemleri gerçekleştirir.

- **Çözüm:** Gerekli özel değeri iletecek bir parametre kullanarak bu yöntemleri birleştirin.

## Neden Refactor

Benzer yöntemleriniz varsa, muhtemelen bunun gerektirdiği tüm sonuçlarla birlikte yinelenen kodunuz vardır.

Ayrıca, bu işlevin başka bir sürümünü eklemeniz gerekiyorsa, başka bir yöntem oluşturmanız gerekecektir. Bunun yerine, mevcut yöntemi farklı bir parametreyle çalıştırabilirsiniz.

## Dezavantajları

- Bazen bu yeniden düzenleme tekniği çok ileri götürülebilir, bu da birden çok basit yöntem yerine uzun ve karmaşık bir ortak yöntemle sonuçlanır.

- İşlevselliğin etkinleştirilmesini/devre dışı bırakılmasını bir parametreye taşırken de dikkatli olun. Bu, sonunda  Replace Parameter with Explicit Methods aracılığıyla ele alınması gereken büyük bir koşullu operatörün oluşturulmasına yol açabilir .

## Yeniden Düzenleme Nasıl Yapılır?

1. Bir parametre ile yeni bir yöntem oluşturun ve Extract Method uygulayarak bunu tüm sınıflar için aynı olan koda taşıyın . Bazen yöntemlerin yalnızca belirli bir bölümünün aslında aynı olduğunu unutmayın. Bu durumda, yeniden düzenleme, yalnızca aynı parçanın yeni bir yönteme çıkarılmasından oluşur.

2. Yeni yöntemin kodunda, özel/farklı değeri bir parametre ile değiştirin.

3. Her eski yöntem için, çağrıldığı yerleri bulun ve bu çağrıları bir parametre içeren yeni yönteme yapılan çağrılarla değiştirin. Ardından eski yöntemi silin.
