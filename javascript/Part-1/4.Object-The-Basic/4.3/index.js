let user = {
  name: "john",
};

user = null;

let user_2 = {
  name: "john",
};
let admin = user;
user_2 = null;

function marry(man, woman) {
  woman.husband = man;
  man.wife = woman;
  return {
    father: man,
    mother: woman,
  };
}

let family = marry(
  {
    name: "john",
  },
  {
    name: "ann",
  }
);

delete family.father;
delete family.mother.husband;
family = null;
