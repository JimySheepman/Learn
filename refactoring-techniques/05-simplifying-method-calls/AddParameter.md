# Add Parameter

- **Sorun:** Bir yöntem, belirli eylemleri gerçekleştirmek için yeterli veriye sahip değil.

- **Çözüm:** Gerekli verileri iletmek için yeni bir parametre oluşturun.

## Neden Refactor

Bir yöntemde değişiklik yapmanız gerekir ve bu değişiklikler, daha önce yöntem için mevcut olmayan bilgi veya verilerin eklenmesini gerektirir.

## Faydalar

Buradaki seçim, yeni bir parametre eklemek ile yöntemin ihtiyaç duyduğu verileri içeren yeni bir özel alan eklemek arasındadır. Her zaman bir nesnede tutmanın bir anlamı olmayan ara sıra veya sık sık değişen verilere ihtiyacınız olduğunda bir parametre tercih edilir. Bu durumda, yeniden düzenleme işe yarayacaktır. Aksi takdirde, yöntemi çağırmadan önce özel bir alan ekleyin ve gerekli verilerle doldurun.

## Dezavantajları

- Yeni bir parametre eklemek, onu kaldırmaktan her zaman daha kolaydır, bu nedenle parametre listeleri sıklıkla balondan grotesk boyutlara kadardır. Bu koku,  Long Parameter List olarak bilinir .

- Yeni bir parametre eklemeniz gerekiyorsa, bu bazen sınıfınızın gerekli verileri içermediği veya mevcut parametrelerin gerekli ilgili verileri içermediği anlamına gelir. Her iki durumda da en iyi çözüm, verileri ana sınıfa veya nesnelerine zaten yöntemin içinden erişilebilen diğer sınıflara taşımayı düşünmektir.

## Yeniden Düzenleme Nasıl Yapılır?

1. Yöntemin bir üst sınıfta mı yoksa alt sınıfta mı tanımlandığını görün. Eğer metot onlarda mevcutsa, bu sınıflarda da tüm adımları tekrarlamanız gerekecektir.

2. Aşağıdaki adım, yeniden düzenleme işlemi sırasında programınızı çalışır durumda tutmak için çok önemlidir. Eskisini kopyalayarak yeni bir metot oluşturun ve buna gerekli parametreyi ekleyin. Eski yöntemin kodunu, yeni yöntemin çağrısıyla değiştirin. Yeni parametreye herhangi bir değer ekleyebilirsiniz ( nullnesneler için veya sayılar için sıfır gibi).

3. Eski yönteme yapılan tüm referansları bulun ve bunları yeni yönteme yapılan referanslarla değiştirin.

4. Eski yöntemi silin. Eski yöntem genel arayüzün bir parçasıysa silme mümkün değildir. Bu durumda, eski yöntemi kullanımdan kaldırıldı olarak işaretleyin.
