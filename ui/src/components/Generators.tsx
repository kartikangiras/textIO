import React, { useState } from 'react';
import { generateUUID, generatePassword } from '../utils/generators';
import ToolButton from './ToolButton';

interface GeneratorsProps {
  onOutput: (output: string) => void;
}
