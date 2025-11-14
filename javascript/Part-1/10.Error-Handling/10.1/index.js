try {
  console.log("Start of try runs");
  lalala;
  console.log("End of try (never reached)");
} catch (err) {
  console.log("Catch is error has occurred! ");
}

try {
  lalala;
} catch (err) {
  console.log(err.name);
  console.log(err.message);
  console.log(err.stack);
  console.log(err);
}

let json = "{ bad json }";

try {

  let user = JSON.parse(json);
  console.log( user.name );

} catch (err) {
  console.log( "Our apologies, the data has errors, we'll try to request it one more time." );
  console.log( err.name );
  console.log( err.message );
}

try {
    console.log( 'try' );
    if (confirm('Make an error?')) BAD_CODE();
  } catch (err) {
    console.log( 'catch' );
  } finally {
    console.log( 'finally' );
  }