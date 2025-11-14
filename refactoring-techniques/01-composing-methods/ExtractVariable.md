# Extract Variable

- **Sorun:** Anlaşılması zor bir ifaden var.

```Java
void renderBanner() {
  if ((platform.toUpperCase().indexOf("MAC") > -1) &&
       (browser.toUpperCase().indexOf("IE") > -1) &&
        wasInitialized() && resize > 0 )
  {
    // do something
  }
}
```

- **Çözüm:** İfadenin sonucunu veya bölümlerini kendi kendini açıklayan ayrı değişkenlere yerleştirin.

```Java
void renderBanner() {
  final boolean isMacOs = platform.toUpperCase().indexOf("MAC") > -1;
  final boolean isIE = browser.toUpperCase().indexOf("IE") > -1;
  final boolean wasResized = resize > 0;

  if (isMacOs && isIE && wasInitialized() && wasResized) {
    // do something
  }
}
```

## Neden Refactor

Değişkenleri çıkarmanın temel nedeni, karmaşık bir ifadeyi ara parçalara bölerek daha anlaşılır kılmaktır. Bunlar şunlar olabilir:

- if()Operatörün veya operatörün bir bölümünün ?:C tabanlı dillerdeki durumu

- Ara sonuçları olmayan uzun bir aritmetik ifade

- Uzun çok parçalı çizgiler

Çıkarılan ifadenin kodunuzdaki başka yerlerde kullanıldığını görürseniz, bir değişkeni ayıklamak, Çıkarma Yöntemini gerçekleştirmenin ilk adımı olabilir.

## Faydalar

Daha okunabilir kod! Çıkarılan değişkenlere, değişkenin amacını yüksek ve net bir şekilde bildiren iyi isimler vermeye çalışın. Daha fazla okunabilirlik, daha az uzun soluklu yorumlar. customerTaxValue, cityUnemploymentRate, clientSalutationString, vb. gibi adlar için gidin .

## Dezavantajları

1. Kodunuzda daha fazla değişken var. Ancak bu, kodunuzu okuma kolaylığı ile dengelenir.

2. Koşullu ifadeleri yeniden düzenlerken, derleyicinin elde edilen değeri oluşturmak için gereken hesaplama miktarını en aza indirmek için büyük olasılıkla onu optimize edeceğini unutmayın. Aşağıdaki bir ifadeniz olduğunu söyleyin if (a() || b()) .... bYöntem adönerse , program yöntemi çağırmaz çünkü hangi değer dönerse dönsün trueelde edilen değer yine öyle olacaktır .trueb

3. Bununla birlikte, bu ifadenin parçalarını değişkenlere çıkarırsanız, her iki yöntem de her zaman çağrılır ve bu, özellikle bu yöntemler bazı ağır işler yapıyorsa, programın performansına zarar verebilir.
