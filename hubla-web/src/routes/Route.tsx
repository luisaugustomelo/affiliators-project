import React, { PropsWithChildren } from 'react';
import { Navigate } from 'react-router-dom';

import { useAuth } from '../hooks/auth';

const AuthLayout: React.FC<PropsWithChildren> = ({ children }) => {
  const { user } = useAuth();
  if (!user) {
    return <Navigate to="/" replace={true} />;
  }
  return <>{children}</>;
};

export default AuthLayout;
