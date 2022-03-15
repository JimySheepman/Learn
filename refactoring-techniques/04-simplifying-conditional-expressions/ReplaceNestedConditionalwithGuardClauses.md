# Replace Nested Conditional with Guard Clauses

- **Sorun:** Bir grup iç içe koşul koşulunuz var ve normal kod yürütme akışını belirlemek zor.

```Java
public double getPayAmount() {
  double result;
  if (isDead){
    result = deadAmount();
  }
  else {
    if (isSeparated){
      result = separatedAmount();
    }
    else {
      if (isRetired){
        result = retiredAmount();
      }
      else{
        result = normalPayAmount();
      }
    }
  }
  return result;
}
```

- **Çözüm:** Tüm özel kontrolleri ve uç durumları ayrı maddelere ayırın ve ana kontrollerden önce yerleştirin. İdeal olarak, birbiri ardına "düz" bir koşul listeniz olmalıdır.

```Java
public double getPayAmount() {
  if (isDead){
    return deadAmount();
  }
  if (isSeparated){
    return separatedAmount();
  }
  if (isRetired){
    return retiredAmount();
  }
  return normalPayAmount();
}
```

## Neden Refactor

“Cehennemden gelen koşullu”yu tespit etmek oldukça kolaydır. Her bir yuvalama seviyesinin girintileri, acı ve ıstırap yönünde sağa işaret eden bir ok oluşturur:

```Java
if () {
    if () {
        do {
            if () {
                if () {
                    if () {
                        ...
                    }
                }
                ...
            }
            ...
        }
        while ();
        ...
    }
    else {
        ...
    }
}
```

Kod yürütmenin "normal" akışı hemen belli olmadığı için, her koşulun ne yaptığını ve nasıl yaptığını anlamak zordur. Bu koşullar, genel yapıyı optimize etmek için herhangi bir düşünce ödenmeden bir geçici önlem olarak eklenen her koşulla birlikte, birdenbire evrimi gösterir.

Durumu basitleştirmek için, özel durumları yürütmeyi hemen sonlandıran ve koruma hükümleri doğruysa boş bir değer döndüren ayrı koşullara ayırın. Aslında buradaki göreviniz yapıyı düz hale getirmek.

## Yeniden Düzenleme Nasıl Yapılır?

Yan etkilerden kurtulmaya çalışın - Separate Query from Modifier bu amaç için yardımcı olabilir. Bu çözüm, aşağıda açıklanan yeniden karıştırma için gerekli olacaktır.

1. Bir istisnanın çağrılmasına veya yöntemden bir değerin anında döndürülmesine yol açan tüm koruma maddelerini ayırın. Bu koşulları yöntemin başına yerleştirin.

2. Yeniden düzenleme tamamlandıktan ve tüm testler başarıyla tamamlandıktan sonra, aynı istisnalara veya döndürülen değerlere yol açan koruma yan tümceleri için Consolidate Conditional Expression kullanıp kullanamayacağınıza bakın .
