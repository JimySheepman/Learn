# Push Down Field

- **Sorun:** Bir alan yalnızca birkaç alt sınıfta mı kullanılıyor?

- **Çözüm:** Alanı bu alt sınıflara taşıyın.

## Neden Refactor

Bir alanın tüm sınıflar için evrensel olarak kullanılması planlanmış olsa da, gerçekte alan yalnızca bazı alt sınıflarda kullanılmaktadır. Bu durum, örneğin, planlanmış özellikler dışa açılmadığında ortaya çıkabilir.

Bu, sınıf hiyerarşilerinin işlevselliğinin bir kısmının çıkarılması (veya kaldırılması) nedeniyle de ortaya çıkabilir.

## Faydalar

- Dahili sınıf tutarlılığını geliştirir. Bir alan, gerçekte kullanıldığı yerde bulunur.

- Aynı anda birden fazla alt sınıfa geçerken alanları birbirinden bağımsız olarak geliştirebilirsiniz. Bu, kod çoğaltması yaratır, evet, bu nedenle alanları yalnızca gerçekten farklı şekillerde kullanmayı düşündüğünüzde aşağı itin.

## Yeniden Düzenleme Nasıl Yapılır?

1. Gerekli tüm alt sınıflarda bir alan bildirin.

2. Alanı üst sınıftan kaldırın.
