import React, { useState } from 'react';
import { Copy, RotateCcw, Check } from 'lucide-react';
import { copyToClipboard } from '../utils';
import ToolButton from './ToolButton';

interface TextAreaProps {
  id: string;
  name?: string;
  value: string;
  onChange: (value: string) => void;
  placeholder: string;
  label: string;
  readOnly?: boolean;
  onClear?: () => void;
}

const TextArea: React.FC<TextAreaProps> = ({
  id,
  value,
  onChange,
  placeholder,
  label,
  readOnly = false,
  onClear
}) => {
  const [isCopied, setIsCopied] = useState(false);

  const handleCopy = async () => {
    if (!value) return;

    const success = await copyToClipboard(value);
    if (success) {
      setIsCopied(true);
      setTimeout(() => setIsCopied(false), 2000);
    }
  };

   return (
    <div className="flex flex-col h-full min-h-0">
      <div className="flex items-center justify-between mb-3 flex-shrink-0">
        <span className="text-sm font-medium text-gray-700 dark:text-gray-300">
          {label}
        </span>
        <div className="flex space-x-2">
          {onClear && (
            <ToolButton
              onClick={onClear}
              variant="secondary"
              disabled={!value}
              className="text-xs px-3 py-1"
            >
              <RotateCcw className="w-3 h-3 mr-1" />
              Clear
            </ToolButton>
          )}
          
          <ToolButton
            onClick={handleCopy}
            variant="secondary"
            disabled={!value}
            className={`text-xs px-3 py-1 transition-all duration-200 ${
              isCopied ? 'text-green-600 bg-green-50 border-green-200' : ''
            }`}
          >
            {isCopied ? (
              <>
                <Check className="w-3 h-3 mr-1" />
                Copied!
              </>
            ) : (
              <>
                <Copy className="w-3 h-3 mr-1" />
                Copy
              </>
            )}
          </ToolButton>
        </div>
      </div>
      
      <textarea
        id={id}
        name="message"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
        readOnly={readOnly}
        className={`flex-1 w-full p-4 border border-gray-300 dark:border-gray-600 rounded-lg
                   bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100
                   placeholder-gray-500 dark:placeholder-gray-400 resize-none
                   focus:ring-2 focus:ring-purple-500 focus:border-transparent
                   min-h-0 font-mono text-sm
                   ${readOnly ? 'bg-gray-50 dark:bg-gray-900 cursor-default' : ''}`}
        style={{ minHeight: '200px' }}
      />
    </div>
  );
};

export default TextArea;