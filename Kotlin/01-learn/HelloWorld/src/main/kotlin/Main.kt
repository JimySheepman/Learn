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
    var string:String = "Hello, World!"
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

    println("Name = "+name)

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
    println("Long Value is " + l )
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
    val escapedString : String  = "I am escaped String!\n"
    var rawString :String  = """This is going to be a
   multi-line string and will
   not have any escape sequence""";

    print(escapedString)
    println(rawString)

    // Kotlin Boolean Data Type
    val A: Boolean = true   // defining a variable with true value
    val B: Boolean = false   // defining a variable with false value
    val boolNull: Boolean? = null

    println("Value of variable A "+ A )
    println("Value of variable B "+ B )
    println("Value of variable boolNull "+ boolNull )

    // Kotlin Array Data Type
    val numbers: IntArray = intArrayOf(1, 2, 3, 4, 5)
    println("Value at 3rd position : " + numbers[2])

    // Kotlin Data Type Conversion
    val x: Int = 100
    val y: Long = x.toLong()

    println(y)
}