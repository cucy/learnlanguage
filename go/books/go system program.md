## Table of Contents

1: GETTING STARTED WITH GO AND UNIX SYSTEMS PROGRAMMING

- The structure of the book
- What is systems programming?
- About Go
- Two useful Go tools
  - Advantages and disadvantages of Go
- The various states of a Unix process
- Exercises
- Summary

2: WRITING PROGRAMS IN GO

- Compiling Go code
- Go environment variables
- Using command-line arguments
  - Finding the sum of the command-line arguments
- User input and output
  - Getting user input 
  - Printing output 
- Go functions
  - Naming the return values of a Go function 
  - Anonymous functions
  - Illustrating Go functions 
  - The defer keyword 
  - Using pointer variables in functions 
- Go data structures
  - Arrays
  - Slices
  - Maps
    - Converting an array into a map
  - Structures
  - Interfaces
- Interfaces
- Creating random numbers
- Exercises
- Summary

3: ADVANCED GO FEATURES

- Error handling in Go
  - Functions can return error variables 
  - About error logging
  - The addCLA.go program revisited
- Pattern matching and regular expressions
  - Printing all the values from a given column of a line 
  - Creating summaries
  - Finding the number of occurrences 
  - Find and replace
- Reflection
  - Calling C code from Go
  - Unsafe code
- Comparing Go to other programming languages
- Analysing software
  - Using the strace(1) command-line utility
  - The DTrace utility 
    - Disabling System Integrity Protection on macOS
- Unreachable code
- Avoiding common Go mistakes
- Exercises
- Summary

4: GO PACKAGES, ALGORITHMS, AND DATA STRUCTURES

- About algorithms
  - The Big O notation
- Sorting algorithms
  - The sort.Slice() function
- Linked lists in Go
- Trees in Go
- Developing a hash table in Go
- About Go packages
  - Using standard Go packages
  - Creating your own packages
    - Private variables and functions
    - The init() function 
  - Using your own Go packages
  - Using external Go packages 
    - The go clean command
- Garbage collection
- Your environment
- Go gets updated frequently!
- Exercises
- Summary

5: FILES AND DIRECTORIES

- Useful Go packages
- Command-line arguments revisited!
  - The flag package
- Dealing with directories
  - About symbolic links
  - Implementing the pwd(1) command
  - Developing the which(1) utility in Go
    - Printing the permission bits of a file or directory
- Dealing with files in Go
  - Deleting a file
  - Renaming and moving files
- Developing find(1) in Go
  - Traversing a directory tree
  - Visiting directories only!
- The first version of find(1)
  - Adding some command-line options
  - Excluding filenames from the find output
  - Excluding a file extension from the find output
- Using regular expressions
  - Creating a copy of a directory structure
- Exercises
- Summary

6: FILE INPUT AND OUTPUT

- About file input and output
  - Byte slices 
  - About binary files
- Useful I/O packages in Go
  - The io package 
  - The bufio package
- File I/O operations
  - Writing to files using fmt.Fprintf()
    - About io.Writer and io.Reade
  - Finding out the third column of a line 
- Copying files in Go
  - There is more than one way to copy a file!
  - Copying text files 
  - Using io.Copy 
  - Reading a file all at once! 
  - An even better file copy program 
  - Benchmarking file copying operations 
- Developing wc(1) in Go
  - Counting words
  - The wc.go code! 
    - Comparing the performance of wc.go and wc(1)
  - Reading a text file character by character
    - Doing some file editing! 
- Interprocess communication
- Sparse files in Go
- Reading and writing data records
- File locking in Go
- A simplified Go version of the dd utility
- Exercises
- Summary

7: WORKING WITH SYSTEM FILES

- Which files are considered system files?
- Logging in Go
- Putting data at the end of a file
- About log files
- More pattern matching examples
- Searching files revisited
- Date and time operations
- Rotating log files
- Creating good random passwords
- Another Go update
- Exercises
- Summary

8: PROCESSES AND SIGNALS

- About Unix processes and signals
- Process management
- Unix signals in Go
- The kill(1) command
- Improving file copying
- Plotting data
- Unix pipes in Go
- Unix sockets in Go
- RPC in Go
- Programming a Unix shell in Go
- Yet another minor Go update
- Exercises
- Summary

9: GOROUTINES - BASIC FEATURES

- About goroutines
- The sync Go packages
- Pipelines
- A better version of wc.go
- Exercises
- Summary

10: GOROUTINES - ADVANCED FEATURES

- The Go scheduler
- The sync Go package
- The select keyword
- Signal channels
- Buffered channels
- About timeouts
- Channels of channels
- Nil channels
- Shared memory
- The dWC.go utility revisited
- Detecting race conditions
- About GOMAXPROCS
- Exercises
- Summary

11: WRITING WEB APPLICATIONS IN GO

- What is a web application?
- About the net/http Go package
- Developing web clients
- A small web server
- The html/template package
- About JSON
- Using MongoDB
- A handy command-line utility
- Exercises
- Summary

12: NETWORK PROGRAMMING

- About network programming
- The net Go standard package
- Unix sockets revisited
- Performing DNS lookups
- Developing a simple TCP server
- Developing a simple TCP client
- Developing a simple UDP server
- Developing a simple UDP client
- A concurrent TCP server
- Remote procedure call (RPC)
- Exercises
- Summary
