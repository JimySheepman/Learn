select title,description from film;
select * from film where length > 60 and length <75;
select * from film where rental_rate =0.99 and replacement_cost =12.99 or replacement_cost =28.99;
select last_name from customer where first_name ='Mary';
select * from  film where not(rental_rate=4.99 or rental_rate=2.99) and not length >50;
