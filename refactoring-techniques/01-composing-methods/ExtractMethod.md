# Extract Method

- **Sorun:** Birlikte gruplanabilecek bir kod parçanız var.

```Java
void printOwing() {
  printBanner();

  // Print details.
  System.out.println("name: " + name);
  System.out.println("amount: " + getOutstanding());
}
```

- **Çözüm:** Bu kodu ayrı bir yeni yönteme (veya işleve) taşıyın ve eski kodu yöntem çağrısıyla değiştirin.

```Java
void printOwing() {
  printBanner();
  printDetails(getOutstanding());
}

void printDetails(double outstanding) {
  System.out.println("name: " + name);
  System.out.println("amount: " + outstanding);
}
```

## Neden Refactor

Bir yöntemde ne kadar çok satır bulunursa, yöntemin ne yaptığını anlamak o kadar zor olur. Bu yeniden düzenlemenin ana nedeni budur.

Çıkarma yöntemleri, kodunuzdaki pürüzlü kenarları ortadan kaldırmanın yanı sıra, diğer birçok yeniden düzenleme yaklaşımında da bir adımdır.
Faydalar

  -  Daha okunabilir kod! Yeni yönteme, yöntemin amacını açıklayan bir ad verdiğinizden emin olun: createOrder(), renderCustomerInfo(), vb.

  - Daha az kod tekrarı. Genellikle bir yöntemde bulunan kod, programınızdaki başka yerlerde yeniden kullanılabilir. Böylece, kopyaları yeni yönteminize yapılan çağrılarla değiştirebilirsiniz.

  - Kodun bağımsız bölümlerini yalıtır, yani hataların daha az olası olduğu anlamına gelir (örneğin, yanlış değişken değiştirilirse).

## Yeniden Düzenleme Nasıl Yapılır?

1. Yeni bir yöntem oluşturun ve amacını açıkça ortaya koyacak şekilde adlandırın.

2. İlgili kod parçasını yeni yönteminize kopyalayın. Parçayı eski konumundan silin ve bunun yerine yeni yöntem için bir çağrı yapın.

3. Bu kod parçasında kullanılan tüm değişkenleri bulun. Parçanın içinde bildirilirlerse ve bunun dışında kullanılmazlarsa, onları değiştirmeden bırakın; yeni yöntem için yerel değişkenler olacaklar.

4. Değişkenler, ayıkladığınız koddan önce bildirilmişse, daha önce içerdikleri değerleri kullanmak için bu değişkenleri yeni yönteminizin parametrelerine iletmeniz gerekir. Bazen Temp'i Query ile Değiştir'e başvurarak bu değişkenlerden kurtulmak daha kolaydır .

5. Çıkarılan kodunuzda yerel bir değişkenin bir şekilde değiştiğini görürseniz, bu değiştirilen değere daha sonra ana yönteminizde ihtiyaç duyulacağı anlamına gelebilir. Tekrar kontrol edin! Ve eğer durum gerçekten buysa, her şeyi çalışır durumda tutmak için bu değişkenin değerini ana yönteme döndürün.

