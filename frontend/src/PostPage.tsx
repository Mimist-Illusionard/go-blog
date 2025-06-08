import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Container, Typography, Card, CardContent, Divider, Box } from '@mui/material';
import './PostFeed.css';

interface Comment {
  id: number;
  author: string;
  text: string;
}

interface Post {
  id: number;
  author: string;
  text: string;
  comments: Comment[];
}

const PostPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [post, setPost] = useState<Post | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch(`http://localhost:9090/posts/${id}`)
      .then(response => {
        if (!response.ok) {
          throw new Error('Ошибка загрузки поста');
        }
        return response.json();
      })
      .then(data => {
        setPost(data);
        setLoading(false);
      })
      .catch(err => {
        setError(err.message);
        setLoading(false);
      });
  }, [id]);

  if (loading) return <p className="center-text">Загрузка...</p>;
  if (error) return <p className="center-text error-text">Ошибка: {error}</p>;
  if (!post) return null;

  return (
    <Container className="dark-background" maxWidth="md">
      <Card className="post-card" sx={{ mb: 4 }}>
        <CardContent>
          <Typography variant="h5" className="post-author" gutterBottom>
            {post.author}
          </Typography>
          <Typography variant="body1" className="post-text">
            {post.text}
          </Typography>
        </CardContent>
      </Card>

      <Typography variant="h6" color="white" gutterBottom>
        Комментарии
      </Typography>

      {post.comments.length === 0 ? (
        <Typography className="post-text">Комментариев пока нет.</Typography>
      ) : (
        post.comments.map(comment => (
          <Box key={comment.id} className="post-card" sx={{ mb: 2, p: 2 }}>
            <Typography variant="subtitle2" className="post-author">
              {comment.author}
            </Typography>
            <Divider sx={{ my: 1, borderColor: '#555' }} />
            <Typography className="post-text">{comment.text}</Typography>
          </Box>
        ))
      )}
    </Container>
  );
};

export default PostPage;
