# Introduce Parameter Object

- **Sorun:** Yöntemleriniz yinelenen bir parametre grubu içerir.

- **Çözüm:** Bu parametreleri bir nesne ile değiştirin.

## Neden Refactor

Özdeş parametre gruplarıyla genellikle birden çok yöntemde karşılaşılır. Bu, hem parametrelerin kendilerinin hem de ilgili işlemlerin kod tekrarına neden olur. Parametreleri tek bir sınıfta birleştirerek, diğer yöntemleri bu koddan kurtararak, bu verileri işleme yöntemlerini oraya da taşıyabilirsiniz.

## Faydalar

- Daha okunabilir kod. Karmaşık parametreler yerine, anlaşılır bir ada sahip tek bir nesne görürsünüz.

- Şuraya buraya dağılmış özdeş parametre grupları, kendi tür kod yinelemelerini yaratırlar: özdeş kod çağrılmazken, sürekli olarak özdeş parametre ve argüman gruplarıyla karşılaşılır.

## Dezavantajları

Yalnızca verileri yeni bir sınıfa taşırsanız ve orada herhangi bir davranış veya ilgili işlemi taşımayı planlamıyorsanız, bu bir Data Class kokusu almaya başlar .

## Yeniden Düzenleme Nasıl Yapılır?

1. Parametre grubunuzu temsil edecek yeni bir sınıf oluşturun. Sınıfı değişmez yapın.

2. Yeniden düzenleme yapmak istediğiniz yöntemde, parametre nesnenizin iletileceği Add Parameter öğesini kullanın. Tüm metot çağrılarında, eski metot parametrelerinden oluşturulan nesneyi bu parametreye iletin.

3. Şimdi eski parametreleri yöntemden tek tek silmeye başlayın, bunları kodda parametre nesnesinin alanlarıyla değiştirin. Her parametre değişiminden sonra programı test edin.

4. Bittiğinde, yöntemin bir bölümünü (hatta bazen tüm yöntemi) bir parametre nesne sınıfına taşımanın bir anlamı olup olmadığına bakın. Öyleyse, Move Method veya Extract Method kullanın .
