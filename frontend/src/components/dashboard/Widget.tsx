import { Card, Typography } from "@mui/material";

type Props = {
  title: string;
  color: string;
  bgcolor: string;
};

const Widget = ({ title, color, bgcolor }: Props) => {
  return (
    <Card
      sx={{
        py: 5,
        boxShadow: 0,
        textAlign: "center",
        color,
        bgcolor,
      }}
    >
      <Typography variant="h3">{title}</Typography>
      <Typography variant="subtitle2" sx={{ opacity: 0.72 }}>
        {title}
      </Typography>
    </Card>
  );
};

export default Widget;
