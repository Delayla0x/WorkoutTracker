import React, { useState } from 'react';
import axios from 'axios';

const WorkoutForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    duration: '',
    intensity: '',
  });

  // State to hold error messages
  const [errorMessage, setErrorMessage] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    setErrorMessage(''); // Reset error message on new submission attempt

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
      console.error('Error creating workout:', error.response ? error.response.data : error);

      // More robust error handling to account for various error scenarios
      if (error.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        if (error.response.data && typeof error.response.data === 'object') {
          setErrorMessage('Error: ' + (error.response.data.message || JSON.stringify(error.response.data)));
        } else {
          setErrorMessage('Error: ' + error.response.statusText);
        }
      } else if (error.request) {
        // The request was made but no response was received
        setErrorMessage('Error: The server did not respond.');
      } else {
        // Something happened in setting up the request that triggered an Error
        setErrorMessage('Error: ' + error.message);
      }
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
      {errorMessage && <div style={{ color: 'red', marginTop: '10px' }}>{errorMessage}</div>}
      <button type="submit">Add Workout</button>
    </form>
  );
};

export default WorkoutForm;