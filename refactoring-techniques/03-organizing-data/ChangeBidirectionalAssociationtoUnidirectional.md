# Change Bidirectional Association to Unidirectional

- **Sorun:** Sınıflar arasında çift yönlü bir ilişkiniz var, ancak sınıflardan biri diğerinin özelliklerini kullanmıyor.

- **Çözüm:** Kullanılmayan ilişkilendirmeyi kaldırın.

## Neden Refactor

Çift yönlü bir ilişkilendirmeyi sürdürmek genellikle tek yönlü bir ilişkiye göre daha zordur ve ilgili nesneleri düzgün bir şekilde oluşturmak ve silmek için ek kod gerektirir. Bu, programı daha karmaşık hale getirir.

Ek olarak, yanlış uygulanan çift yönlü bir ilişki, çöp toplama için sorunlara neden olabilir (sırasıyla kullanılmayan nesneler tarafından belleğin şişmesine neden olur).

Örnek: çöp toplayıcı, artık diğer nesneler tarafından başvurulmayan nesneleri bellekten kaldırır. Diyelim ki bir nesne çifti User- Orderoluşturuldu, kullanıldı ve sonra terk edildi. Ancak bu nesneler hala birbirlerine atıfta bulundukları için bellekten silinmeyecektir. Bununla birlikte, artık kullanılmayan nesne referanslarını otomatik olarak tanımlayan ve bunları bellekten kaldıran programlama dillerindeki gelişmeler sayesinde bu sorun daha az önemli hale geliyor.

Bir de sınıflar arası bağımlılık sorunu var. Çift yönlü bir ilişkide, iki sınıfın birbirinden haberdar olması gerekir, yani ayrı ayrı kullanılamazlar. Bu ilişkilerin birçoğu mevcutsa, programın farklı bölümleri birbirine çok bağımlı hale gelir ve bir bileşendeki herhangi bir değişiklik diğer bileşenleri etkileyebilir.

## Faydalar

- İlişkiye ihtiyaç duymayan sınıfı basitleştirir. Daha az kod, daha az kod bakımına eşittir.

- Sınıflar arasındaki bağımlılığı azaltır. Bir sınıfta yapılan herhangi bir değişiklik yalnızca o sınıfı etkilediğinden bağımsız sınıfların bakımı daha kolaydır.

## Yeniden Düzenleme Nasıl Yapılır?

1. Sınıflarınız için aşağıdakilerden birinin doğru olduğundan emin olun:

    - İlişkilendirme kullanılmaz.

    - İlişkili nesneyi almanın başka bir yolu var, örneğin bir veritabanı sorgusu aracılığıyla.

    - İlişkili nesne, onu kullanan yöntemlere bir argüman olarak iletilebilir.

2. Durumunuza bağlı olarak, başka bir nesneyle ilişki içeren bir alanın kullanımı, nesneyi farklı bir şekilde almak için bir parametre veya yöntem çağrısı ile değiştirilmelidir.

3. İlişkili nesneyi alana atayan kodu silin.

4.Artık kullanılmayan alanı silin.
