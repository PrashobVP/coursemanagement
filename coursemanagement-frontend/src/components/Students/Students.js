// src/components/Students/Students.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Students = () => {
    const [students, setStudents] = useState([]);
    const [newStudent, setNewStudent] = useState("");

    useEffect(() => {
        fetchStudents();
    }, []);

    const fetchStudents = async () => {
        try {
            const response = await axios.get('/student');
            setStudents(response.data);
        } catch (error) {
            console.error("Error fetching students", error);
        }
    };

    const handleAddStudent = async () => {
        try {
            await axios.post('/student', { name: newStudent });
            fetchStudents(); // Refresh the list
        } catch (error) {
            console.error("Error adding student", error);
        }
    };

    return (
        <div>
            <h2>Students</h2>
            <ul>
                {students.map((student, index) => (
                    <li key={index}>{student.name}</li>
                ))}
            </ul>
            <input 
                type="text" 
                value={newStudent} 
                onChange={(e) => setNewStudent(e.target.value)} 
                placeholder="New Student" 
            />
            <button onClick={handleAddStudent}>Add Student</button>
        </div>
    );
};

export default Students;
