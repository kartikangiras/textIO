import { useState } from 'react';
import { formatterView } from './features/formatter/formatterView';

function App() {
  const [activeTool, setActiveTool] = useState([0]);
  const [input, setInput] = useState([]);
  const [output, setOutput] = useState([0]);
  const [theme, isDark] = useState([0]);

  const stats = getTextStats(input);

  const toolNeedsInput = activeTool !== 'generators';

  const renderToolComponent = () => {
    const props = {input, onOutput: setOutput };

    switch(activeTool) {
      case 'cleanup':
        return <TextCleanup {...props} />;
      case 'case':
        return <CaseConverter {...props} />;
      case 'format':
        return <CodeFormatter {...props} />;
      case 'encoding':
        return <EncodingTools {...props} />;
      case 'generators':
        return <Generators onOutput={setOutput} />;
      default:
        return <TextCleanup {...props} />;
    }
  };

   return (
    <div className="h-screen flex flex-col bg-gray-50 dark:bg-gray-900 transition-colors duration-200">
      <Header isDark={isDark} onThemeToggle={toggleTheme} />

      <TabNavigation activeTool={activeTool} onToolChange={setActiveTool} />

      <div className="flex-1 flex flex-col overflow-hidden min-h-0">
        {activeTool === 'generators' ? (
          <div className="flex-1 flex flex-col overflow-hidden">
            <div className="bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 overflow-y-auto flex-shrink-0 max-h-[60vh]">
              <div className="p-6">
                {renderToolComponent()}
              </div>
            </div>

            <div className="flex-1 flex flex-col bg-white dark:bg-gray-800 min-h-0">
              <div className="p-6 flex-1 flex flex-col min-h-0">
                <TextArea
                  value={output}
                  onChange={() => {}} 
                  placeholder="Generated output will appear here..."
                  label="Generated Output"
                  readOnly
                />
              </div>
            </div>
          </div>
        ) : (
          <div className="flex-1 flex flex-col lg:flex-row overflow-hidden">
            <div className="flex-1 flex flex-col bg-white dark:bg-gray-800 border-b lg:border-b-0 lg:border-r border-gray-200 dark:border-gray-700 min-h-0">
              <div className="p-6 flex-1 flex flex-col min-h-0">
                <TextArea
                  value={input}
                  onChange={setInput}
                  placeholder="Paste or type your text here..."
                  label="Input"
                  onClear={() => {
                    setInput('');
                    setOutput('');
                  }}
                />
              </div>
            </div>
            <div className="flex-1 flex flex-col bg-white dark:bg-gray-800 min-h-0">
              <div className="border-b border-gray-200 dark:border-gray-700 overflow-y-auto flex-shrink-0 max-h-[40vh] lg:max-h-[50vh]">
                <div className="p-6">
                  {renderToolComponent()}
                </div>
              </div>

              <div className="flex-1 flex flex-col min-h-0">
                <div className="p-6 flex-1 flex flex-col min-h-0">
                  <TextArea
                    value={output}
                    onChange={() => {}}
                    placeholder="Processed output will appear here..."
                    label="Output"
                    readOnly
                  />
                </div>
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
