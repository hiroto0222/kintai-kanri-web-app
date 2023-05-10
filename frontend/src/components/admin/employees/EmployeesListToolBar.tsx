import { InputAdornment, OutlinedInput, Toolbar } from "@mui/material";
import { alpha, styled } from "@mui/material/styles";

const StyledRoot = styled(Toolbar)(({ theme }) => ({
  height: 70,
  display: "flex",
  justifyContent: "space-between",
  padding: theme.spacing(0, 1, 0, 3),
}));

const StyledSearch = styled(OutlinedInput)(({ theme }) => ({
  width: 240,
  height: 45,
  backgroundColor: theme.palette.grey[100],
  transition: theme.transitions.create(["box-shadow", "width"], {
    easing: theme.transitions.easing.easeInOut,
    duration: theme.transitions.duration.shorter,
  }),
  "&.Mui-focused": {
    width: 320,
    boxShadow: theme.customShadows.z8,
  },
  "& fieldset": {
    borderWidth: `1px !important`,
    borderColor: `${alpha(theme.palette.grey[500], 0.32)} !important`,
  },
}));

type Props = {
  filterName: string;
  onFilterName: (event: React.ChangeEvent<HTMLInputElement>) => void;
};

// TODO: Implement search functionality
const EmployeesListToolbar = () => {
  return (
    <StyledRoot>
      <StyledSearch
        value=""
        // onChange={() => {}}
        placeholder="Search employee..."
        startAdornment={<InputAdornment position="start"></InputAdornment>}
      />
    </StyledRoot>
  );
};

export default EmployeesListToolbar;
