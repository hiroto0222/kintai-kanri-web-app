import { Theme } from "@mui/material/styles";

export default function Chip(theme: Theme) {
  return {
    MuiChip: {
      styleOverrides: {
        colorSuccess: {
          backgroundColor: theme.palette.success.light,
        },
        colorError: { backgroundColor: theme.palette.error.light },
      },
    },
  };
}
