# Extract Class

- **Sorun:** Bir sınıf iki sınıfın işini yaptığında, gariplik ortaya çıkar.

- **Çözüm:** Bunun yerine, yeni bir sınıf oluşturun ve ilgili işlevsellikten sorumlu alanları ve yöntemleri onun içine yerleştirin.

## Neden Refactor

Sınıflar her zaman açık ve anlaşılır şekilde başlar. Diğer sınıfların işine karışmadan işlerini yaparlar ve sanki kendi işlerine bakarlar. Ancak program genişledikçe, bir yöntem ve ardından bir alan eklenir... ve sonunda, bazı sınıflar hiç düşünüldüğünden daha fazla sorumluluk yerine getirir.

## Faydalar

- Bu yeniden düzenleme yöntemi, Tek Sorumluluk İlkesine bağlılığın korunmasına yardımcı olacaktır . Derslerinizin kodu daha açık ve anlaşılır olacaktır.

- Tek sorumluluk sınıfları daha güvenilirdir ve değişikliklere karşı toleranslıdır. Örneğin, on farklı şeyden sorumlu bir sınıfınız olduğunu söyleyin. Bu sınıfı bir şey için daha iyi hale getirmek için değiştirdiğinizde, diğer dokuz şey için onu bozma riskini alırsınız.

## Dezavantajları

- Bu yeniden düzenleme tekniğiyle "aşırıya kaçarsanız", Inline Class'a başvurmanız gerekir .

## Yeniden Düzenleme Nasıl Yapılır?

Başlamadan önce, sınıfın sorumluluklarını tam olarak nasıl bölmek istediğinize karar verin.

1. İlgili işlevselliği içerecek yeni bir sınıf oluşturun.

2. Eski sınıf ile yeni sınıf arasında bir ilişki oluşturun. Optimal olarak, bu ilişki tek yönlüdür; bu, ikinci sınıfın herhangi bir sorun olmadan yeniden kullanılmasına izin verir. Bununla birlikte, iki yönlü bir ilişkinin gerekli olduğunu düşünüyorsanız, bu her zaman kurulabilir.

3. Yeni sınıfa taşımaya karar verdiğiniz her alan ve yöntem için  Move Field ve Move Method kullanın . Çok sayıda hata yapma riskini azaltmak için yöntemler için özel olanlarla başlayın. Her seferinde biraz yer değiştirmeye çalışın ve en sonunda bir hata düzeltme yığınını önlemek için her hareketten sonra sonuçları test edin.

4. Taşınmayı bitirdikten sonra, ortaya çıkan sınıflara bir kez daha bakın. Değişen sorumluluklara sahip eski bir sınıf, daha fazla netlik için yeniden adlandırılabilir. Varsa, iki yönlü sınıf ilişkilerinden kurtulup kurtulamayacağınızı görmek için tekrar kontrol edin.

5. Ayrıca yeni sınıfa dışarıdan erişilebilirliği de düşünün. Özel yaparak, eski sınıftan alanlar aracılığıyla yöneterek sınıfı tamamen istemciden gizleyebilirsiniz. Alternatif olarak, istemcinin değerleri doğrudan değiştirmesine izin vererek bunu herkese açık hale getirebilirsiniz. Buradaki kararınız, yeni sınıftaki değerlerde beklenmedik doğrudan değişiklikler yapıldığında eski sınıfın davranışı için ne kadar güvenli olduğuna bağlıdır
