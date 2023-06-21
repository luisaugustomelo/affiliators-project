import {
  NavLink,
  Title,
  Container,
  Navigation,
  Wrapper,
  Logout,
} from './sytles';
import { FaSignOutAlt } from 'react-icons/fa';

const Header: React.FC = () => {
  return (
    <Container>
      <Wrapper>
        <Title>Hubla</Title>
        <Navigation>
          <NavLink href="#">Home</NavLink>
          <NavLink href="#">About</NavLink>
          <NavLink href="#">Contact</NavLink>
          <Logout>
            <FaSignOutAlt />
          </Logout>
        </Navigation>
      </Wrapper>
    </Container>
  );
};

export default Header;
