## Table of Contents

1: CHAT APPLICATION WITH WEB SOCKETS

- A simple web server
  - Separating views from logic using templates
    - Doing things once 
    - Using your own handlers 
  - Properly building and executing Go programs 
- Modeling a chat room and clients on the server
  - Modeling the client
  - Modeling a room
  - Concurrency programming using idiomatic Go
  - Turning a room into an HTTP handler
  - Using helper functions to remove complexity
  - Creating and using rooms
- Building an HTML and JavaScript chat client
  - Getting more out of templates
- Tracing code to get a look under the hood
  - Writing a package using TDD 
  - Interfaces 
  - Unit tests 
    - Red-green testing 
  - Implementing the interface
    - Unexported types being returned to users 
  - Using our new trace package
  - Making tracing optional 
  - Clean package APIs
- Summary

2: ADDING USER ACCOUNTS

- Handlers all the way down
- Making a pretty social sign-in page
- Endpoints with dynamic paths
- Getting started with OAuth2
  - Open source OAuth2 packages
- Tell the authorization providers about your app
- Implementing external logging in
- Summary

3: THREE WAYS TO IMPLEMENT PROFILE PICTURES

- Avatars from the OAuth2 server
- Implementing Gravatar
- Uploading an avatar picture
- Combining all three implementations
- Summary

4: COMMAND-LINE TOOLS TO FIND DOMAIN NAMES

- Pipe design for command-line tools
- Five simple programs
- Composing all five programs
- Summary

5: BUILDING DISTRIBUTED SYSTEMS AND WORKING WITH FLEXIBLE DATA

- The system design
- Installing the environment
- Reading votes from Twitter
- Counting votes
- Running our solution
- Summary

6: EXPOSING DATA AND FUNCTIONALITY THROUGH A RESTFUL DATA WEB SERVICE API

- RESTful API design
- Sharing data between handlers
- Wrapping handler functions
- Injecting dependencies
- Responding
- Understanding the request
- Serving our API with one function
- Handling endpoints
- A web client that consumes the API
- Running the solution
- Summary

7: RANDOM RECOMMENDATIONS WEB SERVICE

- The project overview
- Representing data in code
- Generating random recommendations
- Summary

8: FILESYSTEM BACKUP

- Solution design
- The backup package
- The user command-line tool
- The daemon backup tool
- Testing our solution
- Summary

9: BUILDING A Q&A APPLICATION FOR GOOGLE APP ENGINE

- The Google App Engine SDK for Go
- Google Cloud Datastore
- Entities and data access
- Google App Engine users
- Transactions in Google Cloud Datastore
- Querying in Google Cloud Datastore
- Votes
- Casting a vote
- Exposing data operations over HTTP
- Running apps with multiple modules
- Deploying apps with multiple modules
- Summary

10: MICRO-SERVICES IN GO WITH THE GO KIT FRAMEWORK

- Introducing gRPC
- Protocol buffers
- Building the service
- Modeling method calls with requests and responses
- An HTTP server in Go kit
- A gRPC server in Go kit
- Creating a server command
- Building a gRPC client
- Rate limiting with service middleware
- Summary

11: DEPLOYING GO APPLICATIONS USING DOCKER

- Using Docker locally
- Deploying Docker images
- Deploying to Digital Ocean
- Summary

Appendix: Good Practices for a Stable Go Environment 350



- Installing Go 
- Configuring Go
- Getting GOPATH right
- Go tools 
- Cleaning up, building, and running tests on save 
- Integrated developer environments 356
  - Sublime Text 3
  - Visual Studio Code
- Summary
