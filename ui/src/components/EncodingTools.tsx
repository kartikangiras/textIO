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

const EncodingTools: React.FC<EncodingToolsProps> = ({ input, onOutput }) => {
  const [error, setError] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);

  const handleEncoding = async (action: string) => {
    setError('');
    setLoading(true);
    
    try {
      let result = '';
      
      switch (action) {
        case 'base64Encode':
          result = encodeBase64(input);
          break;
        case 'base64Decode':
          result = decodeBase64(input);
          break;
        case 'urlEncode':
          result = encodeURL(input);
          break;
        case 'urlDecode':
          result = decodeURL(input);
          break;
        case 'sha256':
          result = await generateSHA256(input);
          break;
        default:
          result = input;
      }
      
      onOutput(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setLoading(false);
    }
  };
}

