# Remove Middle Man

- **Sorun:** Bir sınıfın, yalnızca diğer nesnelere yetki veren çok fazla yöntemi vardır.

- **Çözüm:** Bu yöntemleri silin ve istemciyi doğrudan bitiş yöntemlerini çağırmaya zorlayın.

## Neden Refactor

Bu tekniği açıklamak için,  Hide Delegate'den gelen şu terimleri kullanacağız :

- Sunucu , istemcinin doğrudan erişime sahip olduğu nesnedir.

- Temsilci , istemcinin ihtiyaç duyduğu işlevselliği içeren son nesnedir.

İki tür sorun vardır

1. Sunucu sınıfı kendi başına hiçbir şey yapmaz ve basitçe gereksiz karmaşıklık yaratır. Bu durumda, bu sınıfa hiç ihtiyaç olup olmadığını düşünün.

2. Temsilciye her yeni özellik eklendiğinde , sunucu sınıfında bunun için bir temsilci yöntemi oluşturmanız gerekir . Çok fazla değişiklik yapılırsa, bu oldukça yorucu olacaktır.

## Yeniden Düzenleme Nasıl Yapılır?

1. Sunucu sınıfı nesnesinden temsilci sınıfı nesnesine erişmek için bir alıcı oluşturun .

2. Sunucu sınıfındaki delegasyon yöntemlerine yapılan çağrıları, delege sınıfındaki yöntemler için doğrudan çağrılarla değiştirin .
