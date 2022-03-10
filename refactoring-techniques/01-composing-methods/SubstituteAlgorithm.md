# Substitute Algorithm

- **Sorun:** Yani mevcut bir algoritmayı yenisiyle değiştirmek mi istiyorsunuz?

```Java
String foundPerson(String[] people){
  for (int i = 0; i < people.length; i++) {
    if (people[i].equals("Don")){
      return "Don";
    }
    if (people[i].equals("John")){
      return "John";
    }
    if (people[i].equals("Kent")){
      return "Kent";
    }
  }
  return "";
}
```

- **Çözüm:** Algoritmayı uygulayan yöntemin gövdesini yeni bir algoritmayla değiştirin.

```Java
String foundPerson(String[] people){
  List candidates =
    Arrays.asList(new String[] {"Don", "John", "Kent"});
  for (int i=0; i < people.length; i++) {
    if (candidates.contains(people[i])) {
      return people[i];
    }
  }
  return "";
}
```

## Neden Refactor

1. Kademeli yeniden düzenleme, bir programı iyileştirmenin tek yöntemi değildir. Bazen bir yöntem sorunlarla o kadar karmaşıktır ki, yöntemi yıkmak ve yeni bir başlangıç ​​yapmak daha kolaydır. Ve belki de çok daha basit ve daha verimli bir algoritma buldunuz. Bu durumda, eski algoritmayı yenisiyle değiştirmeniz yeterlidir.

2. Zaman geçtikçe, algoritmanız iyi bilinen bir kitaplık veya çerçeveye dahil edilebilir ve bakımı basitleştirmek için bağımsız uygulamanızdan kurtulmak isteyebilirsiniz.

3. Programınızın gereksinimleri o kadar çok değişebilir ki, mevcut algoritmanız görev için kurtarılamaz.

## Yeniden Düzenleme Nasıl Yapılır?

1. Mevcut algoritmayı mümkün olduğunca basitleştirdiğinizden emin olun. Önemsiz kodu, Extract Method kullanarak diğer yöntemlere taşıyın . Algoritmanızda ne kadar az hareketli parça varsa, değiştirilmesi o kadar kolay olur.

2. Yeni algoritmanızı yeni bir yöntemle oluşturun. Eski algoritmayı yenisiyle değiştirin ve programı test etmeye başlayın.

3. Sonuçlar eşleşmezse eski uygulamaya dönün ve sonuçları karşılaştırın. Tutarsızlığın nedenlerini belirleyin. Nedeni genellikle eski algoritmadaki bir hata olsa da, yenisinde çalışmayan bir şey olması daha olasıdır.

4. Tüm testler başarıyla tamamlandığında, eski algoritmayı tamamen silin!
