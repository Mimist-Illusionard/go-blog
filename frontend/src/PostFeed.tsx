import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { 
  Card, 
  CardContent, 
  Typography, 
  Container, 
  CardActionArea,
  CircularProgress,
  Alert
} from '@mui/material';
import Grid from '@mui/material/Grid';
import './PostFeed.css';

interface Post {
  id: number;
  author: string;
  text: string;
  createdAt?: string;
}

const PostFeed: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await fetch('http://localhost:9090/posts/');
        if (!response.ok) {
          throw new Error('Ошибка загрузки данных');
        }
        const data = await response.json();
        setPosts(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Неизвестная ошибка');
      } finally {
        setLoading(false);
      }
    };

    fetchPosts();
  }, []);

  if (loading) {
    return (
      <Container className="center-text">
        <CircularProgress />
      </Container>
    );
  }

  if (error) {
    return (
      <Container className="center-text">
        <Alert severity="error">{error}</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" className="post-feed-container dark-background">
      <Typography 
        variant="h3" 
        align="center" 
        gutterBottom 
        fontWeight="bold" 
        color="white"
        sx={{ marginTop: 4, marginBottom: 6 }}
      >
        Forum
      </Typography>
      
      <Grid container spacing={4} justifyContent="center">
        {posts.map((post) => (
          <Grid key={post.id} size={{ sm: 6, md: 4 }}>
            <Card className="post-card" sx={{ height: '100%' }}>
              <CardActionArea 
                component={Link}
                to={`/posts/${post.id}`}
                sx={{ height: '100%' }}
              >
                <CardContent sx={{ height: '100%' }}>
                  <Typography 
                    variant="h6" 
                    className="post-author" 
                    gutterBottom
                    sx={{ fontWeight: 'bold' }}
                  >
                    {post.author}
                  </Typography>
                  <Typography 
                    variant="body1" 
                    className="post-text"
                    sx={{ 
                      display: '-webkit-box',
                      WebkitLineClamp: 3,
                      WebkitBoxOrient: 'vertical',
                      overflow: 'hidden',
                      textOverflow: 'ellipsis'
                    }}
                  >
                    {post.text}
                  </Typography>
                  {post.createdAt && (
                    <Typography 
                      variant="caption" 
                      color="text.secondary"
                      sx={{ display: 'block', mt: 2 }}
                    >
                      {new Date(post.createdAt).toLocaleDateString()}
                    </Typography>
                  )}
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