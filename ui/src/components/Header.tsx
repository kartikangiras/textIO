import React from 'react';
import { Moon, Sun, Github } from 'lucide-react';

interface HeaderProps {
  isDark: boolean;
  onThemeToggle: () => void;
}

const Header: React.FC<HeaderProps> = ({ isDark, onThemeToggle }) => {
  return (
    <div className="bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-700 px-6 py-4 flex-shrink-0">
      <div className="flex items-center justify-between">
        {/* Logo and Title */}
        <div className="flex items-center space-x-3">
          <h1 className="text-2xl font-bold bg-gradient-to-r from-purple-600 to-indigo-600 bg-clip-text text-transparent">
            TextForge
          </h1>
          <div className="hidden sm:block">
            <p className="text-sm text-gray-600 dark:text-gray-400">
              Ultimate Text & Data Manipulation Toolkit
            </p>
          </div>
        </div>