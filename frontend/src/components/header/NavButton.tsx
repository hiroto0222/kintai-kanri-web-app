import { IconButton } from "@mui/material";

type Props = {
  onOpenNav: () => void;
};

const NavButton = ({ onOpenNav }: Props) => {
  return (
    <IconButton
      onClick={onOpenNav}
      sx={{
        mr: 1,
        color: "text.primary",
      }}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="24"
        viewBox="0 0 24 24"
        width="24"
      >
        <path d="M0 0h24v24H0z" fill="none" />
        <path d="M3 18h18v-2H3v2zm0-5h18v-2H3v2zm0-7v2h18V6H3z" />
      </svg>
    </IconButton>
  );
};

export default NavButton;
