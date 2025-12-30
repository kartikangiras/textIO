import React, { useState } from 'react';
import { 
  encodeBase64, 
  decodeBase64, 
  encodeURL, 
  decodeURL, 
  generateSHA256 
} from '../utils/formatters';
import ToolButton from './ToolButton';

interface EncodingToolsProps {
  input: string;
  onOutput: (output: string) => void;
}

