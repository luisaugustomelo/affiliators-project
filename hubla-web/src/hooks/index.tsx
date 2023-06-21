import React, { PropsWithChildren } from 'react';

import { AuthProvider } from './auth';
import { ToastProvider } from './toast';

const AppProvider: React.FC<PropsWithChildren> = ({ children }) => (
  <AuthProvider>
    <ToastProvider>{children}</ToastProvider>
  </AuthProvider>
);

export default AppProvider;
