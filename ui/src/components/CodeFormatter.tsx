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