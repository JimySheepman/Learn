# Replace Magic Number with Symbolic Constant

- **Sorun:** Kodunuz, belirli bir anlamı olan bir sayı kullanır.

```Java
double potentialEnergy(double mass, double height) {
  return mass * height * 9.81;
}
```

- **Çözüm:** Bu sayıyı, sayının anlamını açıklayan, insan tarafından okunabilir bir ada sahip bir sabitle değiştirin.

```Java
static final double GRAVITATIONAL_CONSTANT = 9.81;

double potentialEnergy(double mass, double height) {
  return mass * height * GRAVITATIONAL_CONSTANT;
}
```

## Neden Refactor

Sihirli sayı, kaynakta karşılaşılan ancak açık bir anlamı olmayan sayısal bir değerdir. Bu "anti-desen", programı anlamayı ve kodu yeniden düzenlemeyi zorlaştırır.

Ancak bu sihirli sayıyı değiştirmeniz gerektiğinde daha fazla zorluk ortaya çıkar. Bul ve değiştir bunun için çalışmaz: aynı numara farklı yerlerde farklı amaçlar için kullanılabilir, yani bu numarayı kullanan her kod satırını doğrulamanız gerekir.

## Faydalar

- Sembolik sabit, değerinin anlamının canlı belgeleri olarak hizmet edebilir.

- Bir sabitin değerini değiştirmek, başka bir yerde farklı bir amaç için kullanılan aynı sayıyı yanlışlıkla değiştirme riski olmadan, tüm kod tabanı boyunca bu sayıyı aramaktan çok daha kolaydır.

- Koddaki bir sayının veya dizenin yinelenen kullanımını azaltın. Bu, özellikle değer karmaşık ve uzun olduğunda ( 3.14159veya gibi 0xCAFEBABE) önemlidir.

## Bunu bildiğim iyi oldu

**Tüm sayılar sihirli değildir.**

Bir sayının amacı açıksa, onu değiştirmeye gerek yoktur. Klasik bir örnek:

```Java
for (i = 0; i < count; i++) { ... }
```

**Alternatifler:**

1. Bazen sihirli bir sayı, yöntem çağrılarıyla değiştirilebilir. Örneğin, bir koleksiyondaki öğelerin sayısını belirten sihirli bir numaranız varsa, koleksiyonun son öğesini kontrol etmek için onu kullanmanıza gerek yoktur. Bunun yerine, toplama uzunluğunu almak için standart yöntemi kullanın.

2. Sihirli sayılar bazen tür kodu olarak kullanılır. İki tür kullanıcınız olduğunu ve hangisinin hangisi olduğunu belirtmek için bir sınıfta bir sayı alanı kullandığınızı varsayalım: yöneticiler 1ve sıradan kullanıcılar 2.

Bu durumda, tür kodundan kaçınmak için yeniden düzenleme yöntemlerinden birini kullanmalısınız:

- Replace Type Code with Class

- Replace Type Code with Subclasses

- Replace Type Code with State/Strategy

## Yeniden Düzenleme Nasıl Yapılır?

1. Bir sabit tanımlayın ve ona sihirli sayının değerini atayın.

2. Sihirli sayının tüm sözlerini bulun.

3. Bulduğunuz sayıların her biri için, bu özel durumda sihirli sayının sabitin amacına karşılık geldiğini iki kez kontrol edin. Evet ise, sayıyı sabitinizle değiştirin. Bu önemli bir adımdır, çünkü aynı sayı tamamen farklı anlamlara gelebilir (ve duruma göre farklı sabitlerle değiştirilir).
