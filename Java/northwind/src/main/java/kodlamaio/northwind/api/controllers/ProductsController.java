package kodlamaio.northwind.api.controllers;

import kodlamaio.northwind.business.abstracts.ProductService;
import kodlamaio.northwind.entities.concretes.Product;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController // mimari yapımızı belirliyor Restful
@RequestMapping("/api/products") // isteği geldiği yeri belirtir. Ve gelen yere göre controllere çalıştırır.
public class ProductsController {

    private ProductService productService;

    @Autowired // ProductService productService Bu objeyi projede tanıyıp buluyor ve dinamik olarak new' leme işemini yapıyor.
    public ProductsController(ProductService productService) {
        this.productService = productService;
    }

    @GetMapping("/getall") // /api/products/getall diye istek gelrise altaki method çalışır.
    public List<Product> getAll() {
        return this.productService.getAll();
    }
}
