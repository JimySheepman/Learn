# Pull Up Method

- **Sorun:** Alt sınıflarınız benzer işi yapan yöntemlere sahiptir.

- **Çözüm:** Yöntemleri aynı yapın ve ardından bunları ilgili üst sınıfa taşıyın.

## Neden Refactor

Alt sınıflar birbirinden bağımsız olarak büyüdü ve gelişti, aynı (veya neredeyse aynı) alanlara ve yöntemlere neden oldu.

## Faydalar

- Yinelenen kodlardan kurtulur. Bir yöntemde değişiklik yapmanız gerekiyorsa, yöntemin tüm kopyalarını alt sınıflarda aramak zorunda kalmaktansa, bunu tek bir yerde yapmak daha iyidir.

- Bu yeniden düzenleme tekniği, herhangi bir nedenle, bir alt sınıf bir üst sınıf yöntemini yeniden tanımlıyorsa, ancak temelde aynı işi yapıyorsa da kullanılabilir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Üst sınıflarda benzer yöntemleri araştırın. Aynı değillerse, birbirleriyle eşleşecek şekilde biçimlendirin.

2. Yöntemler farklı bir parametre kümesi kullanıyorsa, parametreleri üst sınıfta görmek istediğiniz forma koyun.

3. Yöntemi üst sınıfa kopyalayın. Burada, yöntem kodunun yalnızca alt sınıflarda bulunan ve bu nedenle üst sınıfta bulunmayan alanları ve yöntemleri kullandığını görebilirsiniz. Bunu çözmek için şunları yapabilirsiniz:

    - Alanlar için: Alt sınıflarda alıcılar ve ayarlayıcılar oluşturmak için   Pull Up Field veya Self-Encapsulate Field kullanın; daha sonra bu alıcıları üst sınıfta soyut olarak ilan edin.

    - Yöntemler için: Ya  Pull Up Method kullanın ya da üst sınıfta onlar için soyut yöntemler bildirin (sınıfınızın daha önce soyut hale geleceğini unutmayın).

4. Yöntemleri alt sınıflardan kaldırın.

5. Yöntemin çağrıldığı konumları kontrol edin. Bazı yerlerde, bir alt sınıfın kullanımını üst sınıfla değiştirebilirsiniz.
