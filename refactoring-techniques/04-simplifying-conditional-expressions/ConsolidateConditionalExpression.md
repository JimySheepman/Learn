# Consolidate Conditional Expression

- **Sorun:** Aynı sonuca veya eyleme yol açan birden çok koşulunuz var.

```Java
double disabilityAmount() {
  if (seniority < 2) {
    return 0;
  }
  if (monthsDisabled > 12) {
    return 0;
  }
  if (isPartTime) {
    return 0;
  }
  // Compute the disability amount.
  // ...
}
```

- **Çözüm:** Tüm bu koşul ifadelerini tek bir ifadede birleştirin.

```Java
double disabilityAmount() {
  if (isNotEligibleForDisability()) {
    return 0;
  }
  // Compute the disability amount.
  // ...
}
```

## Neden Refactor

Kodunuz, aynı eylemleri gerçekleştiren birçok alternatif operatör içeriyor. Operatörlerin neden bölündüğü açık değil.

Konsolidasyonun temel amacı, daha fazla netlik için koşulluyu ayrı bir yönteme çıkarmaktır.

## Faydalar

- Yinelenen kontrol akış kodunu ortadan kaldırır. Aynı "hedefe" sahip birden çok koşulu birleştirmek, tek bir eyleme yol açan yalnızca bir karmaşık kontrol yaptığınızı göstermeye yardımcı olur.

- Tüm operatörleri birleştirerek, artık bu karmaşık ifadeyi, koşulun amacını açıklayan bir adla yeni bir yöntemde yalıtabilirsiniz.

## Yeniden Düzenleme Nasıl Yapılır?

Yeniden düzenlemeden önce, koşulların herhangi bir "yan etkisi" olmadığından veya yalnızca değerleri döndürmek yerine başka bir şeyi değiştirmediğinden emin olun. Yan etkiler, operatörün kendi içinde yürütülen kodda gizleniyor olabilir; örneğin, bir koşula bağlı olarak bir değişkene bir şey eklendiğinde.

1. and ve kullanarak koşul ifadelerini tek bir ifadede birleştirin or. Konsolidasyon sırasında genel bir kural olarak:

    - İç içe koşul ifadeleri kullanılarak birleştirilir and.

    - Ardışık koşul cümleleri ile birleştirilir or.

2. Operatör koşullarında Extract Method gerçekleştirin ve yönteme ifadenin amacını yansıtan bir ad verin.
