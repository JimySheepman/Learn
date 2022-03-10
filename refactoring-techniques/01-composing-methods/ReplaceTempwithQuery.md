# Replace Temp with Query

- **Sorun:** Bir ifadenin sonucunu, kodunuzda daha sonra kullanmak üzere yerel bir değişkene yerleştirirsiniz.

```Java
double calculateTotal() {
  double basePrice = quantity * itemPrice;
  if (basePrice > 1000) {
    return basePrice * 0.95;
  }
  else {
    return basePrice * 0.98;
  }
}
```

- **Çözüm:** Tüm ifadeyi ayrı bir yönteme taşıyın ve sonucu döndürün. Değişken kullanmak yerine yöntemi sorgulayın. Gerekirse, yeni yöntemi diğer yöntemlere dahil edin.

```Java
double calculateTotal() {
  if (basePrice() > 1000) {
    return basePrice() * 0.95;
  }
  else {
    return basePrice() * 0.98;
  }
}
double basePrice() {
  return quantity * itemPrice;
}
```

## Neden Refactor

Bu yeniden düzenleme, çok uzun bir yöntemin bir kısmı için  Extract Method uygulanması için zemin hazırlayabilir .

Aynı ifade bazen başka yöntemlerde de bulunabilir, bu da ortak bir yöntem oluşturmayı düşünmek için bir nedendir.
Faydalar

- Kod okunabilirliği. Yöntemin amacını anlamak getTax()satırdan çok daha kolaydır orderPrice() * 0.2.

- Değiştirilen satır birden fazla yöntemde kullanılıyorsa, veri tekilleştirme yoluyla daha ince kod.

## Bunu bildiğim iyi oldu

Bu yeniden düzenleme, bu yaklaşımın bir performans isabetine neden olup olmadığı sorusunu gündeme getirebilir. Dürüst cevap şudur: evet, öyle, çünkü ortaya çıkan kod yeni bir yöntemi sorgulayarak yüklenebilir. Ancak günümüzün hızlı CPU'ları ve mükemmel derleyicileri ile yük neredeyse her zaman minimum düzeyde olacaktır. Buna karşılık, okunabilir kod ve bu yöntemi program kodundaki başka yerlerde yeniden kullanma yeteneği - bu yeniden düzenleme yaklaşımı sayesinde - çok dikkat çekici avantajlardır.

Bununla birlikte, temp değişkeniniz gerçekten zaman alan bir ifadenin sonucunu önbelleğe almak için kullanılıyorsa, ifadeyi yeni bir yönteme çıkardıktan sonra bu yeniden düzenlemeyi durdurmak isteyebilirsiniz.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yöntem içinde değişkene yalnızca bir kez bir değer atandığından emin olun. Değilse , değişkenin yalnızca ifadenizin sonucunu depolamak için kullanılacağından emin olmak için Split Temporary Variable.

2. İlgi ifadesini yeni bir yönteme yerleştirmek için Extract Method kullanın . Bu yöntemin yalnızca bir değer döndürdüğünden ve nesnenin durumunu değiştirmediğinden emin olun. Yöntem, nesnenin görünür durumunu etkiliyorsa, Separate Query from Modifier'yu kullanın .

3. Değişkeni yeni yönteminize bir sorgu ile değiştirin.
