# Tasarım Örüntüleri (Design Patterns)

Tasarım desenleri, yazılım tasarımında yaygın olarak ortaya çıkan sorunlara tipik çözümlerdir. Kodunuzda yinelenen bir tasarım sorununu çözmek için özelleştirebileceğiniz önceden hazırlanmış planlar gibidirler.

Kullanıma hazır işlevler veya kitaplıklarda olduğu gibi, yalnızca bir kalıp bulup programınıza kopyalayamazsınız. Model, belirli bir kod parçası değil, belirli bir sorunu çözmek için genel bir kavramdır. Kalıp detaylarını takip edebilir ve kendi programınızın gerçeklerine uygun bir çözüm uygulayabilirsiniz.

Modeller genellikle algoritmalarla karıştırılır, çünkü her iki kavram da bilinen bazı problemlerin tipik çözümlerini tanımlar. Bir algoritma her zaman bir hedefe ulaşabilecek net bir eylemler kümesi tanımlarken, bir model bir çözümün daha üst düzey bir açıklamasıdır. Aynı desenin iki farklı programa uygulanan kodu farklı olabilir.

Bir algoritmaya benzetme bir yemek tarifidir: her ikisinin de bir hedefe ulaşmak için net adımları vardır. Öte yandan, bir model daha çok bir taslak gibidir: sonucun ve özelliklerinin ne olduğunu görebilirsiniz, ancak kesin uygulama sırası size kalmış.
Desen nelerden oluşur?

Çoğu kalıp, çok resmi olarak tanımlanır, böylece insanlar onları birçok bağlamda yeniden üretebilir. Bir model açıklamasında genellikle bulunan bölümler şunlardır:

- Modelin amacı, hem sorunu hem de çözümü kısaca açıklar.
- Motivasyon, sorunu ve kalıbın mümkün kıldığı çözümü daha da açıklar.
- Sınıfların yapısı, modelin her bir parçasını ve bunların nasıl ilişkili olduğunu gösterir.
- Popüler programlama dillerinden birindeki kod örneği, kalıbın arkasındaki fikri kavramayı kolaylaştırır.

Bazı kalıp katalogları, kalıbın uygulanabilirliği, uygulama adımları ve diğer kalıplarla ilişkiler gibi diğer faydalı detayları listeler.

## Neden kalıpları öğrenmeliyim?

Gerçek şu ki, tek bir model hakkında bilgi sahibi olmadan uzun yıllar programcı olarak çalışmayı başarabilirsiniz. Pek çok insan tam da bunu yapıyor. Bu durumda bile, bilmeden bazı kalıpları uyguluyor olabilirsiniz. Öyleyse neden onları öğrenmek için zaman harcayasınız?

- Tasarım kalıpları, yazılım tasarımındaki yaygın sorunlara denenmiş ve test edilmiş çözümlerden oluşan bir araç takımıdır. Bu sorunlarla hiç karşılaşmasanız bile, kalıpları bilmek yine de yararlıdır çünkü size nesne yönelimli tasarım ilkelerini kullanarak her türlü sorunu nasıl çözeceğinizi öğretir.

- Tasarım kalıpları, sizin ve ekip arkadaşlarınızın daha verimli iletişim kurmak için kullanabileceği ortak bir dili tanımlar. “Ah, bunun için bir Singleton kullan” diyebilirsiniz ve herkes önerinizin arkasındaki fikri anlayacaktır. Deseni ve adını biliyorsanız, singleton'un ne olduğunu açıklamaya gerek yok.

## Kalıpların sınıflandırılması

Tasarım kalıpları, karmaşıklıklarına, ayrıntı düzeylerine ve tasarlanmakta olan tüm sisteme uygulanabilirlik derecelerine göre farklılık gösterir. Yol yapımına benzetmeyi seviyorum: Bazı trafik ışıkları kurarak veya yayalar için yeraltı geçitleriyle çok seviyeli bir kavşak inşa ederek bir kavşağı daha güvenli hale getirebilirsiniz.

En temel ve düşük seviyeli kalıplara genellikle deyimler denir. Genellikle yalnızca tek bir programlama diline uygulanırlar.

En evrensel ve üst düzey desenler mimari desenlerdir. Geliştiriciler bu kalıpları hemen hemen her dilde uygulayabilir. Diğer kalıplardan farklı olarak, tüm bir uygulamanın mimarisini tasarlamak için kullanılabilirler.

Ek olarak, tüm desenler amaçlarına veya amaçlarına göre kategorilere ayrılabilir. Bu kitap üç ana kalıp grubunu kapsar:

### Creational Patterns

Esnekliği ve mevcut kodun yeniden kullanımını artıran nesne oluşturma mekanizmaları sağlar. `new` keyword kullanımını yönetir!

- Abstact Factory
- Builder
- Factory Method
- Prototype
- Singleton

### Structural Patterns

Yapıları esnek ve verimli tutarken nesnelerin ve sınıfların daha büyük yapılarda nasıl birleştirileceğini açıklar. Sınıflar arası kurulan doğru ilişki/bağlantılar üzerinden kurduğu yapılar ile ilgilenir bunları düzenlemeye ve anlamlandırmaya çalışır.

- Adapter
- Bridge
- Composite
- Decorator
- Facade
- Flyweight
- Proxy

### Behavioral Patterns

Etkili iletişimi ve nesneler arasında sorumlulukların atanmasını sağlar. Davranışsal(Behavioral) örüntüler ise daha çok algoritmalar, algoritmaların, davranışların soyutlanması üzerine olan örüntülerdir. Algoritmalar arası iletişim şekli, sorumlulukların paylaşımı veya aktarımının nasıl yapılacağını tanımlar.

- Chain Of Responsibility
- Command
- Interpreter
- Mediator
- Memento
- Observer
- State
- Strategy
- Template Method
- Visitor

## Lisans

[MIT](https://choosealicense.com/licenses/mit/)
