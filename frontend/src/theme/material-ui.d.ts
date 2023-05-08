import "@mui/material/styles";
import { customShadows } from "./customShadows";

declare module "@mui/material/styles" {
  interface Theme {
    customShadows: customShadows; // optional
  }
}
