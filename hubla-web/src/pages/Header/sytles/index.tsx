import styled from 'styled-components';

export const Container = styled.header`
  width: 100%;
  background: #f5f5f5;
  display: flex;
  justify-content: center;
`;

export const Wrapper = styled.div`
  width: 75%;
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  @media (max-width: 768px) {
    width: 80%;
    flex-direction: column;
  }
`;

export const Title = styled.h1`
  font-size: 24px;
  color: rgb(49, 64, 0);

  @media (max-width: 768px) {
    margin-bottom: 20px;
  }
`;

export const Navigation = styled.nav`
  @media (max-width: 768px) {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }
`;

export const NavLink = styled.a`
  margin-left: 20px;
  color: rgb(49, 64, 0);
  text-decoration: none;

  :hover {
    background: rgb(215, 255, 96);
  }
`;

export const Logout = styled.a`
  margin-left: 30px;
  color: rgb(49, 64, 0);
  text-decoration: none;

  :hover {
    background: rgb(215, 255, 96);
  }
`;
