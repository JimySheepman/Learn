# Replace Constructor with Factory Method

- **Sorun:** Nesne alanlarında parametre değerleri ayarlamaktan daha fazlasını yapan karmaşık bir kurucunuz var.

```Java
class Employee {
  Employee(int type) {
    this.type = type;
  }
  // ...
}
```

- **Çözüm:** Bir fabrika yöntemi oluşturun ve bunu yapıcı çağrılarını değiştirmek için kullanın.

```Java
class Employee {
  static Employee create(int type) {
    employee = new Employee(type);
    // do some heavy lifting.
    return employee;
  }
  // ...
}
```

## Neden Refactor

Bu yeniden düzenleme tekniğini kullanmanın en belirgin nedeni, Replace Type Code with Subclasse ile ilgilidir .

Daha önce bir nesnenin oluşturulduğu ve kodlanan türün değerinin kendisine iletildiği kodunuz var. Yeniden düzenleme yöntemini kullandıktan sonra, birkaç alt sınıf belirdi ve bunlardan kodlanmış türün değerine bağlı olarak nesneler oluşturmanız gerekiyor. Orijinal kurucuyu alt sınıf nesnelerini döndürmek için değiştirmek imkansızdır, bunun yerine gerekli sınıfların nesnelerini döndürecek ve ardından orijinal kurucuya yapılan tüm çağrıları değiştirecek statik bir fabrika yöntemi yaratırız.

Fabrika yöntemleri, inşaatçılar göreve bağlı olmadığında başka durumlarda da kullanılabilir.  Change Value to Reference çalışırken önemli olabilirler . Ayrıca, parametre sayısı ve türünün ötesine geçen çeşitli oluşturma modları ayarlamak için de kullanılabilirler.

## Faydalar

- Bir fabrika yöntemi, çağrıldığı sınıfın bir nesnesini mutlaka döndürmez. Genellikle bunlar, yönteme verilen argümanlara göre seçilen alt sınıfları olabilir.

- Bir fabrika yöntemi, neyi ve nasıl döndürdüğünü açıklayan daha iyi bir ada sahip olabilir, örneğin Troops::GetCrew(myTank).

- Bir fabrika yöntemi, her zaman yeni bir örnek oluşturan bir kurucunun aksine, önceden oluşturulmuş bir nesneyi döndürebilir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Bir fabrika yöntemi oluşturun. İçinde geçerli kurucuya bir çağrı yapın.

2. Tüm yapıcı çağrılarını fabrika yöntemine yapılan çağrılarla değiştirin.

3. Yapıcıyı private olarak bildirin.

4. Yapıcı kodunu araştırın ve mevcut sınıfın bir nesnesini oluşturmakla doğrudan ilgili olmayan kodu izole etmeye çalışın, bu kodu fabrika yöntemine taşıyın.
