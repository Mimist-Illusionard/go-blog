import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PostFeed from './PostFeed';
import PostPage from './PostPage';
import AppToolbar from './AppToolbar';
import LoginPage from './LoginPage';
import CreatePostPage from './CreatePostPage';

function App() {
  return (
    <Router>
      <>
        <AppToolbar />
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/" element={<LoginPage />} />
          <Route path="/feed" element={<PostFeed />} />
          <Route path="/posts/:id" element={<PostPage />} />
          <Route path="/create-post" element={<CreatePostPage />} />
        </Routes>
      </>
    </Router>
  );
}

export default App;