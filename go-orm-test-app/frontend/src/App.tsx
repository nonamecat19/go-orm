import { Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import ItemsPage from './pages/ItemsPage';
import ListDetailPage from './pages/ListDetailPage';

export default function App() {
  return (
    <Routes>
      <Route path='/' element={<HomePage />} />
      <Route path='/lists/:id' element={<ListDetailPage />} />
      <Route path='/items' element={<ItemsPage />} />
    </Routes>
  );
}
