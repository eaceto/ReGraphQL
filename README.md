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
GET /persons/{person}
````

## Index
* [Quick start](#quick-start)
* [Contributing](#contributing)
* [License](#license)
* [Author](#author)

### Quick start

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
*File* **starwars.yml**

2. Copy **starwars.yml** into **./config**

3. Run the service (using Docker Compose)
````shell
docker-compose up
````

4. Query your new HTTP endpoint!
````shell
curl 'http://127.0.0.1:8080/graphql/persons/lukeskywalker' --compressed
````

### Contributing
Before contributing to ReGraphQL, please read the instructions detailed in our [contribution guide](CONTRIBUTING.md).

### License
ReGraphQL is released under the MIT license. See [LICENSE](LICENSE) for details.

### Author
Created by [Ezequiel (Kimi) Aceto](https://eaceto.dev).