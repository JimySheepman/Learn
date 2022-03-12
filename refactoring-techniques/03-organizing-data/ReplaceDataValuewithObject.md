# Replace Data Value with Object

- **Sorun:** Bir sınıf (veya sınıflar grubu) bir veri alanı içerir. Alanın kendi davranışı ve ilişkili verileri vardır.

- **Çözüm:** Yeni bir sınıf oluşturun, eski alanı ve davranışını sınıfa yerleştirin ve sınıfın nesnesini orijinal sınıfta saklayın.

## Neden Refactor

Bu yeniden düzenleme temelde Extract Class'ın özel bir durumudur . Onu farklı kılan şey, yeniden düzenlemenin nedenidir.

Extract Class'ta farklı şeylerden sorumlu tek bir sınıfımız var ve sorumluluklarını bölmek istiyoruz .

Bir veri değerinin bir nesneyle değiştirilmesiyle, programın büyümesi nedeniyle artık o kadar basit olmayan ve artık ilişkili veri ve davranışlara sahip olan ilkel bir alanımız (sayı, dize vb.) var. Bir yandan, bu alanlar hakkında kendi başlarına korkutucu bir şey yok. Ancak bu alanlar-ve-davranışlar ailesi, aynı anda birkaç sınıfta bulunabilir ve yinelenen kodlar oluşturabilir.

Dolayısıyla tüm bunlar için yeni bir sınıf oluşturup hem alanı hem de ilgili veri ve davranışları ona taşıyoruz.

## Faydalar

    Sınıflar arasındaki ilişkiyi geliştirir. Veriler ve ilgili davranışlar tek bir sınıfın içindedir.

## Yeniden Düzenleme Nasıl Yapılır?

Yeniden düzenlemeye başlamadan önce, sınıf içinden alana doğrudan referanslar olup olmadığına bakın. Öyleyse , orijinal sınıfta gizlemek için Self Encapsulate Field'ı kullanın.

1. Yeni bir sınıf oluşturun ve alanınızı ve ilgili alıcıyı kopyalayın. Ayrıca, alanın basit değerini kabul eden bir kurucu oluşturun. Orijinal sınıfa gönderilen her yeni alan değeri yeni bir değer nesnesi oluşturacağından, bu sınıfın bir ayarlayıcısı olmayacaktır.

2. Orijinal sınıfta alan türünü yeni sınıfla değiştirin.

3. Orijinal sınıftaki alıcıda, ilişkili nesnenin alıcısını çağırın.

4. Ayarlayıcıda yeni bir değer nesnesi oluşturun. Alan için daha önce burada başlangıç ​​değerleri ayarlanmışsa, yapıcıda yeni bir nesne oluşturmanız da gerekebilir.

## Sonraki adımlar

Bu yeniden düzenleme tekniğini uyguladıktan sonra , nesneyi içeren alana Change Value to Reference uygulamak akıllıca olacaktır. Bu, bir ve aynı değer için düzinelerce nesneyi depolamak yerine, bir değere karşılık gelen tek bir nesneye bir referansın saklanmasına izin verir.

Çoğu zaman, bir nesnenin bir gerçek dünya nesnesinden (kullanıcılar, siparişler, belgeler vb.) sorumlu olmasını istediğinizde bu yaklaşıma ihtiyaç duyulur. Aynı zamanda, bu yaklaşım tarihler, para, aralıklar vb. nesneler için yararlı olmayacaktır.
