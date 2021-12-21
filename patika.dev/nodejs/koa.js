const Koa = require("koa");
const app = new Koa();

// response
app.use((ctx) => {
  // console.log(ctx.request.href); -> full URL
  if (ctx.url === "/") {
    ctx.body = "<h1>index</h1>";
  } else if (ctx.url === "/about") {
    ctx.body = "<h1>About</h1>";
  } else if (ctx.url === "/contact") {
    ctx.body = "<h1>contact</h1>";
  } else {
    ctx.body = "Not Found->404";
  }
});

app.listen(3000);
