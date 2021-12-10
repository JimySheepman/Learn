## Factory Method

Factory Method: `Virtual Constructor Olarak da bilinir`

### Niyet

![Figure 1-1](https://refactoring.guru/images/patterns/content/factory-method/factory-method-en.png "Figure 1-1")

Fabrika Yöntemi, bir üst sınıfta nesneler oluşturmak için bir arabirim sağlayan, ancak alt sınıfların oluşturulacak nesnelerin türünü değiştirmesine izin veren yaratıcı bir tasarım modelidir.

### Sorun

Bir lojistik yönetim uygulaması oluşturduğunuzu hayal edin. Uygulamanızın ilk sürümü yalnızca kamyonlarla taşımayı gerçekleştirebilir, bu nedenle kodunuzun büyük kısmı `Truck` sınıfında bulunur.

Bir süre sonra uygulamanız oldukça popüler hale geliyor. Her gün deniz taşımacılığı şirketlerinden deniz lojistiğini uygulamaya dahil etmek için onlarca talep alıyorsunuz.

![Figure 1-2](https://refactoring.guru/images/patterns/diagrams/factory-method/problem1-en.png "Figure 1-2")

###### Kodun geri kalanı zaten mevcut sınıflara bağlıysa, programa yeni bir sınıf eklemek o kadar kolay değildir

Harika bir haber, değil mi? Peki ya kod? Şu anda kodunuzun çoğu `Truck` sınıfına bağlı. Uygulamaya `Ships` eklemek, tüm kod tabanında değişiklik yapılmasını gerektirir. Ayrıca, daha sonra uygulamaya başka bir ulaşım türü eklemeye karar verirseniz, muhtemelen tüm bu değişiklikleri tekrar yapmanız gerekecektir.

Sonuç olarak, ulaşım nesnelerinin sınıfına bağlı olarak uygulamanın davranışını değiştiren koşullarla dolu oldukça kötü bir kodla karşılaşacaksınız.

### Çözüm

Fabrika Yöntemi modeli, doğrudan nesne oluşturma çağrılarını (yeni operatörü kullanarak) özel bir fabrika yöntemine yapılan çağrılarla değiştirmenizi önerir. Endişelenmeyin: nesneler yine de yeni operatör aracılığıyla oluşturulur, ancak fabrika yönteminden çağrılır. Bir fabrika yöntemiyle döndürülen nesnelere genellikle ürün denir.

![Figure 1-3](https://refactoring.guru/images/patterns/diagrams/factory-method/solution1.png "Figure 1-3")

###### Alt sınıflar, fabrika yöntemiyle döndürülen nesnelerin sınıfını değiştirebilir

İlk bakışta, bu değişiklik anlamsız görünebilir: yapıcı çağrısını programın bir bölümünden diğerine taşıdık. Ancak şunu göz önünde bulundurun: artık bir alt sınıfta fabrika yöntemini geçersiz kılabilir ve yöntem tarafından oluşturulan ürünlerin sınıfını değiştirebilirsiniz.

Yine de küçük bir sınırlama vardır: alt sınıflar, yalnızca bu ürünler ortak bir temel sınıfa veya arabirime sahipse farklı türde ürünler döndürebilir. Ayrıca, temel sınıftaki fabrika yönteminin dönüş türü bu arabirim olarak bildirilmelidir.

![Figure 1-4](https://refactoring.guru/images/patterns/diagrams/factory-method/solution2-en.png "Figure 1-4")

###### Tüm ürünler aynı arayüzü takip etmelidir

Örneğin, hem 'Truck' hem de 'Ship' sınıfları, 'deliver' adlı bir yöntem bildiren 'Transport' arabirimini uygulamalıdır. Her sınıf bu yöntemi farklı şekilde uygular: kamyonlar kargoyu karadan, gemiler kargoyu denizden teslim eder. 'RoadLogistics' sınıfındaki fabrika yöntemi kamyon nesnelerini döndürürken, 'SeaLogistics' sınıfındaki fabrika yöntemi gemileri döndürür.

![Figure 1-5](https://refactoring.guru/images/patterns/diagrams/factory-method/solution3-en.png "Figure 1-5")

###### Tüm ürün sınıfları ortak bir arabirim uyguladığı sürece, nesnelerini kırmadan istemci koduna iletebilirsiniz

Fabrika yöntemini kullanan kod (genellikle müşteri kodu olarak adlandırılır), çeşitli alt sınıflar tarafından döndürülen gerçek ürünler arasında bir fark görmez. Müşteri, tüm ürünleri soyut 'Transport' olarak ele alır. Müşteri, tüm taşıma nesnelerinin teslim yöntemine sahip olması gerektiğini bilir, ancak tam olarak nasıl çalıştığı müşteri için önemli değildir.

### Yapı

![Figure 1-6](https://refactoring.guru/images/patterns/diagrams/factory-method/structure-indexed.png "Figure 1-6")

- `Product` yaratıcısı ve alt sınıfları tarafından üretilebilen tüm nesneler için ortak olan arabirimi bildirir.
- `Concrete Products` ürün arayüzünün farklı uygulamalarıdır.
- Creator sınıfı, yeni ürün nesnelerini döndüren fabrika yöntemini bildirir. Bu yöntemin iade türünün ürün arayüzüyle eşleşmesi önemlidir. Tüm alt sınıfları yöntemin kendi sürümlerini uygulamaya zorlamak için fabrika yöntemini soyut olarak ilan edebilirsiniz. Alternatif olarak, temel fabrika yöntemi bazı varsayılan ürün türlerini döndürebilir. Adına rağmen, ürün oluşturmanın yaratıcının birincil sorumluluğu olmadığını unutmayın. Genellikle, içerik oluşturucu sınıfı zaten ürünlerle ilgili bazı temel iş mantığına sahiptir. Fabrika yöntemi, bu mantığı somut ürün sınıflarından ayırmaya yardımcı olur. İşte bir benzetme: Büyük bir yazılım geliştirme şirketinin programcılar için bir eğitim departmanı olabilir. Ancak, bir bütün olarak şirketin birincil işlevi hala programcı üretmek değil, kod yazmaktır.
- `Concrete Creators` temel fabrika yöntemini geçersiz kılar, böylece farklı bir ürün türü döndürür.

Not: Fabrika yönteminin her zaman yeni örnekler oluşturması gerekmediğini unutmayın. Ayrıca bir önbellekten, nesne havuzundan veya başka bir kaynaktan var olan nesneleri de döndürebilir.

### Pseudocode

Bu örnek, istemci kodunu somut kullanıcı arabirimi sınıflarına bağlamadan platformlar arası kullanıcı arabirimi öğeleri oluşturmak için Fabrika Yönteminin nasıl kullanılabileceğini gösterir.

![Figure 1-6](https://refactoring.guru/images/patterns/diagrams/factory-method/structure-indexed.png "Figure 1-6")

Temel iletişim sınıfı, penceresini oluşturmak için farklı UI öğelerini kullanır. Çeşitli işletim sistemlerinde bu öğeler biraz farklı görünebilir, ancak yine de tutarlı bir şekilde davranmalıdırlar. Windows'ta bir düğme, Linux'ta hala bir düğmedir.

Fabrika yöntemi devreye girdiğinde, her işletim sistemi için iletişim kutusunun mantığını yeniden yazmanıza gerek yoktur. Temel iletişim sınıfı içinde düğmeler üreten bir fabrika yöntemi bildirirsek, daha sonra fabrika yönteminden Windows stili düğmeler döndüren bir iletişim alt sınıfı oluşturabiliriz. Alt sınıf daha sonra iletişim kutusunun kodunun çoğunu temel sınıftan devralır, ancak fabrika yöntemi sayesinde ekranda Windows görünümlü düğmeler oluşturabilir.

Bu kalıbın çalışması için, temel iletişim sınıfının soyut düğmelerle çalışması gerekir: bir temel sınıf veya tüm somut düğmelerin izlediği bir arabirim. Bu şekilde, hangi tür düğmelerle çalışırsa çalışsın, iletişim kutusunun kodu işlevsel kalır.

Elbette bu yaklaşımı diğer UI öğelerine de uygulayabilirsiniz. Ancak, diyaloğa eklediğiniz her yeni fabrika yöntemiyle, `Abstract Factory` modeline yaklaşırsınız. Korkmayın, bu model hakkında daha sonra konuşacağız.

```java

class Dialog is
    abstract method createButton():Button

    method render() is
        Button okButton = createButton()
        okButton.onClick(closeDialog)
        okButton.render()

class WindowsDialog extends Dialog is
    method createButton():Button is
        return new WindowsButton()

class WebDialog extends Dialog is
    method createButton():Button is
        return new HTMLButton()

interface Button is
    method render()
    method onClick(f)

class WindowsButton implements Button is
    method render(a, b) is
    method onClick(f) is

class HTMLButton implements Button is
    method render(a, b) is
    method onClick(f) is


class Application is
    field dialog: Dialog

    method initialize() is
        config = readApplicationConfigFile()

        if (config.OS == "Windows") then
            dialog = new WindowsDialog()
        else if (config.OS == "Web") then
            dialog = new WebDialog()
        else
            throw new Exception("Error! Unknown operating system.")

    method main() is
        this.initialize()
        dialog.render()
```

### Uygulanabilirlik

- Kodunuzun birlikte çalışması gereken nesnelerin tam türlerini ve bağımlılıklarını önceden bilmiyorsanız Fabrika Yöntemini kullanın.

- Fabrika Yöntemi, ürün yapım kodunu, ürünü gerçekten kullanan koddan ayırır. Bu nedenle, ürün yapım kodunu kodun geri kalanından bağımsız olarak genişletmek daha kolaydır. Örneğin, uygulamaya yeni bir ürün türü eklemek için yalnızca yeni bir içerik oluşturucu alt sınıfı oluşturmanız ve içindeki fabrika yöntemini geçersiz kılmanız gerekir.
- Kitaplığınızın veya çerçevenizin kullanıcılarına dahili bileşenlerini genişletmenin bir yolunu sağlamak istediğinizde Fabrika Yöntemini kullanın.
- Kalıtım, muhtemelen bir kitaplığın veya çerçevenin varsayılan davranışını genişletmenin en kolay yoludur. Ancak çerçeve, standart bir bileşen yerine alt sınıfınızın kullanılması gerektiğini nasıl anlar?

Çözüm, çerçeve genelinde bileşenleri oluşturan kodu tek bir fabrika yöntemine indirgemek ve bileşenin kendisini genişletmenin yanı sıra herkesin bu yöntemi geçersiz kılmasına izin vermektir.

Bunun nasıl işe yarayacağını görelim. Açık kaynaklı bir UI çerçevesi kullanarak bir uygulama yazdığınızı hayal edin. Uygulamanızın yuvarlak düğmeleri olmalıdır, ancak çerçeve yalnızca kare düğmeler sağlar.

Standart `Button` sınıfını muhteşem bir `RoundButton` alt sınıfıyla genişletirsiniz. Ancak şimdi ana `UIFramework` sınıfına varsayılan bir alt sınıf yerine yeni düğme alt sınıfını kullanmasını söylemeniz gerekiyor.

Bunu başarmak için, bir temel çerçeve sınıfından bir `UIWithRoundButtons` alt sınıfı yaratır ve onun `createButton` yöntemini geçersiz kılarsınız. Bu yöntem, temel sınıftaki `Button` nesnelerini döndürürken, alt sınıfınızın `RoundButton` nesnelerini döndürmesini sağlarsınız. Şimdi `UIFramework` yerine `UIWithRoundButtons` sınıfını kullanın. Ve bununla ilgili!

- Var olan nesneleri her seferinde yeniden oluşturmak yerine yeniden kullanarak sistem kaynaklarından tasarruf etmek istediğinizde Fabrika Yöntemini kullanın.

- Veritabanı bağlantıları, dosya sistemleri ve ağ kaynakları gibi büyük, kaynak yoğun nesnelerle uğraşırken bu ihtiyacı sıklıkla yaşarsınız.

Mevcut bir nesneyi yeniden kullanmak için ne yapılması gerektiğini düşünelim:

1. İlk olarak, oluşturulan tüm nesneleri takip etmek için bir miktar depolama alanı oluşturmanız gerekir.
2. Birisi bir nesne istediğinde, program o havuzun içinde boş bir nesne aramalıdır.
3. … ve ardından müşteri koduna geri gönderin.
4. Boş nesne yoksa, program yeni bir tane oluşturmalı (ve havuza eklemelidir).

Bu çok fazla kod! Ve programı yinelenen kodlarla kirletmemeniz için hepsi tek bir yere yerleştirilmelidir.

Muhtemelen bu kodun yerleştirilebileceği en bariz ve uygun yer, nesnelerini yeniden kullanmaya çalıştığımız sınıfın kurucusudur. Ancak, bir kurucu tanım gereği her zaman yeni nesneler döndürmelidir. Mevcut örnekleri döndüremez.

Bu nedenle, mevcut nesneleri yeniden kullanmanın yanı sıra yeni nesneler oluşturabilen düzenli bir yönteme sahip olmanız gerekir. Bu bir fabrika yöntemine çok benziyor.

### Nasıl Uygulanır?

1. Tüm ürünlerin aynı arayüzü takip etmesini sağlayın. Bu arabirim, her üründe anlamlı olan yöntemleri bildirmelidir.
2. Creator sınıfının içine boş bir fabrika yöntemi ekleyin. Yöntemin dönüş türü, ortak ürün arabirimiyle eşleşmelidir.
3. Yaratıcının kodunda, ürün kurucularına yapılan tüm referansları bulun. Ürün oluşturma kodunu fabrika yöntemine ayıklarken, bunları tek tek fabrika yöntemine yapılan çağrılarla değiştirin. İade edilen ürünün türünü kontrol etmek için fabrika yöntemine geçici bir parametre eklemeniz gerekebilir. Bu noktada fabrika yönteminin kodu oldukça çirkin görünebilir. Hangi ürün sınıfını somutlaştıracağını seçen büyük bir anahtar operatörüne sahip olabilir. Ama merak etmeyin, yakında düzelteceğiz.

4. Şimdi, fabrika yönteminde listelenen her ürün türü için bir dizi yaratıcı alt sınıf oluşturun. Alt sınıflarda fabrika yöntemini geçersiz kılın ve temel yöntemden uygun yapı kodu bitlerini çıkarın.

5. Çok fazla ürün türü varsa ve hepsi için alt sınıflar oluşturmak mantıklı değilse, alt sınıflarda temel sınıftan kontrol parametresini yeniden kullanabilirsiniz.

Örneğin, aşağıdaki sınıf hiyerarşisine sahip olduğunuzu düşünün: birkaç alt sınıfa sahip temel Posta sınıfı: `AirMail` ve `GroundMail`; `Transport` sınıfları `Plane`, `Truck` ve `Train`'dir. `AirMail` sınıfı yalnızca `Plane` nesnelerini kullanırken, `GroundMail` hem `Truck` hem de `Train` nesneleri ile çalışabilir. Her iki durumu da ele almak için yeni bir alt sınıf (örneğin `TrainMail`) oluşturabilirsiniz, ancak başka bir seçenek daha var. İstemci kodu, hangi ürünü almak istediğini kontrol etmek için `GroundMail` sınıfının `Factory Method` bir argüman iletebilir.

### Artıları

- Yaratıcı ve somut ürünler arasında sıkı bağlantılardan kaçınırsınız.
- Tek Sorumluluk İlkesi. Ürün oluşturma kodunu programda tek bir yere taşıyarak kodun desteklenmesini kolaylaştırabilirsiniz.
- Açık/Kapalı Prensibi. Mevcut müşteri kodunu bozmadan programa yeni ürün türlerini tanıtabilirsiniz.

### Eksileri

- Modeli uygulamak için birçok yeni alt sınıf tanıtmanız gerektiğinden kod daha karmaşık hale gelebilir. En iyi durum senaryosu, kalıbı mevcut bir oluşturucu sınıfları hiyerarşisine dahil ettiğiniz zamandır.
