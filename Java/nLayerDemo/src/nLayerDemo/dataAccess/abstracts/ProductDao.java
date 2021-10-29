package nLayerDemo.dataAccess.abstracts;

import nLayerDemo.entities.concretes.Product;

import java.util.List;

public interface ProductDao {
    void add(Product product);

    void update(Product product);

    void delete(Product product);

    Product get(int id); // 1 ürün getirmsi için Product döndürüyor

    // araya sonradan veri eklemek yeniden referans almasını sağlar.
    // javada bu duruma çözüm olarak ArrayList(dinamik bir yapıdır.) yapısı geliyor
    // ArrayList<Integer> arrayList = new ArrayList<Integer>(); -> syntax ı. Array in base List yapısıdır.
    // List bir interface dir.
    List<Product> getAll(); // listeleme
}
