
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';

import HomePage from './pages/HomePage';
import VideoUpload from './pages/VideoUpload';
function App() {
  return (
    <div className="App">
      <header className="App-header">
      <h1>Video Uploader</h1>
      </header>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="video/*" element={<VideoUpload />} />
        </Routes>
      </BrowserRouter>

    </div>
  );
}

export default App;
