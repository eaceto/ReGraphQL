# ReGraphQL
**A simple (yet effective) REST / HTTP to GraphQL router**

ReGraphQL helps you expose REST/HTTP endpoints and route it to a GraphQL endpoints.
Doing this has the following benefits:

* Queries are stored and controlled server side.
* Can modify and optimise your queries on demand without redoploying your (frontend) clients
* Can use GET (HTTP Method) instead of GraphQL's POST
 
It helps you going...

**From this** 
````graphql
query($person: StarWarsPeople!) {
	getPerson(person: $person) {
		birthYear
		eyeColors
		films {
			title
		}
		gender
		hairColors
		height
		homeworld {
			name
		}
		mass
		name
		skinColors
		species {
			name
		}
		starships {
			name
		}
		vehicles {
			name
		}
	}
}
````

**To**
````http request
GET /person/:person
````

## How to

1. Describe a route in a file using **yaml**, which matches your HTTP endpoint with your GraphQL endpoint and Query 

````yaml
routes:
    - http:
          uri: '/persons/{person}'
          method: GET
      graphql:
          endpoint: 'https://swapi.skyra.pw/'
          query: |
              query($person: StarWarsPeople!) {
                  getPerson(person: $person) {
                      birthYear
                      eyeColors
                      films {
                          title
                      }
                      gender
                      hairColors
                      height
                      homeworld {
                          name
                      }
                      mass
                      name
                      skinColors
                      species {
                          name
                      }
                      starships {
                          name
                      }
                      vehicles {
                          name
                      }
                  }
              }
````

2. Run the service (with Docker, for example)
````shell
docker run --rm -p 8080:8080 .
````

4. Query your new HTTP endpoint!
