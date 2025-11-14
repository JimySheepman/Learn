# Introduce Local Extension

- **Sorun:** Bir yardımcı program sınıfı, ihtiyacınız olan bazı yöntemleri içermez. Ancak bu yöntemleri sınıfa ekleyemezsiniz.

- **Çözüm:** Yöntemleri içeren yeni bir sınıf oluşturun ve onu yardımcı program sınıfının alt öğesi veya sarmalayıcısı yapın.

## Neden Refactor

Kullanmakta olduğunuz sınıf, ihtiyacınız olan yöntemlere sahip değil. Daha da kötüsü, bu yöntemleri ekleyemezsiniz (çünkü sınıflar, örneğin bir üçüncü taraf kitaplığındadır). İki çıkış yolu var:

- İlgili sınıftan, yöntemleri içeren ve diğer her şeyi üst sınıftan miras alan bir alt sınıf oluşturun . Bu yol daha kolaydır, ancak bazen yardımcı program sınıfının kendisi tarafından engellenir ( nedeniyle final).

- Tüm yeni yöntemleri içeren bir sarmalayıcı sınıf oluşturun ve başka bir yerde yardımcı program sınıfından ilgili nesneye temsilci atanır. Bu yöntem, yalnızca sarmalayıcı ve yardımcı program nesnesi arasındaki ilişkiyi sürdürmek için koda değil, aynı zamanda yardımcı program sınıfının genel arabirimini taklit etmek için çok sayıda basit yetkilendirme yöntemine ihtiyacınız olduğundan daha fazla iş gerektirir.

## Faydalar

    Ek yöntemleri ayrı bir uzantı sınıfına (sarmalayıcı veya alt sınıf) taşıyarak, uygun olmayan kodla istemci sınıflarını birbirine karıştırmaktan kaçınırsınız. Program bileşenleri daha tutarlıdır ve daha fazla yeniden kullanılabilir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yeni bir uzantı sınıfı oluşturun:

    - Seçenek A: Onu, yardımcı program sınıfının bir çocuğu yapın.

    - Seçenek B: Bir sarmalayıcı oluşturmaya karar verdiyseniz, yetkilendirmenin yapılacağı yardımcı program sınıfı nesnesini depolamak için içinde bir alan oluşturun. Bu seçeneği kullanırken, yardımcı program sınıfının genel yöntemlerini tekrarlayan ve yardımcı program nesnesinin yöntemlerine basit yetkilendirme içeren yöntemler oluşturmanız gerekecektir.

2. Yardımcı program sınıfının yapıcısının parametrelerini kullanan bir yapıcı oluşturun.

3. Ayrıca parametrelerinde yalnızca orijinal sınıfın nesnesini alan alternatif bir "dönüştürücü" kurucu oluşturun. Bu, orijinal sınıfın nesnelerinin uzantısını değiştirmeye yardımcı olacaktır.

4. Sınıfta yeni genişletilmiş yöntemler oluşturun. Yabancı yöntemleri diğer sınıflardan bu sınıfa taşıyın veya işlevleri uzantıda zaten mevcutsa yabancı yöntemleri silin.

5. İşlevselliğinin gerekli olduğu yerlerde, yardımcı program sınıfının kullanımını yeni uzantı sınıfıyla değiştirin.
