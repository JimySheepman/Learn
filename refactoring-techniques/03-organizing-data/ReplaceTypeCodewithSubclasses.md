# Replace Type Code with Subclasses

- **Sorun:** Program davranışını doğrudan etkileyen kodlanmış bir türünüz var (bu alanın değerleri koşullu olarak çeşitli kodları tetikler).

- **Çözüm:** Kodlanmış türün her değeri için alt sınıflar oluşturun. Ardından ilgili davranışları orijinal sınıftan bu alt sınıflara çıkarın. Kontrol akış kodunu polimorfizm ile değiştirin.

## Neden Refactor

Bu yeniden düzenleme tekniği, Replace Type Code with Class'ta daha karmaşık bir bükülmedir .

İlk yeniden düzenleme yönteminde olduğu gibi, bir alan için izin verilen tüm değerleri oluşturan bir dizi basit değere sahipsiniz. Bu değerler genellikle sabitler olarak belirtilmesine ve anlaşılabilir adlara sahip olmasına rağmen, kullanımları hala ilkel olduklarından kodunuzu hataya açık hale getirir. Örneğin, parametrelerde bu değerlerden birini kabul eden bir yönteminiz var. Belirli bir anda, USER_TYPE_ADMIN değere sahip sabit yerine "ADMIN", yöntem aynı dizeyi küçük harfle ( "admin") alır, bu da yazarın (sizin) amaçlamadığı başka bir şeyin yürütülmesine neden olur.

Burada , koşullu  gibi kontrol akış koduyla ilgileniyoruz `if`,`switch`ve`?:`. $user->type === self::USER_TYPE_ADMIN Başka bir deyişle, bu operatörlerin koşulları içinde kodlanmış değerler ( ) gibi alanlar kullanılır.  Replace Type Code with Class'i kullanacak olsaydık , tüm bu kontrol akışı yapıları en iyi şekilde veri türünden sorumlu bir sınıfa taşınırdı. Sonuçta, bu elbette aynı problemlerle orijinaline çok benzer bir tip sınıfı yaratacaktır.

## Faydalar

- Kontrol akış kodunu silin. switch Orijinal sınıfta bir hantal yerine , kodu uygun alt sınıflara taşıyın. Bu, Tek Sorumluluk İlkesine bağlılığı artırır ve programı genel olarak daha okunabilir hale getirir.

- Kodlanmış bir tür için yeni bir değer eklemeniz gerekiyorsa, tek yapmanız gereken mevcut koda dokunmadan yeni bir alt sınıf eklemektir (bkz. Açık/Kapalı İlke ).

- Tip kodunu sınıflarla değiştirerek, programlama dili düzeyinde metotlar ve alanlar için tip ipuçlarının yolunu açıyoruz. Bu, kodlanmış bir türde bulunan basit sayısal veya dize değerleri kullanılarak mümkün olmazdı.

## Ne Zaman Kullanılmamalı

- Zaten bir sınıf hiyerarşiniz varsa bu teknik geçerli değildir. Nesne yönelimli programlamada kalıtım yoluyla ikili bir hiyerarşi oluşturamazsınız. Yine de tür kodunu kalıtım yerine kompozisyon yoluyla değiştirebilirsiniz. Bunu yapmak için Replace Type Code with State/Strateg'ı kullanın .

- Bir nesne oluşturulduktan sonra tür kodunun değerleri değişebiliyorsa, bu teknikten kaçının. Nesnenin sınıfını bir şekilde anında değiştirmemiz gerekecek, bu mümkün değil. Yine de, bu durumda bir alternatif de Type Code with State/Strategy olacaktır .

## Yeniden Düzenleme Nasıl Yapılır?

1. Tür kodunu içeren alan için bir alıcı oluşturmak için Self Encapsulate Field'ı kullanın .

2. Üst sınıf yapıcısını özel yapın. Üst sınıf oluşturucuyla aynı parametrelerle statik bir fabrika yöntemi oluşturun. Kodlanan türün başlangıç ​​değerlerini alacak parametreyi içermelidir. Bu parametreye bağlı olarak fabrika yöntemi, çeşitli alt sınıflardan nesneler yaratacaktır. Bunu yapmak için, kodunda büyük bir koşul oluşturmanız gerekir, ancak en azından gerçekten gerekli olduğunda tek koşul bu olacaktır; aksi takdirde, alt sınıflar ve polimorfizm yapacaktır.

3. Kodlanmış türün her değeri için benzersiz bir alt sınıf oluşturun. İçinde, kodlanmış türün alıcısını, kodlanmış türün karşılık gelen değerini döndürecek şekilde yeniden tanımlayın.

4. Türü kodlu alanı üst sınıftan silin. Alıcısını soyut yapın.

5. Artık alt sınıflarınız olduğuna göre, alanları ve yöntemleri üst sınıftan ilgili alt sınıflara taşımaya başlayabilirsiniz ( Push Down Field ve Push Down Method yardımıyla ).

6. Mümkün olan her şey taşındığında, tür kodunu bir kez ve tamamen kullanan koşullardan kurtulmak için Replace Conditional with Polymorphism'i kullanın.
