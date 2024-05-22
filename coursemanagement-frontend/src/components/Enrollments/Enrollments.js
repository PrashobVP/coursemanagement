// src/components/Enrollments/Enrollments.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Enrollments = () => {
    const [enrollments, setEnrollments] = useState([]);
    const [newEnrollment, setNewEnrollment] = useState({ studentId: '', courseId: '' });

    useEffect(() => {
        fetchEnrollments();
    }, []);

    const fetchEnrollments = async () => {
        try {
            const response = await axios.get('/enrollment');
            setEnrollments(response.data);
        } catch (error) {
            console.error("Error fetching enrollments", error);
        }
    };

    const handleAddEnrollment = async () => {
        try {
            await axios.post('/enrollment', newEnrollment);
            fetchEnrollments(); // Refresh the list
        } catch (error) {
            console.error("Error adding enrollment", error);
        }
    };

    return (
        <div>
            <h2>Enrollments</h2>
            <ul>
                {enrollments.map((enrollment, index) => (
                    <li key={index}>{`Student ID: ${enrollment.studentId}, Course ID: ${enrollment.courseId}`}</li>
                ))}
            </ul>
            <input 
                type="text" 
                value={newEnrollment.studentId} 
                onChange={(e) => setNewEnrollment({ ...newEnrollment, studentId: e.target.value })} 
                placeholder="Student ID" 
            />
            <input 
                type="text" 
                value={newEnrollment.courseId} 
                onChange={(e) => setNewEnrollment({ ...newEnrollment, courseId: e.target.value })} 
                placeholder="Course ID" 
            />
            <button onClick={handleAddEnrollment}>Add Enrollment</button>
        </div>
    );
};

export default Enrollments;
