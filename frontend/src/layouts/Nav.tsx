import {
  Box,
  Drawer,
  List,
  ListItemButton,
  ListItemText,
  styled,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

type Props = {
  openNav: boolean;
  onCloseNav: () => void;
};

const StyledNavItem = styled(ListItemButton)(({ theme }) => ({
  ...theme.typography.body2,
  height: 48,
  width: 200,
  textAlign: "center",
  position: "relative",
  color: theme.palette.text.secondary,
  borderRadius: theme.shape.borderRadius,
}));

const Nav = ({ openNav, onCloseNav }: Props) => {
  const { t } = useTranslation();
  const navigate = useNavigate();

  const navConfig = [
    {
      title: t("nav.mypage"),
      path: "/dashboard",
    },
    {
      title: t("nav.employees"),
      path: "/dashboard/employees",
    },
  ];

  return (
    <Box component="nav">
      <Drawer
        open={openNav}
        onClose={onCloseNav}
        ModalProps={{
          keepMounted: true,
        }}
      >
        <Box>
          <List sx={{ px: 1, py: 4 }}>
            {navConfig.map((item) => (
              <StyledNavItem onClick={() => navigate(item.path)}>
                <ListItemText primary={item.title} />
              </StyledNavItem>
            ))}
          </List>
        </Box>
      </Drawer>
    </Box>
  );
};

export default Nav;
