
var recipes = require("../recipes.json");
var router = require("express").Router();

router.get("/shopping-list", (req, res) => {
  let ids = req.query.ids;
  if (!ids) {
    res.status(400).send("NOT_FOUND");
    return;
  }
  ids = ids.split(",");
  ids.forEach((id, index) => {
    ids[index] = Number(id);
  });
  const filteredRecipes = recipes.filter((recipe) => {
    return ids.includes(recipe.id);
  });

  if (!filteredRecipes.length) {
    res.status(404).send("NOT_FOUND");
    return;
  }

  let ingrediants = [];

  for (let i = 0; i < filteredRecipes.length; i++) {
    ingrediants = [...ingrediants, ...filteredRecipes[i].ingredients];
  }

  return res.status(200).send(ingrediants);
});

module.exports = router;