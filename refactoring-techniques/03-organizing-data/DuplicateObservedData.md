# Duplicate Observed Data

- **Sorun:** Etki alanı verileri, GUI'den sorumlu sınıflarda depolanıyor mu?

- **Çözüm:** Ardından, etki alanı sınıfı ile GUI arasında bağlantı ve senkronizasyon sağlayarak verileri ayrı sınıflara ayırmak iyi bir fikirdir.

## Neden Refactor

Aynı veriler için birden çok arayüz görünümüne sahip olmak istiyorsunuz (örneğin, hem masaüstü uygulamanız hem de mobil uygulamanız var). GUI'yi etki alanından ayırmayı başaramazsanız, kod tekrarından ve çok sayıda hatadan kaçınmakta çok zorlanacaksınız.

## Faydalar

- Sorumluluğu iş mantığı sınıfları ve sunum sınıfları arasında bölersiniz (bkz. Tek Sorumluluk İlkesi ), bu da programınızı daha okunaklı ve anlaşılır kılar.

- Yeni bir arayüz görünümü eklemeniz gerekiyorsa, yeni sunum sınıfları oluşturun; iş mantığının koduna dokunmanıza gerek yoktur (bkz. Açık/Kapalı İlkesi ).

- Artık farklı kişiler iş mantığı ve kullanıcı arayüzleri üzerinde çalışabilir.

## Ne Zaman Kullanılmamalı

- Self Encapsulate Field şablonu kullanılarak gerçekleştirilen bu yeniden düzenleme tekniği, tüm sınıfların web sunucusuna yapılan sorgular arasında yeniden oluşturulduğu web uygulamaları için geçerli değildir.

- Yine de, iş mantığını ayrı sınıflara ayırmanın genel ilkesi, web uygulamaları için de geçerli olabilir. Ancak bu, sisteminizin nasıl tasarlandığına bağlı olarak farklı yeniden düzenleme teknikleri kullanılarak uygulanacaktır.

## Yeniden Düzenleme Nasıl Yapılır?

1. GUI sınıfındaki etki alanı verilerine doğrudan erişimi gizleyin . Bunun için Self Encapsulate Field kullanmak en iyisidir . Böylece bu veriler için alıcıları ve ayarlayıcıları yaratırsınız.

2. GUI sınıfı olay işleyicilerinde , yeni alan değerleri ayarlamak için ayarlayıcıları kullanın. Bu, bu değerleri ilişkili etki alanı nesnesine iletmenize izin verecektir .

3. Bir etki alanı sınıfı oluşturun ve gerekli alanları GUI sınıfından ona kopyalayın . Tüm bu alanlar için alıcılar ve ayarlayıcılar oluşturun.

4. Bu iki sınıf için bir Gözlemci kalıbı oluşturun:

    - Etki alanı sınıfında , gözlemci nesnelerini ( GUI nesneleri ) ve ayrıca bunları kaydetme, silme ve bildirme yöntemlerini depolamak için bir dizi oluşturun.

    - GUI sınıfında , nesnedeki değişikliklere tepki verecek ve GUI sınıfındaki alanların değerlerini güncelleyecek olan yöntemin yanı sıra etki alanı sınıfına referansları depolamak için bir alan oluşturun . Özyinelemeyi önlemek için değer güncellemelerinin doğrudan yöntemde kurulması gerektiğini unutmayın.update()

    - GUI sınıfı yapıcısında, bir etki alanı sınıfı örneği oluşturun ve oluşturduğunuz alana kaydedin. GUI nesnesini etki alanı nesnesinde bir gözlemci olarak kaydedin .

    - Etki alanı sınıfı alanları için ayarlayıcılarda , yeni değerleri GUI'ye iletmek için gözlemciyi bilgilendirme yöntemini (diğer bir deyişle GUI sınıfında güncelleme yöntemi) çağırın.

    - GUI sınıfı alanlarının ayarlayıcılarını, doğrudan etki alanı nesnesinde yeni değerler ayarlayacak şekilde değiştirin. Değerlerin bir etki alanı sınıfı ayarlayıcı aracılığıyla ayarlanmadığından emin olun; aksi takdirde sonsuz özyineleme ortaya çıkar.
