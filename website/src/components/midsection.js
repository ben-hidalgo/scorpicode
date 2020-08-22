import React from 'react';
import {
	FaReact,
	FaMobileAlt,
	FaOsi
} from 'react-icons/fa';

import './style.scss';

const Midsection = () => (
	<div>
		<section className="section">
			<div className="container">
				<div className="columns is-multiline">
					<div className="column is-one-third">
						<article className="media">
							<figure className="media-left">
								<span className="icon is-medium">
									<FaReact size="29px" color="#5e227f" />
								</span>
							</figure>
							<div className="media-content">
								<div className="content">
									<h1 className="title is-size-4">Gatsby + React</h1>
									<p className="subtitle is-size-5">
										Scorpicode is built with Gatsby and React
									</p>
								</div>
							</div>
						</article>
					</div>
					<div className="column is-one-third">
						<article className="media">
							<figure className="media-left">
								<span className="icon is-medium">
									<FaMobileAlt size="29px" color="blue" />
								</span>
							</figure>
							<div className="media-content">
								<div className="content">
									<h1 className="title is-size-4">Responsive Design</h1>
									<p className="subtitle is-size-5">
										Using the Bulma / Flexbox model and Sass
									</p>
								</div>
							</div>
						</article>
					</div>

					<div className="column is-one-third">
						<article className="media">
							<figure className="media-left">
								<span className="icon is-medium">
									<FaOsi size="29px" className="has-text-primary" />
								</span>
							</figure>
							<div className="media-content">
								<div className="content">
									<h1 className="title is-size-4">Open Source</h1>
									<p className="subtitle is-size-5">
										
									</p>
								</div>
							</div>
						</article>
					</div>
				</div>
			</div>
		</section>
		<section className="section">
			<div className="content">
				<p><a href="https://github.com/ben-hidalgo/scorpicode">Scorpicode</a>  is a technical architecture prototype demonstrating a microservices monorepo.</p>
				<p>
					It contains a React/Mobx frontend, multiple GoLang backend services, a Gatsby website and a GitHub Actions CI/CD pipeline supporting multi-developer, multi-branch development workflows.
				</p>
				<p>There is support for role-based user authentication (signup / signin) using Auth0 and JWT via micro-frontends.</p>
				<p>Kubernetes artifact creation & updates are Helm managed and secrets encryption / decryption uses SOPS yaml files.</p>
				<p>Minikube is used to run development dependencies locally, also with Helm charts "stable."</p>
			</div>
		</section>
	</div>
);

export default Midsection;
