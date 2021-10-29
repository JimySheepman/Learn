package nLayerDemo;

import nLayerDemo.business.abstracts.ProductService;
import nLayerDemo.business.concretes.ProductManager;
import nLayerDemo.core.JLoggerManagerAdapter;
import nLayerDemo.dataAccess.concretes.AbcProductDao;
import nLayerDemo.dataAccess.concretes.HibernateProductDao;
import nLayerDemo.entities.concretes.Product;

public class Main {
    public static void main(String[] args) {
        // entitler hariç new işlemi y apıyorsan ileride problem yaşıyacaksındır
        // gerçek sistemlerde genelde herşeye interface üzerinden erişim sağlıyoruz
        // ProductService productService = new ProductManager(new HibernateProductDao());
        ProductService productService = new ProductManager(new AbcProductDao(), new JLoggerManagerAdapter());
        Product product = new Product(1, 2, "Elma", 12, 54);
        productService.add(product);

    }
}
