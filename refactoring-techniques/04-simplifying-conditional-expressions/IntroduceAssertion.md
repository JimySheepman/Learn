# Introduce Assertion

- **Sorun:** Kodun bir bölümünün doğru çalışması için belirli koşulların veya değerlerin doğru olması gerekir.

```Java
double getExpenseLimit() {
  // Should have either expense limit or
  // a primary project.
  return (expenseLimit != NULL_EXPENSE) ?
    expenseLimit :
    primaryProject.getMemberExpenseLimit();
}
```

- **Çözüm:** Bu varsayımları belirli onaylama kontrolleriyle değiştirin.

```Java
double getExpenseLimit() {
  Assert.isTrue(expenseLimit != NULL_EXPENSE || primaryProject != null);

  return (expenseLimit != NULL_EXPENSE) ?
    expenseLimit:
    primaryProject.getMemberExpenseLimit();
}
```

## Neden Refactor

Kodun bir bölümünün, örneğin bir nesnenin mevcut durumu veya bir parametrenin veya yerel değişkenin değeri hakkında bir şey varsaydığını söyleyin. Genellikle bu varsayım, bir hata olması dışında her zaman doğru olacaktır.

Karşılık gelen iddiaları ekleyerek bu varsayımları açık hale getirin. Yöntem parametrelerinde tür ipucunda olduğu gibi, bu iddialar kodunuz için canlı belgeler işlevi görebilir.

Kodunuzun nerede iddialara ihtiyaç duyduğunu görmek için bir kılavuz olarak, belirli bir yöntemin çalışacağı koşulları tanımlayan yorumları kontrol edin.

## Faydalar

Bir varsayım doğru değilse ve bu nedenle kod yanlış sonuç veriyorsa, bu ölümcül sonuçlara ve veri bozulmasına neden olmadan önce yürütmeyi durdurmak daha iyidir. Bu aynı zamanda, programın testini gerçekleştirmenin yollarını tasarlarken gerekli bir testi yazmayı ihmal ettiğiniz anlamına gelir.

## Dezavantajları

- Bazen bir istisna, basit bir iddiadan daha uygundur. İstisnanın gerekli sınıfını seçebilir ve kalan kodun onu doğru şekilde işlemesine izin verebilirsiniz.

- Bir istisna ne zaman basit bir iddiadan daha iyidir? İstisna, kullanıcının veya sistemin eylemlerinden kaynaklanıyorsa ve istisnayı işleyebilirsiniz. Öte yandan, sıradan adsız ve işlenmeyen istisnalar temel olarak basit iddialara eşdeğerdir - onları ele almazsınız ve bunlara yalnızca asla meydana gelmemesi gereken bir program hatasının sonucu olarak neden olurlar.

## Yeniden Düzenleme Nasıl Yapılır?

Bir koşulun varsayıldığını gördüğünüzde, emin olmak için bu koşul için bir iddia ekleyin.

İddiayı eklemek, programın davranışını değiştirmemelidir.

Kodunuzdaki her şey için iddiaları kullanarak aşırıya kaçmayın . Yalnızca kodun doğru çalışması için gerekli olan koşulları kontrol edin. Belirli bir iddia yanlış olsa bile kodunuz normal çalışıyorsa, iddiayı güvenle kaldırabilirsiniz.