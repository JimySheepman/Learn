# Consolidate Duplicate Conditional Fragments

- **Sorun:** Bir koşulun tüm dallarında aynı kod bulunabilir.

```Java
if (isSpecialDeal()) {
  total = price * 0.95;
  send();
}
else {
  total = price * 0.98;
  send();
}
```

- **Çözüm:** Kodu koşullu dışında taşıyın.

```Java
if (isSpecialDeal()) {
  total = price * 0.95;
}
else {
  total = price * 0.98;
}
send();
```

## Neden Refactor

Yinelenen kod, genellikle kodun koşullu dallar içindeki evriminin bir sonucu olarak, bir koşulun tüm dallarında bulunur. Takım gelişimi buna katkıda bulunan bir faktör olabilir.

## Faydalar

Kod tekilleştirme.

## Yeniden Düzenleme Nasıl Yapılır?

1. Çoğaltılan kod koşullu dalların başındaysa, kodu koşulludan önceki bir yere taşıyın.

2. Kod dalların sonunda yürütülüyorsa, onu koşulludan sonra yerleştirin.

3. Yinelenen kod dalların içinde rastgele yer alıyorsa, sonraki kodun sonucunu değiştirip değiştirmediğine bağlı olarak önce kodu dalın başına veya sonuna taşımayı deneyin.

4. Uygunsa ve yinelenen kod bir satırdan uzunsa, Extract Method kullanmayı deneyin .
