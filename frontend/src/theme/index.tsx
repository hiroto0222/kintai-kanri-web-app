import { CssBaseline } from "@mui/material";
import {
  ThemeProvider as MUIThemeProvider,
  StyledEngineProvider,
  ThemeOptions,
  createTheme,
} from "@mui/material/styles";
import { useMemo } from "react";
import customShadows from "./customShadows";
import GlobalStyles from "./globalStyles";
import componentsOverrides from "./overrides";
import palette from "./pallete";
import shadows from "./shadows";
import typography from "./typography";

type Props = {
  children: React.ReactNode;
};

const ThemeProvider: React.FC<Props> = ({ children }) => {
  const themeOptions: ThemeOptions = useMemo(
    () => ({
      palette,
      shape: { borderRadius: 6 },
      typography,
      shadows: shadows(),
      customShadows: customShadows(),
    }),
    []
  );

  const theme = createTheme(themeOptions);
  theme.components = componentsOverrides(theme);

  return (
    <StyledEngineProvider injectFirst>
      <MUIThemeProvider theme={theme}>
        <CssBaseline />
        <GlobalStyles />
        {children}
      </MUIThemeProvider>
    </StyledEngineProvider>
  );
};

export default ThemeProvider;
