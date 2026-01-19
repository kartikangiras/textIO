import { useState} from 'react';

import TextCleanup from './components/TextCleanup';
import CaseConverter  from './components/CaseConverter';
import CodeFormatter from './components/CodeFormatter';
import EncodingTools  from './components/EncodingTools';
import Generators from './components/Generators';

import  Header  from './components/Header';
import  TabNavigation  from './components/TabNavigation';
import  TextArea  from './components/TextArea';
import  StatsBar  from './components/StatsBar';

function App() {
  const [activeTool, setActiveTool] = useState<string>('cleanup');
  const [input, setInput] = useState<string>('');
  const [output, setOutput] = useState<string>('');
  const [isDark, setIsDark] = useState<boolean>(false);

  const stats = {
    chars: input.length,
    words: input.trim() === '' ? 0 : input.trim().split(/\s+/).length,
    lines: input.split('\n').length,
  };

  const toolNeedsInput = activeTool !== 'generators';

  const toggleTheme = () => {
    setIsDark(!isDark);
    document.documentElement.classList.toggle('dark');
  };

  const renderToolComponent = () => {
    const commonProps = { 
      input: input,          
      onOutput: setOutput    
    };

    switch(activeTool) {
      case 'cleanup':
        return <TextCleanup {...commonProps} />;
      case 'case':
        return <CaseConverter {...commonProps} />;
      case 'format':
        return <CodeFormatter {...commonProps} />;
      case 'encoding':
        return <EncodingTools {...commonProps} />;
      case 'generators':
        return <Generators onOutput={setOutput} />;
      default:
        return <TextCleanup {...commonProps} />;
    }
  };

   return (
    <div className={`h-screen flex flex-col transition-colors duration-200 ${isDark ? 'bg-gray-900 text-white' : 'bg-gray-50 text-gray-900'}`}>
      
      <Header isDark={isDark} onThemeToggle={toggleTheme} />

      <TabNavigation activeTool={activeTool} onToolChange={setActiveTool} />

      <div className="flex-1 flex flex-col overflow-hidden min-h-0">

        {activeTool === 'generators' ? (
          <div className="flex-1 flex flex-col overflow-hidden max-w-4xl mx-auto w-full p-4">
            <div className="bg-white dark:bg-gray-800 rounded-lg shadow p-6 mb-4">
              {renderToolComponent()}
            </div>
            <TextArea
              label="Generated Output"
              value={output}
              readOnly
              placeholder="Result will appear here..." 
              onChange={() => {}}
            />
          </div>
        ) : (
          
          <div className="flex-1 flex flex-col lg:flex-row overflow-hidden">

            <div className="flex-1 p-4 flex flex-col border-r border-gray-200 dark:border-gray-700">
              <TextArea
                label="Input"
                value={input}
                onChange={setInput}
                placeholder="Paste your text here..."
                onClear={() => { setInput(''); setOutput(''); }}
              />
            </div>

            <div className="flex-1 flex flex-col min-h-0">

              <div className="border-b border-gray-200 dark:border-gray-700 p-4 bg-gray-50 dark:bg-gray-800/50">
                {renderToolComponent()}
              </div>
              <div className="flex-1 p-4 flex flex-col">
                <TextArea
                  label="Output"
                  value={output}
                  readOnly
                  placeholder="Processed result..."
                  onChange={() => {}}
                />
              </div>
            </div>
          </div>
        )}
      </div>

      {toolNeedsInput && <StatsBar stats={stats} />}
    </div>
  );
}

export default App;