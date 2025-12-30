import React from 'react';
import { Copy, RotateCcw } from 'lucide-react';
import { copyToClipboard } from '../utils/textUtils';
import ToolButton from './ToolButton';

interface TextAreaProps {
  value: string;
  onChange: (value: string) => void;
  placeholder: string;
  label: string;
  readOnly?: boolean;
  onClear?: () => void;
}

const TextArea: React.FC<TextAreaProps> = ({
  value,
  onChange,
  placeholder,
  label,
  readOnly = false,
  onClear
}) => {
  const handleCopy = async () => {
    if (value) {
      const success = await copyToClipboard(value);
      if (success) {
      }
    }
  };