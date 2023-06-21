import React, { useState } from 'react';
import Header from '../Header';
import DragInDrop from '../DragInDrop';
import FileProcessing from '../FileProcessing';

interface DataContextType {
  id: string;
  setID: React.Dispatch<React.SetStateAction<string>>;
}

const DataContext = React.createContext<DataContextType | undefined>(undefined);

const Dashboard: React.FC = () => {
  const [id, setID] = useState('');
  return (
    <>
      <Header />
      <DataContext.Provider value={{ id, setID }}>
        <DragInDrop setID={setID} />
        <FileProcessing id={id} />
      </DataContext.Provider>
    </>
  );
};

export default Dashboard;
