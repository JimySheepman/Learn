# Encapsulate Collection

- **Sorun:** Bir sınıf, bir koleksiyon alanı ve koleksiyonla çalışmak için basit bir alıcı ve ayarlayıcı içerir.

- **Çözüm:** Alıcı tarafından döndürülen değeri salt okunur yapın ve koleksiyonun öğelerini eklemek/silmek için yöntemler oluşturun.

## Neden Refactor

Bir sınıf, bir nesne koleksiyonu içeren bir alan içerir. Bu koleksiyon bir dizi, liste, küme veya vektör olabilir. Koleksiyonla çalışmak için normal bir alıcı ve ayarlayıcı oluşturuldu.

Ancak koleksiyonlar, diğer veri türleri tarafından kullanılandan biraz farklı bir protokol tarafından kullanılmalıdır. Alıcı yöntemi, toplama nesnesinin kendisini döndürmemelidir, çünkü bu, istemcilerin sahip sınıfının bilgisi olmadan koleksiyon içeriğini değiştirmesine izin verir. Ayrıca, bu, istemcilere nesne verilerinin iç yapılarının çok fazlasını gösterir. Koleksiyon öğelerini alma yöntemi, koleksiyonun değiştirilmesine izin vermeyen veya yapısı hakkında aşırı veri ifşa etmeyen bir değer döndürmelidir.

Ayrıca koleksiyona değer atayan bir metot da olmamalıdır. Bunun yerine eleman ekleme ve silme işlemleri yapılmalıdır. Bu sayede sahip nesne, koleksiyon öğelerinin eklenmesi ve silinmesi üzerinde kontrol sahibi olur.

Böyle bir protokol, bir koleksiyonu uygun şekilde kapsüller, bu da sahip sınıf ile istemci kodu arasındaki ilişki derecesini nihai olarak azaltır.

## Faydalar

- Koleksiyon alanı bir sınıf içinde kapsüllenir. Alıcı çağrıldığında, koleksiyonu içeren sınıfın bilgisi olmadan koleksiyon öğelerinin yanlışlıkla değiştirilmesini veya üzerine yazılmasını önleyen koleksiyonun bir kopyasını döndürür.

- Koleksiyon öğeleri, dizi gibi ilkel bir tür içinde yer alıyorsa, koleksiyonla çalışmak için daha uygun yöntemler oluşturursunuz.

- Koleksiyon öğeleri, ilkel olmayan bir kapsayıcının (standart koleksiyon sınıfı) içinde yer alıyorsa, koleksiyonu kapsülleyerek, koleksiyonun istenmeyen standart yöntemlerine erişimi kısıtlayabilirsiniz (örneğin, yeni öğelerin eklenmesini kısıtlayarak).

## Yeniden Düzenleme Nasıl Yapılır?

1. Koleksiyon öğeleri eklemek ve silmek için yöntemler oluşturun. Parametrelerinde toplama öğelerini kabul etmelidirler.

2. Sınıf yapıcısında bu yapılmazsa, alana ilk değer olarak boş bir koleksiyon atayın.

3. Toplama alanı ayarlayıcısının çağrılarını bulun. Ayarlayıcıyı, öğe ekleme ve silme işlemlerini kullanması için değiştirin veya bu işlemlerin istemci kodunu çağırmasını sağlayın. Ayarlayıcıların yalnızca tüm koleksiyon öğelerini diğerleriyle değiştirmek için kullanılabileceğini unutmayın. Bu nedenle ayarlayıcı adının ( Rename Method) olarak değiştirilmesi tavsiye edilebilir replace.

4. Koleksiyon değiştirildikten sonra koleksiyon alıcısının tüm çağrılarını bulun. Kodu, koleksiyondan öğe eklemek ve silmek için yeni yöntemlerinizi kullanacak şekilde değiştirin.

5. Alıcıyı, koleksiyonun salt okunur bir temsilini döndürecek şekilde değiştirin.

6. Koleksiyon sınıfının içinde daha iyi görünecek kod için koleksiyonu kullanan istemci kodunu inceleyin.
