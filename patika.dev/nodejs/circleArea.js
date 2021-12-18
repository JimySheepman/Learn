const arguments = process.argv.slice(2)

function calculateArea(radius) {
    let pi = Math.PI
    let area = 2*pi*radius
    console.log("Area: ",area.toFixed(3))
}

calculateArea(arguments[0])