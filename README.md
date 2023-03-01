# Boilerplate of Go - CRUD GoLang REST API with Mux, GORM & Meta Pagination

![Boilerplate of Go - CRUD GoLang REST API with Mux, GORM & Meta Pagination](https://miro.medium.com/max/1400/0*ck-mgOkywViHkZx2.png)

We will learn about implementing boilerplate of golang, CRUD using Golang REST API and Gorilla Mux for routing requests, GORM to access the database, and we can use PostgreSQL or MySQL as the database provider.

## Topics :
- Setting up the Golang Project
- Defining the Product Entity
- Connecting to the database
- Routing
- Implementing CRUD in Golang Rest API
	 - Create
	 - Get By ID
	 - Get All Data with pagination
	 - Update
	 - Delete (we don't use delete like destroy data in here, but we use status like active/inactive to keep the data traceable)
- Testing CRUD Operations

## Our Step :
- In the first step, we create an entities to represent the structure of our data in the entity package
- In the second step, we create query a function to handle our query
- Then register the new function in our model interface
- In the third step, we create algorithm function to handle all condition in our flow proccess
- and again, we register it into source contract interface
- In fourth step is create a route function to handle our input payload and send it to our controller
- And the last step is, we need to initiate our model interface, controller interface and handler in router.go before we write our endpoint path and the method

## Notes :
Last but not least, let's start to try to write a function that we think will be used many times in helpers.
Happy coding and keep learning 😜
