import React from 'react';
import { cleanupText } from '../utils/textUtils';
import ToolButton from './ToolButton';

interface TextCleanupProps {
  input: string;
  onOutput: (output: string) => void;
}

