# Split Temporary Variable

- **Sorun:** Bir yöntemin içinde çeşitli ara değerleri depolamak için kullanılan yerel bir değişkeniniz var (döngü değişkenleri hariç).

```Java
double temp = 2 * (height + width);
System.out.println(temp);
temp = height * width;
System.out.println(temp);
```

- **Çözüm:** Farklı değerler için farklı değişkenler kullanın. Her değişken yalnızca belirli bir şeyden sorumlu olmalıdır.

```Java
final double perimeter = 2 * (height + width);
System.out.println(perimeter);
final double area = height * width;
System.out.println(area);
```

## Neden Refactor

Bir işlevin içindeki değişkenlerin sayısını gözden kaçırıyorsanız ve bunları ilgisiz çeşitli amaçlar için yeniden kullanıyorsanız, değişkenleri içeren kodda değişiklik yapmanız gerektiği anda sorunlarla karşılaşacağınızdan emin olabilirsiniz. Doğru değerlerin kullanıldığından emin olmak için her değişken kullanım durumunu yeniden kontrol etmeniz gerekecektir.

## Faydalar

- Program kodunun her bir bileşeni yalnızca bir ve tek bir şeyden sorumlu olmalıdır. Bu, istenmeyen etkilerden korkmadan herhangi bir şeyi kolayca değiştirebileceğiniz için kodu korumayı çok daha kolaylaştırır.

- Kod daha okunabilir hale gelir. Bir değişken uzun zaman önce aceleyle yaratılmışsa, muhtemelen hiçbir şeyi açıklamayan bir isme sahiptir: k, a2, value, vb. Ama bu durumu yeni değişkenleri anlaşılır, açıklayıcı bir şekilde adlandırarak düzeltebilirsiniz. Bu tür adlar customerTaxValue, cityUnemploymentRateve benzerlerine benzeyebilir clientSalutationString.

- Bu yeniden düzenleme tekniği, daha sonra Extract Method kullanmayı düşünüyorsanız kullanışlıdır .

## Yeniden Düzenleme Nasıl Yapılır?

1. Değişkene değer verilen koddaki ilk yeri bulun. Burada değişkeni, atanan değere karşılık gelen bir adla yeniden adlandırmalısınız.

2. Değişkenin bu değerinin kullanıldığı yerlerde eski adı yerine yeni adı kullanın.

3. Değişkene farklı bir değer atandığı yerlerde gerektiği kadar tekrarlayın.
