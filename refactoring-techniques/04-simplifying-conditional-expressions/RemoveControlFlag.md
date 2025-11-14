# Remove Control Flag

- **Sorun:** Birden çok boole ifadesi için kontrol bayrağı görevi gören bir boole değişkeniniz var.

- **Çözüm:** Değişken yerine , breakve continuekullanın return.

## Neden Refactor

Kontrol bayrakları, "uygun" programcıların fonksiyonları için her zaman bir giriş noktasına (fonksiyon bildirim satırı) ve bir çıkış noktasına (fonksiyonun en sonunda) sahip olduğu eski günlere kadar uzanır.

Döngülerdeki ve diğer karmaşık yapılardaki kontrol akışını değiştirmek için özel operatörlerimiz olduğundan, modern programlama dillerinde bu tarz tik artık kullanılmamaktadır:

- break: döngüyü durdurur

- continue: geçerli döngü dalının yürütülmesini durdurur ve sonraki yinelemede döngü koşullarını kontrol etmeye gider

- return: tüm işlevin yürütülmesini durdurur ve operatörde verilmişse sonucunu döndürür

## Faydalar

Kontrol bayrağı kodu genellikle kontrol akışı operatörleriyle yazılan koddan çok daha ağırdır.

## Yeniden Düzenleme Nasıl Yapılır?

1. Döngüden veya geçerli yinelemeden çıkışa neden olan kontrol bayrağına değer atamasını bulun.

2. break Bu bir döngüden çıkış ise ile değiştirin ; continue, bu bir yinelemeden çıkış ise veya returnbu değeri işlevden döndürmeniz gerekiyorsa .

3. Kalan kodu ve kontrol bayrağıyla ilişkili kontrolleri kaldırın.
