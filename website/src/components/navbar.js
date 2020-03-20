import React from 'react';

import './style.scss';
import logo from '../images/scorpicode210.png';

const Navbar = () => (
		<nav className="navbar">
			<div className="container">
				<div className="navbar-brand">
					<a
						className="navbar-item"
						href="."
					>
						<img src={logo} alt="scorpicode logo" />
					</a>
					<a 
						className="navbar-item"
						href="/login"
					>
						Login
					</a>
				</div>
			</div>
		</nav>
);

export default Navbar;
