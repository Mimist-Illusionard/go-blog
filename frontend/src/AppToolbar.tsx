import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@mui/material';
import { useNavigate, useLocation } from 'react-router-dom';
import './AppToolbar.css';

const AppToolbar: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  const handleLogoClick = () => {
    if (!isLoginPage) {
    navigate('/feed');
    }
  };

  const isLoginPage = location.pathname === '/login' || location.pathname === '/';

  return (
    <AppBar position="static" className="app-toolbar">
      <Toolbar className="app-toolbar-inner">
        <Typography
          variant="h6"
          className="app-toolbar-logo"
          onClick={handleLogoClick}
          style={{ cursor: 'pointer' }}
          color="#05f2e9"
        >
          Codebase
        </Typography>

        {!isLoginPage && (
          <Button color="inherit" onClick={handleLogout} className="app-toolbar-logout">
            Выйти
          </Button>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default AppToolbar;
