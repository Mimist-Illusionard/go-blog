import React, { useEffect, useState } from 'react';
import { Card, CardContent, Typography, Container, Grid, CardActionArea } from '@mui/material';
import './PostFeed.css';

interface Post {
  author: string;
  text: string;
}

const PostFeed: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('http://localhost:9090/posts/')
      .then(response => {
        if (!response.ok) {
          throw new Error('Ошибка загрузки данных');
        }
        return response.json();
      })
      .then(data => {
        setPosts(data);
        setLoading(false);
      })
      .catch(err => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  if (loading) return <p className="center-text">Загрузка...</p>;
  if (error) return <p className="center-text error-text">Ошибка: {error}</p>;

  return (
    <Container maxWidth="lg" className="post-feed-container dark-background">
      <Typography variant="h3" align="center" gutterBottom fontWeight="bold" color="white">
        Codebase Forum
      </Typography>
      <Grid container spacing={4} justifyContent="center">
        {posts.map((post, index) => (
          <Grid key={index} size={{ sm: 6, md: 4 }}>
            <Card className="post-card fixed-height">
              <CardActionArea onClick={() => alert(`Открыть пост от ${post.author}`)}>
                <CardContent>
                  <Typography variant="h6" className="post-author" gutterBottom>
                    Автор: {post.author}
                  </Typography>
                  <Typography variant="body1" className="post-text">
                    {post.text}
                  </Typography>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default PostFeed;
