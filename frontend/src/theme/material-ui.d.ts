import "@mui/material/styles";
import { customShadows } from "./customShadows";
import { customPalette } from "./palette";

declare module "@mui/material/styles" {
  interface Theme {
    palette: customPalette;
    customShadows: customShadows; // optional
  }
}
