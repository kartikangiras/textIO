import React, { useState } from 'react';
import { formatJSON, convertKvToJson, minifyCSS } from '../utils/formatters';
import ToolButton from './ToolButton';

interface CodeFormatterProps {
    input: string;
    onOutput: (output: string) => void;
}

const CodeFormatter: React.FC<CodeFormatterProps> = ({ input, onOutput }) => {
  const [error, setError] = useState<string>('');
}

const handleFormat = async (action: string) => {
    setError('');
    try {
      let result = '';
      
      switch (action) {
        case 'jsonBeautify':
          result = formatJSON(input, true);
          break;
        case 'jsonMinify':
          result = formatJSON(input, false);
          break;
        case 'kvToJson':
          result = convertKvToJson(input);
          break;
        case 'cssMinify':
          result = minifyCSS(input);
          break;
        default:
          result = input;
      }
      
      onOutput(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    }
  };