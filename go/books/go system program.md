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
  - Altering existing data
- About log files
  - About logging
  - Logging facilities 
  - Logging levels 
  - The syslog Go package
  - Processing log files
  - File permissions revisited
    - Changing file permissions
    - Finding other kinds of information about files
- More pattern matching examples
  - A simple pattern matching example
  - An advanced example of pattern matching
  - Renaming multiple files using regular expressions
- Searching files revisited
  - Finding the user ID of a user
  - Finding all the groups a user belongs to 
  - Finding files that belong or do not belong to a given user
  - Finding files based on their permissions
- Date and time operations
  - Playing with dates and times
  - Reformatting the times in a log file 
- Rotating log files
- Creating good random passwords
- Another Go update
- Exercises
- Summary

8: PROCESSES AND SIGNALS

- About Unix processes and signals
- Process management
  - About Unix signals 
- Unix signals in Go
- The kill(1) command
  - A simple signal handler in Go
  - Handling three different signals! 
  - Catching every signal that can be handled
  - Rotating log files revisited!
- Improving file copying
- Plotting data
- Unix pipes in Go
  - Reading from standard input
  - Sending data to standard output 
  - Implementing cat(1) in Go 
  - The plotIP.go utility revisited
- Unix sockets in Go
- RPC in Go
- Programming a Unix shell in Go
- Yet another minor Go update
- Exercises
- Summary

9: GOROUTINES - BASIC FEATURES

- About goroutines
  - Concurrency and parallelism
- The sync Go packages
  - A simple example
    - Creating multiple goroutines
  - Waiting for goroutines to finish their jobs
    - Creating a dynamic number of goroutines
  - About channels 
    - Writing to a channel 
    - Reading from a channel
    - Explaining h1s.go
- Pipelines
- A better version of wc.go
  - Calculating totals
  - Doing some benchmarking 
- Exercises
- Summary

10: GOROUTINES - ADVANCED FEATURES

- The Go scheduler
- The sync Go package
- The select keyword
- Signal channels
- Buffered channels
- About timeouts
  - An alternative way to implement timeouts
- Channels of channels
- Nil channels
- Shared memory
  - Using sync.Mutex
  - sing sync.RWMutex
- The dWC.go utility revisited
  - Using a buffered channel
  - Using shared memory
  - More benchmarking 
- Detecting race conditions
- About GOMAXPROCS
- Exercises
- Summary

11: WRITING WEB APPLICATIONS IN GO

- What is a web application?
- About the net/http Go package
- Developing web clients
  - Fetching a single URL
    - Setting a timeout
  - Developing better web clients
- A small web server
- The html/template package
- About JSON
  - Saving JSON data 
  - Parsing JSON data
  - Using Marshal() and Unmarshal()
- Using MongoDB
  - Basic MongoDB administration
  - Using the MongoDB Go driver
  - Creating a Go application that displays MongoDB data
  - Creating an application that displays MySQL data 
- A handy command-line utility
- Exercises
- Summary

12: NETWORK PROGRAMMING

- About network programming
  - About TCP/IP
  - About TCP 
    - The TCP handshake!
  - About UDP and IP
  - About Wireshark and tshark
  - About the netcat utility
- The net Go standard package
- Unix sockets revisited
  - A Unix socket server 
  - A Unix socket client 
- Performing DNS lookups
  - Using an IP address as input
  - Using a host name as input
  - Getting NS records for a domain
- Developing a simple TCP server
- Developing a simple TCP client
  - Using other functions for the TCP server
  - Using alternative functions for the TCP client
- Developing a simple UDP server
- Developing a simple UDP client
- A concurrent TCP server
- Remote procedure call (RPC)
  - An RPC server 
  - An RPC client
- Exercises
- Summary
