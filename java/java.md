<h1> Java</h1> 

<h1>Contents</h1>
<ol>
  <li><a href="#sec1">Java Syntax</a></li>
  <li><a href="#sec2">Java Output</a></li>
  <li><a href="#">Java Comments</a></li>
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
In this tutorial, we will only use println() as it makes it easier to read the output of code.