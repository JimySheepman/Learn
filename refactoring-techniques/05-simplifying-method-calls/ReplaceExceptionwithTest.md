# Replace Exception with Test

- **Sorun:** Basit bir testin işi yapacağı bir yere istisna mı atıyorsunuz?

```Java
double getValueForPeriod(int periodNumber) {
  try {
    return values[periodNumber];
  } catch (ArrayIndexOutOfBoundsException e) {
    return 0;
  }
}
```

- **Çözüm:** İstisnayı bir durum testi ile değiştirin.

```Java
double getValueForPeriod(int periodNumber) {
  if (periodNumber >= values.length) {
    return 0;
  }
  return values[periodNumber];
}
```

## Neden Refactor

Beklenmeyen bir hatayla ilgili düzensiz davranışları işlemek için özel durumlar kullanılmalıdır. Test için bir yedek olarak hizmet etmemelidirler. Çalıştırmadan önce bir koşulu doğrulayarak bir istisnadan kaçınılabilirse, bunu yapın. İstisnalar gerçek hatalar için ayrılmalıdır.

Örneğin, bir mayın tarlasına girdiniz ve orada bir mayını tetiklediniz ve bir istisna ile sonuçlandı; istisna başarıyla halledildi ve mayın tarlasının ötesindeki emniyete hava yoluyla kaldırıldınız. Ancak başlangıçta mayın tarlasının önündeki uyarı işaretini okuyarak tüm bunlardan kaçınabilirdiniz.

## Faydalar

Basit bir koşul, bazen istisna işleme kodundan daha açık olabilir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Bir uç durum için bir koşul oluşturun ve bunu try/catch bloğundan önce taşıyın.

2. catchKodu bu koşullu içindeki bölümden taşıyın .

3. Bölümde, catcholağan bir adsız özel durum oluşturma kodunu yerleştirin ve tüm testleri çalıştırın.

4. tryTestler sırasında herhangi bir istisna atılmadıysa, / catchoperatöründen kurtulun .
