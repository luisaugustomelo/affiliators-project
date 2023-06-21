import React, {
  createContext,
  useCallback,
  useState,
  PropsWithChildren,
  useContext,
} from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';

interface AuthState {
  token: string;
  user: object;
}
interface SignInCredentials {
  email: string;
  password: string;
}

interface AuthContextData {
  user: object;
  signIn(credentials: SignInCredentials): Promise<void>;
  signOut(): void;
}

const AuthContext = createContext<AuthContextData>({} as AuthContextData);

export const AuthProvider: React.FC<PropsWithChildren> = ({ children }) => {
  const navigate = useNavigate();
  const [data, setData] = useState<AuthState>(() => {
    const token = localStorage.getItem('@Hubla:token');
    const user = localStorage.getItem('@Hubla:user');

    if (token && user) {
      return { token, user: JSON.parse(user) };
    }

    return {} as AuthState;
  });

  const signIn = useCallback(
    async ({ email, password }: SignInCredentials) => {
      const response = await api.post('/login', {
        email,
        password,
      });

      const { token, user } = response.data;

      localStorage.setItem('@Hubla:token', token);
      localStorage.setItem('@Hubla:user', JSON.stringify(user));

      setData({ token, user });

      navigate('/dashboard');
    },
    [navigate],
  );

  const signOut = useCallback(async () => {
    localStorage.removeItem('@Hubla:token');
    localStorage.removeItem('@Hubla:user');

    setData({} as AuthState);
  }, []);

  return (
    <AuthContext.Provider value={{ user: data.user, signIn, signOut }}>
      {children}
    </AuthContext.Provider>
  );
};

export function useAuth(): AuthContextData {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }

  return context;
}
