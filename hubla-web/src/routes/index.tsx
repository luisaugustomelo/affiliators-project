import React from 'react';
import { Routes, Route } from 'react-router-dom';

import AuthLayout from './Route';

import SignIn from '../pages/SignIn';
import SignUp from '../pages/SignUp';

import Dashboard from '../pages/Dashboard';

const AppRoutes: React.FC = () => (
  <Routes>
    {/* <Route element={<AuthLayout />} /> */}
    <Route path="/" element={<SignIn />} />
    <Route path="/signup" element={<SignUp />} />
    <Route
      path="/dashboard"
      element={
        <AuthLayout>
          <Dashboard />
        </AuthLayout>
      }
    />
    <Route path="/signin" element={<SignIn />} />
  </Routes>
);

export default AppRoutes;
