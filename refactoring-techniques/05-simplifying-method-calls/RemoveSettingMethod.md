# Remove Setting Method

- **Sorun:** Bir alanın değeri yalnızca oluşturulduğunda ayarlanmalı ve bundan sonra hiçbir zaman değişmemelidir.

- **Çözüm:** Bu nedenle, alanın değerini belirleyen yöntemleri kaldırın.

## Neden Refactor

Bir alanın değerinde herhangi bir değişikliği önlemek istiyorsunuz.

## Yeniden Düzenleme Nasıl Yapılır?

1. Bir alanın değeri yalnızca yapıcıda değiştirilebilir olmalıdır. Yapıcı değeri ayarlamak için bir parametre içermiyorsa, bir tane ekleyin.

2. Tüm ayarlayıcı çağrılarını bulun.

    - Geçerli sınıfın yapıcısı için yapılan bir çağrıdan hemen sonra bir ayarlayıcı çağrısı bulunuyorsa, argümanını yapıcı çağrısına taşıyın ve ayarlayıcıyı kaldırın.

    - Alana doğrudan erişimle yapıcıdaki ayarlayıcı çağrılarını değiştirin.

3. Ayarlayıcıyı silin.
