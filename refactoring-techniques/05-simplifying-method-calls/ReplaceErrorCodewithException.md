# Replace Error Code with Exception

- **Sorun:** Bir yöntem, bir hatayı belirten özel bir değer döndürür?

```Java
int withdraw(int amount) {
  if (amount > _balance) {
    return -1;
  }
  else {
    balance -= amount;
    return 0;
  }
}
```

- **Çözüm:** Bunun yerine bir istisna atın.

```Java
void withdraw(int amount) throws BalanceException {
  if (amount > _balance) {
    throw new BalanceException();
  }
  balance -= amount;
}
```

## Neden Refactor

Hata kodlarını döndürmek, prosedürel programlamadan kalma eski bir kalıntıdır. Modern programlamada, hata işleme, istisnalar olarak adlandırılan özel sınıflar tarafından gerçekleştirilir. Bir sorun oluşursa, istisna işleyicilerinden biri tarafından "yakalanan" bir hatayı "fırlarsınız". Normal koşullarda yok sayılan özel hata işleme kodu yanıt vermek için etkinleştirilir.

## Faydalar

- Çeşitli hata kodlarını kontrol etmek için kodu çok sayıda koşuldan kurtarır. İstisna işleyicileri, normal yürütme yollarını anormal olanlardan ayırmanın çok daha kısa bir yoludur.

- İstisna sınıfları kendi yöntemlerini uygulayabilir, böylece hata işleme işlevinin bir kısmını içerir (örneğin, hata mesajları göndermek için).

- İstisnaların aksine, bir kurucunun yalnızca yeni bir nesne döndürmesi gerektiğinden hata kodları bir kurucuda kullanılamaz.

## Dezavantajları

Bir istisna işleyicisi, goto benzeri bir koltuk değneğine dönüşebilir. Bundan kaçının! Kod yürütmeyi yönetmek için istisnalar kullanmayın. İstisnalar yalnızca bir hatayı veya kritik durumu bildirmek için atılmalıdır.

## Yeniden Düzenleme Nasıl Yapılır?

Bu yeniden düzenleme adımlarını aynı anda yalnızca bir hata kodu için gerçekleştirmeye çalışın. Bu, tüm önemli bilgileri kafanızda tutmanızı ve hatalardan kaçınmanızı kolaylaştıracaktır.

1. Hata kodlarını döndüren bir yönteme yapılan tüm çağrıları bulun ve bir hata kodunu kontrol etmek yerine onu try/ catchbloklarına sarın.

2. Yöntemin içinde, bir hata kodu döndürmek yerine bir istisna atın.

3. Yöntem imzasını, atılan istisna hakkında bilgi içerecek şekilde değiştirin ( @throwsbölüm).
