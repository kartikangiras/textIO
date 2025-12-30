import React, { useState } from 'react';
import { generateUUID, generatePassword } from '../utils/generators';
import ToolButton from './ToolButton';

interface GeneratorsProps {
  onOutput: (output: string) => void;
}

const Generators: React.FC<GeneratorsProps> = ({ onOutput }) => {
  const [passwordLength, setPasswordLength] = useState<number>(16);
  const [passwordOptions, setPasswordOptions] = useState({
    lowercase: true,
    uppercase: true,
    numbers: true,
    symbols: true
  });
  const [error, setError] = useState<string>('');

  const handleUUIDGeneration = (version: 1 | 4) => {
    try {
      const uuid = generateUUID(version);
      onOutput(uuid);
      setError('');
    } catch (err) {
      setError('Failed to generate UUID');
    }
  };
