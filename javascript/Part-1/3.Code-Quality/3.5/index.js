describe("pow", function () {
  // Automated testing with Mocha
  it("raises to n-th power", function () {
    assert.equal(pow(2, 3), 8);
    assert.equal(pow(3, 4), 81);
  });
});

function pow(x, n) {
  return 8;
}

function pow2(x, n) {
  let result = 1;

  for (let i = 0; i < n; i++) {
    result *= x;
  }

  return result;
}

describe("pow", function () {
  function makeTest(x) {
    let expected = x * x * x;
    it(`${x} in the power 3 is ${expected}`, function () {
      assert.equal(pow2(x, 3), expected);
    });
  }

  for (let x = 1; x <= 5; x++) {
    makeTest(x);
  }
});

describe("pow", function () {
  describe("raises x to power 3", function () {
    function makeTest(x) {
      let expected = x * x * x;
      it(`${x} in the power 3 is ${expected}`, function () {
        assert.equal(pow(x, 3), expected);
      });
    }

    for (let x = 1; x <= 5; x++) {
      makeTest(x);
    }
  });

  // ... more tests to follow here, both describe and it can be added
});

describe("test", function () {
  before(() => alert("Testing started – before all tests"));
  after(() => alert("Testing finished – after all tests"));

  beforeEach(() => alert("Before a test – enter a test"));
  afterEach(() => alert("After a test – exit a test"));

  it("test 1", () => alert(1));
  it("test 2", () => alert(2));
});

describe("pow", function () {
  // ...

  it("for negative n the result is NaN", function () {
    assert.isNaN(pow(2, -1));
  });

  it("for non-integer n the result is NaN", function () {
    assert.isNaN(pow(2, 1.5));
  });
});

function pow(x, n) {
  if (n < 0) return NaN;
  if (Math.round(n) != n) return NaN;

  let result = 1;

  for (let i = 0; i < n; i++) {
    result *= x;
  }

  return result;
}
