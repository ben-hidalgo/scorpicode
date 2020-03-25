import React from 'react';
import Helmet from 'react-helmet';

export default () => {
	return <Helmet>
		<meta
			name="viewport"
			content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0"
		/>
		<meta name="description" content="scorpicode frontend" />
		<meta name="keywords" content="golang kubernetes helm react mobx bulma" />
		<title>Scorpicode</title>
		<html lang="en" />
		{/* Google / Search Engine Meta Tags */}
		<meta itemprop="name" content="ben-hidalgo" />
		<meta itemprop="description" content="scorpicode microservice architecture" />
		<meta itemprop="image" content="./images/scorpicode210.png" />
	</Helmet>
}
