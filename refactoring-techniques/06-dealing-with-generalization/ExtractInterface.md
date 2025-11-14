# Extract Interface

- **Sorun:** Birden çok istemci, bir sınıf arabiriminin aynı bölümünü kullanıyor. Başka bir durum: arayüzün iki sınıftaki kısmı aynıdır.

- **Çözüm:** Bu özdeş kısmı kendi arayüzüne taşıyın.

## Neden Refactor

1. Sınıflar farklı durumlarda özel roller oynadığında arayüzler çok uygundur. Hangi rolü açıkça belirtmek için Extract Interface'i kullanın .

2. Başka bir uygun durum, bir sınıfın sunucusunda gerçekleştirdiği işlemleri tanımlamanız gerektiğinde ortaya çıkar. Sonunda birden çok türdeki sunucuların kullanımına izin verilmesi planlanıyorsa, tüm sunucuların arabirimi uygulaması gerekir.

## Bunu bildiğim iyi oldu

Extract Superclass ve Extract Interface arasında belirli bir benzerlik vardır .

Bir arabirimin çıkarılması, ortak kodun değil, yalnızca ortak arabirimlerin yalıtılmasına izin verir. Başka bir deyişle, sınıflar Duplicate Code içeriyorsa , arabirimi çıkarmak, veri tekilleştirmenize yardımcı olmaz.

Yine de bu sorun , çoğaltmayı içeren davranışı ayrı bir bileşene taşımak ve tüm işi bu bileşene devretmek için Extract Class uygulanarak azaltılabilir . Ortak davranışın boyutu büyükse, her zaman Extract Superclass öğesini kullanabilirsiniz . Bu elbette daha da kolay, ancak bu yolu seçerseniz yalnızca bir ebeveyn sınıfı alacağınızı unutmayın.

## Yeniden Düzenleme Nasıl Yapılır?

1. Boş bir arayüz oluşturun.

2. Arayüzde ortak işlemleri bildirin.

3. Arayüzün uygulanması için gerekli sınıfları bildirin.

4. Yeni arabirimi kullanmak için istemci kodundaki tür bildirimlerini değiştirin.
