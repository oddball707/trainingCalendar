import React from 'react';
import MaterialLayout from './components/Layout/MaterialLayout';
import CalendarGenerator from './components/CalendarGenerator';

function App() {
  return (
    <div>
      <MaterialLayout>
        <CalendarGenerator />
      </MaterialLayout>
    </div>
  );
}
export default App;