import React from "react";
import { convertCase } from '../utils/textUtils';
import ToolButton from './ToolButton';

interface caseConverterProps{
    input: string,
    onOutput: (output: string) => void
}

const caseConverter: React.FC<caseConverterProps> = ({input, onOutput}) => {
    const handleConversion = (caseType: string) => {
        const result = convertCase(input, caseType);
        onOutput(result);
    };

    
}