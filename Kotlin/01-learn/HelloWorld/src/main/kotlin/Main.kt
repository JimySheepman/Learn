/*
* Kotlin Hard Keywords
* Following is a list of hard keywords and they cannot be used as identifiers:
*
* as
* as?
* break
* class
* continue
* do
* else
* false
* for
* fun
* if
* in
* !in
* interface
* is
* !is null
* object
* package
* return
* super
* this
* throw
* true
* try
* typealias
* typeof
* val
* var
* when
* while
*
* Kotlin Soft Keywords
* Following is the list of keywords (soft) in the context when they are applicable and can be used as identifiers in other contexts:
*
* by
* catch
* constructor
* delegate
* dynamic
* field
* file
* finally
* get
* import
* init
* param
* property
* receiver
* set
* setparam
* value
* where
*
* Kotlin Modifier Keywords
* Following is the list of tokens which act as keywords in modifier lists of declarations and can be used as identifiers in other contexts:
*
* actual
* abstract
* annotation
* companion
* const
* crossinline
* data
* enum
* expect
* external
* final
* infix
* inline
* inner
* internal
* lateinit
* noinline
* open
* operator
* out
* override
* private
* protected
* public
* reified
* sealed
* suspend
* tailrec
* vararg
* */

fun main() {
    var string: String = "Hello, World!"
    println("$string")

    // this is first commit

    var userName: String = "Alex"
    println(userName)
    userName = "Jhon"
    println(userName)


    var age: Int = 22

    println("Hello $userName your age is $age")

    var name = "test name"
    var height = 12.3

    println(name)
    println(height)

    println("Name = " + name)

    // Read-only Variables
    val readOnly = "test name"
    println(readOnly)

    // Kotlin Number Data Types
    val a: Int = 10000
    val d: Double = 100.00
    val f: Float = 100.00f
    val l: Long = 1000000004
    val s: Short = 10
    val b: Byte = 1

    println("Int Value is " + a)
    println("Double  Value is " + d)
    println("Float Value is " + f)
    println("Long Value is " + l)
    println("Short Value is " + s)
    println("Byte Value is " + b)

    // Kotlin Character Data Type
    val letter: Char    // defining a Char variable
    letter = 'A'        // Assigning a value to it
    println("$letter")

    println('\n') //prints a newline character
    println('\$') //prints a dollar $ character
    println('\\') //prints a back-slash \ character

    // Kotlin String Data Type
    val escapedString: String = "I am escaped String!\n"
    var rawString: String = """This is going to be a
   multi-line string and will
   not have any escape sequence""";

    print(escapedString)
    println(rawString)

    // Kotlin Boolean Data Type
    val A: Boolean = true   // defining a variable with true value
    val B: Boolean = false   // defining a variable with false value
    val boolNull: Boolean? = null

    println("Value of variable A " + A)
    println("Value of variable B " + B)
    println("Value of variable boolNull " + boolNull)

    // Kotlin Array Data Type
    val numbers: IntArray = intArrayOf(1, 2, 3, 4, 5)
    println("Value at 3rd position : " + numbers[2])

    // Kotlin Data Type Conversion
    val x1: Int = 100
    val y1: Long = x1.toLong()

    println(y1)

    var operatorsX: Int = 40
    var operatorsY: Int = 20

    var operatorsZ: Boolean = true
    var operatorsW: Boolean = false
    var operatorsT: Boolean = true

    var x: Int = 60      // 60 = 0011 1100
    var y: Int = 13      // 13 = 0000 1101
    var z: Int

    // Arithmetic Operators
    println("x + y = " + (operatorsX + operatorsY))
    println("x - y = " + (operatorsX - operatorsY))
    println("x / y = " + (operatorsX / operatorsY))
    println("x * y = " + (operatorsX * operatorsY))
    println("x % y = " + (operatorsX % operatorsY))

    // Relational Operators
    println("x > y = " + (operatorsX > operatorsY))
    println("x < y = " + (operatorsX < operatorsY))
    println("x >= y = " + (operatorsX >= operatorsY))
    println("x <= y = " + (operatorsX <= operatorsY))
    println("x == y = " + (operatorsX == operatorsY))
    println("x != y = " + (operatorsX != operatorsY))

    // Assignment Operators
    operatorsX += 5
    println("x += 5 = " + operatorsX)

    operatorsX = 40;
    operatorsX -= 5
    println("x -= 5 = " + operatorsX)

    operatorsX = 40
    operatorsX *= 5
    println("x *= 5 = " + operatorsX)

    operatorsX = 40
    operatorsX /= 5
    println("x /= 5 = " + operatorsX)

    operatorsX = 43
    operatorsX %= 5
    println("x %= 5 = " + operatorsX)

    // Unary Operators
    println("+x = " + (+operatorsX))
    println("-x = " + (-operatorsX))
    println("++x = " + (++operatorsX))
    println("--x = " + (--operatorsX))
    println("!x = " + (!true))

    // Logical Operators
    println("x && y =" + (operatorsZ && operatorsW))
    println("x || y =" + (operatorsZ || operatorsW))
    println("!y =" + (!operatorsW))

    // Bitwise Operations
    // 240 = 1111 0000
    z = x.shl(2)
    println("x.shl(2) = " + z)

    // 15 = 0000 1111
    z = x.shr(2)
    println("x.shr(2) = " + z)

    // 12 = 0000 1100
    z = x.and(y)
    println("x.and(y)  = " + z)

    // 61 = 0011 1101
    z = x.or(y)
    println("x.or(y)  = " + z)

    // 49 = 0011 0001
    z = x.xor(y)
    println("x.xor(y)  = " + z)

    // -61 = 1100 0011
    z = x.inv()
    println("x.inv()  = " + z)

    println("x.and(y) = " + operatorsZ.and(operatorsW))
    println("x.or(y) = " + operatorsZ.or(operatorsW))
    println("x.and(z) = " + operatorsZ.and(operatorsT))

    println("operatorsZ.toString() = " + operatorsZ.toString())
}