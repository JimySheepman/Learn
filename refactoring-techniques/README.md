# Refactoring Techniques

## Composing Methods

Yeniden düzenlemenin çoğu, yöntemleri doğru bir şekilde oluşturmaya ayrılmıştır. Çoğu durumda, aşırı uzun yöntemler tüm kötülüklerin köküdür. Bu yöntemlerin içindeki değişken kodlar, yürütme mantığını gizler ve yöntemin anlaşılmasını ve hatta değiştirilmesini son derece zorlaştırır. Bu gruptaki yeniden düzenleme teknikleri, yöntemleri kolaylaştırır, kod tekrarını ortadan kaldırır ve gelecekteki iyileştirmelerin önünü açar.

* [Extract Method](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/ExtractMethod.md)
* [Inline Method](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/InlineMethod.md)
* [Extract Variable](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/ExtractVariable.md)
* [Inline Temp](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/InlineTemp.md)
* [Replace Temp with Query](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/ReplaceTempwithQuery.md)
* [Split Temporary Variable](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/SplitTemporaryVariable.md)
* [Remove Assignments to Parameters](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/RemoveAssignmentstoParameters.md)
* [Replace Method with Method Object](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/ReplaceMethodwithMethodObject.md)
* [Substitute Algorithm](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/01-composing-methods/SubstituteAlgorithm.md)

## Moving Features between Objects

İşlevselliği farklı sınıflar arasında mükemmel olmayan bir şekilde dağıtmış olsanız bile, hala umut var. Bu yeniden düzenleme teknikleri, işlevlerin sınıflar arasında nasıl güvenli bir şekilde taşınacağını, yeni sınıfların nasıl oluşturulacağını ve uygulama ayrıntılarının genel erişimden nasıl gizleneceğini gösterir.

* [Move Method](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/MoveMethod.md)
* [Move Field](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/MoveField.md)
* [Extract Class](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/ExtractClass.md)
* [Inline Class](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/InlineClass.md)
* [Hide Delegate](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/)
* [Remove Middle Man](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/HideDelegate.md)
* [Introduce Foreign Method](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/RemoveMiddleMan.md)
* [Introduce Local Extension](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/02-moving-features-between-objects/IntroduceLocalExtension.md)

## Organizing Data

Bu yeniden düzenleme teknikleri, ilkel öğeleri zengin sınıf işlevselliğiyle değiştirerek veri işlemeye yardımcı olur.Bir diğer önemli sonuç, sınıfları daha taşınabilir ve yeniden kullanılabilir hale getiren sınıf ilişkilendirmelerinin çözülmesidir.

* [Self Encapsulate Field](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/SelfEncapsulateField.md)
* [Replace Data Value with Object](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ReplaceDataValuewithObject.md)
* [Change Value to Reference](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ChangeValuetoReference.md)
* [Change Reference to Value](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ChangeReferencetoValue.md)
* [Replace Array with Object](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ReplaceArraywithObject.md)
* [Duplicate Observed Data](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/)
* [Change Unidirectional Association to Bidirectional](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ChangeUnidirectionalAssociationtoBidirectional.md)
* [Change Bidirectional Association to Unidirectional](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ChangeBidirectionalAssociationtoUnidirectional.md)
* [Replace Magic Number with Symbolic Constant](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ReplaceMagicNumberwithSymbolicConstant.md)
* [Encapsulate Field](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/EncapsulateField.md)
* [Encapsulate Collection](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/EncapsulateCollection.md)
* [Replace Type Code with Class](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ReplaceTypeCodewithClass.md)
* [Replace Type Code with Subclasses](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ReplaceTypeCodewithSubclasses.md)
* [Replace Type Code with State/Strategy](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ReplaceTypeCodewithStateStrategy.md)
* [Replace Subclass with Fields](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/03-organizing-data/ReplaceSubclasswithFields.md)

## Simplifying Conditional Expressions

Koşullar zaman içinde mantıklarında giderek daha karmaşık hale gelme eğilimindedir ve bununla mücadele etmek için daha fazla teknik vardır.

* [](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/04-simplifying-conditional-expressions/)

## Simplifying Method Calls

Bu teknikler, yöntem çağrılarını daha basit ve anlaşılmasını kolaylaştırır. Bu da, sınıflar arasındaki etkileşim için arayüzleri basitleştirir.

* [](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/05-simplifying-method-calls/)

## Dealing with Generalization

Soyutlamanın, öncelikle sınıf miras hiyerarşisi boyunca işlevsellik taşıma, yeni sınıflar ve arabirimler oluşturma ve kalıtımı yetkilendirme ile değiştirme ve bunun tersi ile ilişkili kendi yeniden düzenleme teknikleri grubu vardır.

* [](https://github.com/JimySheepman/Learn/blob/main/refactoring-techniques/06-dealing-with-generalization/)
