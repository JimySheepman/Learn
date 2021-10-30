package kodlamaio.northwind.entities.abstracts;

import kodlamaio.northwind.entities.concretes.Product;
import org.springframework.data.jpa.repository.JpaRepository;

//sorguları yapmamızı sağlıyor , repositoryden aldık
public interface ProductDao extends JpaRepository<Product,Integer> {
}
