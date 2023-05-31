import { CircularProgress } from "@mui/material";

const Loading = () => {
  return (
    <CircularProgress sx={{ position: "absolute", top: "50%", left: "50%" }} />
  );
};

export default Loading;
