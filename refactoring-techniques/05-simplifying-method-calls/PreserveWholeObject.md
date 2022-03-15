# Preserve Whole Object

- **Sorun:** Bir nesneden birkaç değer alırsınız ve ardından bunları bir yönteme parametre olarak iletirsiniz.

```Java
int low = daysTempRange.getLow();
int high = daysTempRange.getHigh();
boolean withinPlan = plan.withinRange(low, high);
```

- **Çözüm:** Bunun yerine, tüm nesneyi geçirmeyi deneyin.

```Java
boolean withinPlan = plan.withinRange(daysTempRange);
```

## Neden Refactor

Sorun, yönteminiz her çağrılmadan önce, gelecekteki parametre nesnesinin yöntemlerinin çağrılması gerektiğidir. Bu yöntemler veya yöntem için elde edilen veri miktarı değiştirilirse, programda dikkatlice bir düzine yer bulmanız ve bu değişiklikleri her birinde uygulamanız gerekecektir.

Bu yeniden düzenleme tekniğini uyguladıktan sonra, gerekli tüm verileri alma kodu tek bir yerde, yani yöntemin kendisinde saklanacaktır.

## Faydalar

- Karmaşık parametreler yerine, anlaşılır bir ada sahip tek bir nesne görürsünüz.

- Yöntemin bir nesneden daha fazla veriye ihtiyacı varsa, yöntemin kullanıldığı tüm yerleri yeniden yazmanız gerekmez - yalnızca yöntemin kendi içinde.

## Dezavantajları

Bazen bu dönüşüm bir yöntemin daha az esnek olmasına neden olur: önceden yöntem birçok farklı kaynaktan veri alabilirdi, ancak şimdi yeniden düzenleme nedeniyle kullanımını yalnızca belirli bir arabirime sahip nesnelerle sınırlandırıyoruz.

## Yeniden Düzenleme Nasıl Yapılır?

1. Gerekli değerleri alabileceğiniz nesne için yöntemde bir parametre oluşturun.

2. Şimdi eski parametreleri yöntemden birer birer kaldırmaya başlayın, bunları parametre nesnesinin ilgili yöntemlerine yapılan çağrılarla değiştirin. Her parametre değişiminden sonra programı test edin.

3. Yöntem çağrısından önceki parametre nesnesinden alıcı kodunu silin.
