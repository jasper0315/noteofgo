<link rel="stylesheet" type="text/css" href="styles.css">
<h1> Java Tutorials</h1> 

<h1>Contents</h1>
<ol>
    <li><a href="#sec0">Words</a></li>
    <li><a href="#sec1">Java Syntax</a></li>
    <li><a href="#sec2">Java Output</a></li>
    <li><a href="#sec3">Java Comments</a></li>
    <li><a href="#sec4">Java Variables</a></li>
    <li><a href="#sec5">Java Data Types</a></li>
    <li><a href="#sec6">Java Type Casting</a></li>
    <li><a href="#sec7">Java Operators</a></li>
    <li><a href="#sec8">Java Strings</a></li>
    <li><a href="#sec9">Java Math</a></li>
    <li><a href="#sec10">Java Booleans</a></li>
    <li><a href="#sec11">Java If...Else</a></li>
    <li><a href="#sec12">Java Switch</a></li>
    <li><a href="#sec13">Java While Loop</a></li>
    <li><a href="#sec14">Java For Loop</a></li>
    <li><a href="#sec15">Java Break/Continue</a></li>
    <li><a href="#sec16">Java Arrays</a></li>
</ol>


<h1 id=sec0>Words</h1>
<ul>
<li><strong>access modifier :</strong><br>is a keyword in many programming languages, including Java, that <span class="red-text"><u>is used to specify the visibility and accessibility of classes, methods, variables, and other class members.</u></span> Access modifiers control which parts of your code can access or modify a particular class member. They help you define the level of encapsulation and security in your program. In Java, there are four main access modifiers</li>
<li><strong>public :</strong><br>is an access modifier in Java that specifies the visibility and accessibility of a class, method, or variable. When a member is declared as public, <span class="red-text"><u>it can be accessed from anywhere in the code, including from other classes and packages.</u></span></li>
<li><strong>protected :</strong><br><span class="red-text"><u> Members marked as protected are accessible within the same package and also by subclasses, regardless of whether they are in the same package or not.</u></span> This access modifier is often used for allowing controlled access to class members for subclasses.</li>
<li><strong>default (no modifier) :</strong><br> If no access modifier is specified, the default access modifier is applied. <span class="red-text"><u>Members with default access are accessible only within the same package. They are not accessible outside the package,</u></span> even if the class is part of a different package.</li>
<li><strong>private :</strong><br> <span class="red-text"><u>Members marked as private are accessible only within the same class.</u></span> They are not accessible from outside the class. This is the most restrictive access modifier and provides the highest level of encapsulation.</li>
<li><strong>Object</strong><br> is an instance of a class. A class is a blueprint or template that defines the attributes (data) and behaviors (methods) that the objects of that class will have. <span class="red-text"><u>Objects are created based on these class definitions.</u></span></li>
</ul>


<h1 id="sec1">Java Syntax</h1>
<p>
<h2> Java Syntax </h2>
we used the following code to print "Hello World" to the screen:

```java
public class Main {
  public static void main(String[] args) {
    System.out.println("Hello World");
  }
}
```

<h3>Example explained</h3>
Every line of code that runs in Java must be inside a class. In our example, we named the class Main. A class should always start with an uppercase first letter.

    Note:
    Java is case-sensitive: "MyClass" and "myclass" has different meaning.

The name of the java file must match the class name. When saving the file, save it using the class name and add ".java" to the end of the filename. To run the example above on your computer, make sure that Java is properly installed: Go to the Get Started Chapter for how to install Java. The output should be:

    Hello World

<h2>The main Method</h2>
The main() method is required and you will see it in every Java program:

```java
public static void main(String[] args)
```
Any code inside the main() method will be executed. Don't worry about the keywords before and after main. You will get to know them bit by bit while reading this tutorial.

For now, just remember that every Java program has a class name which must match the filename, and that every program must contain the main() method.

<h2>System.out.println()</h2>
Inside the main() method, we can use the println() method to print a line of text to the screen:

```java
public static void main(String[] args) {
  System.out.println("Hello World");
}
```

    Note: 
    The curly braces {} marks the beginning and the end of a block of code.

System is a built-in Java class that contains useful members, such as out, which is short for "output". The println() method, short for "print line", is used to print a value to the screen (or a file).

Don't worry too much about System, out and println(). Just know that you need them together to print stuff to the screen.

You should also note that each code statement must end with a semicolon (;).
</p>


<h1 id = "sec2">Java Output</h1>

<h2>Print Text</h2>
You learned from the previous chapter that you can use the println() method to output values or print text in Java:

```java
ExampleGet your own Java Server
System.out.println("Hello World!");
```

You can add as many println() methods as you want. Note that it will add a new line for each method:

Example

```java
System.out.println("Hello World!");
System.out.println("I am learning Java.");
System.out.println("It is awesome!");
```

<h2>Double Quotes</h2>
When you are working with text, it must be wrapped inside double quotations marks "".

If you forget the double quotes, an error occurs:

Example

```java
System.out.println("This sentence will work!");
System.out.println(This sentence will produce an error);
```

<h2>The Print() Method</h2>
There is also a print() method, which is similar to println().

The only difference is that it does not insert a new line at the end of the output:

Example

```java
System.out.print("Hello World! ");
System.out.print("I will print on the same line.");
```

Note that we add an extra space (after "Hello World!" in the example above), for better readability.


<h2>Print Numbers</h2>
You can also use the println() method to print numbers.
However, unlike text, we don't put numbers inside double quotes:

```java
System.out.println(3);
System.out.println(358);
System.out.println(50000);
```
You can also perform mathematical calculations inside the println() method:

```java
System.out.println(3 + 3);

System.out.println(2 * 5);
```

<h1 id=sec3>Java Comments</h1>
Comments can be used to explain Java code, and to make it more readable. It can also be used to prevent execution when testing alternative code.

<h2>Single-line Comments</h2>
Single-line comments start with two forward slashes (//).
Any text between // and the end of the line is ignored by Java (will not be executed).

This example uses a single-line comment before a line of code:

```java
// This is a comment
System.out.println("Hello World");
```

This example uses a single-line comment at the end of a line of code:

```java
System.out.println("Hello World"); // This is a comment
```

<h2>Java Multi-line Comments</h2>
Multi-line comments start with /* and ends with */.
Any text between /* and */ will be ignored by Java.
This example uses a multi-line comment (a comment block) to explain the code:

```java
/* The code below will print the words Hello World
to the screen, and it is amazing */
System.out.println("Hello World");
```

    Single or multi-line comments?
    It is up to you which you want to use. Normally, we use // for short comments, and /* */ for longer.


<h1 id=sec4>Java Variables</h1>
<h2>Variables</h2>

Variables are containers for storing data values.
In Java, there are different <strong>types</strong> of variables, for example:

<strong>String</strong> - stores text, such as "Hello". String values are surrounded by double quotes

<strong>int</strong> - stores integers (whole numbers), without decimals, such as 123 or -123

<strong>float</strong> - stores floating point numbers, with decimals, such as 19.99 or -19.99

<strong>char</strong> - stores single characters, such as 'a' or 'B'. Char values are surrounded by single quotes

<strong>boolean</strong> - stores values with two states: true or false

<h3>Declaring (Creating) Variables</h3>
To create a variable, you must specify the type and assign it a value:

Syntax

```java
type variableName = value;
```

Where type is one of Java's types (such as int or String), and variableName is the name of the variable (such as x or name). The equal sign is used to assign values to the variable.

To create a variable that should store text, look at the following Example:

Create a variable called name of type String and assign it the value "John":

```java
String name = "John";
System.out.println(name);
```


To create a variable that should store a number, look at the following example:

Example:

Create a variable called myNum of type int and assign it the value 15:

```java
int myNum = 15;
System.out.println(myNum);
```

You can also declare a variable without assigning the value, and assign the value later:

Example:

```java
int myNum;
myNum = 15;
System.out.println(myNum);
```

Note that if you assign a new value to an existing variable, it will overwrite the previous value:

Example

Change the value of myNum from 15 to 20:

```java
int myNum = 15;
myNum = 20;  // myNum is now 20
System.out.println(myNum);
```

<h3>Final Variables</h3>
If you don't want others (or yourself) to overwrite existing values, use the final keyword (this will declare the variable as "final" or "constant", which means unchangeable and read-only):

Example

```java
final int myNum = 15;
myNum = 20;  // will generate an error: cannot assign a value to a final variable
```

<h3>Other Types</h3>
A demonstration of how to declare variables of other types:

Example
```java
int myNum = 5;
float myFloatNum = 5.99f;
char myLetter = 'D';
boolean myBool = true;
String myText = "Hello";
```

<h2>Display Variables</h2>
The println() method is often used to display variables.
To combine both text and a variable, use the + character:

Example

```java
String name = "John";
System.out.println("Hello " + name);
```

You can also use the + character to add a variable to another variable:

Example

```java
String firstName = "John ";
String lastName = "Doe";
String fullName = firstName + lastName;
System.out.println(fullName);
```

For numeric values, the + character works as a mathematical operator (notice that we use int (integer) variables here):

Example

```java
int x = 5;
int y = 6;
System.out.println(x + y); // Print the value of x + y
```

From the example above, you can expect:
<ul>
    <li>x stores the value 5</li>
    <li>y stores the value 6</li>
    <li>Then we use the println() method to display the value of x + y, which is 11</li>
</ul>

<h2>Declare Many Variables</h2>
To declare more than one variable of the <strong>same type</strong>, you can use a comma-separated list:

ExampleGet your own Java Server

Instead of writing:

```java
int x = 5;
int y = 6;
int z = 50;
System.out.println(x + y + z);
```
You can simply write:

```java
int x = 5, y = 6, z = 50;
System.out.println(x + y + z);
```

<h2>One Value to Multiple Variables</h2>
You can also assign the same value to multiple variables in one line:

Example

```java
int x, y, z;
x = y = z = 50;
System.out.println(x + y + z);
```

<h2>Java Identifiers</h2>
<h3>Identifiers</h3>

All Java variables must be identified with unique names. These unique names are called identifiers. Identifiers can be short names (like x and y) or more descriptive names (age, sum, totalVolume).

Note: It is recommended to use descriptive names in order to create understandable and maintainable code:

Example:

```java
// Good
int minutesPerHour = 60;

// OK, but not so easy to understand what m actually is
int m = 60;
```

The general rules for naming variables are:
<ul>
    <li>Names can contain letters, digits, underscores, and dollar signs</li>
    <li>Names must begin with a letter</li>
    <li>Names should start with a lowercase letter and it cannot contain whitespace</li>
    <li>Names can also begin with $ and _ (but we will not use it in this tutorial)</li>
    <li>Names are case sensitive ("myVar" and "myvar" are different variables)</li>
    <li>Reserved words (like Java keywords, such as int or boolean) cannot be used as names</li>
</ul>


<h1 id=sec5>Java Data Types</h1>
<h2>Java Data Types</h2>
As explained in the previous chapter, a variable in Java must be a specified data type:

Example


```java
int myNum = 5;               // Integer (whole number)
float myFloatNum = 5.99f;    // Floating point number
char myLetter = 'D';         // Character
boolean myBool = true;       // Boolean
String myText = "Hello";     // String
```

Data types are divided into two groups:
<ul>
    <li>Primitive data types - includes byte, short, int, long, float, double, boolean and char</li>
    <li>Non-primitive data types - such as String, Arrays and Classes (you will learn more about these in a later chapter)</li>
</ul>

<h3>Primitive Data Types</h3>
A primitive data type specifies the size and type of variable values, and it has no additional methods.

There are eight primitive data types in Java:

<table>
  <tr>
    <th>Data Type</th>
    <th>Size</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>byte</td>
    <td>1 byte</td>
    <td>Stores whole numbers from -128 to 127</td>
  </tr>
  <tr>
    <td>short</td>
    <td>2 bytes</td>
    <td>Stores whole numbers from -32,768 to 32,767</td>
  </tr>
  <tr>
    <td>int</td>
    <td>4 bytes</td>
    <td>Stores whole numbers from -2,147,483,648 to 2,147,483,647</td>
  </tr>
  <tr>
    <td>long</td>
    <td>8 bytes</td>
    <td>Stores whole numbers from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807</td>
  </tr>
  <tr>
    <td>float</td>
    <td>4 bytes</td>
    <td>Stores fractional numbers. Sufficient for storing 6 to 7 decimal digits</td>
  </tr>
  <tr>
    <td>double</td>
    <td>8 bytes</td>
    <td>Stores fractional numbers. Sufficient for storing 15 decimal digits</td>
  </tr>
  <tr>
    <td>boolean</td>
    <td>1 bit</td>
    <td>Stores true or false values</td>
  </tr>
  <tr>
    <td>char</td>
    <td>2 bytes</td>
    <td>Stores a single character/letter or ASCII values</td>
  </tr>
</table>
note that:
<ul>
  <li><strong>int</strong><br> = an abbreviation of integer</li>
  <li><strong>char</strong><br>= character like A, B or C</li>
</ul>

<h2>Java Numbers</h2>
<h3>Numbers</h3>
Primitive number types are divided into two groups:
<ul>
<li><strong>Integer types</strong> stores whole numbers, positive or negative (such as 123 or -456), without decimals. Valid types are <strong>byte</strong>, <strong>short</strong>, <strong>int</strong> and <strong>long</strong>. Which type you should use, depends on the numeric value.</li>

<li><strong>Floating point types</strong> represents numbers with a fractional part, containing one or more decimals. There are two types: <strong>float</strong> and <strong>double</strong>.</li>
</ul>

    Even though there are many numeric types in Java, the most used for numbers are int (for whole numbers) and double (for floating point numbers). However, we will describe them all as you continue to read.

<h3>Integer Types</h3>
<h4>Byte</h4>
The <strong>byte</strong> data type can store whole numbers from -128 to 127. This can be used instead of int or other integer types to save memory when you are certain that the value will be within -128 and 127:

Example

```java
byte myNum = 100;
System.out.println(myNum);
```

<h4>Short</h4>
The <strong>short</strong> data type can store whole numbers from -32768 to 32767:

Example

```java
short myNum = 5000;
System.out.println(myNum);
```

<h4>Int</h4>
The <strong>int</strong> data type can store whole numbers from -2147483648 to 2147483647. In general, and in our tutorial, the int data type is the preferred data type when we create variables with a numeric value.

Example
```java
int myNum = 100000;
System.out.println(myNum);
```

<h4>Long</h4>
The <strong>long</strong> data type can store whole numbers from -9223372036854775808 to 9223372036854775807. This is used when int is not large enough to store the value. Note that you should end the value with an "L":

Example
```java
long myNum = 15000000000L;
System.out.println(myNum);
```

<h3>Floating Point Types</h4>
You should use a floating point type whenever you need a number with a decimal, such as 9.99 or 3.14515.
The <strong>float</strong> and <strong>double</strong> data types can store fractional numbers. Note that you should end the value with an "f" for floats and "d" for doubles:

Float Example
```java
float myNum = 5.75f;
System.out.println(myNum);
```

Double Example
```java
double myNum = 19.99d;
System.out.println(myNum);
```

<h4>Use <strong>float</strong> or <strong>double</strong>?</h4>

The precision of a floating point value indicates how many digits the value can have after the decimal point. The precision of float is only six or seven decimal digits, while <strong>double</strong> variables have a precision of about 15 digits. Therefore it is safer to use double for most calculations.

<h3>Scientific Numbers</h3>
A floating point number can also be a scientific number with an "e" to indicate the power of 10:

Example
```java
float f1 = 35e3f;
double d1 = 12E4d;
System.out.println(f1);
System.out.println(d1);
```

<h2>Java Boolean Data Types</h2>
<h3>Boolean Types</h3>
Very often in programming, you will need a data type that can only have one of two values, like:
<ul>
<li>YES / NO</li>
<li>ON / OFF</li>
<li>TRUE / FALSE</li>
</ul>
For this, Java has a boolean data type, which can only take the values true or false:

Example
```java
boolean isJavaFun = true;
boolean isFishTasty = false;
System.out.println(isJavaFun);     // Outputs true
System.out.println(isFishTasty);   // Outputs false
```

    Boolean values are mostly used for conditional testing.

    You will learn much more about booleans and conditions later in this tutorial.

<h2>Java Characters</h2>
<h3>Characters</h3>

The char data type is used to store a <strong>single</strong> character. The character must be surrounded by single quotes, like 'A' or 'c':

Example
```java
char myGrade = 'B';
System.out.println(myGrade);
```

Alternatively, if you are familiar with ASCII values, you can use those to display certain characters:

Example
```java
char myVar1 = 65, myVar2 = 66, myVar3 = 67;
System.out.println(myVar1);
System.out.println(myVar2);
System.out.println(myVar3);
```
    Note that you can remember the knowledge below for reading the code above.
    ・ 65 = A
    ・ 66 = B
    ・ 67 = C

Tip: A list of all ASCII values can be found in our ASCII Table Reference.

<h3>Strings</h3>
The String data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:

Example
```java
String greeting = "Hello World";
System.out.println(greeting);
```

The String type is so much used and integrated in Java, that some call it "the special <strong>ninth</strong> type".

A String in Java is actually a <strong>non-primitive</strong> data type, because it refers to an object. The String object has methods that are used to perform certain operations on strings. <strong>Don't worry if you don't understand the term "object" just yet.</strong> We will learn more about strings and objects in a later chapter.

<h2>Non-Primitive Data Types</h2>
Non-primitive data types are called reference types because they refer to objects.

The main difference between primitive and non-primitive data types are:
<ul>
<li>Primitive types are predefined (already defined) in Java. Non-primitive types are created by the programmer and is not defined by Java (except for String).</li>
<li>Non-primitive types can be used to call methods to perform certain operations, while primitive types cannot.</li>
<li>A primitive type has always a value, while non-primitive types can be null.</li>
<li>A primitive type starts with a lowercase letter, while non-primitive types starts with an uppercase letter.</li>
</ul>
Examples of non-primitive types are Strings, Arrays, Classes, Interface, etc. You will learn more about these in a later chapter.


<h1 id=sec6>Java Type Casting</h1>
<h2>Java Type Casting</h2>
Type casting is when you assign a value of one primitive data type to another type.

In Java, there are two types of casting:
<ul>
<li><strong>Widening Casting</strong> (automatically) - converting a smaller type to a larger type size<br>
byte -> short -> char -> int -> long -> float -> double</li>

<li><strong>Narrowing Casting</strong> (manually) - converting a larger type to a smaller size type<br>
double -> float -> long -> int -> char -> short -> byte</li>
</ul>

<h2>Widening Casting</h2>
Widening casting is done automatically when passing a smaller size type to a larger size type:

Example
```java
public class Main {
  public static void main(String[] args) {
    int myInt = 9;
    double myDouble = myInt; // Automatic casting: int to double

    System.out.println(myInt);      // Outputs 9
    System.out.println(myDouble);   // Outputs 9.0
  }
}
```

<h2>Narrowing Casting</h2>
Narrowing casting must be done manually by placing the type in parentheses in front of the value:

Example
```java
public class Main {
  public static void main(String[] args) {
    double myDouble = 9.78d;
    int myInt = (int) myDouble; // Manual casting: double to int

    System.out.println(myDouble);   // Outputs 9.78
    System.out.println(myInt);      // Outputs 9
  }
}
```


<h1 id=sec7>Java Operators</h1>
<h2>Java Operators</h2>

Operators are used to perform operations on variables and values.

In the example below, we use the + <strong>operator</strong> to add together two values:

Example
```java
int x = 100 + 50;
```

Although the + operator is often used to add together two values, like in the example above, it can also be used to add together a variable and a value, or a variable and another variable:

Example
```java
int sum1 = 100 + 50;        // 150 (100 + 50)
int sum2 = sum1 + 250;      // 400 (150 + 250)
int sum3 = sum2 + sum2;     // 800 (400 + 400)
```

Java divides the operators into the following groups:
<ul>
  <li>Arithmetic operators</li>
  <li>Assignment operators</li>
  <li>Comparison operators</li>
  <li>Logical operators</li>
  <li>Bitwise operators</li>
</ul>

<h2>Arithmetic Operators</h2>
Arithmetic operators are used to perform common mathematical operations.

<table>
<tr>
  <th>Operator</th>
  <th>Name</th>
  <th>Description</th>
  <th>Example</th>
</tr>
<tr>
  <td> + </td>
  <td>Addition</td>
  <td>Adds together two values</td>
  <td>x + y</td>	
</tr>
<tr>
  <td> - </td>
  <td>Subtraction</td>
  <td>Subtracts one value from another</td>
  <td>x - y</td>
</tr>	
  <td>*</td>
  <td>Multiplication</td>
  <td>Multiplies two values</td> 
  <td>x * y</td>
</tr>
<tr>
  <td>/</td>
  <td>Division</td>
  <td>Divides one value by another</td>
  <td>x / y</td>
</tr>
<tr>
  <td>%</td>
  <td>Modulus</td>
  <td>Returns the division remainder</td>
  <td>x % y</td>
</tr>
<tr>
  <td>++</td>
  <td>Increment</td>
  <td>Increases the value of a variable by 1</td>
  <td>++x</td>
</tr>
<tr>
  <td>--</td>
  <td>Decrement</td>
  <td>Decreases the value of a variable by 1</td>
  <td>--x</td>
</tr>
</table>

<h2>Java Assignment Operators</h2>
Assignment operators are used to assign values to variables.

In the example below, we use the assignment operator (=) to assign the value 10 to a variable called x:

Example
```java
int x = 10;
```

The addition assignment operator (+=) adds a value to a variable:

Example
```java
int x = 10;
x += 5;
```

A list of all assignment operators:
<table>
<tr>
  <th>Operator</th>
  <th>Example</th>
  <th>Same As</th>
</tr>
<tr>
  <td>=</td>
  <td>x = 5</td>
  <td>x = 5</td>
</tr>
<tr>
  <td>+=</td>
  <td>x += 3</td>
  <td>x = x + 3</td>
</tr>
<tr>
  <td>-=</td>
  <td>x -= 3</td>
  <td>x = x - 3</td>
</tr>
<tr>
  <td>*=</td>
  <td>x *= 3</td>
  <td>x = x * 3</td>
</tr>
<tr>
  <td>/=</td>
  <td>x /= 3</td>
  <td>x = x / 3</td>
</tr>
<tr>
  <td>%=</td>
  <td>x %= 3</td>
  <td>x = x % 3</td>
</tr>
<tr>
  <td>&=</td>
  <td>x &= 3</td>
  <td>x = x & 3</td>
</tr>
<tr>
  <td>|=</td>
  <td>x |= 3</td>
  <td>x = x | 3</td>
</tr>
<tr>
  <td>^=</td>
  <td>x ^= 3</td>
  <td>x = x ^ 3</td>
</tr>
<tr>
  <td>>>=</td>
  <td>x >>= 3</td>
  <td>x = x >> 3</td>
</tr>
<tr>
  <td><<=</td>
  <td>x <<= 3</td>
  <td>x = x << 3</td>
</tr>
</table>
note that:
<ul>
  <li><strong>"|="</strong><br>= <span class="red-text"><u>performing a bitwise OR operation between the variable x and the integer value 3,</u></span> and then assigning the result back to the variable x. </li>
  <li><strong>Bitwise</strong><br>= (computing)treating a value as a series of bits rather than a numerical quantity.</li>
  <li><strong>A bitwise OR operators</strong><br>= <span class="red-text"><u>a fundamental bitwise operation used to combine the individual bits of two binary numbers.</u></span> The OR operation takes two binary values (or bits) and produces a result in which a particular bit in the output is set to 1 if at least one of the corresponding bits in the input values is 1. If you wanna know more specifically, you can chech <a href="https://www.upgrad.com/blog/what-is-bitwise-operator-in-java-and-its-classification-with-examples/">this article.</a></strong></u></li>
</ul>

<h2>Java Comparison Operators</h2>
Comparison operators are used to compare two values (or variables). This is important in programming, because it helps us to find answers and make decisions.

The return value of a comparison is either true or false. These values are known as Boolean values, and you will learn more about them in the Booleans and If..Else chapter.

In the following example, we use the greater than operator (>) to find out if 5 is <strong>greater than</strong> 3:

Example
```java
int x = 5;
int y = 3;
System.out.println(x > y); // returns true, because 5 is higher than 3
```

<table>
<tr>
  <th>Operator</th>
  <th>Name</th>
  <th>Example</th>
</tr>
<tr>
  <td>==</td>
  <td>Equal to</td>
  <td>x == y</td>
</tr>
<tr>
  <td>!=</td>
  <td>Not equal</td>
  <td>x != y</td>
</tr>
<tr>
  <td>></td>
  <td>Greater than</td>
  <td>x > y</td>
</tr>
<tr>
  <td><</td>
  <td>Less than</td>
  <td>x < y</td>
</tr>
<tr>
  <td>>=</td>
  <td>Greater than or equal to</td>
  <td>x >= y</td>
</tr>
<tr>
  <td><=</td>
  <td>Less than or equal to</td>
  <td>x <= y</td>
</tr>
</table>

<h2>Java Logical Operators</h2>
You can also test for true or false values with logical operators.<br>Logical operators are used to determine the logic between variables or values:

<table>
<tr>
  <th>Operator</th>
  <th>Name</th>
  <th>Description</th>
  <th>Example</th>
</tr>
<tr>
  <td>&&</td>
  <td>Logical and</td>
  <td>Returns true if both statements are true</td>
  <td>x < 5 &&  x < 10</td>
</tr>
<tr>
  <td>||</td>
  <td>Logical or</td>
  <td>Returns true if one of the statements is true</td>
  <td>x < 5 || x < 4</td>
</tr>
<tr>
  <td>!</td>
  <td>Logical not</td>
  <td>Reverse the result, returns false if the result is true</td>
  <td>!(x < 5 && x < 10)</td>
</tr>
</table>


<h1 id=sec8>Java Strings</h1>
<h2>Java Strings</h2>
<h3>Java Strings</h3>

Strings are used for storing text.

A String variable contains a collection of characters surrounded by double quotes:

Example<br>
Create a variable of type String and assign it a value:
```java
String greeting = "Hello";
```

<h3>String Length</h3>
A String in Java is actually an object, which contain methods that can perform certain operations on strings. For example, the length of a string can be found with the length() method:

Example
```java
String txt = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
System.out.println("The length of the txt string is: " + txt.length());
```

<h3> More String Methods</h3>
There are many string methods available, for example toUpperCase() and toLowerCase():

Example
```java
String txt = "Hello World";
System.out.println(txt.toUpperCase());   // Outputs "HELLO WORLD"
System.out.println(txt.toLowerCase());   // Outputs "hello world"
```

<h3>Finding a Character in a String</h3>
The indexOf() method returns the index (the position) of the first occurrence of a specified text in a string (including whitespace):

Example
```java
String txt = "Please locate where 'locate' occurs!";
System.out.println(txt.indexOf("locate")); // Outputs 7
```

    Java counts positions from zero.
    0 is the first position in a string, 1 is the second, 2 is the third ...

<h3>Complete String Reference</h3>
For a complete reference of String methods, go to our Java String Methods Reference.<br>
The reference contains descriptions and examples of all string methods.

<h2>Java String Concatenation</h2>
<h3>String Concatenation</h3>
The + operator can be used between strings to combine them. This is called <strong>concatenation</strong>:

Example
```java
String firstName = "John";
String lastName = "Doe";
System.out.println(firstName + " " + lastName);
```

    Note that we have added an empty text (" ") to create a space between firstName and lastName on print.

You can also use the concat() method to concatenate two strings:

Example
```java
String firstName = "John ";
String lastName = "Doe";
System.out.println(firstName.concat(lastName));
```

<h2>Java Numbers and Strings</h2>
<h3>Adding Numbers and Strings</h3>

If you add two numbers, the result will be a number:

Example
```java
int x = 10;
int y = 20;
int z = x + y;  // z will be 30 (an integer/number)
```

If you add two strings, the result will be a string concatenation:

Example
```java
String x = "10";
String y = "20";
String z = x + y;  // z will be 1020 (a String)
```

If you add a number and a string, the result will be a string concatenation:

Example
```java
String x = "10";
int y = 20;
String z = x + y;  // z will be 1020 (a String)
```

<h2>Java Special Characters</h2>
<h3>Strings - Special Characters</h3>
Because strings must be written within quotes, Java will misunderstand this string, and generate an error:

```java
String txt = "We are the so-called "Vikings" from the north.";
```
The solution to avoid this problem, is to use the <strong>backslash escape character.</strong>

The backslash (\) escape character turns special characters into string characters:
<table>
<tr>
  <th>Escape character</th>
  <th>Result</th>
  <th>Description</th>
</tr>
<tr>
  <td>\'</td>
  <td>'</td>
  <td>Single quote</td>
</tr>
<tr>
  <td>\"</td>
  <td>"</td>
  <td>Double quote</td>
</tr>
<tr>
  <td>\\</td>
  <td>\</td>
  <td>Backslash</td>
</tr>
</table>
The sequence \"  inserts a double quote in a string:

Example
```java
String txt = "We are the so-called \"Vikings\" from the north.";
```

The sequence \'  inserts a single quote in a string:

Example
```java
String txt = "It\'s alright.";
```

The sequence \\  inserts a single backslash in a string:

Example
```java
String txt = "The character \\ is called backslash.";
```

Other common escape sequences that are valid in Java are:
<table>
<tr>
  <th>Code</th>
  <th>Result</th>
</tr>
<tr>
  <td>\n</td>
  <td>New Line</td>
</tr>
<tr>
  <td>\r</td>
  <td>Carriage Return</td>
</tr>
<tr>
  <td>\t</td>
  <td>Tab</td>
</tr>
<tr>
  <td>\b</td>
  <td>Backspace</td>
</tr>
<tr>
  <td>\f</td>
  <td>Form Feed</td>
</tr>
</table>


<h1 id=sec9>Java Math</h1>
The Java Math class has many methods that allows you to perform mathematical tasks on numbers.

<h2>Math.max(x,y)</h2>
The Math.max(x,y) method can be used to find the highest value of x and y:

Example
```java
Math.max(5, 10);
```

<h2>Math.min(x,y)</h2>
The Math.min(x,y) method can be used to find the lowest value of x and y:

Example
```java
Math.min(5, 10);
```

<h2>Math.sqrt(x)</h2>
The Math.sqrt(x) method returns the square root of x:

Example
```java
Math.sqrt(64);
```

<h2>Math.abs(x)</h2>
The Math.abs(x) method returns the absolute (positive) value of x:

Example
```java
Math.abs(-4.7);
```

<h2>Random Numbers</h2>
Math.random() returns a random number between 0.0 (inclusive), and 1.0 (exclusive):

Example
```java
Math.random();
```

To get more control over the random number, for example, if you only want a random number between 0 and 100, you can use the following formula:

Example
```java
int randomNum = (int)(Math.random() * 101);  // 0 to 100
```

Complete Math Reference
For a complete reference of Math methods, go to our <a href="https://www.w3schools.com/java/java_ref_math.asp">Java Math Methods Reference</a>.


<h1 id=sec10>Java Booleans</h1>
<h2>Java Booleans</h2>
Very often, in programming, you will need a data type that can only have one of two values, like:

<ul>
  <li>YES / NO</li>
  <li>ON / OFF</li>
  <li>TRUE / FALSE</li>
</ul>

For this, Java has a boolean data type, which can store true or false values.

<h2>Boolean Values</h2>
A boolean type is declared with the boolean keyword and can only take the values true or false:

Example
```java
boolean isJavaFun = true;
boolean isFishTasty = false;
System.out.println(isJavaFun);     // Outputs true
System.out.println(isFishTasty);   // Outputs false
```

However, it is more common to return boolean values from boolean expressions, for conditional testing (see below).

<h2>Boolean Expression</h2>
A Boolean expression returns a boolean value: true or false.<br>
This is useful to build logic, and find answers.

For example, you can use a comparison operator, such as the greater than (>) operator, to find out if an expression (or a variable) is true or false:

Example
```java
int x = 10;
int y = 9;
System.out.println(x > y); // returns true, because 10 is higher than 9
```

Or even easier:

Example
```java
System.out.println(10 > 9); // returns true, because 10 is higher than 9
```

In the examples below, we use the equal to (==) operator to evaluate an expression:

Example
```java
int x = 10;
System.out.println(x == 10); // returns true, because the value of x is equal to 10
```

Example
```java
System.out.println(10 == 15); // returns false, because 10 is not equal to 15
```

<h2>Real Life Example</h2>
Let's think of a "real life example" where we need to find out if a person is old enough to vote.

In the example below, we use the >= comparison operator to find out if the age (25) is greater than OR equal to the voting age limit, which is set to 18:

Example
```java
int myAge = 25;
int votingAge = 18;
System.out.println(myAge >= votingAge);
```

Cool, right? An even better approach (since we are on a roll now), would be to wrap the code above in an if...else statement, so we can perform different actions depending on the result:

Example<br>
Output "Old enough to vote!" if myAge is greater than or equal to 18. Otherwise output "Not old enough to vote.":

```java
int myAge = 25;
int votingAge = 18;

if (myAge >= votingAge) {
  System.out.println("Old enough to vote!");
} else {
  System.out.println("Not old enough to vote.");
}
```

    Booleans are the basis for all Java comparisons and conditions.

    You will learn more about conditions (if...else) in the next chapter.


<h1 id=sec11>Java If ...Else</h1>
<h2>If ...Else</h2>
<h3>Java Conditions and If Statements</h3>
You already know that Java supports the usual logical conditions from mathematics:

<ul>
  <li>Less than: a < b</li>
  <li>Less than or equal to: a <= b</li>
  <li>Greater than: a > b</li>
  <li>Greater than or equal to: a >= b</li>
  <li>Equal to a == b</li>
  <li>Not Equal to: a != b</li>
</ul>
You can use these conditions to perform different actions for different decisions.

Java has the following conditional statements:
<ul>
  <li>Use <strong>if</strong> to specify a block of code to be executed, if a specified condition is true</li>
  <li>Use <strong>else</strong> to specify a block of code to be executed, if the same condition is false</li>
  <li>Use <strong>else if</strong> to specify a new condition to test, if the first condition is false</li>
  <li>Use <strong>switch</strong> to specify many alternative blocks of code to be executed</li>
</ul>

<h3>The if Statement</h3>
Use the if statement to specify a block of Java code to be executed if a condition is true.

Syntax
```java
if (condition) {
  // block of code to be executed if the condition is true
}
```
    Note that if is in lowercase letters. Uppercase letters (If or IF) will generate an error.

In the example below, we test two values to find out if 20 is greater than 18. If the condition is true, print some text:

Example
```java
if (20 > 18) {
  System.out.println("20 is greater than 18");
}
```

We can also test variables:

Example
```java
int x = 20;
int y = 18;
if (x > y) {
  System.out.println("x is greater than y");
}
```

<h4>Example explained</h4>
In the example above we use two variables, <strong>x</strong> and <strong>y</strong>, to test whether x is greater than y (using the > operator). As x is 20, and y is 18, and we know that 20 is greater than 18, we print to the screen that "x is greater than y".

<h3>The else Statement</h3>
Use the else statement to specify a block of code to be executed if the condition is false.

Syntax
```java
if (condition) {
  // block of code to be executed if the condition is true
} else {
  // block of code to be executed if the condition is false
}
```
Example
```java
int time = 20;
if (time < 18) {
  System.out.println("Good day.");
} else {
  System.out.println("Good evening.");
}
// Outputs "Good evening."
```
 

<h4>Example explained</h4>
In the example above, time (20) is greater than 18, so the condition is false. Because of this, we move on to the else condition and print to the screen "Good evening". If the time was less than 18, the program would print "Good day".

<h3>The else if Statement</h3>
Use the else if statement to specify a new condition if the first condition is false.

Syntax
```java
if (condition1) {
  // block of code to be executed if condition1 is true
} else if (condition2) {
  // block of code to be executed if the condition1 is false and condition2 is true
} else {
  // block of code to be executed if the condition1 is false and condition2 is false
}
```
Example
```java
int time = 22;
if (time < 10) {
  System.out.println("Good morning.");
} else if (time < 18) {
  System.out.println("Good day.");
} else {
  System.out.println("Good evening.");
}
// Outputs "Good evening."
```

<h4>Example explained</h4>
In the example above, time (22) is greater than 10, so the first condition is false. The next condition, in the else if statement, is also false, so we move on to the else condition since condition1 and condition2 is both false - and print to the screen "Good evening".

However, if the time was 14, our program would print "Good day."

<h2>Short Hand If ...Else</h2>
There is also a short-hand <a href="https://www.w3schools.com/java/java_conditions.asp">if else</a>, which is known as the ternary operator because it consists of three operands.

It can be used to replace multiple lines of code with a single line, and is most often used to replace simple if else statements:

Syntax
```java
variable = (condition) ? expressionTrue :  expressionFalse;
Instead of writing:
```

Example
```java
int time = 20;
if (time < 18) {
  System.out.println("Good day.");
} else {
  System.out.println("Good evening.");
}
```

You can simply write:

Example
```java
int time = 20;
String result = (time < 18) ? "Good day." : "Good evening.";
System.out.println(result);
```


<h1 id=sec12>Java Switch</h1>
<h2>Java Switch Statements</h2>
Instead of writing <strong>many</strong> if..else statements, you can use the switch statement.

The switch statement selects one of many code blocks to be executed:

Syntax
```java
switch(expression) {
  case x:
    // code block
    break;
  case y:
    // code block
    break;
  default:
    // code block
}
```
This is how it works:
<ul>
  <li>The switch expression is evaluated once.</li>
  <li>The value of the expression is compared with the values of each case.</li>
  <li>If there is a match, the associated block of code is executed.</li>
  <li>The break and default keywords are optional, and will be described later in this chapter</li>
</ul>
The example below uses the weekday number to calculate the weekday name:

Example
```java
int day = 4;
switch (day) {
  case 1:
    System.out.println("Monday");
    break;
  case 2:
    System.out.println("Tuesday");
    break;
  case 3:
    System.out.println("Wednesday");
    break;
  case 4:
    System.out.println("Thursday");
    break;
  case 5:
    System.out.println("Friday");
    break;
  case 6:
    System.out.println("Saturday");
    break;
  case 7:
    System.out.println("Sunday");
    break;
}
// Outputs "Thursday" (day 4)
```

<h2>The break Keyword</h2>
When Java reaches a break keyword, it breaks out of the switch block.

This will stop the execution of more code and case testing inside the block.

When a match is found, and the job is done, it's time for a break. There is no need for more testing.

    A break can save a lot of execution time because it "ignores" the execution of all the rest of the code in the switch block.

<h2>The default Keyword</h2>
The default keyword specifies some code to run if there is no case match:

Example
```java
int day = 4;
switch (day) {
  case 6:
    System.out.println("Today is Saturday");
    break;
  case 7:
    System.out.println("Today is Sunday");
    break;
  default:
    System.out.println("Looking forward to the Weekend");
}
// Outputs "Looking forward to the Weekend"
```
 

Note that if the default statement is used as the last statement in a switch block, it does not need a break.


<h1 id=sec13>Java While Loop</h1>
<h2>Loops</h2>

Loops can execute a block of code as long as a specified condition is reached.

Loops are handy because they save time, reduce errors, and they make code more readable.

<h2>Java While Loop</h2>
The while loop loops through a block of code as long as a specified condition is true:

Syntax
```java
while (condition) {
  // code block to be executed
}
```
In the example below, the code in the loop will run, over and over again, as long as a variable (i) is less than 5:

Example
```java
int i = 0;
while (i < 5) {
  System.out.println(i);
  i++;
}
```

    Note: Do not forget to increase the variable used in the condition, otherwise the loop will never end!

<h2>The Do/While Loop</h2>
The do/while loop is a variant of the while loop. This loop will execute the code block once, before checking if the condition is true, then it will repeat the loop as long as the condition is true.

Syntax
```java
do {
  // code block to be executed
}
while (condition);
```
The example below uses a do/while loop. The loop will always be executed at least once, even if the condition is false, because the code block is executed before the condition is tested:

Example
```java
int i = 0;
do {
  System.out.println(i);
  i++;
}
while (i < 5);
```

    Do not forget to increase the variable used in the condition, otherwise the loop will never end!


<h1 id=sec14>Java For Loop</h1>
<h2>Java For Loop</h2>
When you know exactly how many times you want to loop through a block of code, use the for loop instead of a while loop:

SyntaxGet your own Java Server
for (statement 1; statement 2; statement 3) {
  // code block to be executed
}
Statement 1 is executed (one time) before the execution of the code block.

Statement 2 defines the condition for executing the code block.

Statement 3 is executed (every time) after the code block has been executed.

The example below will print the numbers 0 to 4:

Example
for (int i = 0; i < 5; i++) {
  System.out.println(i);
}

Example explained
Statement 1 sets a variable before the loop starts (int i = 0).

Statement 2 defines the condition for the loop to run (i must be less than 5). If the condition is true, the loop will start over again, if it is false, the loop will end.

Statement 3 increases a value (i++) each time the code block in the loop has been executed.

Another Example
This example will only print even values between 0 and 10:

Example
for (int i = 0; i <= 10; i = i + 2) {
  System.out.println(i);
}

Nested Loops
It is also possible to place a loop inside another loop. This is called a nested loop.

The "inner loop" will be executed one time for each iteration of the "outer loop":

Example
// Outer loop
for (int i = 1; i <= 2; i++) {
  System.out.println("Outer: " + i); // Executes 2 times
  
  // Inner loop
  for (int j = 1; j <= 3; j++) {
    System.out.println(" Inner: " + j); // Executes 6 times (2 * 3)
  }
} 