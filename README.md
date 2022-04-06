# ReGraphQL

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/eaceto/ReGraphQL)

![Docker Image Version (latest semver)](https://img.shields.io/docker/v/eaceto/regraphql?color=red&label=Docker%20Image%20version)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/eaceto/regraphql?color=red&label=Docker%20Image%20size)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/eaceto/ReGraphQL/Go?label=GitHub%20CI)

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

# Index
* [Requirements](#requirements)
* [Features](#features)
* [Quick start](#quick-start)
* [Docker Image](#docker-image)
* [Contributing](#contributing)
* [License](#license)
* [Author](#author)

## Requirements

* Go 1.18

## Features
*As per version 1.0.1*

- [x] Maps HTTP params to GraphQL Variables
- [x] Forwards HTTP headers to GraphQL request
- [x] Reads configuration from **.env** file
- [x] Reads configuration from **environment variables**
- [x] Logs using Kubernetes' [**klog**](https://github.com/kubernetes/klog) v2
- [x] Docker Image below 20MB
- [ ] Exposes metrics using [Prometheus](https://prometheus.io/)
- [ ] Exposes Liveness, Readiness and Startup Probes 
- [ ] Implements Hot reload for routes

## Quick start

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

2. Copy **starwars.yml** into **/tmp/config**

3. Run the service (using Docker Compose)
````shell
[sudo] docker-compose up
````

4. Query your new HTTP endpoint!
````shell
curl 'http://127.0.0.1:8080/graphql/persons/lukeskywalker' --compressed
````

## Docker Image
Docker image is based on Google's Distroless. The final image is around 11.2MB and packs only the necessary things to run the service.

````shell
docker pull eaceto/regraphql:1.0.1
````

## Contributing
Before contributing to ReGraphQL, please read the instructions detailed in our [contribution guide](CONTRIBUTING.md).

## License
ReGraphQL is released under the MIT license. See [LICENSE](LICENSE) for details.

## Author
Created by [Ezequiel (Kimi) Aceto](https://eaceto.dev).