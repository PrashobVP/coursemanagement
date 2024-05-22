// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Courses from './components/Courses/Courses';
import Students from './components/Students/Students';
import Teachers from './components/Teachers/Teachers';
import Enrollments from './components/Enrollments/Enrollments';

function App() {
    return (
        <Router>
            <div>
                <Navbar />
                <Routes>
                    <Route path="/courses" element={<Courses />} />
                    <Route path="/students" element={<Students />} />
                    <Route path="/teachers" element={<Teachers />} />
                    <Route path="/enrollments" element={<Enrollments />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
