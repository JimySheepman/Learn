# Rename Method

- **Sorun:** Bir yöntemin adı, yöntemin ne yaptığını açıklamaz.

- **Çözüm:** Yöntemi yeniden adlandırın.

## Neden Refactor

Belki bir yöntemin adı en başından beri kötüydü - örneğin, biri yöntemi aceleyle yarattı ve onu iyi adlandırmaya özen göstermedi.

Veya belki de yöntem ilk başta iyi adlandırılmıştı ancak işlevselliği arttıkça yöntem adı iyi bir tanımlayıcı olmaktan çıktı.

## Faydalar

Kod okunabilirliği. Yeni yönteme ne yaptığını yansıtan bir ad vermeye çalışın. vb. createOrder() gibi bir şey .renderCustomerInfo()

## Yeniden Düzenleme Nasıl Yapılır?

1. Yöntemin bir üst sınıfta mı yoksa alt sınıfta mı tanımlandığını görün. Eğer öyleyse, bu sınıflarda da tüm adımları tekrarlamanız gerekir.

2. Sonraki yöntem, yeniden düzenleme işlemi sırasında programın işlevselliğini korumak için önemlidir. Yeni bir adla yeni bir yöntem oluşturun. Eski yöntemin kodunu ona kopyalayın. Eski yöntemdeki tüm kodu silin ve bunun yerine yeni yöntem için bir çağrı ekleyin.

3. Eski yönteme yapılan tüm referansları bulun ve bunları yenisine yapılan referanslarla değiştirin.

4. Eski yöntemi silin. Eski yöntem bir genel arabirimin parçasıysa, bu adımı gerçekleştirmeyin. Bunun yerine eski yöntemi kullanımdan kaldırıldı olarak işaretleyin.
