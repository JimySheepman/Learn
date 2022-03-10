# Inline Class

- **Sorun:** Bir sınıf neredeyse hiçbir şey yapmaz ve hiçbir şeyden sorumlu değildir ve bunun için hiçbir ek sorumluluk planlanmaz.

- **Çözüm:** Tüm özellikleri sınıftan diğerine taşıyın.

## Neden Refactor

Genellikle bu tekniğe, bir sınıfın özellikleri diğer sınıflara "aktarıldıktan" sonra ihtiyaç duyulur ve bu sınıfa yapacak çok az şey kalır.

## Faydalar

Gereksiz sınıfları ortadan kaldırmak, bilgisayardaki işletim belleğini ve kafanızdaki bant genişliğini boşaltır.

## Yeniden Düzenleme Nasıl Yapılır?

1. Alıcı sınıfta, verici sınıfta bulunan ortak alanları ve yöntemleri oluşturun. Yöntemler, verici sınıfın eşdeğer yöntemlerine başvurmalıdır.

2. Verici sınıfa yapılan tüm başvuruları, alıcı sınıfın alanlarına ve yöntemlerine yapılan başvurularla değiştirin.

3. Şimdi programı test edin ve herhangi bir hata eklenmediğinden emin olun. Testler her şeyin A-OK çalıştığını gösteriyorsa, tüm işlevleri orijinal olandan alıcı sınıfa tamamen nakletmek için Move Method ve Move Field'ı kullanmaya başlayın. Orijinal sınıf tamamen boşalana kadar bunu yapmaya devam edin.

4. Orijinal sınıfı silin.
