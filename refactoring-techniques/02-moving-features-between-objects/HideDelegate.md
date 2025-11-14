# Hide Delegate

- **Sorun:** İstemci, B nesnesini А nesnesinin bir alanından veya yönteminden alır. Ardından istemci, B nesnesinin bir yöntemini çağırır.

- **Çözüm:** A sınıfında, çağrıyı B nesnesine devreden yeni bir yöntem oluşturun. Artık istemci, B sınıfını bilmiyor veya buna bağlı değil.

## Neden Refactor

Başlangıç ​​olarak, terminolojiye bakalım:

- Sunucu , istemcinin doğrudan erişime sahip olduğu nesnedir.

- Temsilci , istemcinin ihtiyaç duyduğu işlevselliği içeren son nesnedir.

Bir istemci başka bir nesneden bir nesne istediğinde, ardından ikinci nesne başka bir nesne istediğinde bir çağrı zinciri görünür ve bu böyle devam eder. Bu çağrı dizileri, istemciyi sınıf yapısı boyunca gezinmeye dahil eder. Bu ilişkilerdeki herhangi bir değişiklik, müşteri tarafında değişiklikler gerektirecektir.

## Faydalar

Yetkilendirmeyi istemciden gizler. İstemci kodunun nesneler arasındaki ilişkilerin ayrıntıları hakkında ne kadar az bilmesi gerekiyorsa, programınızda değişiklik yapmak o kadar kolay olur.

## Dezavantajları

Aşırı sayıda delegasyon yöntemi oluşturmanız gerekiyorsa, sunucu sınıfı gereksiz bir arabulucu olma riskiyle karşı karşıya kalır ve bu da fazladan Middle Man'a yol açar .

Yeniden Düzenleme Nasıl Yapılır?

1. İstemci tarafından çağrılan temsilci sınıfının her yöntemi için , sunucu sınıfında çağrıyı temsilci sınıfına devreden bir yöntem oluşturun .

2. İstemci kodunu, sunucu sınıfının yöntemlerini çağıracak şekilde değiştirin .

3. Değişiklikleriniz istemciyi temsilci sınıfına ihtiyaç duymaktan kurtarırsa, temsilci sınıfına erişim yöntemini sunucu sınıfından kaldırabilirsiniz (başlangıçta temsilci sınıfını almak için kullanılan yöntem ).
