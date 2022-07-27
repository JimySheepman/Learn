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

        CollectionsDemo collectionsDemo = new CollectionsDemo();
        collectionsDemo.printAll();

        HashTableDemo hashTableDemo = new HashTableDemo();
        hashTableDemo.printAll();

        PropDemo propDemo = new PropDemo();
        propDemo.printAll();
    }
}
