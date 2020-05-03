# Chapter 1. let’s get going: Syntax Basics

## Ready, set, Go!

Back in 2007, the search engine Google had a problem. They had to maintain programs with millions of lines of code. Before they could test new changes, they had to compile the code into a runnable form, a process which at the time took the better part of an hour. Needless to say, this was bad for developer productivity.

So Google engineers Robert Griesemer, Rob Pike, and Ken Thompson sketched out some goals for a new language:

- Fast compilation
- Less cumbersome code
- Unused memory freed automatically (garbage collection)
- Easy-to-write software that does several operations simultaneously (concurrency)
- Good support for processors with multiple cores

After a couple years of work, Google had created Go: a language that was fast to write code for and produced programs that were fast to compile and run. The project switched to an open source license in 2009. It’s now free for anyone to use. And you should use it! Go is rapidly gaining popularity thanks to its simplicity and power.

If you’re writing a command-line tool, Go can produce executable files for Windows, macOS, and Linux, all from the same source code. If you’re writing a web server, it can help you handle many users connecting at once. And no matter what you’re writing, it will help you ensure that your code is easier to maintain and add to.

Ready to learn more? Let’s Go!

## The Go Playground

The easiest way to try Go is to visit https://play.golang.org in your web browser. There, the Go team has set up a simple editor where you can enter Go code and run it on their servers. The result is displayed right there in your browser.

![Play Goloang](./../images/chapter1/f0003-01.png)

(Of course, this only works if you have a stable internet connection. If you don’t, see [“Installing Go on your computer”](https://golang.org/doc/install#install) to learn how to download and run the Go compiler directly on your computer. Then run the following examples using the compiler instead.)

### Let’s try it out now!

![Play Goloang](./../images/chapter1/f0003-02.png)

1. Open https://play.golang.org in your browser. (Don’t worry if what you see doesn’t quite match the screenshot; it just means they’ve improved the site since this book was printed!)

2. Delete any code that’s in the editing area, and type this instead:

```go
    package main
    import "fmt"
    func main() {
        fmt.Println("Hello, Go!")
    }
```

> NOTE - Don’t worry, we’ll explain what all this means on the next page!

3. Click the Format button, which will automatically reformat your code according to Go conventions.

4. Click the Run button.

You should see “Hello, Go!” displayed at the bottom of the screen. Congratulations, you’ve just run your first Go program!

![Play Goloang](./../images/chapter1/f0003-03.png)

## What does it all mean?

You’ve just run your first Go program! Now let’s look at the code and figure out what it actually means...

Every Go file starts with a package clause. A **package** is a collection of code that all does similar things, like formatting strings or drawing images. The package clause gives the name of the package that this file’s code will become a part of. In this case, we use the special package main, which is required if this code is going to be run directly (usually from the terminal).

Next, Go files almost always have one or more import statements. Each file needs to **import** other packages before its code can use the code those other packages contain. Loading all the Go code on your computer at once would result in a big, slow program, so instead you specify only the packages you need by importing them.

![Play Goloang](./../images/chapter1/f0004-01.png)

The last part of every Go file is the actual code, which is often split up into one or more functions. A function is a group of one or more lines of code that you can call (run) from other places in your program. When a Go program is run, it looks for a function named main and runs that first, which is why we named this function `main`.

**Don’t worry if you don’t understand all this right now!**
We’ll look at everything in more detail in the next few pages.

## The typical Go file layout

You’ll quickly get used to seeing these three sections, in this order, in almost every Go file you work with:

1. The package clause
2. Any import statements
3. The actual code
   ![Play Goloang](./../images/chapter1/f0004-02.png)

The saying goes, “a place for everything, and everything in its place.” Go is a very _consistent_ language. This is a good thing: you’ll often find you just _know_ where to look in your project for a given piece of code, without having to think about it!

## there are no Dumb Questions

**Q: My other programming language requires that each statement end with a semicolon. Doesn’t Go?**

**A:** You can use semicolons to separate statements in Go, but it’s not required (in fact, it’s generally frowned upon).

**Q: What’s this Format button? Why did we click that before running our code?**

**A:** The Go compiler comes with a standard formatting tool, called go fmt. The Format button is the web version of go fmt.

Whenever you share your code, other Go developers will expect it to be in the standard Go format. That means that things like indentation and spacing will be formatted in a standard way, making it easier for everyone to read. Where other languages achieve this by relying on people manually reformatting their code to conform to a style guide, with Go all you have to do is run go fmt, and it will automatically fix everything for you.

We ran the formatter on every example we created for this book, and you should run it on all your code, too!

## What if something goes wrong?

Go programs have to follow certain rules to avoid confusing the compiler. If we break one of these rules, we’ll get an error message.

Suppose we forgot to add parentheses on our call to the Println function on line 6.

If we try to run this version of the program, we get an error: f0005-01
![Play Goloang](./../images/chapter1/f0005-01.png)
![Play Goloang](./../images/chapter1/f0005-02.png)

Go tells us which source code file and line number we need to go to so we can fix the problem. (The Go Playground saves your code to a temporary file before running it, which is where the _prog.go_ filename comes from.) Then it gives a description of the error. In this case, because we deleted the parentheses, Go can’t tell we’re trying to call the Println function, so it can’t understand why we’re putting "Hello, Go" at the end of line 6.

## Breaking Stuff is Educational!

![Play Goloang](./../images/chapter1/f0006-01.png)

We can get a feel for the rules Go programs have to follow by intentionally breaking our program in various ways. Take this code sample, try making one of the changes below, and run it. Then undo your change and try the next one. See what happens!

```go
    package main
    import "fmt"
    func main() {
        fmt.Println("Hello, Go!")}
```

> NOTE - Try breaking our code sample and see what happens!

| If you do this...                                                       | ...it will fail because...                                                                                  |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| Delete the package clause... ~~package main~~                           | Every Go file has to begin with a package clause.                                                           |
| Delete the import statement... ~~import "fmt"~~                         | Every Go file has to import every package it references.                                                    |
| Import a second (unused) package... ~~import "fmt" import "strings"~~   | Go files must import only the packages they reference. (This helps keep your code compiling fast!)          |
| Rename the main function... func ~~main~~hello                          | Go looks for a function named main to run first.                                                            |
| Change the Println call to lowercase... fmt.Pprintln("Hello, Go!")      | Everything in Go is case-sensitive, so although fmt.Println is valid, there’s no such thing as fmt.println. |
| Delete the package name before Println... ~~fmt~~.Println("Hello, Go!") | The Println function isn’t part of the main package, so Go needs the package name before the function call. |

Let’s try the first one as an example...
![Play Goloang](./../images/chapter1/f0006-03.png)

## Calling functions

Our example includes a call to the fmt package’s Println function. To call a function, type the function name (Println in this case), and a pair of parentheses.

![Play Goloang](./../images/chapter1/f0007-01.png)

![Play Goloang](./../images/chapter1/f0007-02.png)

Like many functions, Println can take one or more arguments: values you want the function to work with. The arguments appear in parentheses after the function name.

![Play Goloang](./../images/chapter1/f0007-03.png)

Println can be called with no arguments, or you can provide several arguments. When we look at other functions later, however, you’ll find that most require a specific number of arguments. If you provide too few or too many, you’ll get an error message saying how many arguments were expected, and you’ll need to fix your code.

## The Println function

Use the Println function when you need to see what your program is doing. Any arguments you pass to it will be printed (displayed) in your terminal, with each argument separated by a space.

After printing all its arguments, Println will skip to a new terminal line. (That’s why “ln” is at the end of its name.)

![Play Goloang](./../images/chapter1/f0007-04.png)

## Using functions from other packages

The code in our first program is all part of the main package, but the Println function is in the fmt package. (The fmt stands for “format.”) To be able to call Println, we first have to import the package containing it.

![Play Goloang](./../images/chapter1/f0008-01.png)

Once we’ve imported the package, we can access any functions it offers by typing the package name, a dot, and the name of the function we want.

![Play Goloang](./../images/chapter1/f0008-02.png)

Here’s a code sample that calls functions from a couple other packages. Because we need to import multiple packages, we switch to an alternate format for the import statement that lets you list multiple packages within parentheses, one package name per line.

![Play Goloang](./../images/chapter1/f0008-03.png)

Once we’ve imported the math and strings packages, we can access the math package’s Floor function with math.Floor, and the strings package’s Title function with strings.Title.

You may have noticed that in spite of including those two function calls in our code, the above sample doesn’t display any output. We’ll look at how to fix that next.
