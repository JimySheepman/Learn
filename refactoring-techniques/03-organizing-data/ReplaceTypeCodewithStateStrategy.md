# Replace Type Code with State/Strategy

**Tür kodu nedir?** Tür kodu, ayrı bir veri türü yerine, bazı varlıklar için izin verilen değerler listesini oluşturan bir dizi sayı veya dizeye sahip olduğunuzda oluşur. Genellikle bu belirli sayılara ve dizelere sabitler aracılığıyla anlaşılır adlar verilir, bu tür kodlara bu kadar çok rastlanmasının nedeni budur.

- **Sorun:** Davranışı etkileyen kodlanmış bir türünüz var ama ondan kurtulmak için alt sınıfları kullanamazsınız.

- **Çözüm:** Tür kodunu bir durum nesnesiyle değiştirin. Bir alan değerini tür koduyla değiştirmek gerekirse, başka bir durum nesnesi “takılır”.

## Neden Refactor

Tür kodunuz var ve bu bir sınıfın davranışını etkiler, bu nedenle  Replace Type Code with Class kullanamayız .

Tip kodu, bir sınıfın davranışını etkiler, ancak mevcut sınıf hiyerarşisi veya diğer nedenlerden dolayı kodlanan tür için alt sınıflar oluşturamıyoruz. Bu, Type Code with Subclasses'ı uygulayamayacağımız anlamına gelir .

## Faydalar

- Bu yeniden düzenleme tekniği, kodlanmış bir türe sahip bir alanın nesnenin ömrü boyunca değerini değiştirdiği durumlardan kurtulmanın bir yoludur. Bu durumda, değerin değiştirilmesi, orijinal sınıfın başvurduğu durum nesnesinin değiştirilmesi yoluyla yapılır.

- Kodlanmış türden yeni bir değer eklemeniz gerekiyorsa, tek yapmanız gereken mevcut kodu değiştirmeden yeni bir durum alt sınıfı eklemektir (bkz. Açık/Kapalı İlke ).

## Dezavantajları

    Basit bir tür kodunuz varsa, ancak yine de bu yeniden düzenleme tekniğini kullanırsanız, birçok ekstra (ve gereksiz) sınıfınız olacaktır.

## Bunu bildiğim iyi oldu

Bu yeniden düzenleme tekniğinin uygulanması, iki tasarım modelinden birini kullanabilir: Durum veya Strateji . Hangi modeli seçerseniz seçin uygulama aynıdır. Peki belirli bir durumda hangi modeli seçmelisiniz?

Algoritma seçimini kontrol eden bir koşulu bölmeye çalışıyorsanız, Strateji'yi kullanın.

Ancak, kodlanmış türün her bir değeri yalnızca bir algoritma seçiminden değil, sınıfın tüm durumundan, sınıf durumu, alan değerleri ve diğer birçok eylemden sorumluysa, Durum iş için daha iyidir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Tür kodunu içeren alan için bir alıcı oluşturmak için Self Encapsulate Field'ı kullanın .

2. Yeni bir sınıf oluşturun ve ona tür kodunun amacına uyan anlaşılır bir ad verin. Bu sınıf devlet (veya strateji ) rolünü oynayacaktır . İçinde, soyut kodlanmış bir alan alıcısı oluşturun.

3. Kodlanmış türün her değeri için durum sınıfının alt sınıflarını oluşturun. Her alt sınıfta, kodlanmış alanın alıcısını, kodlanmış türün karşılık gelen değerini döndürecek şekilde yeniden tanımlayın.

4. Soyut durum sınıfında, kodlanmış türün değerini parametre olarak kabul eden statik bir fabrika yöntemi oluşturun. Bu parametreye bağlı olarak, fabrika yöntemi çeşitli durumlarda nesneler yaratacaktır. Bunun için kodunda büyük bir koşullu oluşturun; yeniden düzenleme tamamlandığında tek olacak.

5. Orijinal sınıfta, kodlanmış alanın türünü durum sınıfına değiştirin. Alanın ayarlayıcısında, yeni durum nesnelerini almak için fabrika durumu yöntemini çağırın.

6. Artık alanları ve yöntemleri üst sınıftan karşılık gelen durum alt sınıflarına taşımaya başlayabilirsiniz (Plan ve Push Down Method kullanarak ).

7. Hareket edebilen her şey taşındığında, tür kodunu bir kez ve herkes için kullanan koşullardan kurtulmak için Replace Conditional with Polymorphism kullanın.
