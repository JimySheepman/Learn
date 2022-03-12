# Replace Subclass with Fields

- **Sorun:** Yalnızca (sürekli geri dönen) yöntemlerinde farklılık gösteren alt sınıflarınız var.

- **Çözüm:** Yöntemleri üst sınıftaki alanlarla değiştirin ve alt sınıfları silin.

## Neden Refactor

Bazen yeniden düzenleme, yalnızca tür kodundan kaçınma biletidir.

Böyle bir durumda, alt sınıfların hiyerarşisi yalnızca belirli yöntemlerle döndürülen değerlerde farklı olabilir. Bu yöntemler, hesaplamanın sonucu bile değildir, ancak yöntemlerin kendilerinde veya yöntemlerin döndürdüğü alanlarda kesin olarak belirtilir. Sınıf mimarisini basitleştirmek için, bu hiyerarşi duruma göre gerekli değerlere sahip bir veya birkaç alan içeren tek bir sınıfa sıkıştırılabilir.

Bu değişiklikler, büyük miktarda işlevselliği bir sınıf hiyerarşisinden başka bir yere taşıdıktan sonra gerekli hale gelebilir. Mevcut hiyerarşi artık o kadar değerli değil ve alt sınıfları artık sadece ölü ağırlıkta.

## Faydalar

Sistem mimarisini basitleştirir. Tüm yapmak istediğiniz farklı yöntemlerde farklı değerler döndürmekse, alt sınıflar oluşturmak aşırıya kaçar.

## Yeniden Düzenleme Nasıl Yapılır?

1. Replace Constructor with Factory Method uygulayın .

2. Alt sınıf oluşturucu çağrılarını üst sınıf fabrika yöntemi çağrılarıyla değiştirin.

3. Üst sınıfta, sabit değerler döndüren alt sınıf yöntemlerinin her birinin değerlerini depolamak için alanlar bildirin.

4. Yeni alanları başlatmak için korumalı bir üst sınıf oluşturucu oluşturun.

5. Varolan alt sınıf oluşturucularını, üst sınıfın yeni oluşturucusunu çağırmaları ve ilgili değerleri ona geçirmeleri için oluşturun veya değiştirin.

6. İlgili alanın değerini döndürmesi için her sabit yöntemi üst sınıfta uygulayın. Ardından yöntemi alt sınıftan kaldırın.

7. Alt sınıf oluşturucunun ek işlevleri varsa, yapıcıyı üst sınıf fabrika yöntemine dahil etmek için  Inline Method kullanın.

8. Alt sınıfı silin.
