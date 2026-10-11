import type { ReactNode } from 'react';
import { Link } from 'react-router-dom';

interface LayoutProps {
  children: ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  return (
    <div className='min-h-screen bg-gray-100'>
      <header className='bg-white shadow'>
        <div className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4'>
          <div className='flex justify-between items-center'>
            <nav className='flex space-x-4'>
              <Link to='/' className='text-gray-600 hover:text-gray-900'>
                Списки
              </Link>
              <Link to='/items' className='text-gray-600 hover:text-gray-900'>
                Всі товари
              </Link>
            </nav>
          </div>
        </div>
      </header>
      <main className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6'>
        {children}
      </main>
    </div>
  );
}
