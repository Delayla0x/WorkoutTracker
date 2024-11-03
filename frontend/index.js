import React, { useCallback, useEffect } from 'react';

const WorkoutTrackerComponent = React.memo(({ onWorkoutClick }) => {
  
});

function WorkoutTrackerParent(props) {
  const handleWorkoutClick = useCallback(() => {
    
  }, []);
  
  return <WorkoutTrackerComponent onWorkoutClick={handleWorkworkoutClick} />;
}

useEffect(() => {
 const workoutTimer = setTimeout(() => {
   
 }, 1000);
 
 return () => clearTimeout(workoutTimer);
}, []);