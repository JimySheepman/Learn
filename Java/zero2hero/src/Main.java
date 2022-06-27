public class Main {
/*
! DERS-1
* Object -> objelerin state ve behavior'ları vardır. Bir class'ın öğesidir.
* Class -> türünün objesini destekleğen ve state/behavior tanımlanan
*           şablonlardır.
* Methods -> methodlar temelde behavior'dır.
* Instace Variables -> objelerin benzersiz variable'lar kümesi vardır.
*                       objenin state'i bu variable'lar atanan değerler
*                       ile oluşur.
* Case Sensitivity bir dildir.
* Genellikle community camel case kullanıyor. Class'ların ilk harfi
* her zaman büyük.
* Method Names küçük harf ile başlamalıdır.
* Program File Name, Program dosyasının adı, sınıf adıyla tam olarak eşleşmelidir.
* main ile başlamak zorundadır.
* variable name -> {[a-z],[A-Z],_,$} ile başlamalıdır.
* Modifiers;
*           Access Modifiers − default, public , protected, private
*           Non-access Modifiers − final, abstract, strictfp
* Variables;
*           Local Variables
*           Class Variables (Static Variables)
*           Instance Variables (Non-static Variables)
*
* Read All Keyword Link:
* https://docs.oracle.com/cd/E13226_01/workshop/docs81/pdf/files/workshop/JavaKeywordReference.pdf
! DERS-2
* java nesne yönelimli bir dildir;
* Polymorphism, Inheritance, Abstraction, Encapsulation, Classes, Objects, Instance,
* Message Passing, Method, sahiptir.
* Local variables -> Yöntemler, yapıcılar veya bloklar içinde tanımlanan değişkenlere yerel değişkenler denir.
* Instance variables -> bir sınıf içindeki ancak herhangi bir yöntemin dışındaki değişkenlerdir.
* Class variables -> static anahtar sözcüğüyle herhangi bir yöntemin dışında bir sınıf içinde bildirilen değişkenlerdir.
* Constructors -> Her sınıfın bir constructor vardır. Bir sınıf için açıkça bir constructor yazmazsak, Java derleyicisi
*                 o sınıf için varsayılan bir constructor oluşturur. Class lar ile aynı adı taşımalıdır.
*                  Constructors'lar paramete alarak yada almayarak oluşturulabilinir.
* Declaration ->  Bir nesne tipine sahip bir değişken adına sahip bir değişken bildirimi.
* Instantiation -> Nesneyi oluşturmak için 'new' anahtar sözcüğü kullanılır.
* Initialization -> 'new' anahtar sözcüğünü bir yapıcıya yapılan çağrı izler. Bu çağrı, yeni nesneyi başlatır.
! DERS-3
* Java'da kullanılabilen iki veri türü vardır -
* Primitive Data Types :Java tarafından desteklenen sekiz ilkel veri türü vardır. İlkel veri türleri, dil tarafından
*                       önceden tanımlanır ve bir anahtar sözcükle adlandırılır. Şimdi sekiz ilkel veri türünü
*                       ayrıntılı olarak inceleyelim.

* Reference/Object Data Types: Referans değişkenleri, sınıfların tanımlı kurucuları kullanılarak oluşturulur.
*                              Nesnelere erişmek için kullanılırlar. Bu değişkenler, değiştirilemeyen belirli
*                              bir tür olarak bildirilir. Örneğin, Çalışan, Köpek Yavrusu vb. Sınıf nesneleri
*                              ve çeşitli dizi değişkenleri, referans veri tipine girer. Herhangi bir başvuru
*                              değişkeninin varsayılan değeri boştur. Bir başvuru değişkeni, bildirilen türdeki
*                              herhangi bir nesneye veya herhangi bir uyumlu türe başvurmak için kullanılabilir.
*                              Örnek: Animal Animal = new Animal("zürafa");
! DERS-4
* Modifier Types ikiye ayrılır;
* Access Control Modifiers ->
*                           public -> her yerden erşilebilir.
*                           private -> Yalnızca Class tarafından görülebilir
*                           protected -> Pakete ve tüm alt sınıflara görünür.
* Non-Access Modifiers ->
*                       static -> Sınıf yöntemleri ve değişkenleri oluşturmak için  kullanılır.
*                       final -> implementations sonlandırmak için sınıf, method ve variable'a kullanılır.
*                       abstract -> Soyut sınıflar ve yöntemler oluşturmak için kullanılır.
*                       synchronized & volatile -> threads için kullanılır.
! DERS-5
* Numbers Class -> Geliştirme aşamasında primitive data type yerine nesneleri kullanmamız gereken
*                  durumlarla karşılaşıyoruz. Bunu başarmak için Java, wrapper classes sağlar.
*                  primitive data type nesneye dönüştürmeye boxing denir ve bu, derleyici tarafından halledilir.
! DERS-9
* Regular Expressions -> 3 class'tan oluşur.
*                        Pattern Class -> normal bir ifadenin derlenmiş bir temsilidir.
*                        Matcher Class -> kalıbı yorumlayan ve bir girdi dizesine karşı eşleştirme işlemleri
*                                         gerçekleştiren motordur.
*                        PatternSyntaxException -> normal ifade modelinde bir sözdizimi hatasını gösteren denetlenmeyen
*                                                  bir istisnadır.
* Regular Expression Syntax ->
*                             ^	Satırın başlangıcıyla eşleşir.
*                             $ 	Satırın sonuyla eşleşir.
*                             . 	Yeni satır dışında herhangi bir tek karakterle eşleşir. M seçeneğini kullanmak ,
*                                   yeni satırla da eşleşmesini sağlar.
*                             [...] 	Parantez içindeki herhangi bir karakterle eşleşir.
*                             [^...] 	Parantez içinde olmayan herhangi bir karakterle eşleşir.
*                             \A 	Tüm dizenin başlangıcı.
*                             \z 	Tüm dizenin sonu.
*                             \Z 	İzin verilen son satır sonlandırıcı dışında tüm dizenin sonu.
*                             re* 	Önceki ifadenin 0 veya daha fazla tekrarı ile eşleşir.
*                             re 	Önceki şeyden 1 veya daha fazlasıyla eşleşir.
*                             re? 	Önceki ifadenin 0 veya 1 tekrarıyla eşleşir.
*                             re { n} 	Önceki ifadenin tam olarak n tekrarı ile eşleşir.
*                             re { n,} 	Önceki ifadenin n veya daha fazla tekrarı ile eşleşir.
*                             re { n, m} 	Önceki ifadenin en az n ve en fazla m tekrarı ile eşleşir.
*                             a | b 	a veya b ile eşleşir.
*                             (re) 	Normal ifadeleri gruplandırır ve eşleşen metni hatırlar.
*                             (?: re) 	Eşleşen metni hatırlamadan normal ifadeleri gruplandırır.
*                             (?> re) 	Geri izleme olmadan bağımsız desenle eşleşir.
*                             \w 	Kelime karakterleriyle eşleşir.
*                             \W 	Sözcük olmayan karakterlerle eşleşir.
*                             \s 	Boşlukla eşleşir. [\t\n\r\f] ile eşdeğerdir.
*                             \S 	Beyaz olmayan boşlukla eşleşir.
*                             \d 	Rakamlarla eşleşir. [0-9] ile eşdeğerdir.
*                             \D 	Rakam olmayanlarla eşleşir.
*                             \A 	Dizenin başlangıcıyla eşleşir.
*                             \Z 	Dizenin sonuyla eşleşir. Yeni satır varsa, satırsonu satırından hemen önce eşleşir.
*                             \z 	Dizenin sonuyla eşleşir.
*                             \G 	Son eşleşmenin bittiği noktayla eşleşir.
*                             \n 	"n" grup numarasını yakalamak için geri referans.
*                             \b 	Köşeli parantezlerin dışındayken sözcük sınırlarıyla eşleşir. Köşeli ayraçların
*                                   içindeyken geri al (0x08) ile eşleşir.
*                             \B 	Sözcük olmayan sınırlarla eşleşir.
*                             \n, \t, vb. 	Yeni satırlar, satır başları, sekmeler vb. ile eşleşir.
*                             \Q 	\E'ye kadar olan tüm karakterlerden kaçın (alıntı yapın).
*                             \E 	\Q ile başlayan alıntıyı bitirir.
*/

    public static void main(String []args){
        // ! DERS-1
        FreshJuice juice = new FreshJuice();
        juice.size = FreshJuice.FreshJuiceSize.MEDIUM;
        System.out.println("Size: "+juice.size);
        System.out.println("Hello, World!");
        // ! DERS-2
        Dog dog = new Dog();
        dog.age=12;
        dog.breed="pug";
        dog.color="white";
        dog.info();
        dog.barking();

        Puppy puppy = new Puppy("tommy");
        puppy.setPuppyAge(2);
        puppy.getPuppyAge();
        System.out.println("variable value:"+puppy.puppyAge);
        //! DERS-3
        AllDataType allDataType = new AllDataType();
        allDataType.printType();

        EmployeeTwo empOne = new EmployeeTwo("Ransika");
        empOne.setSalary(1000);
        empOne.printEmp();
        // ! DERS-4
        Operators operators = new Operators();
        operators.printArithmeticOperators();
        operators.printRelationalOperators();
        operators.printBitwiseOperators();
        operators.printLogicalOperators();
        operators.printAssignmentOperators();
        operators.printMiscellaneousOperators();

        Loops loops = new Loops();
        loops.forLoop();
        loops.whileLoop();
        loops.doWhileLoop();
        loops.enhancedLoop();
        // ! DERS-5
        DecisionMaking decisionMaking = new DecisionMaking();
        decisionMaking.printIfStatement();
        decisionMaking.printIfIfElseElseStatement();
        decisionMaking.printNestedIfStatement();
        decisionMaking.printSwitchStatement();
        // ! DERS-6
        NumbersClass numbersClass = new NumbersClass();
        numbersClass.printAll();
        // ! DERS-7
        StringsClass stringsClass = new StringsClass();
        stringsClass.printAll();
        // ! DERS-8
        Array array = new Array();
        array.printArray();

        DateTime dateTime =new DateTime();
        dateTime.printDate();
        // ! DERS-9
        RegexMatches regexMatches = new RegexMatches();
        regexMatches.printAll();

    }
}
