// src/components/Teachers/Teachers.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Teachers = () => {
    const [teachers, setTeachers] = useState([]);
    const [newTeacher, setNewTeacher] = useState("");

    useEffect(() => {
        fetchTeachers();
    }, []);

    const fetchTeachers = async () => {
        try {
            const response = await axios.get('/teacher');
            setTeachers(response.data);
        } catch (error) {
            console.error("Error fetching teachers", error);
        }
    };

    const handleAddTeacher = async () => {
        try {
            await axios.post('/teacher', { name: newTeacher });
            fetchTeachers(); // Refresh the list
        } catch (error) {
            console.error("Error adding teacher", error);
        }
    };

    return (
        <div>
            <h2>Teachers</h2>
            <ul>
                {teachers.map((teacher, index) => (
                    <li key={index}>{teacher.name}</li>
                ))}
            </ul>
            <input 
                type="text" 
                value={newTeacher} 
                onChange={(e) => setNewTeacher(e.target.value)} 
                placeholder="New Teacher" 
            />
            <button onClick={handleAddTeacher}>Add Teacher</button>
        </div>
    );
};

export default Teachers;
