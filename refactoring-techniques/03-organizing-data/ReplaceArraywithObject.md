# Replace Array with Object

Replace Data Value with Object'a özel bir durumudur .

- **Sorun:** Çeşitli veri türlerini içeren bir diziniz var.

```Java
String[] row = new String[2];
row[0] = "Liverpool";
row[1] = "15";
```

- **Çözüm:** Diziyi, her öğe için ayrı alanları olacak bir nesneyle değiştirin.

```Java
Performance row = new Performance();
row.setName("Liverpool");
row.setWins("15");
```

## Neden Refactor

Diziler, tek bir türden veri ve koleksiyonları depolamak için mükemmel bir araçtır. Ancak posta kutuları gibi bir dizi kullanırsanız, kullanıcı adını kutu 1'de ve kullanıcının adresini kutu 14'te saklarsanız, bir gün yaptığınız için çok mutsuz olacaksınız. Bu yaklaşım, birisi bir şeyi yanlış "kutuya" koyduğunda feci hatalara yol açar ve ayrıca hangi verilerin nerede depolandığını bulmak için zamanınızı gerektirir.

## Faydalar

- Ortaya çıkan sınıfta, daha önce ana sınıfta veya başka bir yerde depolanmış olan tüm ilişkili davranışları yerleştirebilirsiniz.

- Bir sınıfın alanlarını belgelemek, bir dizinin öğelerinden çok daha kolaydır.

## Yeniden Düzenleme Nasıl Yapılır?

1. Dizideki verileri içerecek yeni sınıfı oluşturun. Dizinin kendisini bir ortak alan olarak sınıfa yerleştirin.

2. Bu sınıfın nesnesini orijinal sınıfta saklamak için bir alan oluşturun. Veri dizisini başlattığınız yerde nesnenin kendisini de oluşturmayı unutmayın.

3. Yeni sınıfta, dizi öğelerinin her biri için birer birer erişim yöntemleri oluşturun. Onlara ne yaptıklarını gösteren açıklayıcı isimler verin. Aynı zamanda, ana koddaki bir dizi öğesinin her kullanımını karşılık gelen erişim yöntemiyle değiştirin.

4. Tüm öğeler için erişim yöntemleri oluşturulduğunda diziyi özel yapın.

5. Dizinin her öğesi için sınıfta bir özel alan oluşturun ve erişim yöntemlerini dizi yerine bu alanı kullanacak şekilde değiştirin.

6. Tüm veriler taşındığında diziyi silin.
