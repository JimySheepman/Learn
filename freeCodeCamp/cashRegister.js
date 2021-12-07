const unitValue = [
  ["PENNY", 0.01],
  ["NICKEL", 0.05],
  ["DIME", 0.1],
  ["QUARTER", 0.25],
  ["ONE", 1],
  ["FIVE", 5],
  ["TEN", 10],
  ["TWENTY", 20],
  ["ONE HUNDRED", 100],
];

const cidToObj = (cidArray) => {
  let obj = {};
  let copiedArray = [...cidArray];
  copiedArray.reverse().forEach((item) => {
    obj[item[0]] = item[1];
  });
  return obj;
};

function checkCashRegister(price, cash, cid) {
  const cidObj = cidToObj(cid);
  const unitValueObj = cidToObj(unitValue);
  let res = { status: "", change: [] };
  let updatedCidObj = {};
  let updatedChange = cash - price;
  for (let unit of Object.keys(cidObj)) {
    let tempResAmount = Math.floor(updatedChange / unitValueObj[unit]);
    let currentAmount = Math.round(cidObj[unit] / unitValueObj[unit]);
    let resAmount =
      currentAmount <= tempResAmount ? currentAmount : tempResAmount;
    let resAmountValue = resAmount * unitValueObj[unit];
    updatedCidObj[unit] = parseFloat(
      (cidObj[unit] - resAmountValue).toFixed(2)
    );
    if (resAmount > 0) {
      res.change.push([unit, resAmountValue]);
      updatedChange = parseFloat((updatedChange - resAmountValue).toFixed(2));
    }
  }
  if (updatedChange > 0) {
    res = { status: "INSUFFICIENT_FUNDS", change: [] };
  } else if (
    Object.keys(updatedCidObj).every((unit) => 0 === updatedCidObj[unit])
  ) {
    res = { status: "CLOSED", change: cid };
  } else {
    res.status = "OPEN";
  }
  return res;
}
checkCashRegister(19.5, 500, [
  ["PENNY", 400],
  ["NICKEL"],
  ["DIME", 0],
  ["QUARTER", 0],
  ["ONE", 0],
  ["FIVE", 0],
  ["TEN", 0],
  ["TWENTY", 0],
  ["ONE HUNDRED", 800],
]);
