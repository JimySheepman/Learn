# Replace Parameter with Explicit Methods

- **Sorun:** Bir yöntem, her biri bir parametrenin değerine bağlı olarak çalıştırılan parçalara bölünür.

```Java
void setValue(String name, int value) {
  if (name.equals("height")) {
    height = value;
    return;
  }
  if (name.equals("width")) {
    width = value;
    return;
  }
  Assert.shouldNeverReachHere();
}
```

- **Çözüm:** Yöntemin tek tek parçalarını kendi yöntemlerine ayıklayın ve orijinal yöntem yerine onları çağırın.

```Java
void setHeight(int arg) {
  height = arg;
}
void setWidth(int arg) {
  width = arg;
}
```

## Neden Refactor

Parametreye bağlı değişkenler içeren bir yöntem çok büyüdü. Her dalda önemsiz olmayan kod çalıştırılır ve çok nadiren yeni varyantlar eklenir.

## Faydalar

Kod okunabilirliğini artırır. amacını anlamak çok daha startEngine()kolaydır setValue("engineEnabled", true).

## Ne Zaman Kullanılmamalı

Bir yöntem nadiren değiştiriliyorsa ve içine yeni değişkenler eklenmiyorsa, bir parametreyi açık yöntemlerle değiştirmeyin.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yöntemin her bir varyantı için ayrı bir yöntem oluşturun. Ana yöntemdeki bir parametrenin değerine göre bu yöntemleri çalıştırın.

2. Orijinal yöntemin çağrıldığı tüm yerleri bulun. Bu yerlerde, parametreye bağlı yeni varyantlardan biri için çağrı yapın.

3. Orijinal yönteme çağrı kalmadığında, silin.
