import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import {
  Card,
  CardContent,
  Typography,
  Container,
  CardActionArea,
  CircularProgress,
  Alert,
  Button,
  Box
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
  const navigate = useNavigate();

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await fetch('http://localhost:9090/posts/');
        if (!response.ok) {
          throw new Error('Ошибка загрузки данных');
        }

        const data = await response.json();

        if (!data || !Array.isArray(data)) {
          setPosts([]);
        } else {
          setPosts(data);
        }
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
      <Box display="flex" justifyContent="space-between" alignItems="center" sx={{ mt: 4, mb: 2 }}>
        <Typography variant="h3" fontWeight="bold" color="white">
          Forum
        </Typography>
        <Button
          variant="contained"
          sx={{
            backgroundColor: '#05e9e1',
            color: '#000',
            '&:hover': {
              backgroundColor: '#04d2cc'
            }
          }}
          onClick={() => navigate('/create-post')}
        >
          Создать пост
        </Button>
      </Box>

      {posts.length === 0 ? (
        <Typography variant="h6" color="white" align="center" sx={{ mt: 4 }}>
          Нет постов для отображения.
        </Typography>
      ) : (
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
      )}
    </Container>
  );
};

export default PostFeed;
