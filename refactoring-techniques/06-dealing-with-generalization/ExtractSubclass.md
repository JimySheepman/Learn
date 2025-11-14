# Extract Subclass

- **Sorun:** Bir sınıf, yalnızca belirli durumlarda kullanılan özelliklere sahiptir.

- **Çözüm:** Bir alt sınıf oluşturun ve bu durumlarda kullanın.

## Neden Refactor

Ana sınıfınız, sınıf için belirli bir nadir kullanım durumunu uygulamaya yönelik yöntemler ve alanlara sahiptir. Durum nadir olmakla birlikte, bundan sınıf sorumludur ve ilişkili tüm alanları ve yöntemleri tamamen ayrı bir sınıfa taşımak yanlış olur. Ancak bir alt sınıfa taşınabilirler, bu yeniden düzenleme tekniğinin yardımıyla tam da bunu yapacağız.

## Faydalar

- Hızlı ve kolay bir şekilde bir alt sınıf oluşturur.

- Ana sınıfınız şu anda birden fazla özel durum uyguluyorsa, birkaç ayrı alt sınıf oluşturabilirsiniz.

## Dezavantajları

Görünüşteki basitliğine rağmen, birkaç farklı sınıf hiyerarşisini ayırmanız gerekiyorsa, Kalıtım çıkmaza yol açabilir. Örneğin, Dogsköpeklerin boyutuna ve kürküne bağlı olarak farklı davranışlara sahip bir sınıfınız varsa, iki hiyerarşiyi ortaya koyabilirsiniz:

- boyuta göre: Large, MediumveSmall

- kürk tarafından: SmoothveShaggy

Ve sadece bir sınıftan bir nesne oluşturabildiğiniz için, hem Largeve hem de olan bir köpek yaratmanız gerektiğinde sorunların ortaya çıkması dışında her şey yolunda gidecektir . Bununla birlikte, Devralmak yerine Oluştur'uSmooth kullanarak bu sorunu önleyebilirsiniz (bkz. Strateji modeli). Başka bir deyişle, sınıf, boyut ve kürk olmak üzere iki bileşen alanına sahip olacaktır. Bu alanlara gerekli sınıflardan bileşen nesneleri ekleyeceksiniz. Böylece ve olan bir tane oluşturabilirsiniz .DogDogLargeSizeShaggyFur

## Yeniden Düzenleme Nasıl Yapılır?

1. İlgilenilen sınıftan yeni bir alt sınıf oluşturun.

2. Bir alt sınıftan nesneler oluşturmak için ek verilere ihtiyacınız varsa, bir kurucu oluşturun ve buna gerekli parametreleri ekleyin. Yapıcının üst uygulamasını aramayı unutmayın.

3. Üst sınıfın yapıcısına yapılan tüm çağrıları bulun. Bir alt sınıfın işlevselliği gerekli olduğunda, ana kurucuyu alt sınıf kurucusuyla değiştirin.

4. Gerekli yöntemleri ve alanları üst sınıftan alt sınıfa taşıyın. Bunu Push Down Yöntemi ve Push Down Field ile yapın . İlk önce yöntemleri hareket ettirerek başlamak daha kolaydır. Bu şekilde, alanlar tüm süreç boyunca erişilebilir kalır: taşımadan önce ana sınıftan ve taşıma tamamlandıktan sonra alt sınıfın kendisinden.

5. Alt sınıf hazır olduktan sonra, işlevsellik seçimini kontrol eden tüm eski alanları bulun. Alanların kullanıldığı tüm operatörleri değiştirmek için polimorfizm kullanarak bu alanları silin. Basit bir örnek: Araba sınıfında alanınız vardı isElectricCarve buna bağlı olarak, refuel()yöntemde araba ya gazla doldurulur ya da elektrikle doldurulur. Yeniden düzenleme sonrası, isElectricCaralan kaldırılır ve Carve sınıfları, yöntemin ElectricCarkendi uygulamalarına sahip olur .refuel()