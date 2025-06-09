import React, { useState } from 'react';
import {
  Container,
  Typography,
  TextField,
  Button,
  Box,
  Stack,
  Alert
} from '@mui/material';
import { useNavigate } from 'react-router-dom';

const CreatePostPage: React.FC = () => {
  const [text, setText] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);
  const navigate = useNavigate();

  const handleCreatePost = async () => {
    const userId = Number(localStorage.getItem('userId')) || 1;
    const token = localStorage.getItem('token');

    if (!text.trim()) return;

    try {
      const response = await fetch('http://localhost:9090/posts/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          ...(token ? { Authorization: `Bearer ${token}` } : {})
        },
        body: JSON.stringify({
          user_id: userId,
          text
        })
      });

      if (!response.ok) {
        throw new Error('Не удалось создать пост');
      }

      setSuccess(true);
      setText('');
      setTimeout(() => navigate('/feed'), 1000);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Произошла ошибка');
    }
  };

  return (
    <Container className="dark-background" maxWidth="sm" sx={{ mt: 6 }}>
      <Box className="post-card" sx={{ p: 4 }}>
        <Typography variant="h5" className="post-author" gutterBottom>
          Новый пост
        </Typography>
        <Stack spacing={2}>
          {error && <Alert severity="error">{error}</Alert>}
          {success && <Alert severity="success">Пост опубликован</Alert>}
          <TextField
            multiline
            minRows={4}
            fullWidth
            variant="outlined"
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="О чём вы думаете?.."
            InputProps={{ style: { color: '#D9D9D9' } }}
            InputLabelProps={{ style: { color: '#aaaaaa' } }}
          />
          <Button variant="contained" color="primary" onClick={handleCreatePost}>
            Опубликовать
          </Button>
        </Stack>
      </Box>
    </Container>
  );
};

export default CreatePostPage;
