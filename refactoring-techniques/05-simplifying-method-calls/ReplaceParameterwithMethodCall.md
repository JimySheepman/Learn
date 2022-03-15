# Replace Parameter with Method Call

- **Sorun:** Bir sorgu yöntemini çağırmak ve sonuçlarını başka bir yöntemin parametreleri olarak iletmek, bu yöntem sorguyu doğrudan çağırabilir.

```Java
int basePrice = quantity * itemPrice;
double seasonDiscount = this.getSeasonalDiscount();
double fees = this.getFees();
double finalPrice = discountedPrice(basePrice, seasonDiscount, fees);
```

- **Çözüm:** Değeri bir parametreden geçirmek yerine, yöntem gövdesinin içine bir sorgu çağrısı yerleştirmeyi deneyin.

```Java
int basePrice = quantity * itemPrice;
double finalPrice = discountedPrice(basePrice);
```

## Neden Refactor

Uzun bir parametre listesini anlamak zor. Ek olarak, bu tür yöntemlere yapılan çağrılar, gezinmesi zor olan ancak yönteme iletilmesi gereken sarma ve canlandırıcı değer hesaplamaları ile genellikle bir dizi basamaklamayı andırır. Yani bir metot yardımıyla bir parametre değeri hesaplanabiliyorsa, bunu metodun kendi içinde yapın ve parametreden kurtulun.

## Faydalar

Gereksiz parametrelerden kurtulup metot çağrılarını basitleştiriyoruz. Bu tür parametreler genellikle şimdi olduğu gibi proje için değil, asla gelmeyecek olan gelecekteki ihtiyaçlar için yaratılır.

## Dezavantajları

Diğer ihtiyaçlar için yarın parametreye ihtiyacınız olabilir... yöntemi yeniden yazmanızı sağlar.

## Yeniden Düzenleme Nasıl Yapılır?

1. Değer alma kodunun geçerli yöntemdeki parametreleri kullanmadığından emin olun, çünkü bunlar başka bir yöntemin içinden kullanılamaz. Eğer öyleyse, kodu taşımak mümkün değildir.

2. İlgili kod, tek bir yöntem veya işlev çağrısından daha karmaşıksa, bu kodu yeni bir yöntemde izole etmek ve çağrıyı basitleştirmek için Extract Method kullanın.

3. Ana yöntemin kodunda, değiştirilen parametreye yapılan tüm başvuruları, değeri alan yönteme yapılan çağrılarla değiştirin.

4. Artık kullanılmayan parametreyi ortadan kaldırmak için  Remove Parameter'ı kullanın .
