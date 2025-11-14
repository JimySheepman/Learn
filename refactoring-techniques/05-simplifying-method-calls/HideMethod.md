# Hide Method

- **Sorun:** Bir yöntem, diğer sınıflar tarafından kullanılmaz veya yalnızca kendi sınıf hiyerarşisi içinde kullanılır.

- **Çözüm:** Yöntemi özel veya korumalı yapın.

## Neden Refactor

Çoğu zaman, değerleri alma ve ayarlama yöntemlerini gizleme ihtiyacı, ek davranış sağlayan daha zengin bir arabirimin geliştirilmesinden kaynaklanır, özellikle de yalnızca veri kapsüllemenin biraz ötesine geçen bir sınıfla başladıysanız.

Sınıfa yeni davranış eklendikçe, genel alıcı ve ayarlayıcı yöntemlerinin artık gerekli olmadığını ve gizlenebileceğini görebilirsiniz. Alıcı veya ayarlayıcı yöntemlerini özel yaparsanız ve değişkenlere doğrudan erişim uygularsanız, yöntemi silebilirsiniz.

## Faydalar

- Yöntemleri gizlemek, kodunuzun gelişmesini kolaylaştırır. Özel bir yöntemi değiştirdiğinizde, yöntemin başka hiçbir yerde kullanılamayacağını bildiğiniz için yalnızca mevcut sınıfı nasıl kırmayacağınız konusunda endişelenmeniz gerekir.

- Yöntemleri özel yaparak, sınıfın genel arabiriminin ve genel olarak kalan yöntemlerin önemini vurgularsınız.

## Yeniden Düzenleme Nasıl Yapılır?

1. Düzenli olarak özel hale getirilebilecek yöntemler bulmaya çalışın. Statik kod analizi ve iyi birim test kapsamı, büyük bir destek sağlayabilir.

2. Her yöntemi mümkün olduğunca özel yapın.
