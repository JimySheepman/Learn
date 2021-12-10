# Tasarım deseni nedir?

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

# Kalıpların tarihi

Kalıpları kim icat etti? Bu iyi, ama çok doğru bir soru değil. Tasarım kalıpları belirsiz, karmaşık kavramlar değildir - tam tersi. Desenler, nesne yönelimli tasarımdaki yaygın sorunlara tipik çözümlerdir. Bir çözüm çeşitli projelerde defalarca tekrarlandığında, birileri sonunda ona bir isim verir ve çözümü ayrıntılı olarak açıklar. Temelde bir model bu şekilde keşfedilir.

Model kavramı ilk olarak Christopher Alexander tarafından A Pattern Language: Towns, Buildings, Construction adlı kitapta tanımlanmıştır. Kitap, kentsel çevreyi tasarlamak için bir “dil” tanımlıyor. Bu dilin birimleri kalıplardır. Pencerelerin ne kadar yüksek olması gerektiğini, bir binanın kaç katı olması gerektiğini, bir mahallede ne kadar büyük yeşil alanların olması gerektiğini vb. tanımlayabilirler.

Fikir dört yazar tarafından alındı: Erich Gamma, John Vlissides, Ralph Johnson ve Richard Helm. 1994 yılında, tasarım kalıpları kavramını programlamaya uyguladıkları Design Patterns: Elements of Reusable Object-Oriented Software'i yayınladılar. Kitap, nesne yönelimli tasarımın çeşitli problemlerini çözen 23 model içeriyor ve çok hızlı bir şekilde en çok satan haline geldi. Uzun adı nedeniyle, insanlar onu "dörtlü çetenin kitabı" olarak adlandırmaya başladılar ve kısa süre sonra sadece "GoF kitabı" olarak kısaltıldı.

O zamandan beri, düzinelerce başka nesne yönelimli kalıp keşfedildi. "Desen yaklaşımı" diğer programlama alanlarında çok popüler hale geldi, bu nedenle artık nesne yönelimli tasarımın dışında da birçok başka model var.

# Neden kalıpları öğrenmeliyim?

Gerçek şu ki, tek bir model hakkında bilgi sahibi olmadan uzun yıllar programcı olarak çalışmayı başarabilirsiniz. Pek çok insan tam da bunu yapıyor. Bu durumda bile, bilmeden bazı kalıpları uyguluyor olabilirsiniz. Öyleyse neden onları öğrenmek için zaman harcayasınız?

- Tasarım kalıpları, yazılım tasarımındaki yaygın sorunlara denenmiş ve test edilmiş çözümlerden oluşan bir araç takımıdır. Bu sorunlarla hiç karşılaşmasanız bile, kalıpları bilmek yine de yararlıdır çünkü size nesne yönelimli tasarım ilkelerini kullanarak her türlü sorunu nasıl çözeceğinizi öğretir.

- Tasarım kalıpları, sizin ve ekip arkadaşlarınızın daha verimli iletişim kurmak için kullanabileceği ortak bir dili tanımlar. “Ah, bunun için bir Singleton kullan” diyebilirsiniz ve herkes önerinizin arkasındaki fikri anlayacaktır. Deseni ve adını biliyorsanız, singleton'un ne olduğunu açıklamaya gerek yok.

# Kalıp eleştirisi

Görünüşe göre sadece tembel insanlar tasarım modellerini henüz eleştirmedi. Kalıpları kullanmaya karşı en tipik argümanlara bir göz atalım.

## Zayıf bir programlama dili için Kludges

Genellikle kalıplara duyulan ihtiyaç, insanlar gerekli soyutlama seviyesinden yoksun bir programlama dili veya teknoloji seçtiğinde ortaya çıkar. Bu durumda, kalıplar, dile çok ihtiyaç duyulan süper yetenekleri veren bir çamur haline gelir.

Örneğin, Strateji modeli, çoğu modern programlama dilinde basit bir anonim (lambda) işleviyle uygulanabilir.

## verimsiz çözümler

Modeller, halihazırda yaygın olarak kullanılan yaklaşımları sistematize etmeye çalışır. Bu birleştirme birçokları tarafından bir dogma olarak görülüyor ve kalıpları projelerinin bağlamına uyarlamadan "noktasına kadar" uyguluyorlar.

## haksız kullanım

`Sahip olduğun tek şey bir çekiçse, her şey çivi gibi görünür.`

Bu, kalıplara yeni aşina olan birçok acemi için musallat olan sorundur. Kalıpları öğrendikten sonra, daha basit kodun iyi sonuç vereceği durumlarda bile bunları her yerde uygulamaya çalışırlar.

# Kalıpların sınıflandırılması

Tasarım kalıpları, karmaşıklıklarına, ayrıntı düzeylerine ve tasarlanmakta olan tüm sisteme uygulanabilirlik derecelerine göre farklılık gösterir. Yol yapımına benzetmeyi seviyorum: Bazı trafik ışıkları kurarak veya yayalar için yeraltı geçitleriyle çok seviyeli bir kavşak inşa ederek bir kavşağı daha güvenli hale getirebilirsiniz.

En temel ve düşük seviyeli kalıplara genellikle deyimler denir. Genellikle yalnızca tek bir programlama diline uygulanırlar.

En evrensel ve üst düzey desenler mimari desenlerdir. Geliştiriciler bu kalıpları hemen hemen her dilde uygulayabilir. Diğer kalıplardan farklı olarak, tüm bir uygulamanın mimarisini tasarlamak için kullanılabilirler.

Ek olarak, tüm desenler amaçlarına veya amaçlarına göre kategorilere ayrılabilir. Bu kitap üç ana kalıp grubunu kapsar:

- Creational patterns: esnekliği ve mevcut kodun yeniden kullanımını artıran nesne oluşturma mekanizmaları sağlar.
- Structural patterns: yapıları esnek ve verimli tutarken nesnelerin ve sınıfların daha büyük yapılarda nasıl birleştirileceğini açıklar.
- Behavioral patterns: etkili iletişimi ve nesneler arasında sorumlulukların atanmasını sağlar.

# Lisans

[MIT](https://choosealicense.com/licenses/mit/)
