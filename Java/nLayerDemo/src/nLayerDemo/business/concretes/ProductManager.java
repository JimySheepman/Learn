package nLayerDemo.business.concretes;

import nLayerDemo.business.abstracts.ProductService;
import nLayerDemo.core.LoggerService;
import nLayerDemo.dataAccess.abstracts.ProductDao;
import nLayerDemo.entities.concretes.Product;

import java.util.List;

// microservis mimarisi dış bağımlılıklardan arınma
// unitTest kodu kodla test etmektir.
public class ProductManager implements ProductService {
    //dependence injection
    private ProductDao productDao;
    private LoggerService loggerService;

    public ProductManager(ProductDao productDao, LoggerService loggerService) {
        super();
        this.productDao = productDao;
        this.loggerService = loggerService;
    }


    @Override
    public void add(Product product) {
        if (product.getCategoryId() == 1) {
            System.out.println("Bu kategoride ürün kabul edilmiyor.");
            return;
        }
        this.productDao.add(product);
        this.loggerService.logToSystem("Ürün Eklendi " + product.getName());
    }

    @Override
    public List<Product> getAll() {
        return null;
    }
}
