import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PostFeed from './PostFeed';
import PostPage from './PostPage';
import AppToolbar from './AppToolbar';
import LoginPage from './LoginPage';

function App() {
  return (
    <Router>
      <>
        <AppToolbar />
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/" element={<PostFeed />} />
          <Route path="/posts/:id" element={<PostPage />} />
        </Routes>
      </>
    </Router>
  );
}

export default App;