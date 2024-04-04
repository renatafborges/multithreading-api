<h1 align="center">
   <a href="#"> Address Lookup Using Multiple APIs </a>
</h1>

<h3 align="center">
   This Go program allows users to look up address information based on zip codes using two different APIs: BrasilAPI and ViaCEP. 
   It concurrently queries both APIs and returns the address information from whichever API responds first. It also includes a timeout mechanism to prevent waiting indefinitely for a response.
</h3>

<h4 align="center"> 
	 Status: Finished
</h4>

## Usage
1. Clone the repository.
2. Navigate to the directory containing the main.go file.
3. Run go run main.go.
4. Enter the zip code when prompted.

## Dependencies
This program uses the following dependencies:
1. context: to manage cancellation and timeouts.
2. net/http: to make HTTP requests.
3. io: to read the response body from HTTP requests.

## How It Works
The program prompts the user to enter a zip code for address lookup.
It concurrently queries both the BrasilAPI and ViaCEP using goroutines.
Each goroutine makes an HTTP request to the respective API with the provided zip code.
The program waits for the first response using a select statement with a timeout.
If a response is received within the timeout period, it prints the address information.
If no response is received within the timeout period, it prints a timeout message.

## Functions
main: The entry point of the program. It prompts the user for input, creates a context with a timeout, and launches goroutines to query both APIs.
RequestBrasilApi and RequestViaCep: Functions to query the BrasilAPI and ViaCEP respectively. They handle errors and send the address information or error message to the provided channel.
fetchAddress: Function to make HTTP requests to the specified URL with the provided context. It returns the response body or an error message.

## Error Handling
The program handles errors such as failure to create an HTTP request, failure to fetch the address, and failure to read the response body.
If an error occurs during the address lookup from either API, the program prints an error message.

## Notes
Both APIs provide similar functionality but may have differences in data format or availability.
The program selects the response from the first API to respond within the timeout period, providing a fallback mechanism in case one API is slow or unavailable.
Feel free to use, modify, or extend this program as needed!

## Author
Made with love by Renata Borges 👋🏽 [Get in Touch!](Https://www.linkedin.com/in/renataborgestech)

