# Introduce Foreign Method

- **Sorun:** Bir yardımcı program sınıfı, ihtiyacınız olan yöntemi içermez ve bu yöntemi sınıfa ekleyemezsiniz.

```Java
class Report {
  // ...
  void sendReport() {
    Date nextDay = new Date(previousEnd.getYear(),
      previousEnd.getMonth(), previousEnd.getDate() + 1);
    // ...
  }
}
```

- **Çözüm:** Yöntemi bir istemci sınıfına ekleyin ve yardımcı program sınıfının bir nesnesini ona argüman olarak iletin.

```Java
class Report {
  // ...
  void sendReport() {
    Date newStart = nextDay(previousEnd);
    // ...
  }
  private static Date nextDay(Date arg) {
    return new Date(arg.getYear(), arg.getMonth(), arg.getDate() + 1);
  }
}
```

## Neden Refactor

Belirli bir sınıfın verilerini ve yöntemlerini kullanan kodunuz var. Kodun, sınıftaki yeni bir yöntemde çok daha iyi görüneceğini ve çalışacağını anlıyorsunuz. Ancak, örneğin sınıf bir üçüncü taraf kitaplığında bulunduğundan, yöntemi sınıfa ekleyemezsiniz.

Yönteme taşımak istediğiniz kod, programınızdaki farklı yerlerde birkaç kez tekrarlandığında, bu yeniden düzenlemenin büyük bir getirisi vardır.

Yeni yöntemin parametrelerine yardımcı program sınıfının bir nesnesini ilettiğiniz için, tüm alanlarına erişiminiz olur. Yöntemin içinde, sanki yöntem yardımcı program sınıfının bir parçasıymış gibi, neredeyse istediğiniz her şeyi yapabilirsiniz.

## Faydalar

Kod tekrarını kaldırır. Kodunuz birkaç yerde tekrarlanıyorsa, bu kod parçalarını bir yöntem çağrısı ile değiştirebilirsiniz. Bu, yabancı yöntemin optimal olmayan bir yerde bulunduğu göz önüne alındığında bile, çoğaltmadan daha iyidir.

## Dezavantajları

Bir istemci sınıfında bir yardımcı sınıf yöntemine sahip olmanın nedenleri, sizden sonra kodu koruyan kişi için her zaman açık olmayacaktır. Metot diğer sınıflarda kullanılabiliyorsa, yardımcı program sınıfı için bir sarmalayıcı oluşturup metodu oraya yerleştirerek fayda sağlayabilirsiniz. Bu, bu tür birkaç yardımcı yöntem olduğunda da faydalıdır.  Introduce Local Extensionbu konuda yardımcı olabilir.

## Yeniden Düzenleme Nasıl Yapılır?

1. İstemci sınıfında yeni bir yöntem oluşturun.

2. Bu yöntemde, yardımcı program sınıfının nesnesinin iletileceği bir parametre oluşturun. Bu nesne client sınıfından alınabiliyorsa, böyle bir parametre oluşturmanız gerekmez.

3. İlgili kod parçalarını bu yönteme çıkarın ve bunları yöntem çağrılarıyla değiştirin.

4. Yabancı yöntem etiketini, daha sonra mümkün olursa, bu yöntemi bir yardımcı program sınıfına yerleştirme tavsiyesiyle birlikte yöntemin yorumlarına bıraktığınızdan emin olun . Bu, gelecekte yazılımın bakımını yapacak olanlar için bu yöntemin neden bu özel sınıfta bulunduğunu anlamayı kolaylaştıracaktır.
