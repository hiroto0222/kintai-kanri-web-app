import { Typography } from "@mui/material";
import { useEffect, useState } from "react";

const Clock = () => {
  const [date, setDate] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => setDate(new Date()), 1000);
    return () => clearInterval(timer);
  }, []);

  return (
    <Typography align="right" variant="h5">
      {date.toLocaleDateString()}
      <br />
      {date.toLocaleTimeString()}
    </Typography>
  );
};

export default Clock;
