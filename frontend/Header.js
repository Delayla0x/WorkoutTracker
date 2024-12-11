import React from 'react';
import { Link } from 'react-router-dom';

const Header = () => {
  return (
    <header style={{display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '1rem', backgroundColor: '#f0f0f0'}}>
      <h1>Workout Tracker</h1>
      <nav>
        <ul style={{listStyle: 'none', display: 'flex', gap: '1rem'}}>
          <li><Link to="/view-workouts">View Workouts</Link></li>
          <li><Link to="/add-workout">Add New Workout</Link></li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;