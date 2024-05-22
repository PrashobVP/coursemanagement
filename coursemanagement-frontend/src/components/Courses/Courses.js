// src/components/Courses/Courses.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Courses = () => {
    const [courses, setCourses] = useState([]);
    const [newCourse, setNewCourse] = useState("");

    useEffect(() => {
        fetchCourses();
    }, []);

    const fetchCourses = async () => {
        try {
            const response = await axios.get('/course');
            setCourses(response.data);
        } catch (error) {
            console.error("Error fetching courses", error);
        }
    };

    const handleAddCourse = async () => {
        try {
            await axios.post('/course', { name: newCourse });
            fetchCourses(); // Refresh the list
        } catch (error) {
            console.error("Error adding course", error);
        }
    };

    return (
        <div>
            <h2>Courses</h2>
            <ul>
                {courses.map((course, index) => (
                    <li key={index}>{course.name}</li>
                ))}
            </ul>
            <input 
                type="text" 
                value={newCourse} 
                onChange={(e) => setNewCourse(e.target.value)} 
                placeholder="New Course" 
            />
            <button onClick={handleAddCourse}>Add Course</button>
        </div>
    );
};

export default Courses;
