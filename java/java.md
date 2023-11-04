<h1> Java Tutorials</h1> 

<h1>Contents</h1>
<ol>
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