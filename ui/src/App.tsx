import { useState } from 'react';
import { formatterView } from './features/formatter/formatterView';

function App() {
  const [activeTool, setActiveTool] = useState([0]);
  const [input, setInput] = useState([]);
  const [output, setOutput] = useState([0]);
  const [theme, isDark] = useState([0]);





  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <header className="mb-8 text-center">
        <h1 className="text-3xl font-bold text-gray-800">TextForge</h1>
        <p className="text-gray-600">Golang Backend + React Frontend</p>
      </header>
      
      <main className="max-w-4xl mx-auto">
        {}
        <FormatterView />
      </main>
    </div>
  );
}

export default App;