import React, { useState } from 'react';
import axios from 'axios';

const WorkoutForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    duration: '',
    intensity: '',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault(); 

    try {
      const endpoint = process.env.REACT_APP_API_URL + '/workouts';
      const response = await axios.post(endpoint, formData);

      setFormData({
        name: '',
        duration: '',
        intensity: '',
      });

      console.log('Workout created successfully', response.data);
    } catch (error) {
      console.error('Error creating workout:', error.response.data);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="name">Workout Name</label>
        <input
          type="text"
          id="name"
          name="name"
          value={formData.name}
          onChange={handleChange}
        />
      </div>
      <div>
        <label htmlFor="duration">Duration (minutes)</label>
        <input
          type="number"
          id="duration"
          name="duration"
          value={formData.duration}
          onChange={handleChange}
        />
      </div>
      <div>
        <label htmlFor="intensity">Intensity</label>
        <select
          id="intensity"
          name="intensity"
          value={formData.intensity}
          onChange={handleChange}
        >
          <option value="">Select intensity</option>
          <option value="low">Low</option>
          <option value="medium">Medium</option>
          <option value="high">High</option>
        </select>
      </div>
      <button type="submit">Add Workout</button>
    </form>
  );
};

export default WorkoutForm;