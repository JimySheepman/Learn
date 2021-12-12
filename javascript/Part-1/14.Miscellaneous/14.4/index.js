let user = {
  name: "John",
  hi() {
    alert(this.name);
  },
};
let hi = user.hi;
hi(); // Error, because this is undefined
