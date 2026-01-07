import React from "react";

interface statsBarProps {
    stats: TextStats;
}


  return (
    <div className="bg-white dark:bg-gray-900 border-t border-gray-200 dark:border-gray-700 px-6 py-3">
      <div className="flex flex-wrap items-center justify-center lg:justify-start gap-6">
        {statItems.map((item) => {
          return (
            <div key={item.label} className="flex items-center space-x-1">
              <span className="text-sm font-medium text-gray-700 dark:text-gray-300">
                {item.label}:
              </span>
              <span className="text-sm font-bold text-purple-600 dark:text-purple-400">
                {item.value.toLocaleString()}
              </span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default StatsBar;