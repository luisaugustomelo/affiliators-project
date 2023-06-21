import styled from 'styled-components';
import { rgb, shade } from 'polished';

export const Container = styled.button`
  background: rgb(215, 255, 96);
  height: 56px;
  border-radius: 10px;
  border: 0;
  padding: 0 16px;
  color: #232129;
  width: 100%;
  font-weight: 500;
  margin-top: 16px;
  transition: background-color 0.2s;

  &:hover {
    background: ${shade(0.2, rgb(215, 255, 96))};
  }
`;
