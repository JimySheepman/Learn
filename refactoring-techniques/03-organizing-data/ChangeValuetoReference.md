# Change Value to Reference

- **Sorun:** Yani, tek bir nesneyle değiştirmeniz gereken, tek bir sınıfın birçok özdeş örneğine sahipsiniz.

- **Çözüm:** Aynı nesneleri tek bir referans nesnesine dönüştürün.

## Neden Refactor

Birçok sistemde nesneler ya değerler ya da referanslar olarak sınıflandırılabilir.

- References : bir gerçek dünya nesnesi programdaki yalnızca bir nesneye karşılık geldiğinde. Referanslar genellikle kullanıcı/sipariş/ürün/vb. nesneler.

- Values : bir gerçek dünya nesnesi, programdaki birden çok nesneye karşılık gelir. Bu nesneler tarihler, telefon numaraları, adresler, renkler ve benzerleri olabilir.

Referansa karşı değer seçimi her zaman net değildir. Bazen az miktarda değişmeyen veri içeren basit bir değer vardır. Daha sonra değiştirilebilir veriler eklemek ve bu değişiklikleri nesneye her erişildiğinde geçmek gerekli hale gelir. Bu durumda onu bir referansa dönüştürmek gerekli hale gelir.

## Faydalar

Bir nesne, belirli bir varlık hakkında en güncel bilgileri içerir. Programın bir bölümünde nesne değiştirilirse, bu değişikliklere programın nesneyi kullanan diğer bölümlerinden erişilebilir.

## Dezavantajları

Referansların uygulanması çok daha zordur.

## Yeniden Düzenleme Nasıl Yapılır?

1. Referansların oluşturulacağı sınıfta  Replace Constructor with Factory Method'i kullanın .

2. Referanslara erişim sağlamaktan hangi nesnenin sorumlu olacağını belirleyin. Yeni bir nesne oluşturmak yerine, ihtiyacınız olduğunda onu bir depolama nesnesinden veya statik sözlük alanından almanız gerekir.

3. Referansların önceden mi yoksa dinamik olarak mı oluşturulacağını belirleyin. Nesneler önceden oluşturulmuşsa, kullanmadan önce yüklediğinizden emin olun.

4. Bir referans döndürecek şekilde fabrika yöntemini değiştirin. Nesneler önceden oluşturulmuşsa, var olmayan bir nesne istendiğinde hataların nasıl ele alınacağına karar verin. Ayrıca , yöntemin yalnızca mevcut nesneleri döndürdüğünü bildirmek için  Rename Method Yöntemini kullanmanız gerekebilir .
