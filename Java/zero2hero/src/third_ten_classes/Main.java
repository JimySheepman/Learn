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
* The SortedMap :

The SortedMap interface extends Map. It ensures that the entries are maintained in an ascending key order.

Several methods throw a NoSuchElementException when no items are in the invoking map. A ClassCastException is thrown when an object is incompatible with the elements in a map. A NullPointerException is thrown if an attempt is made to use a null object when null is not allowed in the map.
* The LinkedList :
The LinkedList class extends AbstractSequentialList and implements the List interface. It provides a linked-list data structure.
* The ArrayList :
The ArrayList class extends AbstractList and implements the List interface. ArrayList supports dynamic arrays that can grow as needed.

Standard Java arrays are of a fixed length. After arrays are created, they cannot grow or shrink, which means that you must know in advance how many elements an array will hold.

Array lists are created with an initial size. When this size is exceeded, the collection is automatically enlarged. When objects are removed, the array may be shrunk.
* The HashSet :
HashSet extends AbstractSet and implements the Set interface. It creates a collection that uses a hash table for storage.

A hash table stores information by using a mechanism called hashing. In hashing, the informational content of a key is used to determine a unique value, called its hash code.

The hash code is then used as the index at which the data associated with the key is stored. The transformation of the key into its hash code is performed automatically.
* The LinkedHashSet :
This class extends HashSet, but adds no members of its own.

LinkedHashSet maintains a linked list of the entries in the set, in the order in which they were inserted. This allows insertion-order iteration over the set.

That is, when cycling through a LinkedHashSet using an iterator, the elements will be returned in the order in which they were inserted.

The hash code is then used as the index at which the data associated with the key is stored. The transformation of the key into its hash code is performed automatically.
* The TreeSet :
TreeSet provides an implementation of the Set interface that uses a tree for storage. Objects are stored in a sorted and ascending order.

Access and retrieval times are quite fast, which makes TreeSet an excellent choice when storing large amounts of sorted information that must be found quickly.
* The TreeMap :
The TreeMap class implements the Map interface by using a tree. A TreeMap provides an efficient means of storing key/value pairs in sorted order, and allows rapid retrieval.

You should note that, unlike a hash map, a tree map guarantees that its elements will be sorted in an ascending key order.
* The WeakHashMap :
WeakHashMap is an implementation of the Map interface that stores only weak references to its keys. Storing only weak references allows a key-value pair to be garbage-collected when its key is no longer referenced outside of the WeakHashMap.

This class provides the easiest way to harness the power of weak references. It is useful for implementing "registry-like" data structures, where the utility of an entry vanishes when its key is no longer reachable by any thread.

The WeakHashMap functions identically to the HashMap with one very important exception: if the Java memory manager no longer has a strong reference to the object specified as a key, then the entry in the map will be removed.

Weak Reference − If the only references to an object are weak references, the garbage collector can reclaim the object's memory at any time.it doesn't have to wait until the system runs out of memory. Usually, it will be freed the next time the garbage collector runs.
* LinkedHashMap :
This class extends HashMap and maintains a linked list of the entries in the map, in the order in which they were inserted. This allows insertion-order iteration over the map. That is, when iterating a LinkedHashMap, the elements will be returned in the order in which they were inserted.

You can also create a LinkedHashMap that returns its elements in the order in which they were last accessed.
* The IdentityHashMap :
This class implements AbstractMap. It is similar to HashMap except that it uses reference equality when comparing the elements.

This class is not a general-purpose Map implementation. While this class implements the Map interface, it intentionally violates Map's general contract, which mandates the use of the equals method when comparing objects.

This class is designed for use only in rare cases wherein reference-equality semantics are required. This class provides constant-time performance for the basic operations (get and put), assuming the system identity hash function (System.identityHashCode(Object)) disperses elements properly among the buckets.

This class has one tuning parameter (which affects performance but not semantics): expected maximum size. This parameter is the maximum number of key-value mappings that the map is expected to hold.
* The Iterator :
Often, you will want to cycle through the elements in a collection. For example, you might want to display each element. The easiest way to do this is to employ an iterator, which is an object that implements either the Iterator or the ListIterator interface.

Iterator enables you to cycle through a collection, obtaining or removing elements. ListIterator extends Iterator to allow bidirectional traversal of a list, and the modification of elements.

Before you can access a collection through an iterator, you must obtain one. Each of the collection classes provides an iterator( ) method that returns an iterator to the start of the collection. By using this iterator object, you can access each element in the collection, one element at a time.

In general, to use an iterator to cycle through the contents of a collection, follow these steps −

    Obtain an iterator to the start of the collection by calling the collection's iterator( ) method.

    Set up a loop that makes a call to hasNext( ). Have the loop iterate as long as hasNext( ) returns true.

    Within the loop, obtain each element by calling next( ).

For collections that implement List, you can also obtain an iterator by calling ListIterator.
* The Comparator :
Both TreeSet and TreeMap store elements in sorted order. However, it is the comparator that defines precisely what sorted order means.
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
    }
}
