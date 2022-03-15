# Separate Query from Modifier

- **Sorun:** Bir değer döndüren ama aynı zamanda bir nesnenin içindeki bir şeyi değiştiren bir yönteminiz var mı?

- **Çözüm:** Yöntemi iki ayrı yönteme ayırın. Beklediğiniz gibi, bunlardan biri değeri döndürmeli, diğeri ise nesneyi değiştirmelidir.

## Neden Refactor

Bu faktoring tekniği, Komut ve Sorgu Sorumluluğu Ayrımını uygular . Bu ilke, bir nesnenin içindeki bir şeyi değiştiren koddan veri almaktan sorumlu kodu ayırmamızı söyler.

Veri alma koduna sorgu adı verilir . Bir nesnenin görünür durumundaki şeyleri değiştirme koduna değiştirici adı verilir . Bir sorgu ve değiştirici birleştirildiğinde, koşulunda değişiklik yapmadan veri almanın bir yolu yoktur. Başka bir deyişle, bir soru soruyorsunuz ve yanıt alınırken bile değiştirebiliyorsunuz. Bu sorun, sorguyu çağıran kişi, genellikle çalışma zamanı hatalarına yol açan yöntemin “yan etkileri” hakkında bilgi sahibi olmadığında daha da şiddetli hale gelir.

Ancak yan etkilerin yalnızca bir nesnenin görünür durumunu değiştiren değiştiriciler söz konusu olduğunda tehlikeli olduğunu unutmayın. Bunlar, örneğin, bir nesnenin genel arayüzünden erişilebilen alanlar, bir veritabanındaki giriş, dosyalar vb. olabilir. Bir değiştirici yalnızca karmaşık bir işlemi önbelleğe alır ve bir sınıfın özel alanı içinde kaydederse, herhangi bir yan etkiye neden olamaz. Etkileri.

## Faydalar

Programınızın durumunu değiştirmeyen bir sorgunuz varsa, yalnızca yöntemi çağırmanız nedeniyle sonuçta meydana gelen istenmeyen değişiklikler konusunda endişelenmenize gerek kalmadan istediğiniz kadar çağırabilirsiniz.

## Dezavantajları

Bazı durumlarda, bir komut gerçekleştirdikten sonra veri almak uygundur. Örneğin, bir veritabanından bir şey silerken, kaç satırın silindiğini bilmek istersiniz.

## Yeniden Düzenleme Nasıl Yapılır?

1. Orijinal yöntemin yaptığını döndürmek için yeni bir sorgu yöntemi oluşturun.

2. Orijinal yöntemi, yalnızca yeni sorgu yönteminin çağrılmasının sonucunu döndürecek şekilde değiştirin .

3. Orijinal yönteme yapılan tüm başvuruları, sorgu yöntemine yapılan bir çağrıyla değiştirin . Bu satırdan hemen önce değiştirici yöntemine bir çağrı yapın . Bu, orijinal yöntemin bir koşullu operatör veya döngü durumunda kullanılması durumunda sizi yan etkilerden kurtaracaktır.

4. Artık uygun bir değiştirici yöntem haline gelen orijinal yöntemdeki değer döndüren koddan kurtulun .
