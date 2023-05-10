import { styled } from "@mui/material/styles";
import { Outlet } from "react-router-dom";

const StyledRoot = styled("div")({
  display: "flex",
  minHeight: "100%",
  overflow: "hidden",
});

const Main = styled("div")(({ theme }) => ({
  flexGrow: 1,
  overflow: "auto",
  minHeight: "100%",
  paddingTop: 24,
  paddingBottom: theme.spacing(10),
  [theme.breakpoints.up("lg")]: {
    paddingTop: 40 + 24,
    paddingLeft: theme.spacing(10),
    paddingRight: theme.spacing(10),
  },
}));

export const DashboardLayout = () => {
  return (
    <StyledRoot>
      <Main>
        <Outlet />
      </Main>
    </StyledRoot>
  );
};
