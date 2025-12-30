import React from 'react';
import { ToolType } from '../types';

interface TabNavigationProps {
  activeTool: ToolType;
  onToolChange: (tool: ToolType) => void;
}

const tools = [
  { id: 'cleanup', name: 'Text Cleanup', description: 'Clean and analyze text' },
  { id: 'case', name: 'Case Converter', description: 'Convert text cases' },
  { id: 'format', name: 'Code Formatter', description: 'Format JSON, CSS & more' },
  { id: 'encoding', name: 'Encoding & Hash', description: 'Encode, decode & hash' },
  { id: 'generators', name: 'Generators', description: 'Generate UUIDs & passwords' },
];
