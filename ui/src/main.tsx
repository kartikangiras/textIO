import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.tsx';
import './index.css';

const initializeTheme = () => {
  const saved = localStorage.getItem('textforge-theme');
  const isDark = saved ? JSON.parse(saved) : window.matchMedia('(prefers-color-scheme: dark)').matches;
  
  if (isDark) {
    document.documentElement.classList.add('dark');
  } else {
    document.documentElement.classList.remove('dark');
  }
};

initializeTheme();

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>
);