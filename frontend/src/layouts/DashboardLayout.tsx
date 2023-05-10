import { styled } from "@mui/material/styles";
import { useState } from "react";
import { Outlet } from "react-router-dom";
import { HEADER_HEIGHT } from ".";
import Header from "./Header";
import Nav from "./Nav";

const StyledRoot = styled("div")({
  display: "flex",
  minHeight: "100%",
  overflow: "hidden",
});

const Main = styled("div")(({ theme }) => ({
  flexGrow: 1,
  overflow: "auto",
  minHeight: "100%",
  paddingTop: HEADER_HEIGHT,
  paddingBottom: theme.spacing(10),
  [theme.breakpoints.up("md")]: {
    paddingTop: 40 + HEADER_HEIGHT,
    paddingLeft: theme.spacing(10),
    paddingRight: theme.spacing(10),
  },
  [theme.breakpoints.up("sm")]: {
    paddingTop: 40 + HEADER_HEIGHT,
  },
}));

type Props = {
  isAdmin: boolean;
};

export const DashboardLayout = ({ isAdmin }: Props) => {
  const [open, setOpen] = useState(false);

  return (
    <StyledRoot>
      {isAdmin && (
        <>
          <Header onOpenNav={() => setOpen(true)} />
          <Nav openNav={open} onCloseNav={() => setOpen(false)} />
        </>
      )}
      <Main>
        <Outlet />
      </Main>
    </StyledRoot>
  );
};
