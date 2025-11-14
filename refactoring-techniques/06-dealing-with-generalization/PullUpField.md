# Pull Up Field

- **Sorun:** İki sınıf aynı alana sahiptir.

- **Çözüm:** Alanı alt sınıflardan kaldırın ve üst sınıfa taşıyın.

## Neden Refactor

Alt sınıflar ayrı ayrı büyüdü ve gelişti, aynı (veya neredeyse aynı) alanların ve yöntemlerin ortaya çıkmasına neden oldu.

## Faydalar

- Alt sınıflardaki alanların tekrarını ortadan kaldırır.

- Varsa, yinelenen yöntemlerin alt sınıflardan bir üst sınıfa sonradan taşınmasını kolaylaştırır.

## Yeniden Düzenleme Nasıl Yapılır?

1. Alanların alt sınıflarda aynı ihtiyaçlar için kullanıldığından emin olun.

2. Alanların adları farklıysa, onlara aynı adı verin ve mevcut koddaki alanlara yapılan tüm başvuruları değiştirin.

3. Üst sınıfta aynı ada sahip bir alan oluşturun. Alanlar özelse, üst sınıf alanının korunması gerektiğini unutmayın.

4. Alanları alt sınıflardan kaldırın.

5. Erişim yöntemlerinin arkasına gizlemek için yeni alan için Self Encapsulate Field kullanmayı düşünebilirsiniz .
