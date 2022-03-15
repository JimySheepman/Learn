# Decompose Conditional

- **Sorun:** Karmaşık bir koşul koşulunuz var ( if-then/ elseveya switch).

```Java
if (date.before(SUMMER_START) || date.after(SUMMER_END)) {
  charge = quantity * winterRate + winterServiceCharge;
}
else {
  charge = quantity * summerRate;
}
```

- **Çözüm:** Koşulun karmaşık kısımlarını ayrı yöntemlere ayırın: koşul then ve else.

```Java
if (isSummer(date)) {
  charge = summerCharge(quantity);
}
else {
  charge = winterCharge(quantity);
}
```

## Neden Refactor

Bir kod parçası ne kadar uzunsa, anlaşılması o kadar zor olur. Kod koşullarla dolduğunda anlamak daha da zorlaşıyor:

- Bloktaki kodun ne yaptığını bulmakla meşgulken then, ilgili koşulun ne olduğunu unutuyorsunuz.

- Ayrıştırmakla meşgulken else, içindeki kodun ne yaptığını unutursunuz then.

## Faydalar

- Koşullu kodu açıkça adlandırılmış yöntemlere ayıklayarak, daha sonra kodu koruyacak olan kişinin (örneğin, bundan iki ay sonra sizin gibi!) hayatını kolaylaştırırsınız.

- Bu yeniden düzenleme tekniği, koşullarda kısa ifadeler için de geçerlidir. Dize isSalaryDay(), tarihleri ​​karşılaştırma kodundan çok daha güzel ve daha açıklayıcıdır.

## Yeniden Düzenleme Nasıl Yapılır?

- Koşul koşulunu Extract Method aracılığıyla ayrı bir yönteme çıkarın .

- then İşlemi ve elseblokları için tekrarlayın .
