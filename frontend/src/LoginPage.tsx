import React, { useState } from 'react';
import { TextField, Button, Container, Typography, Box, Stack } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import './LoginPage.css';

const LoginPage: React.FC = () => {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:9090/auth/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ login, password })
      });

      if (!response.ok) {
        throw new Error('Неверный логин или пароль');
      }

      const data = await response.json();
      localStorage.setItem('token', data.token);
      localStorage.setItem('userId', data.user_id);
      navigate('/feed');
    } catch (error) {
      alert(error instanceof Error ? error.message : 'Произошла ошибка');
    }
  };

  const handleRegister = async () => {
    try {
      const response = await fetch('http://localhost:9090/users/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ login, password })
      });

      if (!response.ok) {
        throw new Error('Ошибка регистрации');
      }

      const data = await response.json();
      localStorage.setItem('token', data.token);
      localStorage.setItem('userId', data.user_id);
      navigate('/');
    } catch (error) {
      alert(error instanceof Error ? error.message : 'Произошла ошибка при регистрации');
    }
  };

  return (
    <Container maxWidth="xs" className="login-container">
      <Box className="login-box">
        <Typography variant="h5" align="center" gutterBottom className="login-text">
          Вход в систему
        </Typography>
        <form onSubmit={handleSubmit} className="login-form">
          <TextField
            label="Логин"
            variant="outlined"
            fullWidth
            margin="normal"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
          />
          <TextField
            label="Пароль"
            type="password"
            variant="outlined"
            fullWidth
            margin="normal"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Stack spacing={2} mt={2}>
            <Button type="submit" variant="contained" color="primary" fullWidth>
              Войти
            </Button>
            <Button variant="outlined" color="primary" fullWidth onClick={handleRegister}>
              Зарегистрироваться
            </Button>
          </Stack>
        </form>
      </Box>
    </Container>
  );
};

export default LoginPage;
