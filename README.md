# Scorpicode

Scorpicode is a technical architecture prototype demonstrating a microservices monorepo.  

It contains a React/Mobx frontend, multiple GoLang backend services, a Gatsby website and a GitHub Actions CI/CD pipeline supporting multi-developer, multi-branch development workflows.

There is support for role-based user authentication (signup / signin) using Auth0 and JWT via micro-frontends.

Kubernetes artifact creation & updates are Helm managed and secrets encryption / decryption uses SOPS yaml files.

Minikube is used to run development dependencies locally, also with Helm charts "stable."

## Getting Started

> git clone https://github.com/ben-hidalgo/scorpicode.git

> cd scorpicode

The [Makefile](Makefile) contains many helpful targets.

### Prerequisites

Scorpicode uses 

* GoLang 1.13 
* npm 6.14.8
* sops 3.6.0
* minikube 1.11.0

## Local Runtimes

### Website

> make start-website

> open http://localhost:8000/

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
* https://mobx.js.org/getting-started.html
* https://github.com/gothinkster/react-mobx-realworld-example-app
* https://twitchtv.github.io/twirp/docs/best_practices.html
* https://github.com/golang/go/wiki/CodeReviewComments
* https://github.com/golang/go/wiki/TableDrivenTests
* https://mobx-react.js.org/recipes-context
* https://auth0.com/docs/authorization/concepts/sample-use-cases-rules
* https://www.gatsbyjs.org/starters/amandeepmittal/gatsby-bulma-quickstart/
* https://github.com/gorilla/websocket/blob/master/examples/filewatch/main.go
* https://help.github.com/en/actions/reference/context-and-expression-syntax-for-github-actions
