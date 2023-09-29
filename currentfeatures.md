## Current
SDK Wrapper around lotr api being consumed

MOVIE Endpoints
```
movie.GetMovies() -> Return All Movies
movie.GetMovieByID(movieID) -> Return Movie By ID
```
QUOTE Endpoints
```
quote.getQuotes() -> Return All Quotes
quote.GetQuoteByID(quoteID) -> Return Quote By ID
quote.GetQuotesByMovieID(movieID) ->  Return All Quotes By MovieID
```
Built on src request package main router handling-
Sub packages handle specific endpoints and call request package
