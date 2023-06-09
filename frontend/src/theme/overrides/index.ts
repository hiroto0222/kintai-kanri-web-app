import { Theme } from "@mui/material/styles";
import Autocomplete from "./Autocomplete";
import Backdrop from "./Backdrop";
import Button from "./Button";
import Card from "./Card";
import Chip from "./Chip";
import Input from "./Input";
import Paper from "./Paper";
import Table from "./Table";
import Tooltip from "./Tooltip";
import Typography from "./Typography";

export default function componentsOverrides(theme: Theme) {
  return Object.assign(
    Card(theme),
    Table(theme),
    Input(theme),
    Paper(),
    Button(theme),
    Tooltip(theme),
    Backdrop(theme),
    Typography(theme),
    Autocomplete(theme),
    Chip(theme)
  );
}
