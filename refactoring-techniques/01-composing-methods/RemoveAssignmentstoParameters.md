# Remove Assignments to Parameters

- **Sorun:** Yöntemin gövdesindeki bir parametreye bir değer atanır.

```Java
int discount(int inputVal, int quantity) {
  if (quantity > 50) {
    inputVal -= 2;
  }
  // ...
}
```

- **Çözüm:** Parametre yerine yerel bir değişken kullanın.

```Java
int discount(int inputVal, int quantity) {
  int result = inputVal;
  if (quantity > 50) {
    result -= 2;
  }
  // ...
}
```

## Neden Refactor

Bu yeniden düzenlemenin nedenleri  Split Temporary Variable ile aynıdır , ancak bu durumda yerel bir değişkenle değil, bir parametreyle ilgileniyoruz.

İlk olarak referans yoluyla bir parametre iletilirse, daha sonra parametre değeri yöntem içinde değiştirildikten sonra, bu değer bu yöntemin çağrılmasını talep eden argümana iletilir. Çok sık olarak, bu kazara meydana gelir ve talihsiz etkilere yol açar. Programlama dilinizde parametreler genellikle değere göre (başvuruya göre değil) iletilse bile, bu kodlama tuhaflığı, buna alışkın olmayanları yabancılaştırabilir.

İkincisi, tek bir parametreye farklı değerlerin birden fazla atanması, herhangi bir zamanda parametrede hangi verilerin bulunması gerektiğini bilmenizi zorlaştırır. Parametreniz ve içeriği belgelenirse sorun daha da kötüleşir, ancak gerçek değer yöntem içinde beklenenden farklı olabilir.

## Faydalar

- Programın her elemanı sadece bir şeyden sorumlu olmalıdır. Bu, kodu herhangi bir yan etki olmadan güvenle değiştirebileceğiniz için kod bakımını ileriye doğru çok daha kolay hale getirir.

- Bu yeniden düzenleme, repetitive code to separate methods çıkarılmasına yardımcı olur .

## Yeniden Düzenleme Nasıl Yapılır?

1. Yerel bir değişken oluşturun ve parametrenizin başlangıç ​​değerini atayın.

2. Bu satırı izleyen tüm yöntem kodlarında parametreyi yeni yerel değişkeninizle değiştirin.
