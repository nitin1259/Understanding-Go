# Chapter 2. which code runs next?: Conditionals and Loops

**Every program has parts that apply only in certain situations**. “This code should run if there’s an error. Otherwise, that other code should run.” Almost every program contains code that should be run only when a certain `condition` is true. So almost every programming language provides **conditional statements** that let you determine whether to run segments of code. Go is no exception.

You may also need some parts of your code to run repeatedly. Like most languages, Go provides **loops** that run sections of code more than once. We’ll learn to use both conditionals and loops in this chapter!

## Calling methods

In Go, it’s possible to define **methods**: functions that are associated with values of a given type. Go methods are kind of like the methods that you may have seen attached to “objects” in other languages, but they’re a bit simpler.

We’ll be taking a detailed look at how methods work in Chapter 9. But we need to use a couple methods to make our examples for this chapter work, so let’s look at some brief examples of calling methods now.

The `time` package has a Time type that represents a date (year, month, and day) and time (hour, minute, second, etc.). Each time.Time value has a Year method that returns the year. The code below uses this method to print the current year:

image
The `time.Now` function returns a new `Time` value for the current date and time, which we store in the now variable. Then, we call the Year method on the value that now refers to:

image
The Year method returns an integer with the year, which we then print.

**Methods are functions that are associated with values of a particular type.**

The strings package has a Replacer type that can search through a string for a substring, and replace each occurrence of that substring with another string. The code below replaces every # symbol in a string with the letter o:

image
The strings.NewReplacer function takes arguments with a string to replace ("#"), and a string to replace it with ("o"), and returns a strings.Replacer. When we pass a string to the Replacer value’s Replace method, it returns a string with those replacements made.

image
**The dot indicates that the thing on its right belongs to the thing on its left.**

Whereas the functions we saw earlier belonged to a `package`, the methods belong to an individual `value`. That value is what appears to the left of the dot.

image
Making the grade
In this chapter, we’re going to look at features of Go that let you decide whether to run some code or not, based on a condition. Let’s look at a situation where we might need that ability...

We need to write a program that allows a student to type in their percentage grade and tells them whether they passed or not. Passing or failing follows a simple formula: a grade of 60% or more is passing, and less than 60% is failing. So our program will need to give one response if the percentage users enter is 60 or greater, and a different response otherwise.

Comments
Let’s create a new file, pass_fail.go, to hold our program. We’re going to take care of a detail we omitted in our previous programs, and add a description of what the program does at the top.

image
Most Go programs include descriptions in their source code of what they do, intended for people maintaining the program to read. These comments are ignored by the compiler.

The most common form of comment is marked with two slash characters (//). Everything from the slashes to the end of the line is treated as part of the comment. A // comment can appear on a line by itself, or following a line of code.

// The total number of widgets in the system.
var TotalCount int // Can only be a whole number.
The less frequently used form of comments, block comments, spans multiple lines. Block comments start with /_ and end with _/, and everything between those markers (including newlines) is part of the comment.

/_
Package widget includes all the functions used
for processing widgets.
_/
Getting a grade from the user
Now let’s add some actual code to our pass_fail.go program. The first thing it needs to do is allow the user to input a percentage grade. We want them to type a number and press Enter, and we’ll store the number they typed in a variable. Let’s add code to handle this. (Note: this code will not actually compile as shown; we’ll talk about the reason in a moment!)

image
First, we need to let the user know to enter something, so we use the fmt.Print function to display a prompt. (Unlike the Println function, Print doesn’t skip to a new terminal line after printing a message, which lets us keep the prompt and the user’s entry on the same line.)

Next, we need a way to read (receive and store) input from the program’s standard input, which all keyboard input goes to. The line reader := bufio.NewReader(os.Stdin) stores a bufio.Reader in the reader variable that can do that for us.

image
To actually get the user’s input, we call the ReadString method on the Reader. The ReadString method requires an argument with a rune (character) that marks the end of the input. We want to read everything the user types up until they press Enter, so we give ReadString a newline rune.

Once we have the user input, we simply print it.

That’s the plan, anyway. But if we try to compile or run this program, we’ll get an error:

image
image RELAX
**Don’t worry too much about the details of how bufio.Reader works.**

All you really need to know at this point is that it lets us read input from the keyboard.

Multiple return values from a function or method
We’re trying to read the user’s keyboard input, but we’re getting an error. The compiler is reporting a problem in this line of code:

image
The problem is that the ReadString method is trying to return two values, and we’ve only provided one variable to assign a value to.

In most programming languages, functions and methods can only have a single return value, but in Go, they can return any number of values. The most common use of multiple return values in Go is to return an additional error value that can be consulted to find out if anything went wrong while the function or method was running. A few examples:

image
**Go doesn’t allow us to declare a variable unless we use it.**

Go requires that every variable that gets declared must also get used somewhere in your program. If we add an err variable and then don’t check it, our code won’t compile. Unused variables often indicate a bug, so this is an example of Go helping you detect and fix bugs!

image
Option 1: Ignore the error return value with the blank identifier
The ReadString method returns a second value along with the user’s input, and we need to do something with that second value. We’ve tried just adding a second variable and ignoring it, but our code still won’t compile.

image
When we have a value that would normally be assigned to a variable, but that we don’t intend to use, we can use Go’s blank identifier. Assigning a value to the blank identifier essentially discards it (while making it obvious to others reading your code that you are doing so). To use the blank identifier, simply type a single underscore ( \_ ) character in an assignment statement where you would normally type a variable name.

Let’s try using the blank identifier in place of our old err variable:

image
Now we’ll try the change out. In your terminal, change to the directory where you saved pass_fail.go, and run the program with:

image
When you type a grade (or any other string) at the prompt and press Enter, your entry will be echoed back to you. Our program is working!

Option 2: Handle the error
image
**That’s true. If an error actually occurred, this program wouldn’t tell us!**

If we got an error back from the ReadString method, the blank identifier would just cause the error to be ignored, and our program would proceed anyway, possibly with invalid data.

image
In this case, it would be more appropriate to alert the user and stop the program if there was an error.

The log package has a Fatal function that can do both of these operations for us at once: log a message to the terminal and stop the program. (“Fatal” in this context means reporting an error that “kills” your program.)

Let’s get rid of the blank identifier and replace it with an err variable so that we’re recording the error again. Then, we’ll use the Fatal function to log the error and halt the program.

image
But if we try running this updated program, we’ll see there’s a new problem...

Conditionals
If our program encounters a problem reading input from the keyboard, we’ve set it up to report the error and stop running. But now, it stops running even when everything’s working correctly!

image
Functions and methods like ReadString return an error value of nil, which basically means “there’s nothing there.” In other words, if err is nil, it means there was no error. But our program is set up to simply report the nil error! What we should do is exit the program only if the err variable has a value other than nil.

We can do this using conditionals: statements that cause a block of code (one or more statements surrounded by {} curly braces) to be executed only if a condition is met.

image
An expression is evaluated, and if its result is true, the code in the conditional block body is executed. If it’s false, the conditional block is skipped.

image
As with most other languages, Go supports multiple branches in the conditional. These statements take the form if...else if...else.

if grade == 100 {
fmt.Println("Perfect!")
} else if grade >= 60 {
fmt.Println("You pass.")
} else {
fmt.Println("You fail!")
}
Conditionals rely on a Boolean expression (one that evaluates to true or false) to decide whether the code they contain should be executed.

image
When you need to execute code only if a condition is false, you can use !, the Boolean negation operator, which lets you take a true value and make it false, or a false value and make it true.

image
If you want to run some code only if two conditions are both true, you can use the && (“and”) operator. If you want it to run if either of two conditions is true, you can use the || (“or”) operator.

image
there are no Dumb Questions
Q: My other programming language requires that an if statement’s condition be surrounded with parentheses. Doesn’t Go?

A: No, and in fact the go fmt tool will remove any parentheses you add, unless you’re using them to set order of operations.

image EXERCISE
Because they’re in conditional blocks, only some of the Println calls in the code below will be executed. Write down what the output would be.

image
image Answers in “image Exercise Solution”.

Logging a fatal error, conditionally
Our grading program is reporting an error and exiting, even if it reads input from the keyboard successfully.

image
We know that if the value in our err variable is nil, it means reading from the keyboard was successful. Now that we know about if statements, let’s try updating our code to log an error and exit only if err is not nil.

image
If we rerun our program, we’ll see that it’s working again. And now, if there are any errors when reading user input, we’ll see those as well!

image
Code Magnets
image
A Go program that prints the size of a file is on the fridge. It calls the os.Stat function, which returns an os.FileInfo value, and possibly an error value. Then it calls the Size method on the FileInfo value to get the file size.

But the original program uses the \_ blank identifier to ignore the error value from os.Stat. If an error occurs (which could happen if the file doesn’t exist), this will cause the program to fail.

Reconstruct the extra code snippets to make a program that works just like the original one, but also checks for an error from os.Stat. If the error from os.Stat is not nil, the error should be reported, and the program should exit. Discard the magnet with the \_ blank identifier; it won’t be used in the finished program.

image
image Answers in “Code Magnets Solution”.

Avoid shadowing names
image
fmt.Print("Enter a grade: ")
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
if err != nil {
log.Fatal(err)
}
Naming a variable error would be a bad idea, because it would shadow the name of a type called error.

When you declare a variable, you should make sure it doesn’t have the same name as any existing functions, packages, types, or other variables. If something by the same name exists in the enclosing scope (we’ll talk about scopes shortly), your variable will shadow it—that is, take precedence over it. And all too often, that’s a bad thing.

Here, we declare a variable named int that shadows a type name, a variable named append that shadows a built-in function name (we’ll see the append function in Chapter 6), and a variable named fmt that shadows an imported package name. Those names are awkward, but they don’t cause any errors by themselves...

image
image
...But if we try to access the type, function, or package the variables are shadowing, we’ll get the value in the variable instead. In this case, it results in compile errors:

image
To avoid confusion for yourself and your fellow developers, you should avoid shadowing names wherever possible. In this case, fixing the issue is as simple as choosing nonconflicting names for the variables:

image
As we’ll see in Chapter 3, Go has a built-in type named error. So that’s why, when declaring variables meant to hold errors, we’ve been naming them err instead of error—we want to avoid shadowing the name of the error type with our variable name.

image
If you do name your variables error, your code will probably still work. That is, until you forget that the error type name is shadowed, you try to use the type, and you get the variable instead. Don’t take that chance; use the name err for your error variables!

Converting strings to numbers
Conditional statements will also let us evaluate the entered grade. Let’s add an if/else statement to determine whether the grade is passing or failing. If the entered percentage grade is 60 or greater, we’ll set the status to "passing". Otherwise, we’ll set it to "failing".

// package and import statements omitted
func main() {
fmt.Print("Enter a grade: ")
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
if err != nil {
log.Fatal(err)
}
if input >= 60 {
status := "passing"
} else {
status := "failing"
}
}
In its current form, though, this gets us a compilation error.

image
Here’s the problem: input from the keyboard is read in as a string. Go can only compare numbers to other numbers; we can’t compare a number with a string. And there’s no direct type conversion from string to a number:

image
We have a pair of issues to address here:

The input string still has a newline character on the end, from when the user pressed the Enter key while entering it. We need to strip that off.

The remainder of the string needs to be converted to a floating-point number.

Removing the newline character from the end of the input string will be easy. The strings package has a TrimSpace function that will remove all whitespace characters (newlines, tabs, and regular spaces) from the start and end of a string.

image
So, we can get rid of the newline on input by passing it to TrimSpace, and assigning the return value back to the input variable.

input = strings.TrimSpace(input)
All that should remain in the input string now is the number the user entered. We can use the strconv package’s ParseFloat function to convert it to a float64 value.

image
You pass ParseFloat a string that you want to convert to a number, as well as the number of bits of precision the result should have. Since we’re converting to a float64 value, we pass the number 64. (In addition to float64, Go offers a less precise float32 type, but you shouldn’t use that unless you have a good reason.)

ParseFloat converts the string to a number, and returns it as a float64 value. Like ReadString, it also has a second return value, an error, which will be nil unless there was some problem converting the string. (For example, a string that can’t be converted to a number. We don’t know of a numeric equivalent to "hello"...)

image RELAX
This whole “bits of precision” thing isn’t that important right now.

It’s basically just a measure of how much computer memory a floating-point number takes up. As long as you know that you want a float64, and so you should pass 64 as the second argument to ParseFloat, you’ll be fine.

Let’s update pass_fail.go with calls to TrimSpace and ParseFloat:

image
First, we add the appropriate packages to the import section. We add code to remove the newline character from the input string. Then we pass input to ParseFloat, and store the resulting float64 value in a new variable, grade.

Just as we did with ReadString, we test whether ParseFloat returns an error value. If it does, we report it and stop the program.

Finally, we update the conditional statement to test the number in grade, rather than the string in input. That should fix the error stemming from comparing a string to a number.

If we try to run the updated program, we no longer get the mismatched types string and int error. So it looks like we’ve fixed that issue. But we’ve got a couple more errors to address. We’ll look at those next.

image
Blocks
We’ve converted the user’s grade input to a float64 value, and added it to a conditional to determine if it’s passing or failing. But we’re getting a couple more compile errors:

image
As we’ve seen previously, declaring a variable like status without using it afterward is an error in Go. It seems a little strange that we’re getting the error twice, but let’s disregard that for now. We’ll add a call to Println to print the percentage grade we were given, and the value of status.

image
But now we get a new error, saying that the status variable is undefined when we attempt to use it in our Println statement! What’s going on?

Go code can be divided up into blocks, segments of code. Blocks are usually surrounded by curly braces ({}), although there are also blocks at the source code file and package levels. Blocks can be nested inside one another.

image
The bodies of functions and conditionals are both blocks as well. Understanding this will be key to solving our problem with the status variable...

Blocks and variable scope
Each variable you declare has a scope: a portion of your code that it’s “visible” within. A declared variable can be accessed anywhere within its scope, but if you try to access it outside that scope, you’ll get an error.

A variable’s scope consists of the block it’s declared in and any blocks nested within that block.

image
Here are the scopes of the variables in the code above:

The scope of packageVar is the entire main package. You can access packageVar anywhere within any function you define in the package.

The scope of functionVar is the entire function it’s declared in, including the if block nested within that function.

The scope of conditionalVar is limited to the if block. When we try to access conditionalVar after the closing } brace of the if block, we’ll get an error saying that conditionalVar is undefined!

Now that we understand variable scope, we can explain why our status variable was undefined in the grading program. We declared status in our conditional blocks. (In fact, we declared it twice, since there are two separate blocks. That’s why we got two status declared and not used errors.) But then we tried to access status outside those blocks, where it was no longer in scope.

image
The solution is to move the declaration of the status variable out of the conditional blocks, and up to the function block. Once we do that, the status variable will be in scope both within the nested conditional blocks and at the end of the function block.

image
image WATCH IT!
Don’t forget to change the short variable declarations within the nested blocks to assignment statements!

If you don’t change both occurrences of := to =, you’ll accidentally create new variables named status within the nested conditional blocks, which will then be out of scope at the end of the enclosing function block!

We’ve finished the grading program!
That was it! Our pass_fail.go program is ready for action! Let’s take one more look at the complete code:

image
You can try running the finished program as many times as you like. Enter a percentage grade under 60, and it will report a failing status. Enter a grade over 60, and it will report that it’s passing. Looks like everything’s working!

image
image EXERCISE
Some of the lines of code below will result in a compile error, because they refer to a variable that is out of scope. Cross out the lines that have errors.

image
image Answers in “image Exercise Solution”.

Only one variable in a short variable declaration has to be new
image
It’s true that when the same variable name is declared twice in the same scope, we get a compile error:

image
But as long as at least one variable name in a short variable declaration is new, it’s allowed. The new variable names are treated as a declaration, and the existing names are treated as an assignment.

image
There’s a reason for this special handling: a lot of Go functions return multiple values. It would be a pain if you had to declare all the variables separately just because you want to reuse one of them.

image
Instead, Go lets you use a short variable declaration for everything, even if for one of the variables, it’s actually an assignment.

image
Let’s build a game
We’re going to wrap up this chapter by building a simple game. If that sounds daunting, don’t worry; you’ve already learned most of the skills you’re going to need! Along the way, we’ll learn about loops, which will allow the player to take multiple turns.

Let’s look at everything we’ll need to do:

image
NOTE
This example debuted in Head First Ruby. (Another fine book that you should also buy!) It worked so well that we’re using it again here.

image
Figure 2-1. Gary Richardott Game Designer
Let’s create a new source file, named guess.go.

It looks like our first requirement is to generate a random number. Let’s get started!

Package names vs. import paths
The math/rand package has a Intn function that can generate a random number for us, so we’ll need to import math/rand. Then we’ll call rand.Intn to generate the random number.

image
image
One is the package’s import path, and the other is the package’s name.

When we say math/rand we’re referring to the package’s import path, not its name. An import path is just a unique string that identifies a package and that you use in an import statement. Once you’ve imported the package, you can refer to it by its package name.

For every package we’ve used so far, the import path has been identical to the package name. Here are a few examples:

Import path Package name
"fmt" fmt
"log" log
"strings" strings
But the import path and package name don’t have to be identical. Many Go packages fall into similar categories, like compression or complex math. So they’re grouped together under similar import path prefixes, such as "archive/" or "math/". (Think of them as being similar to the paths of directories on your hard drive.)

Import path Package name
"archive" archive
"archive/tar" tar
"archive/zip" zip
"math" math
"math/cmplx" cmplx
"math/rand" rand
The Go language doesn’t require that a package name have anything to do with its import path. But by convention, the last (or only) segment of the import path is also used as the package name. So if the import path is "archive", the package name will be archive, and if the import path is "archive/zip", the package name will be zip.

Import path Package name
"archive" archive
"archive/tar" tar
"archive/zip" zip
"math" math
"math/cmplx" cmplx
"math/rand" rand
So, that’s why our import statement uses a path of "math/rand", but our main function just uses the package name: rand.

image
Generating a random number
Pass a number to rand.Intn, and it will return a random integer between 0 and the number you provided. In other words, if we pass an argument of 100, we’ll get a random number in the range 0–99. Since we need a number in the range 1–100, we’ll just add 1 to whatever random value we get. We’ll store the result in a variable, target. We’ll do more with target later, but for now we’ll just print it.

image
If we try running our program right now, we’ll get a random number. But we just get the same random number over and over! The problem is, random numbers generated by computers aren’t really that random. But there’s a way to increase that randomness...

image
To get different random numbers, we need to pass a value to the rand.Seed function. That will “seed” the random number generator—that is, give it a value that it will use to generate other random values. But if we keep giving it the same seed value, it will keep giving us the same random values, and we’ll be back where we started.

We saw earlier that the time.Now function will give us a Time value representing the current date and time. We can use that to get a different seed value every time we run our program.

image
The rand.Seed function expects an integer, so we can’t pass it a Time value directly. Instead, we call the Unix method on the Time, which will convert it to an integer. (Specifically, it will convert it to Unix time format, which is an integer with the number of seconds since January 1, 1970. But you don’t really need to remember that.) We pass that integer to rand.Seed.

We also add a couple Println calls to let the user know we’ve chosen a random number. But aside from that, we can leave the rest of our code, including the call to rand.Intn, as is. Seeding the generator should be the only change we need to make.

Now, each time we run our program, we’ll see our message, along with a random number. It looks like our changes are successful!

image
Getting an integer from the keyboard
Our first requirement is complete! Next we need to get the user’s guess via the keyboard.

That should work in much the same way as when we read in a percentage grade from the keyboard for our grading program.

image
There will be only one difference: instead of converting the input to a float64, we need to convert it to an int (since our guessing game uses only whole numbers). So we’ll pass the string read from the keyboard to the strconv package’s Atoi (string to integer) function instead of its ParseFloat function. Atoi will give us an integer as its return value. (Just like ParseFloat, Atoi might also give us an error if it can’t convert the string. If that happens, we again report the error and exit.)

image
Comparing the guess to the target
Another requirement finished. And this next one will be easy... We just need to compare the user’s guess to the randomly generated number, and tell them whether it was higher or lower.

image
If guess is less than target, we need to print a message saying the guess was low. Otherwise, if guess is greater than target, we should print a message saying the guess was high. Sounds like we need an if...else if statement. We’ll add it below the other code in our main function.

image
Now try running our updated program from the terminal. It’s still set up to print target each time it runs, which will be useful for debugging. Just enter a number lower than target, and you should be told your guess was low. If you rerun the program, you’ll get a new target value. Enter a number higher than that, and you’ll be told your guess was high.

image
Loops
Another requirement down! Let’s look at the next one.

image
Currently, the player only gets to guess once, but we need to allow them to guess up to 10 times.

The code to prompt for a guess is already in place. We just need to run it more than once. We can use a loop to execute a block of code repeatedly. If you’ve used other programming languages, you’ve probably encountered loops. When you need one or more statements executed over and over, you place them inside a loop.

image
Loops always begin with the for keyword. In one common kind of loop, for is followed by three segments of code that control the loop:

An initialization (or init) statement that is usually used to initialize a variable

A condition expression that determines when to break out of the loop

A post statement that runs after each iteration of the loop

Often, the initialization statement is used to initialize a variable, the condition expression keeps the loop running until that variable reaches a certain value, and the post statement is used to update the value of that variable. For example, in this snippet, the t variable is initialized to 3, the condition keeps the loop going while t > 0, and the post statement subtracts 1 from t each time the loop runs. Eventually, t reaches 0 and the loop ends.

image
The ++ and -- statements are frequently used in loop post statements. Each time they’re evaluated, ++ adds 1 to a variable’s value, and -- subtracts 1.

image
Used in a loop, ++ and -- are convenient for counting up or down.

image
Go also includes the assignment operators += and -=. They take the value in a variable, add or subtract another value, and then assign the result back to the variable.

image
+= and -= can be used in a loop to count in increments other than 1.

image
When the loop finishes, execution will resume with whatever statement follows the loop block. But the loop will keep going as long as the condition expression evaluates to true. It’s possible to abuse this; here are examples of a loop that will run forever, and a loop that will never run at all:

image
image WATCH IT!
It’s possible for a loop to run forever, in which case your program will never stop on its own.

If this happens, with the terminal active, hold the Control key and press C to halt your program.

Init and post statements are optional
If you want, you can leave out the init and post statements from a for loop, leaving only the condition expression (although you still need to make sure the condition eventually evaluates to false, or you could have an infinite loop on your hands).

image
Loops and scope
Just like with conditionals, the scope of any variables declared within a loop’s block is limited to that block (although the init statement, condition expression, and post statement can be considered part of that scope as well).

image
Also as with conditionals, any variable declared before the loop will still be in scope within the loop’s control statements and block, and will still be in scope after the loop exits.

image
Breaking Stuff is Educational!
image
Here’s a program that uses a loop to count to 3. Try making one of the changes below and run it. Then undo your change and try the next one. See what happens!

image
If you do this... ...it will break because...
Add parentheses after the for keyword
for (x := 1; x <= 3; x++) Some other languages require parentheses around a for loop’s control statements, but not only does Go not require them, it doesn’t allow them.
Delete the : from the init statement
x = 1 Unless you’re assigning to a variable that’s already been declared in the enclosing scope (which you usually won’t be), the init statement needs to be a declaration, not an assignment.
Remove the = from the condition expression
x < 3 The expression x < 3 becomes false when x reaches 3 (whereas x <= 3 would still be true). So the loop would only count to 2.
Reverse the comparison in the condition expression
x >= 3 Because the condition is already false when the loop begins (x is initialized to 1, which is less than 3), the loop will never run.
Change the post statement from x++ to x--
x-- The x variable will start counting down from 1 (1, 0, -1, -2, etc.), and since it will never be greater than 3, the loop will never end.
Move the fmt.Println(x) statement outside the loop’s block Variables declared in the init statement or within the loop block are only in scope within the loop block.
image EXERCISE
Look carefully at the init statement, condition expression, and post statement for each of these loops. Then write what you think the output will be for each one.

NOTE
(We’ve done the first one for you.)

image
image Answers in “image Exercise Solution”.

Using a loop in our guessing game
Our game still only prompts the user for a guess once. Let’s add a loop around the code that prompts the user for a guess and tells them if it was low or high, so that the user can guess 10 times.

We’ll use an int variable named guesses to track the number of guesses the player has made. In our loop’s init statement, we’ll initialize guesses to 0. We’ll add 1 to guesses with each iteration of the loop, and we’ll stop the loop when guesses reaches 10.

We’ll also add a Println statement at the top of the loop’s block to tell the user how many guesses they have left.

image
Now that our loop is in place, if we run our game again, we’ll get asked 10 times what our guess is!

image
Since the code to prompt for a guess and state whether it was high or low is inside the loop, it gets run repeatedly. After 10 guesses, the loop (and the game) will end.

But the loop always runs 10 times, even if the player guesses correctly! Fixing that will be our next requirement.

Skipping parts of a loop with “continue” and “break”
The hard part is done! We only have a couple requirements left to go.

Right now, the loop that prompts the user for a guess always runs 10 times. Even if the player guesses correctly, we don’t tell them so, and we don’t stop the loop. Our next task is to fix that.

image
Go provides two keywords that control the flow of a loop. The first, continue, immediately skips to the next iteration of a loop, without running any further code in the loop block.

image
In the above example, the string "after continue" never gets printed, because the continue keyword always skips back to the top of the loop before the second call to Println can be run.

The second keyword, break, immediately breaks out of a loop. No further code within the loop block is executed, and no further iterations are run. Execution moves to the first statement following the loop.

image
Here, in the first iteration of the loop, the string "before break" gets printed, but then the break statement immediately breaks out of the loop, without printing the "after break" string, and without running the loop again (even though it normally would have run two more times). Execution instead moves to the statement following the loop.

The break keyword seems like it would be applicable to our current problem: we need to break out of our loop when the player guesses correctly. Let’s try using it in our game...

Breaking out of our guessing loop
We’re using an if...else if conditional to tell the player the status of their guess. If the player guesses a number too high or too low, we currently print a message telling them so.

It stands to reason that if the guess is neither too high nor too low, it must be correct. So let’s add an else branch onto the conditional, that will run in the event of a correct guess. Inside the block for the else branch, we’ll tell the player they were right, and then use the break statement to stop the guessing loop.

image
Now, when the player guesses correctly, they’ll see a congratulatory message, and the loop will exit without repeating the full 10 times.

image
That’s another requirement complete!

Revealing the target
image
We’re so close! Just one more requirement left!

If the player makes 10 guesses without finding the target number, the loop will exit. In that event, we need to print a message saying they lost, and tell them what the target was.

But we also exit the loop if the player guesses correctly. We don’t want to say the player has lost when they’ve already won!

So, before our guessing loop, we’ll declare a success variable that holds a bool. (We need to declare it before the loop so that it’s still in scope after the loop ends.) We’ll initialize success to a default value of false. Then, if the player guesses correctly, we’ll set success to true, indicating we don’t need to print the failure message.

image
After the loop, we add an if block that prints the failure message. But an if block only runs if its condition evaluates to true, and we only want to print the failure message if success is false. So we add the Boolean negation operator (!). As we saw earlier, ! turns true values false and false values true.

The result is that the failure message will be printed if success is false, but won’t be printed if success is true.

The finishing touches
image
Congratulations, that’s the last requirement!

Let’s take care of a couple final issues with our code, and then try out our game!

First, as we mentioned, it’s typical to add a comment at the top of each Go program describing what it does. Let’s add one now.

image
Our program is also encouraging cheaters by printing the target number at the start of every game. Let’s remove the Println call that does that.

image
We’re finally ready to try running our complete code!

First, we’ll run out of guesses on purpose to ensure the target number gets displayed...

image
Then we’ll try guessing successfully.

Our game is working great!

image
Congratulations, your game is complete!
image
Using conditionals and loops, you’ve written a complete game in Go! Pour yourself a cold drink—you’ve earned it!

Here’s our complete guess.go source code!

image
Your Go Toolbox
image
That’s it for Chapter 2! You’ve added conditionals and loops to your toolbox.

image
NOTE
Loops

Loops cause a block of code to execute repeatedly.

One common kind of loop starts with the keyword “for”, followed by an init statement that initializes a variable, a condition expression that determines when to break out of the loop, and a post statement that runs after each iteration of the loop.

BULLET POINTS
A method is a kind of function that’s associated with values of a given type.

Go treats everything from a // marker to the end of the line as a comment—and ignores it.

Multiline comments start with /_ and end with _/. Everything in between, including newlines, is ignored.

It’s conventional to include a comment at the top of every program, explaining what it does.

Unlike most programming languages, Go allows multiple return values from a function or method call.

One common use of multiple return values is to return the function’s main result, and then a second value indicating whether there was an error.

To discard a value without using it, use the \_ blank identifier. The blank identifier can be used in place of any variable in any assignment statement.

Avoid giving variables the same name as types, functions, or packages; it causes the variable to shadow (override) the item with the same name.

Functions, conditionals, and loops all have blocks of code that appear within {} braces.

Their code doesn’t appear within {} braces, but files and packages also comprise blocks.

The scope of a variable is limited to the block it is defined within, and all blocks nested within that block.

In addition to a name, a package may have an import path that is required when it is imported.

The continue keyword skips to the next iteration of a loop.

The break keyword exits out of a loop entirely.
