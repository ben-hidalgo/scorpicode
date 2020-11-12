# Scorpicode

Scorpicode is a technical architecture prototype demonstrating microservices in a monorepo.  

It contains a React/Mobx frontend, multiple GoLang backend services, a Gatsby website and a GitHub Actions CI/CD pipeline supporting multi-developer, multi-branch development workflows.

There is support for role-based user authentication (signup / signin) using Auth0 and JWT utilizing micro-frontends.

Kubernetes artifact creation & updates are Helm managed and secrets encryption / decryption uses SOPS yaml files.

Terraform is used to run development dependencies locally.

## Getting Started

> git clone https://github.com/ben-hidalgo/scorpicode.git

> cd scorpicode

The [Makefile](Makefile) contains many helpful targets.  There are significant setup steps not yet documented...

### Prerequisites

Scorpicode uses 

* GoLang 1.13 
* npm 6.14.8
* sops 3.6.0
* Terraform v0.13.0

### Website

> make start-website

> open http://localhost:8000/

## Next steps

- [ ] Health checks
- [ ] TLS / https SSL certificates 
- [ ] Cluster friendly web socket routing
- [ ] Debug rolling upgrade issue on "website" by checksum

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
