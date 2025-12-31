import React from 'react';
import { cleanupText } from '../utils/textUtils';
import ToolButton from './ToolButton';

interface TextCleanupProps {
  input: string;
  onOutput: (output: string) => void;
}

const TextCleanup: React.FC<TextCleanupProps> = ({ input, onOutput }) => {
  const handleCleanup = (action: string) => {
    const result = cleanupText(input, action);
    onOutput(result);
  };

  const tools = [
    {
      action: 'removeExtraSpaces',
      label: 'Remove Extra Spaces',
      description: 'Collapse multiple spaces into single spaces'
    },
    {
      action: 'removeLineBreaks',
      label: 'Remove Line Breaks',
      description: 'Convert text to single line'
    },
    {
      action: 'removeAllSpaces',
      label: 'Remove All Spaces',
      description: 'Remove all whitespace characters'
    },
    {
      action: 'trimLines',
      label: 'Trim Each Line',
      description: 'Remove leading/trailing spaces from each line'
    }
  ];