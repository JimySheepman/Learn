# Introduce Null Object

- **Sorun:** Gerçek nesneler yerine bazı yöntemler döndüğünden , kodunuzda nullbirçok denetiminiz vardır.null

```Java
if (customer == null) {
  plan = BillingPlan.basic();
}
else {
  plan = customer.getPlan();
}
```

- **Çözüm:** Bunun yerine null, varsayılan davranışı sergileyen boş bir nesne döndürün.

```Java
class NullCustomer extends Customer {
  boolean isNull() {
    return true;
  }
  Plan getPlan() {
    return new NullPlan();
  }
  // Some other NULL functionality.
}

// Replace null values with Null-object.
customer = (order.customer != null) ?
  order.customer : new NullCustomer();

// Use Null-object as if it's normal subclass.
plan = customer.getPlan();
```

## Neden Refactor

nullKodunuzu daha uzun ve çirkin hale getirmek için düzinelerce kontrol .

## Dezavantajları

Koşullardan kurtulmanın bedeli yeni bir sınıf daha yaratıyor.

## Yeniden Düzenleme Nasıl Yapılır?

1. Söz konusu sınıftan, boş nesne rolünü gerçekleştirecek bir alt sınıf oluşturun.

2. Her iki sınıfta da, bir boş nesne ve gerçek bir sınıf için isNull()döndürülecek olan yöntemi oluşturun.truefalse

3. null Gerçek bir nesne yerine kodun dönebileceği tüm yerleri bulun . Kodu, boş bir nesne döndürecek şekilde değiştirin.

4. Gerçek sınıfın değişkenlerinin ile karşılaştırıldığı tüm yerleri bulun null. Bu kontrolleri için bir çağrı ile değiştirin isNull().

5.
    - Bir değer eşit olmadığında orijinal sınıfın yöntemleri bu koşullandırmalarda çalıştırılırsa, nullbu yöntemleri boş sınıfta yeniden tanımlayın elseve koşulun oradaki kısmından kodu ekleyin. Daha sonra tüm koşullu silebilir ve farklı davranış polimorfizm yoluyla uygulanacaktır.

    - İşler o kadar basit değilse ve yöntemler yeniden tanımlanamıyorsa, bir değer durumunda gerçekleştirilmesi gereken operatörleri nullboş nesnenin yeni yöntemlerine basitçe ayıklayıp çıkaramayacağınıza bakın. elseVarsayılan olarak işlemler olarak eski kod yerine bu yöntemleri çağırın .
