// src/components/Navbar.js
import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => (
    <nav>
        <ul>
            <li><Link to="/courses">Courses</Link></li>
            <li><Link to="/students">Students</Link></li>
            <li><Link to="/teachers">Teachers</Link></li>
            <li><Link to="/enrollments">Enrollments</Link></li>
        </ul>
    </nav>
);

export default Navbar;
