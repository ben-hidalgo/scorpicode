import React from 'react';

import './style.scss';

import Navbar from './navbar';

const Header = ({ siteTitle }) => (
	<section className="hero gradientBg" >
		<Navbar />
	</section>
);

export default Header;
