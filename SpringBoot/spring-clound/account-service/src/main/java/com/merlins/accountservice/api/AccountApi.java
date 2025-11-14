package com.merlins.accountservice.api;


import com.merlins.accountservice.entity.Account;
import com.merlins.accountservice.service.AccountService;
import org.springframework.boot.autoconfigure.data.web.SpringDataWebProperties;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("account")
public class AccountApi {

    private  final AccountService accountService;

    public AccountApi(AccountService accountService) {
        this.accountService = accountService;
    }

    @GetMapping("/{id}")
    public ResponseEntity<Account> get(@PathVariable("id") String id){
        return ResponseEntity.ok( accountService.get(id));
    }

    @PostMapping
    public ResponseEntity<Account> save(@RequestBody Account account){
        return ResponseEntity.ok( accountService.save(account));
    }

    @PutMapping("/{id}")
    public ResponseEntity<Account> update(@RequestBody String id){
        return ResponseEntity.ok( accountService.update(id));
    }
    @DeleteMapping("/{id}")
    public void delete(@PathVariable("id") String id){
        accountService.delete(id);
    }

    public ResponseEntity<Account> pagination(SpringDataWebProperties.Pageable pageable){
        return ResponseEntity.ok( accountService.pagination(pageable));
    }
}
