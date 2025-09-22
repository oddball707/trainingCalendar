import React from 'react';
import MaterialLayout from './components/Layout/MaterialLayout';
import CalendarGenerator from './components/CalendarGenerator/CalendarGenerator';

function CalendarApp() {
  return (
    <MaterialLayout>
      <CalendarGenerator />
    </MaterialLayout>
  );
}

export default CalendarApp;
