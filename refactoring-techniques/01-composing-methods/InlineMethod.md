# Inline Method

- **Sorun:** Bir yöntem gövdesi, yöntemin kendisinden daha belirgin olduğunda, bu tekniği kullanın.

```Java
class PizzaDelivery {
  // ...
  int getRating() {
    return moreThanFiveLateDeliveries() ? 2 : 1;
  }
  boolean moreThanFiveLateDeliveries() {
    return numberOfLateDeliveries > 5;
  }
}
```

- **Çözüm:** Yönteme yapılan çağrıları yöntemin içeriğiyle değiştirin ve yöntemin kendisini silin.

```Java
class PizzaDelivery {
  // ...
  int getRating() {
    return numberOfLateDeliveries > 5 ? 2 : 1;
  }
}
```

## Neden Refactor

Bir yöntem basitçe başka bir yönteme yetki verir. Kendi içinde, bu delegasyon sorun değil. Ancak bu tür birçok yöntem olduğunda, çözülmesi zor, kafa karıştırıcı bir karışıklık haline gelirler.

Genellikle yöntemler başlangıçta çok kısa değildir , ancak programda değişiklikler yapıldıkça bu şekilde olur. Bu nedenle, kullanım sürelerini doldurmuş yöntemlerden kurtulmaktan çekinmeyin.

## Faydalar

Gereksiz yöntemlerin sayısını en aza indirerek kodu daha basit hale getirirsiniz.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yöntemin alt sınıflarda yeniden tanımlanmadığından emin olun. Yöntem yeniden tanımlanırsa, bu teknikten kaçının.

2. Yönteme yapılan tüm çağrıları bulun. Bu çağrıları yöntemin içeriğiyle değiştirin.

3. Yöntemi silin.
