# Inline Temp

- **Sorun:** Basit bir ifadenin sonucuna atanan geçici bir değişkeniniz var ve başka bir şey değil.

```Java
boolean hasDiscount(Order order) {
  double basePrice = order.basePrice();
  return basePrice > 1000;
}
```

- **Çözüm:** Değişkene yapılan başvuruları ifadenin kendisiyle değiştirin.

```Java
boolean hasDiscount(Order order) {
  return order.basePrice() > 1000;
}
```

## Neden Refactor

Satır içi yerel değişkenler hemen hemen her zaman Replace Temp with Query'in bir parçası olarak veya  Extract Method önünü açmak için kullanılır.

## Faydalar

Bu yeniden düzenleme tekniği kendi başına neredeyse hiçbir fayda sağlamaz. Bununla birlikte, değişkene bir yöntemin sonucu atanmışsa, gereksiz değişkenden kurtularak programın okunabilirliğini marjinal olarak artırabilirsiniz.

## Dezavantajları

Bazen görünüşte işe yaramaz geçici sıcaklıklar, birkaç kez yeniden kullanılan pahalı bir işlemin sonucunu önbelleğe almak için kullanılır. Bu nedenle, bu yeniden düzenleme tekniğini kullanmadan önce, basitliğin performans pahasına olmayacağından emin olun.

## Yeniden Düzenleme Nasıl Yapılır?

1. Değişkeni kullanan tüm yerleri bulun. Değişken yerine, kendisine atanmış olan ifadeyi kullanın.

2. Değişkenin bildirimini ve atama satırını silin.
