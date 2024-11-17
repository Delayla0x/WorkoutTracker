import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Header from './components/Header';
import ViewWorkouts from './components/ViewWorkouts';
import AddWorkout from './components/AddWorkout';

function App() {
  return (
    <Router>
      <div>
        <Header />
        <Routes>
          <Route path="/" element={<ViewWorkouts />} />
          <Route path="/add" element={<AddWorkout />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;