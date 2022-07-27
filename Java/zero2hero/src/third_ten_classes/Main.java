package third_ten_classes;

/*
! Data Structures
* The Enumeration : Bir nesneler koleksiyonundaki öğeleri sıralayabileceğiniz (birer birer elde edebileceğiniz)
*                             yöntemleri tanımlar. Yerine Iterator kullanılıyor artık
* The BitSet :  Bit değerlerini tutan özel bir dizi türü oluşturur. BitSet dizisinin boyutu gerektiği gibi artabilir.
*               Bu onu bir bit vektörüne benzer hale getirir.
* The Vector : Vektör, dinamik bir dizi uygular. ArrayList'e benzer, ancak iki farkla:
*                -   Vektör senkronize edilir.
*                -   Vector, koleksiyon çerçevesinin parçası olmayan birçok eski yöntemi içerir.
*              Dizinin boyutunu önceden bilmiyorsanız veya bir programın ömrü boyunca boyutları değiştirebilecek bir
*              diziye ihtiyacınız varsa, Vector çok yararlı olduğunu kanıtlar.
* The Stack : Yığın, standart bir son giren ilk çıkar yığını uygulayan bir Vector alt sınıfıdır. Yığın, yalnızca boş
*             bir yığın oluşturan varsayılan kurucuyu tanımlar. Yığın, Vector tarafından tanımlanan tüm yöntemleri
*             içerir ve kendi yöntemlerini de ekler.
* The Dictionary : Sözlük, bir anahtar/değer depolama havuzunu temsil eden ve Harita gibi çalışan soyut bir sınıftır.
*                  Bir anahtar ve değer verildiğinde, değeri bir Dictionary nesnesinde saklayabilirsiniz. Değer
*                  kaydedildikten sonra, anahtarını kullanarak onu geri alabilirsiniz. Bu nedenle, bir harita gibi,
*                  bir sözlük de anahtar/değer çiftlerinin bir listesi olarak düşünülebilir.
* The Hashtable : HashMap gibi Hashtable da anahtar/değer çiftlerini bir hash tablosunda saklar. Bir Hashtable
*                 kullanırken, anahtar olarak kullanılan bir nesneyi ve bu anahtara bağlanmasını istediğiniz değeri
*                 belirtirsiniz. Anahtar daha sonra özetlenir ve elde edilen karma kod, değerin tabloda depolandığı
*                 dizin olarak kullanılır.
* The Properties : Properties, Hashtable'ın bir alt sınıfıdır. Anahtarın bir Dize olduğu ve değerin de bir Dize olduğu
*                  değerlerin listesini tutmak için kullanılır.
! Collection
* The List : Liste arabirimi, Koleksiyon'u genişletir ve bir dizi öğeyi depolayan bir koleksiyonun davranışını bildirir.
*           Öğeler, sıfır tabanlı bir dizin kullanılarak listedeki konumlarına göre eklenebilir veya erişilebilir.
*           Bir liste yinelenen öğeler içerebilir.Koleksiyon tarafından tanımlanan yöntemlere ek olarak List, aşağıdaki
*           tabloda özetlenen kendi yöntemlerini tanımlar. Liste yöntemlerinin birçoğu, koleksiyon değiştirilemezse bir
*           UnsupportedOperationException oluşturur ve bir nesne diğeriyle uyumlu olmadığında bir ClassCastException
*           oluşturulur.
* The Set : Küme, yinelenen öğeler içeremeyen bir Koleksiyondur. Matematiksel küme soyutlamasını modeller. Set arabirimi
*           yalnızca Koleksiyondan devralınan yöntemleri içerir ve yinelenen öğelerin yasaklanması kısıtlamasını ekler.
*           Set ayrıca, equals ve hashCode işlemlerinin davranışına ilişkin daha güçlü bir sözleşme ekleyerek, uygulama
*           türleri farklı olsa bile Set örneklerinin anlamlı bir şekilde karşılaştırılmasına olanak tanır.
* The SortedSet : SortedSet arabirimi, Küme'yi genişletir ve artan düzende sıralanmış bir kümenin davranışını bildirir.
*                 Çağıran kümede hiçbir öğe bulunmadığında birkaç yöntem bir NoSuchElementException oluşturur. Bir nesne
*                 bir kümedeki öğelerle uyumlu olmadığında bir ClassCastException oluşturulur. NullPointerException, boş
*                 bir nesne kullanmaya çalışılırsa ve sette null öğesine izin verilmezse atılır.
* The Map : Harita arayüzü, benzersiz anahtarları değerlere eşler. Anahtar, daha sonraki bir tarihte bir değer almak
*           için kullandığınız bir nesnedir. Bir anahtar ve bir değer verildiğinde, değeri bir Map nesnesinde
*           saklayabilirsiniz. Değer kaydedildikten sonra, anahtarını kullanarak geri alabilirsiniz. Çağıran haritada
*           hiçbir öğe olmadığında birkaç yöntem NoSuchElementException oluşturur. Bir nesne bir haritadaki öğelerle
*           uyumlu olmadığında bir ClassCastException oluşturulur. NullPointerException, boş bir nesne kullanılmaya
*           çalışılırsa ve haritada null öğesine izin verilmezse atılır. Değiştirilemeyen bir haritayı değiştirme
*           girişiminde bulunulduğunda bir UnsupportedOperationException oluşturulur.
* The SortedMap : SortedMap arayüzü Harita'yı genişletir. Girişlerin artan bir anahtar sırayla tutulmasını sağlar.
*                 Çağıran haritada hiçbir öğe olmadığında birkaç yöntem bir NoSuchElementException oluşturur. Bir nesne
*                 bir haritadaki öğelerle uyumlu olmadığında bir ClassCastException oluşturulur. Haritada null öğesine
*                 izin verilmediğinde boş bir nesne kullanılmaya çalışılırsa bir NullPointerException oluşturulur.
* The LinkedList : LinkedList sınıfı, AbstractSequentialList'i genişletir ve List arabirimini uygular. Bağlantılı liste
*                  veri yapısı sağlar.
* The ArrayList : ArrayList sınıfı, AbstractList'i genişletir ve List arabirimini uygular. ArrayList, gerektiğinde
*                 büyüyebilen dinamik dizileri destekler. Standart Java dizileri sabit uzunluktadır. Diziler
*                 oluşturulduktan sonra büyüyemez veya küçülemezler, bu da bir dizinin kaç eleman tutacağını önceden
*                 bilmeniz gerektiği anlamına gelir. Dizi listeleri bir başlangıç boyutuyla oluşturulur. Bu boyut
*                 aşıldığında koleksiyon otomatik olarak büyütülür. Nesneler kaldırıldığında dizi küçülebilir.
* The HashSet : HashSet, AbstractSet'i genişletir ve Set arabirimini uygular. Depolama için bir karma tablo kullanan bir
*               koleksiyon oluşturur. Bir karma tablosu, karma adı verilen bir mekanizma kullanarak bilgileri depolar.
*               Karma işleminde, bir anahtarın bilgi içeriği, karma kodu adı verilen benzersiz bir değeri belirlemek
*               için kullanılır. Karma kod daha sonra anahtarla ilişkili verilerin depolandığı dizin olarak kullanılır.
*               Anahtarın hash koduna dönüştürülmesi otomatik olarak gerçekleştirilir.
* The LinkedHashSet : Bu sınıf HashSet'i genişletir, ancak kendi üyesini eklemez. LinkedHashSet, kümedeki girişlerin
*                     eklendikleri sırayla bağlantılı bir listesini tutar. Bu, küme üzerinde ekleme siparişi
*                     yinelemesine izin verir. Diğer bir deyişle, bir yineleyici kullanarak bir LinkedHashSet üzerinden
*                     geçiş yaparken, öğeler eklendikleri sırayla döndürülür. Karma kod daha sonra anahtarla ilişkili
*                     verilerin depolandığı dizin olarak kullanılır. Anahtarın hash koduna dönüştürülmesi otomatik
*                     olarak gerçekleştirilir.
* The TreeSet : TreeSet, depolama için bir ağaç kullanan Set arabiriminin bir uygulamasını sağlar. Nesneler sıralı ve
*               artan bir düzende saklanır. Erişim ve geri alma süreleri oldukça hızlıdır, bu da hızla bulunması gereken
*               büyük miktarda sıralanmış bilgiyi depolarken TreeSet'i mükemmel bir seçim haline getirir.
* The TreeMap : TreeMap sınıfı, bir ağaç kullanarak Harita arabirimini uygular. Bir TreeMap, anahtar/değer çiftlerini
*               sıralı bir şekilde depolamak için verimli bir yol sağlar ve hızlı bir şekilde alınmasını sağlar. Bir
*               karma haritadan farklı olarak, bir ağaç haritasının, öğelerinin artan anahtar düzeninde sıralanmasını
*               garanti ettiğini unutmayın.
* The WeakHashMap : WeakHashMap, anahtarlarına yalnızca zayıf referansları depolayan Harita arabiriminin bir
*                   uygulamasıdır. Yalnızca zayıf referansları depolamak, bir anahtar/değer çiftinin anahtarına artık
*                   WeakHashMap dışında referans verilmediğinde çöp olarak toplanmasına izin verir. Bu sınıf, zayıf
*                   referansların gücünden yararlanmanın en kolay yolunu sağlar. Bir girdinin faydasının, anahtarına
*                   artık herhangi bir iş parçacığı tarafından erişilemediğinde ortadan kalktığı "kayıt defteri benzeri"
*                   veri yapılarını uygulamak için kullanışlıdır. WeakHashMap, çok önemli bir istisna dışında HashMap
*                   ile aynı şekilde çalışır: Java bellek yöneticisinin artık anahtar olarak belirtilen nesneye güçlü
*                   bir referansı yoksa, haritadaki giriş kaldırılacaktır. Zayıf Referans - Bir nesneye yapılan tek
*                   referans zayıf referanslarsa, çöp toplayıcı herhangi bir zamanda nesnenin hafızasını geri alabilir.
*                   Sistemin hafızasının dolmasını beklemek zorunda değildir. Genellikle, çöp toplayıcı bir sonraki
*                   çalıştığında serbest bırakılır.
* LinkedHashMap : Bu sınıf, HashMap'i genişletir ve haritadaki girdilerin eklendikleri sırayla bağlantılı bir listesini
*                 tutar. Bu, harita üzerinde ekleme siparişi yinelemeye izin verir. Diğer bir deyişle, bir LinkedHashMap
*                 yinelenirken, öğeler eklendikleri sırayla döndürülür. Ayrıca, öğelerini en son erişildikleri sırayla
*                 döndüren bir LinkedHashMap oluşturabilirsiniz.
* The IdentityHashMap : Bu sınıf, AbstractMap'i uygular. Öğeleri karşılaştırırken referans eşitliği kullanması dışında
*                       HashMap'e benzer. Bu sınıf, genel amaçlı bir Harita uygulaması değildir. Bu sınıf, Harita
*                       arabirimini uygularken, nesneleri karşılaştırırken eşittir yönteminin kullanılmasını zorunlu
*                       kılan Harita'nın genel sözleşmesini kasıtlı olarak ihlal eder.Bu sınıf, yalnızca referans
*                       eşitliği semantiğinin gerekli olduğu nadir durumlarda kullanılmak üzere tasarlanmıştır.
*                       Bu sınıf, sistem kimliği karma işlevinin (System.identityHashCode(Object)) öğeleri paketler
*                       arasında düzgün bir şekilde dağıttığını varsayarak, temel işlemler (al ve koy) için sabit
*                       zamanlı performans sağlar. Bu sınıfın bir ayar parametresi vardır (performansı etkiler ancak
*                       anlambilimi etkilemez): beklenen maksimum boyut. Bu parametre, eşlemenin tutması beklenen
*                       maksimum anahtar/değer eşleme sayısıdır.
* The Iterator : Çoğu zaman, bir koleksiyondaki öğeler arasında geçiş yapmak isteyeceksiniz. Örneğin, her bir öğeyi
*                görüntülemek isteyebilirsiniz. Bunu yapmanın en kolay yolu, Iterator veya ListIterator arabirimini
*                uygulayan bir nesne olan bir yineleyici kullanmaktır. Yineleyici, öğeleri alarak veya kaldırarak bir
*                koleksiyon arasında geçiş yapmanızı sağlar. ListIterator, yineleyiciyi bir listenin çift yönlü geçişine
*                ve öğelerin değiştirilmesine izin verecek şekilde genişletir. Bir yineleyici aracılığıyla bir
*                koleksiyona erişmeden önce bir tane edinmelisiniz. Koleksiyon sınıflarının her biri, koleksiyonun
*                başına bir yineleyici döndüren bir yineleyici() yöntemi sağlar. Bu yineleyici nesneyi kullanarak,
*                koleksiyondaki her öğeye her seferinde bir öğe olmak üzere erişebilirsiniz. Genel olarak, bir
*                koleksiyonun içeriği arasında geçiş yapmak üzere bir yineleyici kullanmak için şu adımları izleyin:
*                Koleksiyonun yineleyici() yöntemini çağırarak koleksiyonun başlangıcına bir yineleyici edinin.
*                hasNext()'e çağrı yapan bir döngü kurun. hasNext() true döndürdüğü sürece döngünün yinelenmesini
*                sağlayın. Döngü içinde, next() öğesini çağırarak her öğeyi elde edin. List uygulayan koleksiyonlar için
*                ListIterator'ı çağırarak da bir yineleyici elde edebilirsiniz.
* The Comparator : Hem TreeSet hem de TreeMap, öğeleri sıralı düzende saklar. Ancak, sıralı düzenin ne anlama geldiğini
*                  tam olarak tanımlayan karşılaştırıcıdır.
*/
public class Main {

    public static void main(String[] args) {
        // ! Data Structures
        EnumerationTester enumerationTester = new EnumerationTester();
        enumerationTester.printAll();

        BitSetDemo bitSetDemo = new BitSetDemo();
        bitSetDemo.printAll();


        VectorDemo vectorDemo = new VectorDemo();
        vectorDemo.printAll();

        StackDemo stackDemo = new StackDemo();
        stackDemo.printAll();

        HashTableDemo hashTableDemo = new HashTableDemo();
        hashTableDemo.printAll();

        PropDemo propDemo = new PropDemo();
        propDemo.printAll();

        CollectionsDemo collectionsDemo = new CollectionsDemo();
        collectionsDemo.printAll();

        GenericMethodTest genericMethodTest = new GenericMethodTest();
        genericMethodTest.printAll();
    }
}
