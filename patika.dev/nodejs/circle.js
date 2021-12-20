function circleArea(radius) {
  let pi = Math.PI;
  let area = radius * pi * radius;
  console.log("Area: ", area.toFixed(3));
}

function circleCircumference(radius) {
  let pi = Math.PI;
  let area = 2 * pi * radius;
  console.log("Circumference: ", area.toFixed(3));
}

module.exports = {
  circleArea,
  circleCircumference,
};
