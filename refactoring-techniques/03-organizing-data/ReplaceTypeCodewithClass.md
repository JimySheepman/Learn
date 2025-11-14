# Replace Type Code with Class

- **Sorun:** Bir sınıfın, tür kodunu içeren bir alanı vardır. Bu tür değerler operatör koşullarında kullanılmaz ve programın davranışını etkilemez.

- **Çözüm:** Yeni bir sınıf oluşturun ve tür kodu değerleri yerine nesnelerini kullanın

## Neden Refactor

Tür kodunun en yaygın nedenlerinden biri, bir veritabanında bazı karmaşık kavramların bir sayı veya dizeyle kodlandığı alanlar olduğunda, veritabanlarıyla çalışmaktır.

Örneğin, yönetici, düzenleyici veya sıradan kullanıcı olsun, her kullanıcının erişim ayrıcalıkları hakkında bilgi içeren Useralana sahip bir sınıfınız var. user_roleYani bu durumda, bu bilgi alana sırasıyla A, E, ve olarak kodlanır U.

Bu yaklaşımın eksiklikleri nelerdir? Alan belirleyiciler genellikle hangi değerin gönderildiğini kontrol etmez, bu da birisi bu alanlara istenmeyen veya yanlış değerler gönderdiğinde büyük sorunlara neden olabilir.

Ayrıca bu alanlar için tip doğrulaması yapılamaz. Onlara herhangi bir sayı veya dize göndermek mümkündür, bunlar IDE'niz tarafından tip denetimi yapılmaz ve hatta programınızın çalışmasına (ve daha sonra çökmesine) izin verir.

## Faydalar

- Kodlanmış türler olan ilkel değer kümelerini, nesne yönelimli programlamanın sunduğu tüm avantajlarla birlikte tam teşekküllü sınıflara dönüştürmek istiyoruz.

- Tür kodunu sınıflarla değiştirerek, programlama dili düzeyinde yöntemlere ve alanlara iletilen değerler için tür ipuçlarına izin veriyoruz.

- Örneğin, derleyici daha önce bir yönteme bir değer iletildiğinde sayısal sabitiniz ile bazı rastgele sayılar arasındaki farkı görmemişken, şimdi belirtilen tür sınıfına uymayan veriler iletildiğinde, uyarılırsınız. IDE'nizin içindeki hata.

- Böylece kodu türün sınıflarına taşımayı mümkün kılıyoruz. Tüm program boyunca tür değerleriyle karmaşık işlemler yapmanız gerekiyorsa, artık bu kod bir veya birden çok tür sınıfında "canlı" olabilir.

## Ne Zaman Kullanılmamalı

Kodlanmış bir türün değerleri, kontrol akış yapıları ( if, switch, vb.) içinde kullanılıyorsa ve bir sınıf davranışını kontrol ediyorsa, tür kodu için iki yeniden düzenleme tekniğinden birini kullanmalısınız:

- Replace Type Code with Subclasses

- Replace Type Code with State/Strategy

## Yeniden Düzenleme Nasıl Yapılır?

1. Yeni bir sınıf oluşturun ve kodlanmış türün amacına karşılık gelen yeni bir ad verin. Burada buna type class diyeceğiz .

2. Tip kodunu içeren alanı tip sınıfına kopyalayın ve onu özel yapın. Ardından alan için bir alıcı oluşturun. Bu alan için yalnızca yapıcıdan bir değer ayarlanacaktır.

3. Kodlanmış türün her değeri için class türünde statik bir yöntem oluşturun . Kodlanmış türün bu değerine karşılık gelen yeni bir tür sınıfı nesnesi yaratacaktır.

4. Orijinal sınıfta, kodlanmış alanın türünü class type ile değiştirin . Alan ayarlayıcının yanı sıra yapıcıda bu türden yeni bir nesne oluşturun. Alan alıcısını, tür sınıfı alıcısını çağıracak şekilde değiştirin .

5. Kodlanmış türdeki değerlerden söz edilenleri, ilgili tür sınıfı statik yöntemlerin çağrılarıyla değiştirin.

6. Orijinal sınıftan kodlanmış tür sabitlerini kaldırın.
