# Change Unidirectional Association to Bidirectional

- **Sorun:** Her birinin diğerinin özelliklerini kullanması gereken iki sınıfınız var, ancak aralarındaki ilişki sadece tek yönlü.

- **Çözüm:** Eksik ilişkilendirmeyi, ihtiyacı olan sınıfa ekleyin.

## Neden Refactor

Başlangıçta sınıfların tek yönlü bir ilişkisi vardı. Ancak zamanla, müşteri kodunun ilişkilendirmenin her iki tarafına da erişmesi gerekiyordu.

## Faydalar

Bir sınıfın bir ters ilişkilendirmeye ihtiyacı varsa, bunu basitçe hesaplayabilirsiniz. Ancak bu hesaplamalar karmaşıksa, ters ilişkilendirmeyi sürdürmek daha iyidir.

## Dezavantajları

- Çift yönlü çağrışımları uygulamak ve sürdürmek, tek yönlü olanlardan çok daha zordur.

- Çift yönlü ilişkilendirmeler sınıfları birbirine bağımlı hale getirir. Tek yönlü bir ilişkilendirme ile biri diğerinden bağımsız olarak kullanılabilir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Ters ilişkilendirmeyi tutmak için bir alan ekleyin.

2. Hangi sınıfın “baskın” olacağına karar verin. Bu sınıf, öğeler eklendikçe veya değiştikçe ilişkilendirmeyi oluşturan veya güncelleyen, kendi sınıfında ilişkilendirmeyi kuran ve ilişkili nesnede ilişki kurmak için yardımcı yöntemleri çağıran yöntemleri içerecektir.

3. "Baskın olmayan" sınıfta ilişki kurmak için bir yardımcı program yöntemi oluşturun. Yöntem, alanı tamamlamak için parametrelerde verilenleri kullanmalıdır. Daha sonra başka amaçlarla kullanılmaması için yönteme açık bir ad verin.

4. Tek yönlü ilişkiyi denetlemeye yönelik eski yöntemler "baskın" sınıftaysa, bunları ilişkili nesneden yardımcı program yöntemlerine yapılan çağrılarla tamamlayın.

5. İlişkiyi denetlemeye yönelik eski yöntemler "baskın olmayan" sınıftaysa, yöntemleri "baskın" sınıfta oluşturun, çağırın ve yürütmeyi onlara devredin.
