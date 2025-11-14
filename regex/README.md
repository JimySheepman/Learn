# Regex

```go
// ^[a-z0-9_-]{3,15}$
asdd merlins12331234 asd
asd
asda
aa
c1s=
// the The
The fat cat sat on the mat the .
// .ar
The car parked in the garage.
// [Tt]he
The fat cat sat on the mat the .
//ar[.]
A garage is a good place to park a car.
// [^c]ar
The car parked in the garage.
// [a-z]*
The car parked in the garage
// \s*cat\s
The fat cat sat on the concatenation.
// c.+t
The fat cat sat on the mat.
// [T]he
The car is parked in the garage.
// [T]?he
The car is parked in the garage.
// [0-9]{2,3}
The number was 9.9997 but we rounded it off to 10.0.
// [0-9]{2,}
The number was 9.9997 but we rounded it off to 10.0.
// [0-9]{3}
The number was 9.9997 but we rounded it off to 10.0.
// (c|g|p)ar
The car is parked in the garage.
// (?:c|g|p)ar
The car is parked in the garage.
// (T|t)he|car
The car is parked in the garage.
// (f|c|m)at\.?
The fat cat sat on the mat.
// (T|t)he
The car is parked in the garage.
// ^(T|t)he
The car is parked in the garage.
// (at\.)
The fat cat. sat. on the mat.
// (at\.)$
The fat cat. sat. on the mat.
// [T|t]he(?=\sfat)
The fat cat sat on the mat.
// [T|t]he(?!\sfat)
The fat cat sat on the mat.
// (?<=[T|t]he\s)(fat|mat)
The fat cat sat on the mat.
// (?<![T|t]he\s)(cat)
The cat sat on cat.
// The
The fat cat sat on the mat.
// /The/gi
The fat cat sat on the mat.
// /.(at)/ 
The fat cat sat on the mat.
// /.(at)/g
The fat cat sat on the mat.
// /.at(.)?$/
The fat
cat sat
on the mat.
// /.at(.)?$/gm 
The fat
cat sat
on the mat.
// /(.*at)/
The fat cat sat on the mat. 
// /(.*?at)/
The fat cat sat on the mat. 
```
