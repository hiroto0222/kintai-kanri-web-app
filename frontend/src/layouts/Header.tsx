import { AppBar, Toolbar, alpha } from "@mui/material";
import { styled } from "@mui/material/styles";
import { HEADER_HEIGHT } from ".";
import NavButton from "../components/header/NavButton";

const StyledAppBar = styled(AppBar)(({ theme }) => ({
  backgroundColor: alpha(theme.palette.background.default, 0.95),
  boxShadow: "none",
}));

type Props = {
  onOpenNav: () => void;
};

const Header = ({ onOpenNav }: Props) => {
  return (
    <StyledAppBar>
      <Toolbar sx={{ minHeight: HEADER_HEIGHT }}>
        <NavButton onOpenNav={onOpenNav} />
      </Toolbar>
    </StyledAppBar>
  );
};

export default Header;
