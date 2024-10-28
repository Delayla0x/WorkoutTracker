const MyComponent = React.memo(function MyComponent(props) {
  /* render using props */
});

import React, { useCallback } from 'react';

const MyComponent = React.memo(({ onClick }) => {
  /* Use onClick without redeclaring it on every render */
});

function ParentComponent(props) {
  const handleClick = useCallback(() => {
    /* handle click event */
  }, []); // dependencies array

  return <MyComponent onClick={handleClick} />;
}

useEffect(() => {
 const timer = setTimeout(() => {
   // Do something
 }, 1000);

 return () => clearTimeout(timer); // Cleanup
}, []);