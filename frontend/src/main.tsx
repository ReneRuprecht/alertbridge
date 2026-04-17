import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { BrowserRouter, Route, Routes } from 'react-router'
import AlertHistoryView from './pages/AlertHistoryTableView.tsx'

createRoot(document.getElementById('root')!).render(
    <BrowserRouter>
        <Routes>
            <Route path='/' element={<App />} />
            <Route path='/:alertInstance' element={<AlertHistoryView />} />
        </Routes>
    </BrowserRouter>,
)
